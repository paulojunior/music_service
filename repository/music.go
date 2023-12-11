package repository

import (
	"gorm.io/gorm"

	"github.com/paulojunior/code-challange/entity"
)

type MusicRepository struct {
	DB *gorm.DB
}

func NewMusicRepository(DB *gorm.DB) *MusicRepository {
	return &MusicRepository{
		DB: DB,
	}
}

func (r *MusicRepository) FindByISRC(isrc string) (musics []entity.Music, err error) {
	err = r.DB.Where("isrc = ? AND deleted_at IS NULL", isrc).Find(&musics).Error
	return musics, err
}

func (r *MusicRepository) FindByArtistName(artistName string) (musics []entity.Music, err error) {
	err = r.DB.Where("LOWER(?) = ANY(SELECT LOWER(unnest(artists))) AND deleted_at IS NULL", artistName).Find(&musics).Error
	return musics, err
}

func (r *MusicRepository) Insert(music entity.Music) error {
	return r.DB.Model(&entity.Music{}).Create(&music).Error
}
