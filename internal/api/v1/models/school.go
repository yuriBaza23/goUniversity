package models

import "errors"

var SchoolProps []string = []string{"public", "private"}

type School struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Logo string `json:"logo"`
}

func (s *School) VerifyType() error {
	for _, prop := range SchoolProps {
		if prop == s.Type {
			return nil
		}
	}
	return errors.New("invalid school type")
}
