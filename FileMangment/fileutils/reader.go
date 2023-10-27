package fileutils

import "os"

func ReadTextFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err == nil {
		return string(content), nil
	} else {
		return "", err
	}
}
