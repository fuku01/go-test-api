FROM golang:1.20.6-alpine

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

# コンテナ内に`app`ディレクトリを作成
RUN mkdir /go/src/app

# ワーキングディレクトリの設定
WORKDIR /go/src/app

# パッケージの更新とgitのインストール
RUN apk update && apk add --no-cache git make

# ホストのファイルをコンテナの作業ディレクトリに移行
COPY . .

# AIRのインストール
RUN go install github.com/cosmtrek/air@latest

# go.modを参照し、go.sumファイルの更新を行う
RUN go mod tidy

# ポート番号を指定（必要に応じて変更）
EXPOSE 8000

# localではホットリロードを有効にしたいのでairで起動する
CMD ["air", "-c", ".air.toml"]