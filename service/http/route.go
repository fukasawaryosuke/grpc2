package http

import (
	"fmt"
	"log"
	"net"
	"github.com/fukasawaryosuke/serve_streaming_grpc_app/service/usage/usecase"

	dessertGrpc "github.com/fukasawaryosuke/serve_streaming_grpc_app/grpc"

	"github.com/labstack/echo/v4"
)

func UsageRoutes(g *echo.Group, handler IUsageHandler) {
	g.GET("/sampleGrpc", handler.SampleGrpc)
}

func InitializeUsageRoutes(e *echo.Echo) {
	usageUsecase := usecase.NewUsageUsecase()
	usageHandler := NewUsageHandler(usageUsecase)

	usageGroup := e.Group("/usage")

	UsageRoutes(usageGroup, usageHandler)

	go startGrpcServer()
}

// gRPCサーバーを起動する関数
func startGrpcServer() {
	// gRPCサーバをlocalhost:10001で起動します
    lis, err := net.Listen("tcp", "localhost:10001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// grpc/usecase.goで定義したサーバを動かす処理を起動
    s := dessertGrpc.NewServer()

	fmt.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v", err)
	}
}
