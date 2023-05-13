package models

import "testing"

func TestSvaeUser(t *testing.T) {
	SaveUser(nil)
}

func TestGetUserId(t *testing.T) {
	GetUserId(12)
}
