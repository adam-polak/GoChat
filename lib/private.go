package lib

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetSecret(key string) (value string, err error) {

	f, err := os.Open("../secrets.json")
	check(err)

	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}

		// Remove key from string
		s, found := strings.CutPrefix(s, "\""+key+"\"")
		if !found {
			continue
		}

		// Replace white space
		s = strings.ReplaceAll(s, " ", "")

		// Remove quotes
		s = strings.Replace(s, "\"", "", 2)

		// Remove ':'
		s = strings.Replace(s, ":", "", 1)

		// Return value associated with input key
		return s, nil
	}

	return "", errors.New("key not found in secrets.json")
}
