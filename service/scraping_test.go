package service

import (
	"testing"
)

func TestGetThreadURL(t *testing.T) {
	InsertComment(1)
}

func TestGetCurrentMaxComment(t *testing.T) {
	getCurrentMaxComment(1)
}
