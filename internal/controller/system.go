package controller

import "github.com/hangnadi/simple-api-project-2/internal/repository"

// NewSystem New system service
func NewSystem(systemRP repository.System) System {

	svc := &systemSvc{
		systemRP: systemRP,
	}

	return svc
}

func (s *systemSvc) LogOpenFile() {
	s.systemRP.LogOpenFile()
}
