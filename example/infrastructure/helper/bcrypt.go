package helper

type Bcryptor interface {
	Encrypt(s string) (string, error)
	CheckHash(pass, hpass string) error
}

type bcryptor struct{}

func (b bcryptor) Encrypt(s string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (b bcryptor) CheckHash(pass, hpass string) error {
	//TODO implement me
	panic("implement me")
}

func newBcryptor() Bcryptor {
	return &bcryptor{}
}
