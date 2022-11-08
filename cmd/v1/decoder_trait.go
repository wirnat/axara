package v1

type decoderTrait struct {
	decoder
}

func NewDecoderTrait(builder ModuleBuilder) *decoderTrait {
	d := NewDecoder(builder)
	return &decoderTrait{decoder: *d}
}

func (d decoder) DecodeTrait(trait ModuleTrait) (r ModuleTrait) {
	r.Dir = d.Decode(trait.Dir)
	r.FileName = d.Decode(trait.FileName)
	r.Name = d.Decode(trait.Name)
	r.Template = d.Decode(trait.Template)
	return
}
