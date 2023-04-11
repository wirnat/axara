package v1

type decoderTrait struct {
	decoder
}

func NewDecoderTrait(c Constructor) *decoderTrait {
	d := NewDecoder(c)
	return &decoderTrait{decoder: *d}
}

func (d decoder) DecodeTrait(trait Job, mt *ModelTrait) (r Job) {
	r.Dir = d.Decode(trait.Dir, mt)
	r.FileName = d.Decode(trait.FileName, mt)
	r.Name = d.Decode(trait.Name, mt)
	r.Template = d.Decode(trait.Template, mt)
	r.GenerateIn = d.Decode(trait.GenerateIn, mt)
	r.Active = trait.Active
	return
}
