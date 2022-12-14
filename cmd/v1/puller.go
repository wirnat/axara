package v1

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/dgraph-io/badger"
	"github.com/google/go-github/github"
	"github.com/wirnat/axara/cmd/v1/key"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Puller interface {
	Pull(p string, targetDir string) error
}

const (
	basePath = ""
)

type gitPuller struct {
	client *github.Client
}

func NewGitPuller() (g *gitPuller) {
	g = &gitPuller{}

	storage := "/.axara"
	opts := badger.DefaultOptions(storage)
	opts.Logger = nil

	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.View(func(txn *badger.Txn) error {
		token, err := txn.Get([]byte(key.GitKey))
		if err != nil {
			log.Fatal(err)
		}
		err = token.Value(func(val []byte) error {
			transport := &oauth2.Transport{
				Source: oauth2.StaticTokenSource(
					&oauth2.Token{
						AccessToken: string(val),
					},
				),
				Base: &transportHeaders{
					modifiedSince: "",
				},
			}
			g.client = github.NewClient(&http.Client{
				Transport: transport,
			})
			return nil
		})
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	return
}

type cred struct {
	owner     string
	repo      string
	path      string
	targetDir string
}

func extractLink(path string) (l cred) {
	paths := strings.SplitAfter(path, "github.com/")
	if len(paths) > 1 {
		gitPath := paths[1]
		gp := strings.Split(gitPath, "/")
		l.owner = gp[0]
		l.repo = gp[1]
		l.path = ""
	}

	return
}

func (g gitPuller) Pull(p string, targetDir string) error {
	if targetDir == "" {
		targetDir = "aksara-storage"
	}
	cr := extractLink(p)
	cr.targetDir = targetDir
	err := g.getContent(cr)
	if err != nil {
		return err
	}
	return nil
}

func (g gitPuller) getContent(cr cred) error {
	_, directoryContent, _, err := g.client.Repositories.GetContents(context.Background(), cr.owner, cr.repo, cr.path, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, c := range directoryContent {
		local := filepath.Join(basePath, *c.Path)
		switch *c.Type {
		case "file":
			_, err := os.Stat(local)
			if err == nil {
				b, err1 := ioutil.ReadFile(local)
				if err1 == nil {
					sha := calculateGitSHA1(b)
					if *c.SHA == hex.EncodeToString(sha) {
						continue
					}
				}
			}
			g.downloadContents(c, local, cr.owner, cr.repo, cr.targetDir)
		case "dir":
			cr.path = *c.Path
			g.getContent(cr)

		}
	}
	return nil
}

func (g gitPuller) downloadContents(content *github.RepositoryContent, localPath string, owner string, repo string, targetDir string) {
	if content.Content != nil {
		fmt.Println("content:", *content.Content)
	}

	rc, err := g.client.Repositories.DownloadContents(context.Background(), owner, repo, *content.Path, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rc.Close()

	ss.Title = fmt.Sprintf("Download %v... ", localPath)

	dataB, err := ioutil.ReadAll(rc)
	if err != nil {
		return
	}

	err = os.MkdirAll(targetDir+"/"+filepath.Dir(localPath), os.ModePerm)
	if err != nil {
		fmt.Println(fmt.Println("	??? create directory failed"))
		return
	}

	fl, err := os.Create(targetDir + "/" + localPath)
	if err != nil {
		fmt.Println(fmt.Println("	??? create file failed"))
		return
	}
	defer fl.Close()

	n, err := fl.Write(dataB)
	if err != nil {
		fmt.Println(fmt.Println("	??? write file failed"))
		return
	}

	if n != *content.Size {
		fmt.Println(fmt.Sprintf("number of bytes differ, %d vs %d", n, *content.Size))
	}
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
