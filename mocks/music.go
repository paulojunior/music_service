package mocks

import (
	"github.com/paulojunior/code-challange/entity"
	"github.com/stretchr/testify/mock"
)

type MockMusicRepository struct {
	mock.Mock
}

func (m *MockMusicRepository) FindByISRC(isrc string) ([]entity.Music, error) {
	args := m.Called(isrc)
	return args.Get(0).([]entity.Music), nil
}

func (m *MockMusicRepository) FindByArtistName(artistName string) ([]entity.Music, error) {
	args := m.Called(artistName)
	return args.Get(0).([]entity.Music), nil
}

func (m *MockMusicRepository) Insert(music entity.Music) error {
	return nil
}
