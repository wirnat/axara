package generator_v2

import (
	"bufio"
	"fmt"
	v1 "github.com/wirnat/axara/cmd/v1"
	er "github.com/wirnat/axara/cmd/v1/errors"
	"github.com/wirnat/axara/cmd/v1/global"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"text/template"
)

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

	err = g.decodeConstructor(&c, nil)
	if err != nil {
		return err
	}

	//read model
	global.Spinner.Title = "Read model path..."
	files, err := ioutil.ReadDir(c.ModelPath)
	if len(files) < 1 || err != nil {
		return er.NoModelFound
	}

	for _, job := range c.Jobs {
		if job.SingleExecute {
			g.ExecOne(job, c, nil)
		} else {
			g.ExecPerModel(job, c)
		}
	}

	return nil
}

func (g generator) ExecOne(job v1.Job, c v1.Constructor, mt *v1.ModelTrait) error {
	if !job.Active {
		return nil
	}

	job = g.decodeJob(job, mt)

	err := os.MkdirAll(g.Decoder.Decode(job.Dir, mt), os.ModePerm)
	if err != nil {
		fmt.Println("	❌" + err.Error())
		return err
	}

	validTag := false

	if global.Tags != nil {
		for _, inputTag := range global.Tags {
			for _, tag := range job.Tags {
				if tag == inputTag {
					validTag = true
				}
			}
		}
	} else {
		goto startExecute
	}

	if !validTag {
		return nil
	}

startExecute:

	generatedFile := fmt.Sprintf("%v/%v", job.Dir, job.FileName)
	generatedFile = g.Decoder.Decode(generatedFile, mt)

	moduleBuilder := v1.ModuleBuilder{
		Constructor: c,
		ModelTrait:  mt,
	}

	tmt, err := template.ParseFiles(job.Template)
	if err != nil {
		fmt.Println("❌ " + err.Error())
		return err
	}

	if job.GenerateIn != "" {
		err = injectCode(job, generatedFile, *tmt, moduleBuilder)
		if err != nil {
			fmt.Println("	❌ " + err.Error())
			return err
		}
		return nil
	}

	fileTrait, err := os.Create(generatedFile)
	if err != nil {
		fmt.Println("❌ " + err.Error())
		return err

	}

	err = tmt.Execute(fileTrait, moduleBuilder)
	if err != nil {
		fmt.Println("	❌ " + err.Error())
		return err
	}
	err = fileTrait.Close()
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("	✅  %v \n", job.Name))
	return nil
}

func (g generator) ExecPerModel(job v1.Job, c v1.Constructor) error {
	var mt []*v1.ModelTrait

	global.Spinner.Title = "Read model path..."
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
		err = g.decodeConstructor(&c, m)
		g.ExecOne(job, c, m)
	}

	return nil
}

func (g generator) decodeJob(job v1.Job, mt *v1.ModelTrait) v1.Job {
	return v1.Job{
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
}

func (g generator) decodeConstructor(c *v1.Constructor, mt *v1.ModelTrait) (err error) {
	if mt != nil {
		if c.Meta == nil {
			c.Meta = map[string]string{}
		}
		metas, err := g.ReaderMeta.GetMeta(mt.FileInfo, *c, mt.Model)
		if err != nil {
			return err
		}
		for key, val := range metas {
			c.Meta[key] = val
		}
	}

	c.GitAccessKey = g.Decode(c.GitAccessKey, mt)
	c.Key = g.Decode(c.Key, mt)
	c.ModelPath = g.Decode(c.ModelPath, mt)
	c.ModuleName = g.Decode(c.ModuleName, mt)

	var jobs []v1.Job

	for _, job := range c.Jobs {
		jobs = append(jobs, g.decodeJob(job, mt))
	}

	for key, val := range c.Meta {
		c.Meta[key] = g.Decode(val, mt)
	}

	c.Jobs = jobs

	return
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
