package maps

import "errors"

// The key type is special.
// It can only be a comparable type because without the ability to tell if 2 keys are equal.

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
