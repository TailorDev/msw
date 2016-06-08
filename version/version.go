package version

import (
	"bytes"
	"fmt"
)

const Version = "1.0.0"

const VersionPrerelease = "dev"

func FormattedVersion() string {
	var versionString bytes.Buffer
	fmt.Fprintf(&versionString, "%s", Version)

	if VersionPrerelease != "" {
		fmt.Fprintf(&versionString, "-%s", VersionPrerelease)
	}

	return versionString.String()
}
