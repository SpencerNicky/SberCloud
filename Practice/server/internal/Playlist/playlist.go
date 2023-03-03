package Playlist

import (
	"container/list"
	"sync"
	"time"
)

type Song struct {
	Name     string
	Duration time.Duration
}

type Playlist struct {
	songs         *list.List
	currentSong   *list.Element
	isPlaying     bool
	pauseDuration time.Duration
	mu            sync.Mutex
}

type PlaylistControl interface {
	Play()
	Pause()
	AddSong(song Song)
	Next()
	Prev()
}

func NewPlaylist() *Playlist {
	return &Playlist{
		songs: list.New(),
	}
}

func (p *Playlist) Play() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.currentSong == nil {
		p.currentSong = p.songs.Front()
	}
	if p.isPlaying {
		return
	}
	p.isPlaying = true
	go func() {
		for {
			if !p.isPlaying {
				break
			}
			if p.pauseDuration > 0 {
				time.Sleep(p.pauseDuration)
				p.pauseDuration = 0
			}
			if p.currentSong == nil {
				p.isPlaying = false
				break
			}
			song := p.currentSong.Value.(Song)
			time.Sleep(song.Duration)
			p.Next()
		}
	}()
}

func (p *Playlist) Pause() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if !p.isPlaying {
		return
	}
	p.pauseDuration = p.currentSong.Value.(Song).Duration - p.pauseDuration
	p.isPlaying = false
}

func (p *Playlist) AddSong(song Song) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.songs.PushBack(song)
	if p.currentSong == nil {
		p.currentSong = p.songs.Front()
	}
}

func (p *Playlist) Next() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.currentSong == nil {
		return
	}
	if next := p.currentSong.Next(); next != nil {
		p.currentSong = next
	} else {
		p.currentSong = p.songs.Front()
	}
	p.isPlaying = true
}

func (p *Playlist) Prev() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.currentSong == nil {
		return
	}
	if prev := p.currentSong.Prev(); prev != nil {
		p.currentSong = prev
	} else {
		p.currentSong = p.songs.Back()
	}
	p.isPlaying = true
}
