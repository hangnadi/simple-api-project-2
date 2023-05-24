package repository_test

import (
	"testing"

	. "github.com/hangnadi/simple-api-project-2/internal/repository"
)

func TestSystemLogOpenFile(t *testing.T) {
	systemRepo := NewSystem("")
	systemRepo.LogOpenFile()
}
