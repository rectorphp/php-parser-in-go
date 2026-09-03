package version

import (
	"errors"
	"strconv"
	"strings"
)

type Version struct {
	Major, Minor uint64
}

var (
	// ErrInvalidSemVer is returned if a version can not be parsed
	ErrInvalidSemVer = errors.New("invalid semantic version")

	// ErrUnsupportedVer is returned if a version out of supported range
	ErrUnsupportedVer = errors.New("the version is out of supported range")

	php5RangeStart = &Version{Major: 5}
	php5RangeEnd   = &Version{Major: 5, Minor: 6}

	php7RangeStart = &Version{Major: 7}
	php7RangeEnd   = &Version{Major: 7, Minor: 4}

	php8RangeStart = &Version{Major: 8}
)

func New(versionString string) (*Version, error) {
	// Split the parts into [0]Major, [1]Minor
	parts := strings.SplitN(versionString, ".", 2)
	if len(parts) != 2 {
		return nil, ErrInvalidSemVer
	}

	var newVersion = new(Version)
	var err error

	newVersion.Major, err = strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return nil, err
	}

	newVersion.Minor, err = strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return nil, err
	}

	return newVersion, nil
}

func (version *Version) Validate() error {
	if version.InRange(php5RangeStart, php5RangeEnd) ||
		version.InRange(php7RangeStart, php7RangeEnd) ||
		version.GreaterOrEqual(php8RangeStart) {
		return nil
	}

	return ErrUnsupportedVer
}

// Less tests if one version is less than another one
func (version *Version) Less(other *Version) bool {
	return version.Compare(other) < 0
}

// LessOrEqual tests if one version is less than another one or equal
func (version *Version) LessOrEqual(other *Version) bool {
	return version.Compare(other) <= 0
}

// Greater tests if one version is greater than another one
func (version *Version) Greater(other *Version) bool {
	return version.Compare(other) > 0
}

// GreaterOrEqual tests if one version is greater than another one or equal
func (version *Version) GreaterOrEqual(other *Version) bool {
	return version.Compare(other) >= 0
}

// GreaterOrEqual tests if one version is greater than another one or equal
func (version *Version) InRange(start, end *Version) bool {
	return version.Compare(start) >= 0 && version.Compare(end) <= 0
}

// Compare compares this version to another one. It returns -1, 0, or 1 if
// the version smaller, equal, or larger than the other version.
func (version *Version) Compare(other *Version) int {
	if comparisonResult := compareSegment(version.Major, other.Major); comparisonResult != 0 {
		return comparisonResult
	}

	return compareSegment(version.Minor, other.Minor)
}

func compareSegment(first, second uint64) int {
	if first < second {
		return -1
	}
	if first > second {
		return 1
	}

	return 0
}
