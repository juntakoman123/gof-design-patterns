package main

import "fmt"

type Handler interface {
	SetNext(Handler)
	Handle(level string, message string)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(h Handler) {
	b.next = h
}

func (b *BaseHandler) CallNext(level, msg string) {
	if b.next != nil {
		b.next.Handle(level, msg)
	}
}

type InfoHandler struct {
	BaseHandler
}

func (h *InfoHandler) Handle(level, msg string) {
	if level == "info" {
		fmt.Println("[INFO]:", msg)
	} else {
		h.CallNext(level, msg)
	}
}

type ErrorHandler struct {
	BaseHandler
}

func (h *ErrorHandler) Handle(level, msg string) {
	if level == "error" {
		fmt.Println("[ERROR]:", msg)
	} else {
		h.CallNext(level, msg)
	}
}

func main() {
	info := &InfoHandler{}
	err := &ErrorHandler{}

	info.SetNext(err)

	// クライアントは最初のhandlerに投げるだけでOK
	info.Handle("info", "これは情報です")
	info.Handle("error", "これはエラーです")
	info.Handle("debug", "これは無視されます") // どこも処理しない
}
