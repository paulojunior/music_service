package integration

import (
	"reflect"
	"testing"
)

func TestTracksIntegration_GetMostPopularTrack(t *testing.T) {
	tests := []struct {
		name string
		args []Item
		want Item
	}{
		{
			name: "ShouldReturnTheMostPopularTrack",
			args: []Item{
				{
					Name:       "ABC",
					Popularity: 100,
				},
				{
					Name:       "DEF",
					Popularity: 99,
				},
			},
			want: Item{
				Name:       "ABC",
				Popularity: 100,
			},
		},
	}
	for _, tt := range tests {
		integrationService := SpotifyIntegration{}

		t.Run(tt.name, func(t *testing.T) {
			if got := integrationService.GetMostPopularTrack(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TracksIntegration.GetMostPopularTrack() = %v, want %v", got, tt.want)
			}
		})
	}
}
