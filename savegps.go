package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"github.com/labstack/echo/v4"
)

func SaveGPS(c echo.Context) error {
	// ファイルを開く
	file, err := os.OpenFile("public/gps.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// 書き込むデータを作成
	gps := []string{c.QueryParam("lat"), c.QueryParam("lng")}

	// CSV形式で追加書き込み
	writer := csv.NewWriter(file)
	writer.Write(gps)
	writer.Flush()

	return c.String(http.StatusOK, "OK")
}