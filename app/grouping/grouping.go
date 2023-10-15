package grouping

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
)

// 0から始まるn個の連番をランダムに返すイテレータを返却する
func generateRandomIntIterator(n int) func() (int, error) {
	// n個の要素を持つスライスを生成する
	seq := make([]int, n)
	log.Printf("seq: %v\n", seq)

	// 0, 1, 2, .. n の連番を要素として持つようにスライスを初期化する
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

// イテレータからn個の値を取り出して昇順ソートしてから指定された区切り文字で分割された文字列にして返却する
func takeNSortedSeparated(iter func() (int, error), n int, delimiter string) (string, error) {
	// イテレータからn個の値を取り出してseqスライスに格納する
	seq := make([]int, n)
	log.Printf("seq: %v\n", seq)

	for i := 0; i < n; i++ {
		v, err := iter()
		if err != nil {
			return "", err
		}
		seq[i] = v
	}
	log.Printf("seq: %v\n", seq)

	// seqスライスを昇順ソートする
	sort.Slice(seq, func(i, j int) bool {
		return seq[i] < seq[j]
	})
	log.Printf("seq: %v\n", seq)

	// seqスライスの要素を区切り文字で連結した文字列を返却する
	s := ""
	for i := 0; i < n; i++ {
		s += fmt.Sprintf("%d%s", seq[i], delimiter)
	}
	// 末尾の区切り文字を削除する
	s = s[:len(s)-len(delimiter)]

	return s, nil
}

// n個の連番をランダムに並び替えた配列をg個のグループに分けて、各グループの要素を区切り文字で分割された文字列にしたイテレータを返却する
func GenerateGroupedRandomSeqIterator(n int, g int, delimiter string) (func() (string, error), error) {
	// n は自然数であることを前提とする
	if n <= 0 {
		return nil, fmt.Errorf("n must be positive, but %d", n)
	}
	// g は自然数であることを前提とする
	if g <= 0 {
		return nil, fmt.Errorf("g must be positive, but %d", g)
	}

	// 最小グループサイズを計算する
	minGroupSize := n / g
	// 最大要素数のグループ数を計算する
	maxElementsGroupCount := n % g
	log.Printf("minGroupSize: %d\n", minGroupSize)
	log.Printf("maxElementsGroupCount: %d\n", maxElementsGroupCount)

	iter := generateRandomIntIterator(n)

	outputGroupCount := -1
	// グループの数だけ区切り文字で連結した文字列を返すイテレータを返却する
	return func() (string, error) {
		outputGroupCount++
		// イテレータのループ上限はグループ数に等しい
		if outputGroupCount >= g {
			return "", fmt.Errorf("outputGroupCount out of range")
		}

		// 1グループあたり要素数のデフォルトは最小グループサイズ
		groupSize := minGroupSize
		// 最大要素数のグループ数に達するまでは、1グループあたり要素数を1つ増やす
		if outputGroupCount < maxElementsGroupCount {
			groupSize++
		}
		// 1グループ分の文字列要素を取得して返却する
		s, err := takeNSortedSeparated(iter, groupSize, delimiter)
		if err != nil {
			return "", err
		}
		return s, nil
	}, nil
}
