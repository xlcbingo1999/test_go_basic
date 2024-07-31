package Struct

// 外观模式, 本质上就是一个系统里面会有很多的子系统, 这些子系统会分别管理复杂的逻辑

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	security     *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func NewWalletFacade() *WalletFacade {
	return &WalletFacade{
		account:      &Account{},
		wallet:       &Wallet{},
		security:     &SecurityCode{},
		notification: &Notification{},
		ledger:       &Ledger{},
	}
}

type Account struct{}
type Wallet struct{}
type SecurityCode struct{}
type Notification struct{}
type Ledger struct{}

func RunPreview() {
	_ = NewWalletFacade()
}
