package repository_test

import (
	"testing"

	. "github.com/kokka-team/nakama-investor-be/internal/repository"
)

func TestSystemLogOpenFile(t *testing.T) {
	systemRepo := NewSystem("")
	systemRepo.LogOpenFile()
}
