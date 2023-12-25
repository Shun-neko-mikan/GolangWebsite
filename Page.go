package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/kkdai/youtube/v2"
)

type CommonData struct {
	Title string
}

func viewMainPage(c echo.Context) error {
	var common = CommonData{Title: "Main Page"}

	data := struct {
		CommonData
		Constent string
	}{
		CommonData: common,
		Constent:   "Hello, Test!",
	}
	return c.Render(http.StatusOK, "mainPage", data)
}

func viewDetailPae(c echo.Context) error {
	code := c.QueryParam("code")
	var common = CommonData{Title: "Detail Page"}

	data := struct {
		CommonData
		Code string
	}{
		CommonData: common,
		Code: code,
	}
	return c.Render(http.StatusOK, "detailPage", data)
}

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

func GetYoutube(c echo.Context) error {
	id := c.Param("id")
	client := youtube.Client{}
	video, err := client.GetVideo(id)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, "NG")
	}

	stream, _, err := client.GetStream(video, &video.Formats[0])
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, "NG")
	}

	file, err := os.Create("public/static/video.mp4")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, "NG")
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusOK, "NG")
	}

	return c.Render(http.StatusOK, "youtube", id)
}