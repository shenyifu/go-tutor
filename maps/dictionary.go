package maps

import "errors"

type Dictionary map[string]string

var NotFoundErr = errors.New("not found")

//Search return value of key
func (d Dictionary) Search(key string) (string, error) {
	got, ok := d[key]
	if !ok {
		return "", NotFoundErr
	}
	return got, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
