package helper

import (
	"github.com/pkg/errors"
	"regexp"
)

func ExtractFromRegex(regex string, data string) ([]string, error) {
	if regex == "" {
		return nil, errors.New("regex must be provied")
	}

	r, err := regexp.Compile(regex)
	if err != nil {
		return nil, err
	}

	match := r.FindStringSubmatch(data)
	return match, nil
}
