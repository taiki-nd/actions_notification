FROM golang:1.22.2 as builder

WORKDIR /src

# モジュールファイルをコピー
COPY go.mod go.sum ./
# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションのビルド
RUN CGO_ENABLED=0 go build -o /bin/action ./cmd


FROM alpine:3.19
COPY --from=builder /bin/action /bin/action

ENTRYPOINT ["/bin/action"]
