# README
[![release](https://github.com/toriwasa/grouping/actions/workflows/release.yml/badge.svg)](https://github.com/toriwasa/grouping/actions/workflows/release.yml)

## Overview
- n個の連番をランダムにg個のグループに振り分けるCLIツール
- グループごとに指定した区切り文字で連結した文字列を標準出力する

## Usage
- https://github.com/toriwasa/grouping/releases/latest から実行環境に応じた実行ファイルをダウンロードして解凍する
- 例: Windows(64bit)であれば grouping_X.Y.Z_windows_amd64.zip
- 解凍した実行ファイルをターミナル(Powershellやbashなど)で以下のように実行する

```txt
Show Help
$ grouping --help

Usage:
$ grouping -n <number of elements> -g <number of groups> -d <delimiter> -v

Example:
$ grouping -n 10 -g 4 -d ","
3,4,9
2,5,6
7,8
0,1
```

## How to Develop
- VSCode で このリポジトリを開く
- .devcontainer 配下のファイルを利用して開発コンテナを開く
- 起動したコンテナ上のbashで以下の起動コマンドを実行するとカレントディレクトリのコードが実行される

```bash
go run cmd/grouping/main.go --help
```

## How to Build
- VSCode で このリポジトリを開く
- .devcontainer 配下のファイルを利用して開発コンテナを開く
- 起動したコンテナ上のbashで以下のビルドコマンドを実行する

```bash
# Linux向けビルド
go build -o grouping cmd/grouping/main.go
# Winodws向けビルド
GOOS=windows GOARCH=amd64 go build -o grouping.exe cmd/grouping/main.go
```

## How to Test
- VSCode で このリポジトリを開く
- .devcontainer 配下のファイルを利用して開発コンテナを開く
- 起動したコンテナ上のbashで以下のテストコマンドを実行する

```bash
go test -v ./...
```
