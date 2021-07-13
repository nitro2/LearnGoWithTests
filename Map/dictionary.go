package dictionary

import "errors"

var ErrorNotFound = errors.New("could not find the word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(s string) (string, error) {
	definition, ok := d[s]
	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d *Dictionary) Add(key, value string) error {
	d[key] = value
	return nil
}
