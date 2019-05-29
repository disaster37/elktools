package helper

import (
	"github.com/pkg/errors"
	"regexp"
)

func ExtractFromRegex(regex string, data string) ([]string, error) {
	if regex == "" {
		return nil, errors.New("regex must be provied")
	}

	r, err := regex.Compile(regex)
	if err != nil {
		return nil, err
	}

	match := re.FindStringSubmatch(data)
	return match, nil
}
