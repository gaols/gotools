package gotools

import (
	"io/ioutil"
	"testing"
)

func TestIsValidDigitalVersion(t *testing.T) {
	if !IsValidDigitalVersion("2") {
		t.Fatal("parse version failed: 2")
	}
	if !IsValidDigitalVersion("2.1") {
		t.Fatal("parse version failed: 2.1")
	}
	if !IsValidDigitalVersion("2.1.0") {
		t.Fatal("parse version failed: 2.1.0")
	}
	if !IsValidDigitalVersion("2.100") {
		t.Fatal("parse version failed: 2.100")
	}
	if !IsValidDigitalVersion("0.0.0") {
		t.Fatal("parse version failed: 0.0.0")
	}

	if IsValidDigitalVersion("x") {
		t.Fatal("parse version failed: x")
	}
	if IsValidDigitalVersion("2.x") {
		t.Fatal("parse version failed: 2.x")
	}
	if IsValidDigitalVersion("01") {
		t.Fatal("parse version failed: 2.x")
	}
	if IsValidDigitalVersion("0.01") {
		t.Fatal("parse version failed: 2.x")
	}
	if IsValidDigitalVersion("0.0.01") {
		t.Fatal("parse version failed: 2.x")
	}
}

func TestVersionSource_Bumps(t *testing.T) {
	f, _ := ioutil.TempFile("/tmp", "abc")
	defer f.Close()
	s := &VersionSource{InitialVersion: "1", SourcePath: f.Name()}
	testBump(s.BumpMajor, t, s, "2")
	testBump(s.BumpPatch, t, s, "2.0.1")
	testBump(s.BumpMinor, t, s, "2.1.1")
	testBump(s.BumpMinor, t, s, "2.2.1")
	testBump(s.BumpMinor, t, s, "2.3.1")
	testBump(s.BumpMinor, t, s, "2.4.1")
	testBump(s.BumpMinor, t, s, "2.5.1")
	testBump(s.BumpMinor, t, s, "2.6.1")
	testBump(s.BumpMinor, t, s, "2.7.1")
	testBump(s.BumpMinor, t, s, "2.8.1")
	testBump(s.BumpMinor, t, s, "2.9.1")
	testBump(s.BumpMinor, t, s, "2.10.1")
	testBump(s.BumpPatch, t, s, "2.10.2")
	testBump(s.BumpMajor, t, s, "3.10.2")
}

func testBump(bumpFunc func() error, t *testing.T, s *VersionSource, expectedVersion string) {
	err := bumpFunc()
	if err != nil {
		t.Fatal(err)
	}
	if v, err := s.Version(); v != expectedVersion {
		if err != nil {
			t.Fatal("get version failed")
		}
		t.Fatal("bump version failed")
	}
}
