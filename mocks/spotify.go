package mocks

import (
	"github.com/paulojunior/code-challange/integration"
	"github.com/stretchr/testify/mock"
)

type SpotifyIntegrationMock struct {
	mock.Mock
}

func (_m *SpotifyIntegrationMock) GetTrackByISRC(ISRC string) (integration.Item, error) {
	return integration.Item{}, nil
}
