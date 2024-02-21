package common

import (
	"fmt"
	"os"
	"strings"
)

func GetRelativePath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	rootDir := "marketing"
	idx := strings.Index(dir, rootDir)
	if idx < 0 {
		return "", fmt.Errorf("can't find root directory, dir is %s", dir)
	}
	return dir[:idx+len(rootDir)], nil
}
