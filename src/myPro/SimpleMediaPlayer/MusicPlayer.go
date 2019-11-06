package SimpleMediaPlayer

import (
	"errors"
	"fmt"
)

type Music struct {
	Name, Artist, Source, Type string // 音乐名字，作者名字，音乐位置，音乐类型
}

type MusicManager struct {
	music []Music
}

func NewMusicManager() *MusicManager { // 生成音乐管理器
	return &MusicManager{music: make([]Music, 0)}
}

func (m *MusicManager) Len () int { // 获取音乐的数量
	return len(m.music)
}

func (m *MusicManager) Get (tj interface{}) (*Music, error) { // 查找音乐
	if m.Len() == 0 {
		return nil, errors.New("no music")
	}
	switch tj.(type) {
	case string:
		fmt.Println("根据音乐名字查找")
		var result interface{}
		for _, value := range m.music{
			if value.Name == tj {
				result = &value
			}
		}
		if result != nil {
			return result.(*Music), nil
		} else {
			return nil, errors.New("not find music")
		}
	default:
		return nil, errors.New("请输入正确的查找条件")
	}
}

func (m *MusicManager) Add (music *Music)  { // 添加音乐
	m.music = append(m.music, *music)
}

func (m *MusicManager) Remove (name string) *Music { // 移除音乐
	index := -1
	for key, value := range m.music {
		if value.Name == name {
			index = key
		}
	}
	if index < 0 || index >= m.Len() {
		return nil
	}
	removeMusic := m.music[index]
	if index > 0 && index < m.Len() - 1 {
		// 中间删除
		m.music = append(m.music[:index - 1], m.music[index + 1:]...)
	} else if index == 0 {
		// 头上删除
		m.music = append(m.music[1:])
	} else {
		// 尾部删除
		m.music = append(m.music[:m.Len() - 1])
	}
	if m.Len() == 0 {
		m.music = make([]Music, 0)
	}
	return &removeMusic
}