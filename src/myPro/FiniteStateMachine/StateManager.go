/**
 状态管理器
 */
package FiniteStateMachine

import (
	"errors"
)

// 状态没有找到的错误
var ErrStateNotFound = errors.New("state not found")
// 禁止在同状态间转移
var ErrForbidSameStateTransit = errors.New("forbid same state transit")
// 不能转移到指定状态
var ErrCannotTransitToState = errors.New("cannot transit to state")

type StateManager struct {
	stateByName map[string]State // 记录状态的map

	Change func(from, to State) // 状态转换

	curState State  // 当前状态
}

func (m *StateManager) Add (s State) (err error) {
	if s == nil {
		err = errors.New("请添加状态")
	}
	name := StateName(s)
	s.(interface{ // 使用类型断言
		SetName(s string)
	}).SetName(name)
	if m.Get(s.Name()) != nil {
		err = errors.New("要添加的状态已存在")
	}
	m.stateByName[s.Name()] = s
	return err
}

func (m *StateManager) Get (name string) State {
	if s, ok := m.stateByName[name]; ok {
		return s
	}
	return nil
}

func (m *StateManager) CurState () State { // 获取当前状态
	return m.curState
}

func (m *StateManager) CanCurrTransitTo (to string) bool { // 是否能转到目标状态
	if m.curState == nil {
		return  true
	}
	if m.curState.Name() == to && !m.curState.EnableSameTransit() { // 如果相同状态间不能转换，则不能转换
		return false
	}

	return m.curState.CanTransitTo()
}

func (m *StateManager) Transit (to string) (err error) {
	next := m.Get(to)
	if next == nil {
		err = ErrStateNotFound
	}

	per := m.curState

	if m.curState != nil {

		if m.curState.Name() == to && !m.curState.EnableSameTransit() {
			return ErrForbidSameStateTransit
		}

		canTransit := m.CanCurrTransitTo(to)
		if !canTransit {
			err = ErrCannotTransitToState
		}
		m.curState.onEnd()
	}

	m.curState = next

	next.onBegin()

	if m.Change != nil {
		m.Change(per, m.curState)
	}
	return nil
}

/**
 获取一个管理器
 */
func GetStateManager() *StateManager {
	return &StateManager{stateByName: make(map[string]State)}
}