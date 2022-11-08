package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const (
	owner    = "wirnat"
	repo     = "aksara-cli-doc"
	basePath = ""
)

var client *github.Client

type transportHeaders struct {
	modifiedSince string
}

func (t *transportHeaders) RoundTrip(req *http.Request) (*http.Response, error) {

	// Determine the last modified date based on the transportHeader options
	// Do not add any headers if blank or zero
	if t.modifiedSince != "" {
		req.Header.Set("If-Modified-Since", t.modifiedSince)
	}

	return http.DefaultTransport.RoundTrip(req)
}

func main() {
	transport := &oauth2.Transport{
		Source: oauth2.StaticTokenSource(
			&oauth2.Token{
				AccessToken: "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",
			},
		),
		Base: &transportHeaders{
			modifiedSince: "",
		},
	}
	client = github.NewClient(&http.Client{
		Transport: transport,
	})
	getContents("")
}

func getContents(path string) {
	fmt.Println("\n\n")

	fileContent, directoryContent, resp, err := client.Repositories.GetContents(context.Background(), owner, repo, path, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", fileContent)
	fmt.Printf("%#v\n", directoryContent)
	fmt.Printf("%#v\n", resp)

	for _, c := range directoryContent {
		fmt.Println(*c.Type, *c.Path, *c.Size, *c.SHA)

		local := filepath.Join(basePath, *c.Path)
		fmt.Println("local:", local)

		switch *c.Type {
		case "file":
			_, err := os.Stat(local)
			if err == nil {
				b, err1 := ioutil.ReadFile(local)
				if err1 == nil {
					sha := calculateGitSHA1(b)
					if *c.SHA == hex.EncodeToString(sha) {
						fmt.Println("no need to update this file, the SHA is the same")
						continue
					}
				}
			}
			downloadContents(c, local)
		case "dir":
			getContents(filepath.Join(path, *c.Path))
		}
	}
}

func downloadContents(content *github.RepositoryContent, localPath string) {
	if content.Content != nil {
		fmt.Println("content:", *content.Content)
	}

	rc, err := client.Repositories.DownloadContents(context.Background(), owner, repo, *content.Path, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rc.Close()

	_, err = ioutil.ReadAll(rc)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.MkdirAll(filepath.Dir(localPath), os.ModePerm)
	if err != nil {
		fmt.Printf("	❌ create directory failed")
		return
	}

	_, err = os.Create(localPath)
	if err != nil {
		fmt.Printf("	❌ create file failed")
		return
	}

	//wd, _ := os.Getwd()
	//wd = wd + "/remote"
	//wdA := fmt.Sprintf("%v/%v", wd, strings.ReplaceAll(filepath.Dir(localPath), ".", ""))
	//wdF := fmt.Sprintf("%v/%v", wd, strings.ReplaceAll(localPath, ".", ""))
	//
	//err = os.MkdirAll(wdA, 0777)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println("Writing the file:", localPath)
	//f, err := os.Create(wdF)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer f.Close()
	//n, err := f.Write(b)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//if n != *content.Size {
	//	fmt.Printf("number of bytes differ, %d vs %d\n", n, *content.Size)
	//}
}

// calculateGitSHA1 computes the github sha1 from a slice of bytes.
// The bytes are prepended with: "blob " + filesize + "\0" before runing through sha1.
func calculateGitSHA1(contents []byte) []byte {
	contentLen := len(contents)
	blobSlice := []byte("blob " + strconv.Itoa(contentLen))
	blobSlice = append(blobSlice, '\x00')
	blobSlice = append(blobSlice, contents...)
	h := sha1.New()
	h.Write(blobSlice)
	bs := h.Sum(nil)
	return bs
}
