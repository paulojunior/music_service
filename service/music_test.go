package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/paulojunior/code-challange/entity"
	"github.com/paulojunior/code-challange/mocks"
	"github.com/stretchr/testify/mock"
)

func TestMusicService_FindByISRC(t *testing.T) {
	tests := []struct {
		name       string
		wantMusics []entity.Music
		isrc       string
		wantErr    bool
	}{
		{
			name:       "Valid ISRC",
			wantMusics: []entity.Music{},
			isrc:       "ABC123",
			wantErr:    false,
		},
		{
			name:       "Invalid ISRC",
			wantMusics: nil,
			isrc:       "",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			MockMusicRepository := mocks.MockMusicRepository{}

			MockMusicRepository.On("FindByISRC", mock.Anything).Return(tt.wantMusics, tt.wantErr)

			service := NewMusicService(&MockMusicRepository)

			gotMusics, err := service.FindByISRC(tt.isrc)

			if (err != nil) != tt.wantErr {
				t.Errorf("MusicService.FindByISRC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotMusics, tt.wantMusics) {
				t.Errorf("MusicService.FindByISRC() = %v, want %v", gotMusics, tt.wantMusics)
			}
		})
	}
}

func TestMusicService_FindByArtistName(t *testing.T) {
	tests := []struct {
		name          string
		wantMusics    []entity.Music
		artistName    string
		repositoryErr error
		wantErr       bool
	}{
		{
			name:          "Valid Artist Name",
			wantMusics:    []entity.Music{},
			artistName:    "John Doe",
			repositoryErr: nil,
			wantErr:       false,
		},
		{
			name:          "Empty Artist Name",
			wantMusics:    nil,
			artistName:    "",
			repositoryErr: nil,
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMusicRepository := mocks.MockMusicRepository{}
			mockMusicRepository.On("FindByArtistName", mock.Anything).Return(tt.wantMusics, tt.repositoryErr)

			service := NewMusicService(&mockMusicRepository)

			gotMusics, err := service.FindByArtistName(tt.artistName)

			if (err != nil) != tt.wantErr {
				t.Errorf("MusicService.FindByArtistName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotMusics, tt.wantMusics) {
				t.Errorf("MusicService.FindByArtistName() = %v, want %v", gotMusics, tt.wantMusics)
			}
		})
	}
}

func TestMusicService_Insert(t *testing.T) {
	tests := []struct {
		name           string
		inputMusic     entity.Music
		validateErr    error
		repositoryErr  error
		wantValidation bool
		wantRepository bool
		wantErr        bool
	}{
		{
			name:          "Valid Music",
			inputMusic:    entity.Music{Title: "Song Title", Artists: []string{"John Doe"}, ISRC: "123456789", ImageURL: "https.images.google.com/123"},
			repositoryErr: nil,
			wantErr:       false,
		},
		{
			name:        "Invalid Music",
			inputMusic:  entity.Music{},
			validateErr: errors.New("validation error"),
			wantErr:     true,
		},
		{
			name:          "Repository Error",
			inputMusic:    entity.Music{Title: "Song Title", Artists: []string{"John Doe"}},
			repositoryErr: errors.New("repository error"),
			wantErr:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockMusicRepository := mocks.MockMusicRepository{}
			mockMusicRepository.On("Insert", mock.Anything).Return(tt.repositoryErr)

			service := NewMusicService(&mockMusicRepository)

			err := service.Insert(tt.inputMusic)

			if (err != nil) != tt.wantErr {
				t.Errorf("MusicService.Insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
