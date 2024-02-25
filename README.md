# カレントディレクトリを公開するファイルサーバー

## Feature

- ファイルをアップロード
- ファイルの一覧を表示
- ファイルのダウンロード

## How to use

実行すると http://localhost:20768 にファイルサーバーが起動します

## Download

バイナリファイルをこちらからダウンロードできます

https://github.com/kajikentaro/file-server/releases/latest

Go言語環境がインストールされている場合には、以下コマンドを使ってインストールすることもできます

```
$ go install github.com/kajikentaro/file-server@latest 
```


## Options

```
$ file-server -h
Usage of ./file-server:
  -d string
        Directory to serve or store uploaded files (default "./")
  -p int
        Port number to listen on (default 20768)
```



## Build from source

以下コマンドでビルドできます

```
# Windows用
GOOS=windows GOARCH=amd64 go build -o file-server-1.0-windows-amd64.exe main.go

# Linux用
GOOS=linux GOARCH=amd64 go build -o file-server-1.0-linux-amd64 main.go
```
