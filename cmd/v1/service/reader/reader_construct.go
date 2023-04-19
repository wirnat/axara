package reader

import (
	"fmt"
	v1 "github.com/wirnat/axara/cmd/v1"
	"github.com/wirnat/axara/cmd/v1/service/decoder"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
)

type readerConstruct struct{}

func NewReaderConstruct() *readerConstruct {
	return &readerConstruct{}
}

func (r readerConstruct) Read(p string) (b *v1.Constructor, err error) {
	b, err = marshal(p)
	if err != nil {
		return nil, err
	}

	d := decoder.NewDecoder(b)
	for _, imt := range b.IncludeJobs {
		c := new(v1.Constructor)
		c, err = marshal(d.Decode(imt, nil))
		if err != nil {
			return nil, err
		}

		b.Jobs = append(b.Jobs, c.Jobs...)
	}
	return
}

//marshal get data from include yaml file
func marshal(p string) (c *v1.Constructor, err error) {
	targetConfig := fmt.Sprintf("%v", p)
	data, errFile := os.Open(targetConfig)
	if errFile != nil {
		return &v1.Constructor{}, errFile
	}
	defer data.Close()

	byteValue, err := ioutil.ReadAll(data)
	if err != nil {
		return &v1.Constructor{}, errFile
	}
	isJson := strings.Contains(p, ".json")
	if isJson {
		err = yaml.Unmarshal(byteValue, &c)
		if err != nil {
			return nil, err
		}
	} else {
		err = yaml.Unmarshal(byteValue, &c)
		if err != nil {
			return nil, err
		}
	}
	return
}
