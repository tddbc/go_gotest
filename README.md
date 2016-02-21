# TDDBC GO Project
_このプロジェクトは、TDDBC(Test Driven Development Boot Camp)で使用する GO 言語用のプロジェクトテンプレートです。_

[![Build Status](https://travis-ci.org/tddbc/go_gotest.svg?branch=master)](https://travis-ci.org/tddbc/go_gotest)

## はじめに
サンプルプロジェクトでは、HelloWorldクラスを作成し、Sayメソッドが呼ばれたときに"Hello TDD BootCamp!!"を返すプログラムがついてきます。

## セットアップ
### GO言語のセットアップ
[本家サイト](https://golang.org/doc/install)を参考に環境をセットアップしてください。

**注意**
セットアップの際、環境変数に GOPATH の設定を忘れてしまいうケースが多いようです。
また、`$GOPATH/bin` を環境変数 PATH に含めてください。


#### MacPorts で簡単セットアップ
port コマンドを実行して、GO言語をインストールすることができます。
コマンドは、以下の通りです。

```bash
sudo port selfupdate
sudo port upgrade outdated 
sudo port install go
```

#### Homebrew で簡単セットアップ
brew コマンドを実行して、GO言語をインストールすることができます。
コマンドは、以下の通りです。

```bash
sudo brew update
sudo brew install go
```

#### apt-get で簡単セットアップ
apt-get コマンドを実行して、GO言語をインストールすることができます。
コマンドは、以下の通りです。

```bash
sudo apt-get update
sudo apt-get install go
```

### 依存ライブラリのセットアップ
[testify](https://github.com/stretchr/testify) を使っているので別途セットアップが必要です

```sh
go get github.com/stretchr/testify/assert
```

## テストの実行方法
```
go test
```
