package gateways

type BotGateway interface {
	GetStockQuote(stockCode string) (float64, error)
}
