package gotools

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	// TypeStdout is type of stdout
	TypeStdout = 0
	// TypeStderr is type of stderr
	TypeStderr = 1
)

// Local run the command in localhost
// https://studygolang.com/articles/4004   <- run shell command and read output line by line
// https://studygolang.com/articles/7767   <- run command without known args
func Local(localCmd string, paras ...interface{}) (string, error) {
	localCmd = fmt.Sprintf(localCmd, paras...)
	cmd := exec.Command("/bin/bash", "-c", localCmd)
	ret, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(ret), nil
}

// RtLocal run the command in localhost and get command output realtime.
func RtLocal(localCmd string, lineHandler func(line string, lineType int8), paras ...interface{}) error {
	localCmd = fmt.Sprintf(localCmd, paras...)
	cmd := exec.Command("/bin/bash", "-c", localCmd)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	err = cmd.Start()
	if err != nil {
		return err
	}

	go func() {
		defer stdout.Close()
		defer stderr.Close()
		stdoutScanner := bufio.NewScanner(stdout)
		stderrScanner := bufio.NewScanner(stderr)
		for stdoutScanner.Scan() {
			lineHandler(stdoutScanner.Text(), TypeStdout)
		}
		for stderrScanner.Scan() {
			lineHandler(stderrScanner.Text(), TypeStderr)
		}
	}()

	return cmd.Wait()
}

// Tar pack the targetPath and put tarball to tgzPath, targetPath and tgzPath should both the absolute path.
func Tar(tgzPath, targetPath string) error {
	if !IsDir(targetPath) && !IsRegular(targetPath) {
		return errors.New("invalid pack path: " + targetPath)
	}

	targetPathDir := filepath.Dir(RemoveTrailingSlash(targetPath))
	target := filepath.Base(RemoveTrailingSlash(targetPath))
	_, err := Local("tar czf %s -C %s %s", tgzPath, targetPathDir, target)
	return err
}

// UnTar unpack the tarball specified by tgzPath and extract it to the path specified by targetPath
func UnTar(tgzPath, targetPath string) error {
	if !IsDir(targetPath) {
		return errors.New("tar extract path invalid: " + targetPath)
	}

	if !IsRegular(tgzPath) {
		return errors.New("tar path invalid: " + tgzPath)
	}

	_, err := Local("tar xf %s -C %s", tgzPath, targetPath)
	return err
}

type GrepConfig struct {
	// 命中行上下文行数
	Context                int
	FromStart              bool
	AlwaysIncludeFirstLine bool
}

// grep val from src, val cannot contain line break.
// from start means if any match found, the first match will starting from the very start line.
func Grep(src, val string, config *GrepConfig) []string {
	vals := strings.Split(src, "\n")
	if len(vals) == 0 {
		return []string{""}
	}

	var pivots []int
	for i, v := range vals {
		if strings.Contains(v, val) {
			pivots = append(pivots, i)
		}
	}
	var lineIdx []int
	if config.AlwaysIncludeFirstLine {
		lineIdx = append(lineIdx, 0)
	}
	fromStart := config.FromStart
	for _, p := range pivots {
		s := p - config.Context
		e := p + config.Context
		if s < 0 {
			s = 0
		}
		if e > len(vals)-1 {
			e = len(vals) - 1
		}
		if fromStart {
			s = 0
			fromStart = false
		}
		for i := s; i <= e; i++ {
			if !ContainsAnyInt(lineIdx, i) {
				lineIdx = append(lineIdx, i)
			}
		}
	}
	var ret []string
	for _, v := range lineIdx {
		ret = append(ret, vals[v])
	}
	return ret
}
