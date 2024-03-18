# serve_streaming_grpc_app

## ディレクトリ構成

```
- main.go -- APIのエンドポイント
- service
    - http
        - handler.go -- ユースケースを呼ぶためのハンドラ
        - route.go -- ハンドラの設定とgRPCサーバの設定
    - usage
      - usecase
          - usage_usecase.go -- クライアント側のサーバ処理を担当
- grpc
    - usecase.go -- gRPCのサーバ処理を担当
    - dessert.proto -- プロトコル設定ファイル
- pkg
    - grpc
        - dessert.pb.go -- プロトコル設定ファイルから生成されたファイルでメッセージが定義
        - dessert_grpc.pb.go -- プロトコル設定ファイルから生成されたファイルでクライアントとサーバのインターフェースが定義
```

### protoc

https://grpc.io/docs/protoc-installation/

```
protoc --go_out=. --go-grpc_out=. --proto_path=grpc ./grpc/dessert.proto
```

- --go_outでメッセージの定義ファイルの出力先を指定
- --go-grpc_outでサーバのインターフェイスの定義ファイルの出力先を指定
- --proto_pathでプロトコルファイルの存在するディレクトリを指定

### go

```
go mod init ＜プロジェクト名＞
go mod tidy
go run main.go
```

### リクエスト

```
curl -X GET http://localhost:8080/usage/sampleGrpc
```
