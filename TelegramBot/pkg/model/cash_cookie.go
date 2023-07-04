package model

type CashCookie struct {
	cash map[int64]string
}

var cashCookie *CashCookie = nil

func (c *CashCookie) init() {
	c.cash = make(map[int64]string)
}

func initCash() {
	cashCookie = new(CashCookie)
	cashCookie.init()
}

func GetCookie(chatID int64) string {
	if cashCookie == nil {
		initCash()
	}
	return cashCookie.cash[chatID]
}

func SetCookie(chatID int64, cookie string) {
	if cashCookie == nil {
		initCash()
	}
	cashCookie.cash[chatID] = cookie
}
