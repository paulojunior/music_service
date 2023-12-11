package service

import (
	"fmt"

	contract "github.com/paulojunior/code-challange/contract/repository"
	"github.com/paulojunior/code-challange/entity"
	"github.com/pkg/errors"
)

type MusicService struct {
	repository contract.MusicRepository
}

func NewMusicService(repo contract.MusicRepository) *MusicService {
	return &MusicService{
		repository: repo,
	}
}

func (s *MusicService) FindByISRC(isrc string) (musics []entity.Music, err error) {
	if isrc == "" {
		return nil, fmt.Errorf("invalid ISRC: %v", isrc)
	}

	musics, err = s.repository.FindByISRC(isrc)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find music by ISRC")
	}

	return musics, nil
}

func (s *MusicService) FindByArtistName(artistName string) (musics []entity.Music, err error) {
	if artistName == "" {
		return nil, fmt.Errorf("invalid artist name: %s", artistName)
	}

	musics, err = s.repository.FindByArtistName(artistName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find music by artist name")
	}

	return musics, nil
}

func (s *MusicService) Insert(music entity.Music) error {
	err := music.Validate()
	if err != nil {
		return errors.Wrap(err, "music validation failed")
	}

	err = s.repository.Insert(music)
	if err != nil {
		return errors.Wrap(err, "failed to insert music")
	}

	return nil
}
