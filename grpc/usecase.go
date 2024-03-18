package grpc

import (
	"time"

	pb "github.com/fukasawaryosuke/serve_streaming_grpc_app/pkg/grpc"

	"google.golang.org/grpc"
)

type DessertStreamServer struct {
	pb.DessertServiceServer
}

// NewServer　デザートを返却するgRPCサーバーを作成します。
func NewServer() *grpc.Server {
	s := grpc.NewServer()

	// gRPCサーバーにDessertStreamServerを登録。
	pb.RegisterDessertServiceServer(s, &DessertStreamServer{})
	return s
}

// GetDessertStream デザート情報をストリームで送信します。
func (s *DessertStreamServer) GetDessertStream(req *pb.DessertRequest, stream pb.DessertService_GetDessertStreamServer) error {
	desserts := []string{"チーズケーキ", "ティラミス", "マカロン", "エクレア", "カンノーリ", "パンナコッタ", "モンブラン", "クレープ", "シュークリーム", "フルーツタルト"}

	for _, dessertName := range desserts {
		time.Sleep(500 * time.Millisecond)

		// gRPCのストリームを介してクライアントにデータを送信
		// stream.Sendを使うことで呼び出し元の関数にデータを送信することができる
		err := stream.Send(&pb.DessertResponse{
			Description: "美味しい" + dessertName + "です",
			Name:        dessertName,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
