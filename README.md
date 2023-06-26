# 最も単純なファイルサーバー

## Feature

- ファイルをアップロード
- ファイルの一覧を表示
- ファイルのダウンロード

## How to use

実行すると http://localhost:20768 にファイルサーバーが起動します

## Development

以下コマンドでビルドできます

```
# Windows用
GOOS=windows GOARCH=amd64 go build -o file-server-1.0-windows-amd64.exe main.go

# Linux用
GOOS=linux GOARCH=amd64 go build -o file-server-1.0-linux-amd64 main.go
```
