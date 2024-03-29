package database

const (
	FourWeeks = "four weeks"
	SixMonths = "six months"
	AllTime   = "all time"
)

type Track struct {
	TrackName      string `json:"track_name" db:"track_name"`
	ArtistName     string `json:"artist_name" db:"artist_name"`
	SpotifyTrackId string `json:"spotify_track_id" db:"spotify_track_id"`
	TrackImageURL  string `json:"track_image_url" db:"track_image_url"`
	TrackRank      int    `json:"track_rank" db:"track_rank"`
}

type Artist struct {
	ArtistName      string `json:"artist_name"`
	SpotifyArtistId string `json:"spotify_track_id"`
	ArtistImageURL  string `json:"track_image_url"`
	ArtistRank      int    `json:"track_rank"`
}

type Tracks []Track
type TopTracks struct {
	CreatedAt int64  `json:"created_at" db:"created_at"`
	TimeSpan  string `json:"time_span" db:"time_span"`
	Tracks    Tracks `json:"items" db:"items"`
}

type Artists []Artist
type TopArtists struct {
	CreatedAt int64   `json:"created_at" db:"created_at"`
	TimeSpan  string  `json:"time_span" db:"time_span"`
	Artists   Artists `json:"items" db:"items"`
}

type TopInfoResponse struct {
	TopTracks  []TopTracks  `json:"top_tracks"`
	TopArtists []TopArtists `json:"top_artists"`
}
