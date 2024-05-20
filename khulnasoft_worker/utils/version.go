package utils

import (
	"fmt"

	"github.com/khulnasoft/kengine/khulnasoft_utils/log"
	"golang.org/x/mod/semver"
)

var (
	Version   string
	Commit    string
	BuildTime string
)

func init() {
	if len(Version) > 0 && Version[0] != 'v' {
		Version = fmt.Sprintf("v%s", Version)
	}
	if !semver.IsValid(Version) {
		log.Fatal().Msgf("Provided console version %s is not valid", Version)
	}
}
