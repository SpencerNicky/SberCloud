package main

import (
	playlist "SberCloudTest/Practice/server/internal/Playlist"
	"time"
)

func main() {
	myPlaylist := playlist.NewPlaylist()

	myPlaylist.AddSong(playlist.Song{
		Name:     "Song 1",
		Duration: 5 * time.Second,
	})
	myPlaylist.AddSong(playlist.Song{
		Name:     "Song 2",
		Duration: 10 * time.Second,
	})
	myPlaylist.AddSong(playlist.Song{
		Name:     "Song 3",
		Duration: 7 * time.Second,
	})

	myPlaylist.Play()

	time.Sleep(3 * time.Second)

	myPlaylist.Pause()

	time.Sleep(2 * time.Second)
}
