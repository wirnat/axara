package v1

type Decoder interface {
	Decode(code string, mt *ModelTrait) (encoded string)
	DecodeBuilder(builder ModuleBuilder) (r ModuleBuilder)
	DecodeTrait(trait Job, mt *ModelTrait) (r Job)
}
