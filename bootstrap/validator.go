package bootstrap

import (
	"github.com/markgravity/golang-ic/lib/validators"
)

func RegisterValidators() {
	validators.Init()
}
