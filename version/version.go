package version

import (
	"bytes"
	"fmt"
)

// Version is the version of the app
const Version = "1.1.1"

// VersionPrerelease is the state of the app
const VersionPrerelease = "dev"

// FormattedVersion is used to format the full version of the app
func FormattedVersion() string {
	var versionString bytes.Buffer
	fmt.Fprintf(&versionString, "%s", Version)

	if VersionPrerelease != "" {
		fmt.Fprintf(&versionString, "-%s", VersionPrerelease)
	}

	return versionString.String()
}
