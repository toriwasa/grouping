package grouping

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

func Test_パラメータの生成ができる(t *testing.T) {
	// Arrange
	// 期待値パラメータの定義
	expected := parameter{
		1,
		1,
		",",
	}

	// Act
	// パラメータの生成
	actual, err := NewParameter(1, 1, ",")

	// Assert
	// エラーが発生しないことを検証する
	if err != nil {
		t.Fatalf("正常にパラメータが生成されるはずなのにエラーが発生しました: %s", err)
	}
	// 期待値と実際値が一致することを検証する
	if expected != actual {
		t.Errorf("expected: %+v, actual: %+v", expected, actual)
	}
}

func Test_nが0以下の場合はパラメータ生成時にエラーが発生する(t *testing.T) {
	// Arrange
	// テストテーブルの定義
	testCases := []struct {
		n int
	}{
		{-1},
		{0},
	}

	for _, testCase := range testCases {

		// Act
		// パラメータの生成
		_, err := NewParameter(testCase.n, 4, ",")

		// Assert
		// エラーが発生することを検証する
		if err == nil {
			t.Errorf("n が 0 以下の場合はエラーが発生するはずなのにエラーが発生しませんでした, n: %d", testCase.n)
		}
	}
}

func Test_gが0以下の場合はパラメータ生成時にエラーが発生する(t *testing.T) {
	// Arrange
	// テストテーブルの定義
	testCases := []struct {
		g int
	}{
		{-1},
		{0},
	}

	for _, testCase := range testCases {
		// Act
		// パラメータの生成
		_, err := NewParameter(10, testCase.g, ",")

		// Assert
		// エラーが発生することを検証する
		if err == nil {
			t.Errorf("g が 0 以下の場合はエラーが発生するはずなのにエラーが発生しませんでした, g: %d", testCase.g)
		}
	}
}

func Test_区切り文字が空文字の場合はパラメータ生成時にエラーが発生する(t *testing.T) {
	// Arrange
	// 期待値エラーの定義
	expected := "delimiter must not be empty"

	// Act
	// パラメータの生成
	_, err := NewParameter(10, 4, "")

	// Assert
	// エラーが発生することを検証する
	if err == nil {
		t.Fatalf("区切り文字が空文字の場合はエラーが発生するはずなのにエラーが発生しませんでした")
	}
	// エラー内容が期待値と一致することを検証する
	if expected != err.Error() {
		t.Errorf("expected: %s, actual: %s", expected, err.Error())
	}
}

// 検証したいふるまい: n個の連番をランダムに並び替えた配列をg個のグループに分けて、各グループの要素を区切り文字で分割された文字列にしたイテレータを返却する
func Test_n個の連番をランダムに並び替えた配列をg個のグループに分けて各グループの要素を区切り文字で分割された文字列にしたイテレータを返却する(t *testing.T) {
	// Arrange
	// 文字列を区切り文字で分割してint型スライスに追加するヘルパー関数の定義
	splitToInts := func(s string, d string) []int {
		var ints []int
		for _, v := range strings.Split(s, d) {
			i, err := strconv.Atoi(v)
			if err != nil {
				t.Fatalf("成功するはずの文字列をint型に変換する処理でエラーが発生しました: %s", err)
			}
			ints = append(ints, i)
		}
		return ints
	}

	// 入力パラメータの定義
	n := 10
	g := 4
	d := ","
	p, err := NewParameter(n, g, d)
	if err != nil {
		t.Fatalf("成功するはずのパラメータ生成処理でエラーが発生しました: %s", err)
	}

	// 期待値の定義
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Act
	// イテレータの生成
	iter := GenerateGroupedRandomSeqIterator(p)

	// Assert
	var actualInts []int
	actualLines := 0
	t.Log("イテレータの内容を区切り文字で分割してint型スライスに追加する")
	for {
		s, err := iter()
		if err != nil {
			break
		}
		// イテレータの内容があるなら行数をカウントする
		actualLines++
		// ... は可変長引数を表す
		actualInts = append(actualInts, splitToInts(s, d)...)
	}

	t.Log("スライスの要素数が n と一致することを検証する")
	if n != len(actualInts) {
		t.Fatalf("expected: %d, actual: %d", n, len(actualInts))
	}

	t.Log("イテレータの行数が g と一致することを検証する")
	if g != actualLines {
		t.Fatalf("expected: %d, actual: %d", g, actualLines)
	}

	t.Log("返却された文字列に含まれる数値が過不足なく期待値と一致することを検証する")
	// int型スライスを昇順ソートする
	sort.Ints(actualInts)
	if len(expected) != len(actualInts) {
		t.Fatalf("expected: %v, actual: %v", expected, actualInts)
	}
	for i, v := range expected {
		if v != actualInts[i] {
			t.Fatalf("expected: %v, actual: %v", expected, actualInts)
		}
	}
}
