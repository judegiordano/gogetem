package uuid

import u "github.com/google/uuid"

func New() (string, error) {
	s, err := u.NewV7()
	if err != nil {
		return "", err
	}
	return s.String(), nil
}
