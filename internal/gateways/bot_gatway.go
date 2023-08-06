package gateway

type BotGateway interface {
	GetStockQuote(stockCode string) (float64, error)
}
