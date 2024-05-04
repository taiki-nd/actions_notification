# Dockerfile
FROM golang:1.22

WORKDIR /app

# スクリプトや必要なファイルをコピー
COPY . .

# 初期コマンド実行
RUN go mod download

# 実行コマンド
CMD ["go", "run", "./main.go"]
