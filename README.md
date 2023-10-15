# README

## Overview
- n個の連番をランダムにg個のグループに振り分けるCLIツール
- グループごとに指定した区切り文字で連結した文字列を標準出力する

## Usage

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
