package generator_v2

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/iancoleman/strcase"
	v1 "github.com/wirnat/axara/cmd/v1"
	er "github.com/wirnat/axara/cmd/v1/errors"
	"github.com/wirnat/axara/cmd/v1/global"
	"gopkg.in/src-d/go-git.v4"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"
)

var jobDone int

type generator struct {
	v1.ReaderModel
	v1.Decoder
	v1.ReaderMeta
}

func NewGenerator(readerModel v1.ReaderModel, decoder v1.Decoder, readerMeta v1.ReaderMeta) *generator {
	return &generator{ReaderModel: readerModel, Decoder: decoder, ReaderMeta: readerMeta}
}

func (g generator) Generate(c v1.Constructor) (err error) {
	if c.Key != "ᬅᬓ᭄ᬱᬭ" {
		return er.InvalidKey
	}

	if c.Jobs == nil {
		return er.NothingTodo
	}

	if len(global.Tags) < 1 {
		return fmt.Errorf("please run specific tags, ex: -g repository, usecase")
	}

	//read model
	files, err := ioutil.ReadDir(c.ModelPath)
	if len(files) < 1 || err != nil {
		log.Fatal(err.Error())
	}
	totalJob := len(c.Jobs)
	startTime := time.Now()

	for i, job := range c.Jobs {
		if job.SingleExecute {
			err = g.ExecOne(job, c, nil)
			if err != nil {
				log.Fatalf(" ❌ " + err.Error())
			}
		} else {
			err = g.ExecPerModel(job, c)
			if err != nil {
				log.Fatalf(" ❌  " + err.Error())
			}
		}
		//print progress bar cmd
		fmt.Print("Write the code: [")
		for j := 0; j <= totalJob; j++ {
			if j <= i {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\r")
	}
	elapsedTime := time.Since(startTime)
	elapsedSeconds := int(elapsedTime.Milliseconds())

	if jobDone == 0 {
		fmt.Printf("%v job done, maybe your tag not found in jobs", jobDone)
	} else {
		fmt.Printf("%v job done, in %v microseconds ", jobDone, elapsedSeconds)
	}

	return nil
}

func (g generator) ExecOne(job v1.Job, c v1.Constructor, mt *v1.ModelTrait) error {
	if !job.Active {
		return nil
	}
	job = g.decodeJob(job, mt)

	validTag := false

	//check tags
	if len(global.Tags) > 0 {
		for _, inputTag := range global.Tags {
			if job.Tags == nil {
				validTag = false
			}
			for _, tag := range job.Tags {
				if tag == inputTag {
					validTag = true
				}
			}
		}

		if !validTag {
			return nil
		}
	}

	//check override
	generatedFile := fmt.Sprintf("%v/%v", job.Dir, job.FileName)
	generatedFile = g.Decoder.Decode(generatedFile, mt)
	canOverride := g.validateOverride(generatedFile)
	if !canOverride {
		return nil
	}

	//isCleanCommit := g.validateCommit()
	//if !isCleanCommit {
	//	return fmt.Errorf("uncommit change found, please commit your project before start generate the template")
	//}

	err := os.MkdirAll(g.Decoder.Decode(job.Dir, mt), os.ModePerm)
	if err != nil {
		return err
	}

	moduleBuilder := v1.ModuleBuilder{
		Constructor: c,
		ModelTrait:  mt,
	}

	moduleBuilder = g.DecodeBuilder(moduleBuilder)

	tmt, err := template.ParseFiles(job.Template)
	if err != nil {
		return err
	}

	if job.GenerateIn != "" {
		err = injectCode(job, generatedFile, *tmt, moduleBuilder)
		if err != nil {
			return err
		}
		jobDone++
		return nil
	}

	fileTrait, err := os.Create(generatedFile)
	if err != nil {
		return err
	}

	err = tmt.Execute(fileTrait, moduleBuilder)
	if err != nil {
		return err
	}
	err = fileTrait.Close()
	if err != nil {
		panic(err)
	}

	jobDone++
	return nil
}

func (g generator) ExecPerModel(job v1.Job, c v1.Constructor) error {
	var mt []*v1.ModelTrait

	files, err := ioutil.ReadDir(c.ModelPath)
	if len(files) < 1 || err != nil {
		return er.NoModelFound
	}

	//get meta from model and get scanned model trait
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		modelTrait, err := g.GetModelTrait(file, c)
		if err != nil {
			return err
		}

		if modelTrait != nil && modelTrait.Model != "" {
			modelTrait.FileInfo = file
			mt = append(mt, modelTrait)
		}
	}

	//execute per model
	for _, m := range mt {
		builder := g.DecodeBuilder(v1.ModuleBuilder{
			Constructor: c,
			ModelTrait:  m,
		})
		err = g.ExecOne(job, builder.Constructor, builder.ModelTrait)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g generator) decodeJob(job v1.Job, mt *v1.ModelTrait) v1.Job {
	j := v1.Job{
		Name:          g.Decode(job.Name, mt),
		Dir:           g.Decode(job.Dir, mt),
		FileName:      g.Decode(job.FileName, mt),
		Template:      g.Decode(job.Template, mt),
		Active:        job.Active,
		CMD:           job.CMD,
		Tags:          job.Tags,
		GenerateIn:    g.Decode(job.GenerateIn, mt),
		SingleExecute: job.SingleExecute,
	}
	return j
}

func injectCode(job v1.Job, generatedFile string, tmt template.Template, builder v1.ModuleBuilder) (err error) {
	f, err := os.Open(generatedFile)
	if err != nil {
		return err
	}

	defer f.Close()

	// Membaca isi file baris per baris dan menulis kembali ke file yang lama
	tmpFile, err := os.CreateTemp(job.Dir, "test_*.txt")
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	//Collect all data from executed model model
	scanner := bufio.NewScanner(f)
	writer := bufio.NewWriter(tmpFile)

	for scanner.Scan() {
		line := scanner.Text()
		generateCommentFound := strings.Contains(line, "@Generate")
		if generateCommentFound {
			initiator := strings.Fields(line)
			if len(initiator) != 2 {
				return err
			}
			generateTag := initiator[1]
			re, err := regexp.Compile(`[^\w]`)
			if err != nil {
				return err
			}
			generateTag = string(re.ReplaceAll([]byte(generateTag), []byte("")))
			if generateTag == job.GenerateIn {
				var buf strings.Builder
				if err := tmt.Execute(&buf, builder); err != nil {
					return err
				}
				line = buf.String()
			}
		}
		fmt.Fprintln(writer, line)

		err = writer.Flush()
		if err != nil {
			return err
		}
	}

	// Menimpa file lama dengan file yang telah diperbarui
	if err = os.Rename(tmpFile.Name(), generatedFile); err != nil {
		return err
	}
	return nil
}

func (g generator) validateOverride(generatedFile string) bool {
	if !global.OverrideAll {
		if _, err := os.Stat(generatedFile); !errors.Is(err, os.ErrNotExist) {
			var input string

			fmt.Println(generatedFile+" is already exist, do you want to override?", "Y=Yes", "N=No", "YA=Yes for all")
			_, err := fmt.Scanln(&input)
			if err != nil {
				fmt.Println("	something is wrong")
				return true
			}

			input = strcase.ToSnake(input)
			if input == "ya" {
				global.OverrideAll = true
				return true
			}
			if input == "no" || input == "n" {
				return false
			}
			if input != "yes" && input != "y" && input != "ya" {
				return true
			}
		}
	}

	return true
}

func (g generator) validateCommit() bool {
	// Buka repositori Git
	r, err := git.PlainOpen(".")
	if err != nil {
		log.Fatal(err)
	}

	// Ambil working tree
	w, err := r.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	// Cek status perubahan dalam working tree
	status, err := w.Status()
	if err != nil {
		log.Fatal(err)
	}

	// Cek apakah ada perubahan yang belum dicommit
	hasUncommittedChanges := false
	for _, s := range status {
		if s.Staging != git.Unmodified || s.Worktree != git.Unmodified {
			hasUncommittedChanges = true
			break
		}
	}

	if hasUncommittedChanges {
		return false
	} else {
		return true
	}
}
