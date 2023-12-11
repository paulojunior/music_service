package contract

import "github.com/paulojunior/code-challange/entity"

type MusicService interface {
	FindByISRC(isrc string) (musics []entity.Music, err error)
	FindByArtistName(artistName string) (musics []entity.Music, err error)
	Insert(music entity.Music) error
}
