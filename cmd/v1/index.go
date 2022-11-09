package v1

type app struct {
	path string
}

func NewApp(path string) *app {
	return &app{path: path}
}

func (a app) Generate() error {
	constructor, err := NewConstructor(a.path)
	if err != nil {
		return err
	}

	modelReader := NewModelFileReader()
	metaReader := NewReaderMeta()
	gitPuller := NewGitPuller()
	generator := NewGenerator(modelReader, metaReader, gitPuller)
	err = generator.Generate(*constructor)
	if err != nil {
		return err
	}
	return nil
}
