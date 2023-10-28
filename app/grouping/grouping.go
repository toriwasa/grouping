package grouping

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

// int型の値を返却するイテレータ型
type intIterator func() (int, error)

// 0から始まるn個の連番をランダムに返すイテレータを返却する
func generateRandomIntIterator(n int) intIterator {
	// n個の要素を持つスライスを生成する
	seq := make([]int, n)
	log.Printf("seq: %v\n", seq)

	// 0, 1, 2, .. n-1 のn個の連番を要素として持つようにスライスを初期化する
	for i := range seq {
		seq[i] = i
	}
	log.Printf("seq: %v\n", seq)

	// スライスの中身をランダムに並び替える
	rand.Shuffle(len(seq), func(i, j int) {
		seq[i], seq[j] = seq[j], seq[i]
	})
	log.Printf("seq: %v\n", seq)

	// インデックスを管理する変数を定義する
	index := -1
	// スライスの要素を1つ返す関数を返却する
	return func() (int, error) {
		index++
		if index >= len(seq) {
			return 0, fmt.Errorf("index out of range")
		}
		return seq[index], nil
	}
}

// IntIterator から n 個の要素を返すイテレータを返却するメソッド
func (iter intIterator) take(n int) intIterator {
	// インデックスを管理する変数を定義する
	index := 0
	// イテレータから n 個の要素を返すイテレータを返却する
	return func() (int, error) {
		index++
		if index > n {
			return 0, fmt.Errorf("index out of range")
		}
		return iter()
	}
}

// IntIterator の中身を全て昇順ソートしたイテレータを返却するメソッド
func (iter intIterator) sorted() intIterator {
	// 要素数0のint型スライスを生成する
	seq := make([]int, 0)
	// イテレータから要素を取り出してスライスに追加していく
	for {
		v, err := iter()
		if err != nil {
			break
		}
		seq = append(seq, v)
	}

	// スライスを昇順ソートする
	sort.Ints(seq)

	// インデックスを管理する変数を定義する
	index := -1
	// スライスの要素を1つずつ返すイテレータを返却する
	return func() (int, error) {
		index++
		if index >= len(seq) {
			return 0, fmt.Errorf("index out of range")
		}
		return seq[index], nil
	}
}

// IntIterator の中身を全て区切り文字で連結した文字列を返却するメソッド
func (iter intIterator) join(delimiter string) string {

	// iter から要素を取り出して区切り文字で連結した文字列を返却する
	s := ""
	for {
		v, err := iter()
		if err != nil {
			break
		}
		s += fmt.Sprintf("%d%s", v, delimiter)
	}
	// 末尾の区切り文字を削除する
	s = s[:len(s)-len(delimiter)]

	return s
}

// このパッケージのメイン処理の実行に必要なパラメータを表す構造体。
// パッケージ外からは NewParameter 関数を通じて生成する
type parameter struct {
	// 生成する連番の最大値
	maxNumber int
	// 生成するグループの数
	groupCount int
	// 区切り文字
	delimiter string
}

// パラメータを生成する唯一の公開関数
// n: 生成する連番の最大値
// g: 生成するグループの数
// delimiter: 区切り文字
func NewParameter(n int, g int, delimiter string) (parameter, error) {
	// n は自然数であることを前提とする
	if n <= 0 {
		return parameter{}, fmt.Errorf("n must be positive, but %d", n)
	}
	// g は自然数であることを前提とする
	if g <= 0 {
		return parameter{}, fmt.Errorf("g must be positive, but %d", g)
	}
	// delimiter は空文字でないことを前提とする
	if delimiter == "" {
		return parameter{}, fmt.Errorf("delimiter must not be empty")
	}

	return parameter{
		n,
		g,
		delimiter,
	}, nil
}

// n個の連番をランダムに並び替えた配列をg個のグループに分けて、各グループの要素を区切り文字で分割された文字列にしたイテレータを返却する
func GenerateGroupedRandomSeqIterator(p parameter) func() (string, error) {

	// 最小グループサイズを計算する
	minGroupSize := p.maxNumber / p.groupCount
	// 最大要素数のグループ数を計算する
	maxElementsGroupCount := p.maxNumber % p.groupCount
	log.Printf("minGroupSize: %d\n", minGroupSize)
	log.Printf("maxElementsGroupCount: %d\n", maxElementsGroupCount)

	iter := generateRandomIntIterator(p.maxNumber)

	outputGroupCount := -1
	// グループの数だけ区切り文字で連結した文字列を返すイテレータを返却する
	return func() (string, error) {
		outputGroupCount++
		// イテレータのループ上限はグループ数に等しい
		if outputGroupCount >= p.groupCount {
			return "", fmt.Errorf("outputGroupCount out of range")
		}

		// 1グループあたり要素数のデフォルトは最小グループサイズ
		groupSize := minGroupSize
		// 最大要素数のグループ数に達するまでは、1グループあたり要素数を1つ増やす
		if outputGroupCount < maxElementsGroupCount {
			groupSize++
		}
		// 1グループ分の文字列要素を取得して返却する
		s := iter.take(groupSize).sorted().join(p.delimiter)
		return s, nil
	}
}
