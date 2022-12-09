package key

import (
	"errors"
	"fmt"
	"os"
)

//Storage is local database that's created by badger
var Storage = func() string {
	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	storage := fmt.Sprintf("%v/.axara", userDir)
	if _, err := os.Stat(storage); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(storage, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return storage
}
