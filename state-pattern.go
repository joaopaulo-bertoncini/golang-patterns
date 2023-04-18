package main

import "fmt"

// What is the State Pattern?
// The State Pattern is a behavioral design pattern that allows an object
// to alter its behavior when its internal state changes. This pattern is
// useful when an object needs to change its behavior based on its current
// state, but you don’t want to use a long list of conditional statements
// to handle all the possible states.

type Player struct {
	currentState State
	idleState    State
	walkingState State
	jumpingState State
	x            int
	y            int
}

func (p *Player) setState(state State) {
	p.currentState = state
}

func (p *Player) moveUp() {
	p.currentState.moveUp()
}

func (p *Player) moveDown() {
	p.currentState.moveDown()
}

func (p *Player) moveLeft() {
	p.currentState.moveLeft()
}

func (p *Player) moveRight() {
	p.currentState.moveRight()
}

func (p *Player) jump() {
	p.currentState.jump()
}

type State interface {
	moveUp()
	moveDown()
	moveLeft()
	moveRight()
	jump()
}

type IdleState struct {
	player *Player
}

func (s *IdleState) moveUp() {}

func (s *IdleState) moveDown() {}

func (s *IdleState) moveLeft() {}

func (s *IdleState) moveRight() {}

func (s *IdleState) jump() {
	fmt.Println("Jumped!")
	s.player.setState(s.player.jumpingState)
}

type WalkingState struct {
	player *Player
}

func (s *WalkingState) moveUp() {
	s.player.y -= 1
}

func (s *WalkingState) moveDown() {
	s.player.y += 1
}

func (s *WalkingState) moveLeft() {
	s.player.x -= 1
}

func (s *WalkingState) moveRight() {
	s.player.x += 1
}

func (s *WalkingState) jump() {
	fmt.Println("Jumped!")
	s.player.setState(s.player.jumpingState)
}

type JumpingState struct {
	player *Player
}

func (s *JumpingState) moveUp() {
	s.player.y -= 2
}

func (s *JumpingState) moveDown() {
	s.player.y += 2
}

func (s *JumpingState) moveLeft() {
	s.player.x -= 2
}

func (s *JumpingState) moveRight() {
	s.player.x += 2
}

func (s *JumpingState) jump() {}

func main() {
	idleState := &IdleState{}
	walkingState := &WalkingState{}
	jumpingState := &JumpingState{}
	player := &Player{
		currentState: idleState,
		idleState:    idleState,
		walkingState: walkingState,
		jumpingState: jumpingState,
		x:            0,
		y:            0,
	}
	player.moveRight() // x: 1, y: 0
	player.moveUp()    // x: 1, y: -1
	player.jump()      // Jumped!
	player.moveRight() // x: 3, y: -3
}
