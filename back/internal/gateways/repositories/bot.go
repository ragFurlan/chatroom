package repository

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type StockBotGateway struct {
	UrlStock string `json:"stock"`
}

func NewStockBotGateway(urlStock string) *StockBotGateway {
	return &StockBotGateway{
		UrlStock: urlStock,
	}
}

func (bg *StockBotGateway) GetStockQuote(stockCode string) (float64, error) {
	log.Printf("service: GetStockQuote - stockCode: %s", stockCode)

	url := fmt.Sprintf("https://stooq.com/q/l/?s=%s.us&f=sd2t2ohlcv&h&e=csv", stockCode)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("service: GetStockQuote - method: get - err: %s", err)
		return 0, err
	}
	defer resp.Body.Close()

	reader := csv.NewReader(resp.Body)
	_, err = reader.Read()
	if err != nil {
		log.Printf("service: GetStockQuote - method: read - err: %s", err)
		return 0, fmt.Errorf("Error reading header")
	}

	row, err := reader.Read()
	if err != nil {
		log.Printf("service: GetStockQuote - method: read - err: %s", err)
		return 0, err
	}

	if row[6] == "N/D" {
		return 0, fmt.Errorf("This action does not exist")
	}

	value, err := strconv.ParseFloat(row[6], 64)
	if err != nil {
		log.Printf("service: GetStockQuote - method: ParseFloat - err: %s", err)
		return 0, err
	}

	return value, nil
}
