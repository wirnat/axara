package v0

import "github.com/sirupsen/logrus"

type app struct {
	path string
}

func NewApp(path string) *app {
	return &app{
		path: path,
	}
}

func (a app) Generate() error {
	conf, err := OpenConfig(a.path)
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	if conf.Key == "ᬅᬓ᭄ᬱᬭ" {
		reader := NewModelScanner(conf)
		blueprint, err := reader.Scan()
		if err != nil {
			logrus.Fatalln(err.Error())
		}

		decoder := NewDecoder(blueprint)
		blueprint, err = decoder.Decode()
		if err != nil {
			logrus.Fatalln(err.Error())
		}

		err = blueprint.Generate()
		if err != nil {
			logrus.Fatalln(err.Error())
		}
	} else {
		logrus.Fatalln("invalid key in json file")
	}
	return nil
}
