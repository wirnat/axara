package ztring

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToHyphenated(t *testing.T) {
	u := ConvertToHyphenated("user_company")
	assert.Equal(t, "user-company", u)
}

func TestPluralize(t *testing.T) {
	u := Pluralize("user company")
	assert.Equal(t, "user-companies", u)
}
