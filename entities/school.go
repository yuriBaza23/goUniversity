package entities

import "errors"

var SchoolProps []string = []string{"public", "private"}

type School struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func (s *School) VerifyType() error {
	for _, prop := range SchoolProps {
		if prop == s.Type {
			return nil
		}
	}
	return errors.New("invalid school type")
}
