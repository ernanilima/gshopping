package config_test

import (
	"testing"

	"github.com/ernanilima/gshopping/src/app/config"
	"github.com/ernanilima/gshopping/src/test"
)

func TestStartConfig(t *testing.T) {

	config.StartConfig(test.RootDir())

}
