package lib

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func GetSecret(key string) (value string, err error) {
	f, err := os.Open("secrets.json")
	if err != nil {
		return "", err
	}

	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}

		if !strings.Contains(s, "\""+key+"\"") {
			continue
		}

		// Split string into key and value
		arr := strings.Split(s, ":")

		// Remove whitespace from first key half, and add back second half
		s = strings.ReplaceAll(arr[0], " ", "") + arr[1]

		// Remove key from string
		s, _ = strings.CutPrefix(s, "\""+key+"\"")

		// Remove quotes
		s = strings.Replace(s, "\"", "", 2)

		return s, nil
	}

	return "", errors.New("key not found in secrets.json")
}
