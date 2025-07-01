package main

import (
	"fmt"
)

// Observerインターフェース
type Observer interface {
	Update(data string)
}

// Subject（通知元）インターフェース
type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyAll()
}

// 実装: ConcreteSubject
type NewsPublisher struct {
	observers []Observer
	news      string
}

func (p *NewsPublisher) Register(o Observer) {
	p.observers = append(p.observers, o)
}

func (p *NewsPublisher) Unregister(o Observer) {
	for i, obs := range p.observers {
		if obs == o {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
			break
		}
	}
}

func (p *NewsPublisher) NotifyAll() {
	for _, o := range p.observers {
		o.Update(p.news)
	}
}

func (p *NewsPublisher) Publish(news string) {
	p.news = news // 状態変更
	p.NotifyAll()
}

// 実装: ConcreteObserver
type Subscriber struct {
	name string
}

func (s *Subscriber) Update(data string) {
	fmt.Printf("%s received news: %s\n", s.name, data)
}

func main() {
	publisher := &NewsPublisher{}

	alice := &Subscriber{name: "Alice"}
	bob := &Subscriber{name: "Bob"}

	publisher.Register(alice)
	publisher.Register(bob)

	publisher.Publish("Go 1.22 is released!")

	publisher.Unregister(bob)

	publisher.Publish("Go generics improved!")
}
