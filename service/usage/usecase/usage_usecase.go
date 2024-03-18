package usecase

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"time"

	pb "github.com/fukasawaryosuke/serve_streaming_grpc_app/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IUsageUsecase interface {
	GetDessertStream()
}

type usageUsecase struct{}

func NewUsageUsecase() IUsageUsecase {
	return &usageUsecase{}
}
func (uu *usageUsecase) GetDessertStream() {
	startTime := time.Now() // 関数実行開始時刻を記録

	fmt.Printf("---- Start GetDessertStream ----\n\n")
	  port := "localhost:10001"
	  // gRPCサーバのコネクション
    conn, err := grpc.Dial(
    	port,
			// コネクションでSSL/TLSを使用しない
    	grpc.WithTransportCredentials(insecure.NewCredentials()),
			// コネクションが確立されるまで待機する(同期処理をする)
    	grpc.WithBlock(),
    )
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// コネクションを使って、gRPCのサービスのクライアントを受け取る
	client := pb.NewDessertServiceClient(conn)

	// 使用したいストリームを取得しておく
	stream, err := client.GetDessertStream(context.Background(), &pb.DessertRequest{
		Name: "アップルパイ", // このサンプルではリクエストの値は使わない
		Id:   1,
	})
	if err != nil {
		fmt.Printf("could not get dessert stream: %v", err)
	}

	// 非同期処理のための待機グループを作成
	var wg sync.WaitGroup

	// ディナーデータの処理（例：カレー）
	wg.Add(1)
	go func() {
		defer wg.Done()
		// ディナーデータの配列
		dinners := []string{"カレー", "スパゲッティ", "寿司", "ラーメン"}

		// ディナーデータの処理
		for _, dinner := range dinners {
			time.Sleep(1 * time.Second) // 時間待ちを1秒に変更
			fmt.Printf("晩御飯データ :: %s\n", dinner)
		}
	}()

	// デザートデータの受信
	// gRPCの処理側でSend()されたデータをここで受け取る。（デザートのデータを10回受け取る）
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			dessert, err := stream.Recv()
			// ストリーム処理が終了した場合はエラーが「io.EOF」になる
			if err == io.EOF {
				fmt.Println("gRPCストリーム処理完了！")
				break
			}
			if err != nil {
				fmt.Printf("デザートの取得でエラー発生！ : %v", err)
				return
			}
			// ランダムなランクを生成
			rank := rand.Intn(5) + 1
			fmt.Printf("gRPC通信で受け取ったデザート名: %s, 説明: %s, ランク: %d\n", dessert.Name, dessert.Description, rank)
		}
	}()

	// 待機グループの完了を待つ
	wg.Wait()

	defer fmt.Println("\n\n---- End GetDessertStream ----\n")

	endTime := time.Now()              // 関数実行終了時刻を記録
	duration := endTime.Sub(startTime) // 実行時間を計算
	fmt.Printf("実行時間: %v\n", duration)
}
