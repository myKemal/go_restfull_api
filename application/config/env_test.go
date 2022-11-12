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

func TestConfig_GetPort_DefaultWhenNotProvided(t *testing.T) {
	//given
	os.Unsetenv("APPSERVER")

	//when
	port := GetPort()

	assert.Equal(t, "8080", port)
}

func TestConfig_GetPort_FromEnvironment(t *testing.T) {
	//given
	os.Setenv("APPSERVER", "101")
	defer os.Unsetenv("APPSERVER")

	//when
	port := GetPort()

	assert.Equal(t, "101", port)
}
