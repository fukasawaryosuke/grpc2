package main

import (
	"github.com/labstack/echo/v4"
	"github.com/fukasawaryosuke/serve_streaming_grpc_app/service/http"
)

func main() {
	// Echoインスタンスの作成
	e := echo.New()

	// HTTPルーティングの初期化
	http.InitializeUsageRoutes(e)

	// HTTPサーバーを起動
  e.Start(":8080")
}
