package main

import "fmt"

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Println("Email送信:", message)
}

type LogNotifier struct {
	Wrapped Notifier
}

func (l *LogNotifier) Send(message string) {
	fmt.Println("[LOG] メッセージ送信前:", message)
	l.Wrapped.Send(message)
}

type EncryptedNotifier struct {
	Wrapped Notifier
}

func (e *EncryptedNotifier) Send(message string) {
	encrypted := "***" + message + "***" // 簡易的な「暗号化」
	e.Wrapped.Send(encrypted)
}

func main() {
	// 元の機能
	base := &EmailNotifier{}

	// デコレーターを順にラップ
	withLog := &LogNotifier{Wrapped: base}
	withEncrypt := &EncryptedNotifier{Wrapped: withLog}

	// 実行
	withEncrypt.Send("こんにちは！")
}
