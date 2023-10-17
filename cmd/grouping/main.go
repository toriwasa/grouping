package main

import (
	"flag"
	"io"
	"log"

	"github.com/toriwasa/grouping/app/grouping"
)

func main() {
	// DEBUGログのフォーマットを設定
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// DEBUGログのプレフィックスを設定
	log.SetPrefix("DEBUG: ")

	// コマンドライン引数を解析する。 -n, -g, -d, -v というオプションを定義する
	var n, g int
	var d string
	var isVerbose bool
	// n と g はそれぞれ要素数とグループ数を表す
	// デフォルト値はそれぞれ 10 と 4 である
	flag.IntVar(&n, "n", 10, "number of elements")
	flag.IntVar(&g, "g", 4, "number of groups")

	// d は区切り文字を表す
	// デフォルト値はタブ文字である
	flag.StringVar(&d, "d", "\t", "delimiter")

	// v はログを冗長に出力するモードを表す
	// デフォルト値は false である
	flag.BoolVar(&isVerbose, "v", false, "output verbose log")

	// --help オプションをカスタマイズする
	flag.Usage = func() {
		println("Usage: grouping -n <number of elements> -g <number of groups> -d <delimiter> -v")
		println("Example: grouping -n 10 -g 4 -d \",\"")
		println("Description: generate random sequence and group them")
		println("Options:")
		flag.PrintDefaults()
	}

	// コマンドライン引数を解析する
	flag.Parse()

	// verbose モードでない場合はログを出力しない
	if !isVerbose {
		log.SetOutput(io.Discard)
	}

	// コマンドライン引数を出力する
	log.Printf("n: %d, g: %d, d: \"%s\", v: %t\n", n, g, d, isVerbose)

	// コマンドライン引数を元にパラメータを生成する
	p, err := grouping.NewParameter(n, g, d)
	log.Printf("p: %+v\n", p)

	// パラメータ生成時にエラーが発生した場合は終了する
	if err != nil {
		panic(err)
	}

	// 引数に応じた文字列のイテレータを生成する
	iter := grouping.GenerateGroupedRandomSeqIterator(p)
	// イテレータから値を取り出して出力する
	for {
		s, err := iter()
		if err != nil {
			break
		}
		println(s)
	}
}
