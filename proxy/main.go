package main

import "fmt"

type DataFetcher interface {
	Fetch() string
}

type RealFetcher struct{}

func (r *RealFetcher) Fetch() string {
	fmt.Println("リアルなデータ取得中...")
	// 本来は重い処理（DB/APIアクセスなど）
	return "データ: ユーザー情報"
}

type CachingFetcher struct {
	real    DataFetcher
	cached  string
	fetched bool
}

func (c *CachingFetcher) Fetch() string {
	if c.fetched {
		fmt.Println("キャッシュから取得")
		return c.cached
	}
	fmt.Println("Proxy経由で初回取得")
	c.cached = c.real.Fetch()
	c.fetched = true
	return c.cached
}

func main() {
	real := &RealFetcher{}
	proxy := &CachingFetcher{real: real}

	fmt.Println(proxy.Fetch()) // Proxy経由で初回取得 → リアル取得
	fmt.Println(proxy.Fetch()) // キャッシュから取得
}
