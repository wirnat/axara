package v1

type decoderBuilder struct {
	Decoder
}

func NewDecoderBuilder(builder ModuleBuilder) *decoderBuilder {
	decoder := NewDecoder(builder)
	return &decoderBuilder{Decoder: decoder}
}

func (b decoderBuilder) DecodeBuilder() (r ModuleBuilder) {
	builder := b.Decoder.GetBuilder()
	r = ModuleBuilder{
		ModelTrait: builder.ModelTrait,
		Constructor: Constructor{
			Key:           b.Decode(builder.Key),
			ModelPath:     b.Decode(builder.ModelPath),
			ResultPath:    b.Decode(builder.ResultPath),
			ModuleName:    b.Decode(builder.ModuleName),
			ExecuteModels: nil,
			ModuleTraits:  nil,
			Meta:          map[string]string{},
		},
	}

	for _, e := range builder.ExecuteModels {
		r.ExecuteModels = append(r.ExecuteModels, b.Decode(e))
	}

	for _, trait := range builder.ModuleTraits {
		moduleTrait := ModuleTrait{
			Name:     b.Decode(trait.Name),
			Dir:      b.Decode(trait.Dir),
			FileName: b.Decode(trait.FileName),
			Template: b.Decode(trait.Template),
		}
		r.ModuleTraits = append(r.ModuleTraits, moduleTrait)
	}

	for key, val := range builder.Constructor.Meta {
		r.Meta[key] = b.Decode(val)
	}

	return
}
