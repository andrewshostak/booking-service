package service

import (
	"github.com/stretchr/testify/mock"
	"time"
)

var _ LaunchpadChecker = &LaunchpadRepositoryMock{}

type LaunchpadRepositoryMock struct {
	mock.Mock
}

func (m *LaunchpadRepositoryMock) IsLaunchpadAvailable(launchpadId string, launchDate time.Time) (bool, error) {
	args := m.Called(launchpadId, launchDate)
	return args.Bool(0), args.Error(1)
}
