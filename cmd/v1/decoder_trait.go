package v1

type decoderTrait struct {
	decoder
}

func NewDecoderTrait(c Constructor) *decoderTrait {
	d := NewDecoder(c)
	return &decoderTrait{decoder: *d}
}

func (d decoder) DecodeTrait(trait ModuleTrait, mt *ModelTrait) (r ModuleTrait) {
	r.Dir = d.Decode(trait.Dir, mt)
	r.FileName = d.Decode(trait.FileName, mt)
	r.Name = d.Decode(trait.Name, mt)
	r.Template = d.Decode(trait.Template, mt)
	return
}
