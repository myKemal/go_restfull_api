package config

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"os"
)

func TestCofig_EnvMongoURI(t *testing.T) {
	//given
	os.Setenv("MONGOURI", "user:pass@tcp(blabla)/xxx?charset=utf8")
	defer os.Unsetenv("MONGOURI")

	//when
	connectionString := EnvMongoURI()

	assert.Equal(t, "user:pass@tcp(blabla)/xxx?charset=utf8", connectionString)
}
