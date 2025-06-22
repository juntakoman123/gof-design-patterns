package main

import "fmt"

// 商品在庫システム
type InventorySystem struct{}

func (i *InventorySystem) CheckStock(item string) bool {
	fmt.Println("在庫を確認:", item)
	return true // 在庫あり
}

// 支払いシステム
type PaymentSystem struct{}

func (p *PaymentSystem) ProcessPayment(amount int) bool {
	fmt.Println("支払い処理中... 金額:", amount)
	return true // 支払い成功
}

// 配送システム
type ShippingSystem struct{}

func (s *ShippingSystem) ArrangeShipping(item string) {
	fmt.Println("配送を手配:", item)
}

type OrderFacade struct {
	inventory *InventorySystem
	payment   *PaymentSystem
	shipping  *ShippingSystem
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		inventory: &InventorySystem{},
		payment:   &PaymentSystem{},
		shipping:  &ShippingSystem{},
	}
}

func (f *OrderFacade) PlaceOrder(item string, amount int) {
	fmt.Println("=== 注文処理開始 ===")

	if !f.inventory.CheckStock(item) {
		fmt.Println("在庫がありません")
		return
	}

	if !f.payment.ProcessPayment(amount) {
		fmt.Println("支払いに失敗しました")
		return
	}

	f.shipping.ArrangeShipping(item)
	fmt.Println("=== 注文完了 ===")
}

func main() {
	facade := NewOrderFacade()

	// クライアントはこの1メソッドだけ呼べばOK
	facade.PlaceOrder("ノートパソコン", 120000)
}
