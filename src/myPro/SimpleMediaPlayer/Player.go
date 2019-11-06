package SimpleMediaPlayer

import (
	"fmt"
	"time"
)

type Player interface {
	Play(source string)
}

func Play(source, musicType string)  {
	var p Player
	switch musicType {
	case "mp3":
		p = &Mp3Player{}
	case "wav":
		p = &WavPlayer{}
	}
	p.Play(source)
}

type Mp3Player struct {
	stat, pro int
}

func (m *Mp3Player) Play (source string)  { // mp3音乐播放器
	fmt.Println("playing music", source)
	m.pro = 0
	for m.pro < 100 {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print(".")
		m.pro += 10
	}
	fmt.Println("\nmusic finished playing", source)
}

type WavPlayer struct {
	stat, pro int
}

func (m *WavPlayer) Play (source string)  { // wav音乐播放器
	fmt.Println("playing music", source)
	m.pro = 0
	for m.pro < 100 {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print(".")
		m.pro += 10
	}
	fmt.Println("\nmusic finished playing", source)
}
