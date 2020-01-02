# APIサーバーの環境構築

## Golangの実行環境構築

1. Golangのダウンロード＆インストール

Mac
```sh
$ brew install go
```

Linux
```sh
$ sudo apt install golang-go
```

2. PATHを通す

bash
```sh
# ~/.bashrc に下記のコードを追加して保存
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
```

fish
```sh
# ~/.config/fish/config.fish に下記のコードを追加して保存
set -x GOPATH $HOME/go
set -x PATH $PATH $GOPATH/bin
```

3. 動作確認

```sh
#  下記のようにgoコマンドを叩いて、出力があれば環境構築完了
$ go version
go version go1.13.5 darwin/amd64
```
