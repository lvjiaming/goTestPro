/**
 有限状态机
State
 */
package FiniteStateMachine

import "reflect"

type State interface {
	Name() string  // 名字
	EnableSameTransit() bool // 是否在同状态间可以转换
	CanTransitTo() bool // 是否能转换到另一个状态
	onBegin()  // 转换开始
	onEnd() // 转换失败
}

func StateName(s State) string {
	if s == nil {
		return "none"
	}
	return reflect.TypeOf(s).Elem().Name()
}

/**
 自定义的状态
 */
type SelfState struct {
	name string
}

func (s *SelfState) Name () string {
	return s.name
}

func (s *SelfState) SetName (n string)  {
	s.name = n
}

func (s *SelfState) EnableSameTransit () bool {
	return false // 同状态间不能转换
}

func (s *SelfState) onBegin ()  {

}

func (s *SelfState) onEnd ()  {

}

func (s *SelfState) CanTransitTo () bool {
	return true  // 默认能转换到任何状态
}