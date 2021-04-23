// Package main x-stock is my stock bot
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/axiaoxin-com/logging"
	"github.com/axiaoxin-com/x-stock/exporter"
	"github.com/axiaoxin-com/x-stock/parser"
)

func main() {
	logging.SetLevel("info")
	ctx := context.Background()
	stocks, err := parser.AutoFilterStocks(ctx)
	if err != nil {
		logging.Fatal(ctx, err.Error())
	}
	data := exporter.InitExportData(ctx, stocks)
	data.SortByPriceSpace()
	filename := fmt.Sprintf("./docs/selected_stocks_%s.json", time.Now().Format("20060102"))
	_, err = data.ExportJSON(ctx, filename)
	if err != nil {
		logging.Fatal(ctx, err.Error())
	}
}
