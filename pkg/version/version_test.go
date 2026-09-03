package version_test

import (
	"gotest.tools/assert"
	"testing"

	"github.com/rectorphp/php-parser-in-go/pkg/version"
)

func Test(test *testing.T) {
	ver, err := version.New("7.4")
	assert.NilError(test, err)

	assert.Equal(test, *ver, version.Version{
		Major: 7,
		Minor: 4,
	})
}

func TestLeadingZero(test *testing.T) {
	ver, err := version.New("07.04")
	assert.NilError(test, err)

	assert.Equal(test, *ver, version.Version{
		Major: 7,
		Minor: 4,
	})
}

func TestValidate(test *testing.T) {
	supported := []string{"5.6", "7.4", "8.0", "8.2", "8.4", "9.0"}
	for _, versionString := range supported {
		parsed, err := version.New(versionString)
		assert.NilError(test, err)
		assert.NilError(test, parsed.Validate())
	}

	unsupported := []string{"4.9", "6.0"}
	for _, versionString := range unsupported {
		parsed, err := version.New(versionString)
		assert.NilError(test, err)
		assert.Error(test, parsed.Validate(), "the version is out of supported range")
	}
}

func TestInRange(test *testing.T) {
	versionValue, err := version.New("7.0")
	assert.NilError(test, err)

	version2, err := version.New("7.4")
	assert.NilError(test, err)

	ver, err := version.New("7.0")
	assert.NilError(test, err)
	assert.Assert(test, ver.InRange(versionValue, version2))

	ver, err = version.New("7.2")
	assert.NilError(test, err)
	assert.Assert(test, ver.InRange(versionValue, version2))

	ver, err = version.New("7.4")
	assert.NilError(test, err)
	assert.Assert(test, ver.InRange(versionValue, version2))
}
