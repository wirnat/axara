package seeder

import (
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	"math/rand"
	"reflect"
	"time"
)

type day int

func (d day) String() string {
	return []string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}[d]
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func inserted(n int, l []int) bool {
	for _, v := range l {
		if n == v {
			return true
		}
	}
	return false
}

func customGenerator() {
	_ = faker.AddProvider("rand_uuid", func(v reflect.Value) (interface{}, error) {
		return uuid.New().String(), nil
	})

	_ = faker.AddProvider("open_service", func(v reflect.Value) (interface{}, error) {
		rand.Seed(time.Now().UnixNano())
		hmin := 8
		hmax := 12
		jam := rand.Intn(hmax-hmin) + hmin

		mmin := 0
		mmax := 59
		min := rand.Intn(mmax-mmin) + mmin
		op := time.Date(2022, 12, 31, jam, min, 0, 0, time.Local)
		return op, nil
	})

	_ = faker.AddProvider("close_service", func(v reflect.Value) (interface{}, error) {
		rand.Seed(time.Now().UnixNano())
		hmin := 18
		hmax := 23
		jam := rand.Intn(hmax-hmin) + hmin

		mmin := 0
		mmax := 59
		min := rand.Intn(mmax-mmin) + mmin
		op := time.Date(2022, 12, 31, jam, min, 0, 0, time.Local)
		return op, nil
	})
}
