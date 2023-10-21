# はじめに
今回はプログラミング初心者がGolangを勉強したいと思った時にどのような流れで勉強したのかとアウトプットを含めて備忘録的にこのREADME.mdを書いている.

# 環境
今回はDockerなどは使わずにローカル環境動かしていく.
```
ハードウェア M1 macbook air (arm64)
エディタ cursor(VScodeベースのChatGPTを搭載したエディタ)
golang version 1.21.3
```

# go mod initのためにgit initを行う.
最近のgolangはgo_envを使わずにgo modを使用する.

しかし`go mod`を行うためには`git init`をする必要がある.

忘れずにしよう.

私はここで3時間ほどハマってしまった.

# go mod initをしよう.

`git init`を行えたらに`go mod init`をしよう.

npm initみたいなものだ.

そしたら`go.mod`ファイルが生成するはずなので確認しよう.

そしたらもうgoで実際にコードを書くことができる.

# Hello world
実際にコードを書いてみよう.

まずはファイル作成.

名前の最後に.goをつければgoのファイルになる.

私は今回main.goという名前は使わずにhello.goという名前にした.
```go
package main

import "fmt"

func main() {
  fmt.Println("Hello world")
}
```
このコードを実際に動かすには`go run hello.go`をターミナルで入力し実行する.

※packageの後の名前は任意の名前なのかそれともmainなどの予約後でないといけないのかは後でかく.

ちなみに
```go
package hello

import "fmt"

func hello() {
  fmt.Println("Hello world")
}
```
だと次のエラーが出た.

```
# command-line-arguments
runtime.main_main·f: function main is undeclared in the main package
```
cursorに流してみると次のようになる.(cursorくんは非常に優秀でChatGPTを有しているためすぐにエラー内容を聞ける.)

```
このエラーメッセージは、mainパッケージ内でmain関数が宣言されていないために発生しています。Goのプログラムでは、mainパッケージ内にmain関数を定義する必要があります。

コードを以下のように修正して、main関数を追加してください：
```
とりあえずpackageと関数一つはmainにしようという話.

