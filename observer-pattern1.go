package main

// Example 1: Simple Event Notification

import (
 "fmt"
 "sync"
)

// Subject struct
type Subject struct {
 observers []chan string
}

// NewSubject creates a new Subject
func NewSubject() *Subject {
 return &Subject{
  observers: make([]chan string, 0),
 }
}

// Subscribe method for adding new observer
func (s *Subject) Subscribe(observer chan string) {
 s.observers = append(s.observers, observer)
}

// NotifyObservers method for notifying all observers
func (s *Subject) NotifyObservers(message string) {
 for _, observer := range s.observers {
  observer <- message // sending the message to the observer
 }
}

func main() {
 subject := NewSubject()

 // creating channels for observers
 observer1 := make(chan string)
 observer2 := make(chan string)

 subject.Subscribe(observer1)
 subject.Subscribe(observer2)

 var wg sync.WaitGroup // using WaitGroup
 wg.Add(2)             // setting counter to the number of observers

 go func() {
  for {
   msg := <-observer1
   fmt.Println("Observer 1 received:", msg)
   wg.Done() // decrement counter when observer 1 is done
  }
 }()

 go func() {
  for {
   msg := <-observer2
   fmt.Println("Observer 2 received:", msg)
   wg.Done() // decrement counter when observer 2 is done
  }
 }()

 subject.NotifyO