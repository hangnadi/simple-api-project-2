package usecase

import (
	"github.com/hangnadi/simple-api-project-2/internal/controller"
)

// NewSystem new system usecase
func NewSystem(systemSvc controller.System) System {

	ucase := &systemUsecase{
		system: systemSvc,
	}

	return ucase
}

// LogOpenFile to logs
func (u *systemUsecase) LogOpenFile() {
	u.system.LogOpenFile()
}
