package v1

type decoderBuilder struct {
	Decoder
}

func NewDecoderBuilder(c Constructor) *decoderBuilder {
	decoder := NewDecoder(c)
	return &decoderBuilder{Decoder: decoder}
}

func (b decoderBuilder) DecodeBuilder(builder ModuleBuilder) (r ModuleBuilder) {
	r = ModuleBuilder{
		ModelTrait: builder.ModelTrait,
		Constructor: Constructor{
			Key:           b.Decode(builder.Key, &builder.ModelTrait),
			ModelPath:     b.Decode(builder.ModelPath, &builder.ModelTrait),
			ResultPath:    b.Decode(builder.ResultPath, &builder.ModelTrait),
			ModuleName:    b.Decode(builder.ModuleName, &builder.ModelTrait),
			ExecuteModels: nil,
			ModuleTraits:  nil,
			Meta:          map[string]string{},
		},
	}

	for _, e := range builder.ExecuteModels {
		r.ExecuteModels = append(r.ExecuteModels, b.Decode(e, &builder.ModelTrait))
	}

	for _, trait := range builder.ModuleTraits {
		moduleTrait := ModuleTrait{
			Name:     b.Decode(trait.Name, &builder.ModelTrait),
			Dir:      b.Decode(trait.Dir, &builder.ModelTrait),
			FileName: b.Decode(trait.FileName, &builder.ModelTrait),
			Template: b.Decode(trait.Template, &builder.ModelTrait),
		}
		r.ModuleTraits = append(r.ModuleTraits, moduleTrait)
	}

	for key, val := range builder.Constructor.Meta {
		r.Meta[key] = b.Decode(val, &builder.ModelTrait)
	}

	return
}
