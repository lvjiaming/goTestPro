package SimpleMediaPlayer

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

var musicManager *MusicManager

func musicCatch(dealWith []string)  { // 对音乐管理器的操作（增删查）
	if len(dealWith) < 2 {
		err := errors.New("please input right command")
		fmt.Println(err)
		return
	}
	switch dealWith[1] {
	case "Add":
		if len(dealWith) == 6 {
			musicManager.Add(&Music{dealWith[2], dealWith[3], dealWith[4], dealWith[5]})
			fmt.Println("curMusic count: ", musicManager.Len())
		} else {
			err := errors.New("command: music Add <Id> <name> <artist> <source> <type>")
			fmt.Println(err)
		}
	case "Remove":
		if len(dealWith) == 3 {
			musicManager.Remove(dealWith[2])
			fmt.Println("curMusic count: ", musicManager.Len())
		} else {
			err := errors.New("command: music Remove <name>")
			fmt.Println(err)
		}
	default:
		err := errors.New("command: music Add or Remove, not " + dealWith[1])
		fmt.Println(err)
	}
}

func playCommand(command []string)  {
	if len(command) != 2 {
		err := errors.New("command: play <name>")
		fmt.Println(err)
		return
	}
	music, err := musicManager.Get(command[1])
	if err != nil {
		err := errors.New("not find music")
		fmt.Println(err)
		return
	}
	Play(music.Source, music.Type)
}

func MusicPlay()  {
	fmt.Println(`
        Enter following commands to control the player:
        music Add <name><artist><source><type> -- Add a music to the music lib
        music Remove <name> -- Remove the specified music from the lib
        play <name> -- Play the specified music
    `)
	musicManager = NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for  {
		fmt.Println("Enter Command ->")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		command := strings.Split(line, " ") // 将命令以空格为间隔符变换成切片
		switch command[0] {
		case "music":
			musicCatch(command)
		case "play":
			playCommand(command)
		default:
			fmt.Println("Unrecognized command: ", command[0])
		}
	}
}