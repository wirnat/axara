package decoder

import (
	"github.com/wirnat/axara/cmd/v1"
)

func (d decoder) DecodeTrait(trait v1.Job, mt *v1.ModelTrait) (r v1.Job) {
	r.Dir = d.Decode(trait.Dir, mt)
	r.FileName = d.Decode(trait.FileName, mt)
	r.Name = d.Decode(trait.Name, mt)
	r.Template = d.Decode(trait.Template, mt)
	r.GenerateIn = d.Decode(trait.GenerateIn, mt)
	r.Active = trait.Active
	return
}
