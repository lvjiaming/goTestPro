package FiniteStateMachine

import "fmt"

type Idle struct {
	SelfState
}

func (i *Idle) onBegin ()  {
	fmt.Println("IdleState begin")
}

func (i *Idle) onEnd ()  {
	fmt.Println("IdleState end")
}

func (i *Idle) EnableSameTransit () bool {
	return true
}

type Move struct {
	SelfState
}

func (m *Move) onBegin ()  {
	fmt.Println("MoveState begin")
}

func (m *Move) onEnd ()  {
	fmt.Println("MoveState end")
}

func (m *Move) EnableSameTransit () bool {
	return true
}

type Jump struct {
	SelfState
}

func (j *Jump) onBegin ()  {
	fmt.Println("JumpState begin")
}

func (j *Jump) onEnd ()  {
	fmt.Println("JumpState end")
}
