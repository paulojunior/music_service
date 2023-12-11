package contract

import "github.com/paulojunior/code-challange/integration"

type SpotifyIntegration interface {
	GetTrackByISRC(ISRC string) (integration.Item, error)
}
