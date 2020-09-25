package gotools

import (
	"errors"
	"fmt"
	"github.com/gaols/gotools"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// valid version should be: m / m.n.p / m.n
var versionPatt = regexp.MustCompile(`^(0|[1-9]\d*)(\.(0|[1-9]\d*)){0,2}$`)

type VersionSource struct {
	SourcePath     string // version file path, try to use a file named 'version' in current dir if not explicitly set
	InitialVersion string // If explicitly set, use it as the current version.
	ver            string // the current version
}

func (source *VersionSource) BumpMajor() error {
	if err := initVersion(source); err != nil {
		return err
	}
	vps := strings.Split(source.ver, ".")
	v, _ := strconv.ParseInt(vps[0], 10, 32)
	vps[0] = fmt.Sprintf("%d", v+1)
	source.ver = strings.Join(vps, ".")
	return source.persist()
}

func (source *VersionSource) BumpMinor() error {
	if err := initVersion(source); err != nil {
		return err
	}
	vps := strings.Split(source.ver, ".")
	if len(vps) > 1 {
		v, _ := strconv.ParseInt(vps[1], 10, 32)
		vps[1] = fmt.Sprintf("%d", v+1)
		source.ver = strings.Join(vps, ".")
	} else {
		source.ver = vps[0] + ".1"
	}

	return source.persist()
}

func (source *VersionSource) BumpPatch() error {
	if err := initVersion(source); err != nil {
		return err
	}
	vps := strings.Split(source.ver, ".")
	if len(vps) > 2 {
		v, _ := strconv.ParseInt(vps[2], 10, 32)
		vps[2] = fmt.Sprintf("%d", v+1)
		source.ver = strings.Join(vps, ".")
	} else {
		if len(vps) > 1 {
			source.ver = strings.Join(vps, ".") + ".1"
		} else {
			source.ver = vps[0] + ".0.1"
		}
	}
	return source.persist()
}

func (source *VersionSource) persist() error {
	fileInfo, err := os.Stat(source.SourcePath)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(source.SourcePath, []byte(source.ver), fileInfo.Mode())
}

func (source *VersionSource) Version() (string, error) {
	err := source.check()
	if err != nil {
		return "", err
	}
	content, err := ioutil.ReadFile(source.SourcePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// initVersion init version object:
// if InitialVersion is explicitly specified and version source file is empty, use it as the current file version.
// otherwise, try to read the source file to get the current version,
// if the version file is empty, use "0.0.0" as the current version,
// if the version in file is not a valid version, error will return.
func initVersion(source *VersionSource) error {
	if err := source.check(); err != nil {
		return err
	}

	content, err := ioutil.ReadFile(source.SourcePath)
	if err != nil {
		return err
	}

	strContent := string(content)
	if gotools.IsBlank(strContent) {
		if gotools.IsNotBlank(source.InitialVersion) {
			source.ver = source.InitialVersion
		} else {
			source.ver = "0.0.0"
		}
		return nil
	}

	if !IsValidDigitalVersion(strContent) {
		err := errors.New("Invalid version: " + strContent)
		return err
	}

	source.ver = strings.TrimSpace(strContent)
	return nil
}

func (source *VersionSource) check() error {
	if gotools.IsBlank(source.SourcePath) {
		currentPath, err := os.Getwd()
		if err != nil {
			return err
		}
		source.SourcePath = path.Join(currentPath, "version")
	}
	if !gotools.IsFileExists(source.SourcePath) {
		return errors.New(fmt.Sprintf("Version file not found: %s ", source.SourcePath))
	}
	if gotools.IsNotBlank(source.InitialVersion) && !IsValidDigitalVersion(source.InitialVersion) {
		return errors.New("Invalid initial version: " + source.InitialVersion)
	}

	return nil
}

func IsValidDigitalVersion(version string) bool {
	return gotools.IsNotBlank(version) && versionPatt.MatchString(version)
}
