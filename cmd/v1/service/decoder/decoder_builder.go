package decoder

import (
	"github.com/wirnat/axara/cmd/v1"
)

func (d decoder) DecodeBuilder(builder v1.ModuleBuilder) (r v1.ModuleBuilder) {
	r = v1.ModuleBuilder{
		ModelTrait: builder.ModelTrait,
		Constructor: v1.Constructor{
			Key:        d.Decode(builder.Key, builder.ModelTrait),
			ModelPath:  d.Decode(builder.ModelPath, builder.ModelTrait),
			ModuleName: d.Decode(builder.ModuleName, builder.ModelTrait),
			Jobs:       nil,
			Meta:       map[string]string{},
		},
	}

	for _, trait := range builder.Jobs {
		moduleTrait := v1.Job{
			Name:          d.Decode(trait.Name, builder.ModelTrait),
			Dir:           d.Decode(trait.Dir, builder.ModelTrait),
			FileName:      d.Decode(trait.FileName, builder.ModelTrait),
			Template:      d.Decode(trait.Template, builder.ModelTrait),
			GenerateIn:    d.Decode(trait.GenerateIn, builder.ModelTrait),
			SingleExecute: false,
		}
		r.Jobs = append(r.Jobs, moduleTrait)
	}

	for key, val := range builder.Constructor.Meta {
		r.Meta[key] = d.Decode(val, builder.ModelTrait)
	}

	return
}
