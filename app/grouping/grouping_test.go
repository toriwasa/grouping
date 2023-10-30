package grouping

import (
	"slices"
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
		1,
		0,
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

func Test_nがgより小さい場合はパラメータ生成時にエラーが発生する(t *testing.T) {
	// Arrange
	// テストテーブルの定義
	testCases := []struct {
		n int
		g int
	}{
		{1, 2},
		{2, 3},
		{3, 4},
	}

	for _, testCase := range testCases {
		// Act
		// パラメータの生成
		_, err := NewParameter(testCase.n, testCase.g, ",")

		// Assert
		// エラーが発生することを検証する
		if err == nil {
			t.Errorf("n が g より小さい場合はエラーが発生するはずなのにエラーが発生しませんでした, n: %d, g: %d", testCase.n, testCase.g)
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

func Test_区切り文字が数値の場合はパラメータ生成時にエラーが発生する(t *testing.T) {
	// Arrange
	// テストテーブルの定義
	testCases := []struct {
		d string
	}{
		{"0"},
		{"1"},
		{"000"},
		{"009"},
		{"123"},
	}

	for _, testCase := range testCases {
		// Act
		// パラメータの生成
		_, err := NewParameter(10, 4, testCase.d)

		// Assert
		// エラーが発生することを検証する
		if err == nil {
			t.Errorf("区切り文字が数値の場合はエラーが発生するはずなのにエラーが発生しませんでした, d: %s", testCase.d)
		}
	}
}

// 検証したいふるまい: n個の連番をランダムに並び替えた配列をg個のグループに分けて、各グループの要素を区切り文字で分割された文字列にしたイテレータを返却する
func Test_n個の連番をランダムに並び替えた配列をg個のグループに分けて各グループの要素を区切り文字で分割された文字列にしたイテレータを返却する(t *testing.T) {
	// Arrange
	// 文字列を区切り文字で分割してint型スライスとして返却するヘルパー関数の定義
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

	// 実際値を表す構造体の定義
	type actualResult struct {
		// イテレータから取り出した文字列を区切り文字で分割して1つのint型スライスにまとめたもの
		intSlice []int
		// イテレータから取り出した文字列を区切り文字で分割した際の各行のint型要素数をスライスにしたもの
		rowIntCountSlice []int
		// イテレータの行数
		rowCount int
	}
	// string型を返すイテレータを元に実際値の構造体を生成するヘルパー関数の定義
	generateActual := func(iter func() (string, error), delimiter string) actualResult {
		// 実際値を表す構造体の初期化
		var actual actualResult
		for {
			s, err := iter()
			if err != nil {
				break
			}
			// イテレータの内容をint型スライスに変換する
			tmpIntSlice := splitToInts(s, delimiter)
			// int型スライスの要素数を記録しておく
			actual.rowIntCountSlice = append(actual.rowIntCountSlice, len(tmpIntSlice))
			// int型スライスを既存スライスに連結する
			actual.intSlice = append(actual.intSlice, tmpIntSlice...)
			// イテレータの行数をカウントする
			actual.rowCount++
		}
		return actual
	}

	// テストテーブルの定義
	testCases := []struct {
		n                int
		g                int
		d                string
		expectedIntSlice []int
	}{
		{10, 4, ",", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{10, 4, "1.0", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{3, 3, ":::", []int{0, 1, 2}},
		{7, 6, "@", []int{0, 1, 2, 3, 4, 5, 6}},
	}

	for _, testCase := range testCases {
		// Arrange
		n := testCase.n
		g := testCase.g
		d := testCase.d
		expectedIntSlice := testCase.expectedIntSlice

		// 入力パラメータの定義
		p, err := NewParameter(n, g, d)
		if err != nil {
			t.Fatalf("成功するはずのパラメータ生成処理でエラーが発生しました: %s", err)
		}

		// Act
		// イテレータの生成
		iter := GenerateGroupedRandomSeqIterator(p)

		// Assert
		// イテレータから実際値を取り出す
		actual := generateActual(iter, d)
		t.Logf("actual: %+v", actual)

		t.Log("スライスの要素数が n と一致することを検証する")
		if n != len(actual.intSlice) {
			t.Fatalf("expected: %d, actual: %d", n, len(actual.intSlice))
		}

		t.Log("イテレータの行数が g と一致することを検証する")
		if g != actual.rowCount {
			t.Fatalf("expected: %d, actual: %d", g, actual.rowCount)
		}

		t.Log("各グループ要素数の最大値と最小値の差は0または1であることを検証する")
		// actualIntLengthSlice の最大値および最小値取得
		max := slices.Max(actual.rowIntCountSlice)
		min := slices.Min(actual.rowIntCountSlice)
		t.Logf("max: %d, min: %d", max, min)
		// 最大値と最小値の差が0または1であることを検証する
		if max-min != 0 && max-min != 1 {
			t.Fatalf("expected: %d or %d, actual: %d", 0, 1, max-min)
		}

		t.Log("返却された文字列に含まれる数値が過不足なく期待値と一致することを検証する")
		// int型スライスを昇順ソートする
		sort.Ints(actual.intSlice)
		if len(expectedIntSlice) != len(actual.intSlice) {
			t.Fatalf("expected: %v, actual: %v", expectedIntSlice, actual.intSlice)
		}
		for i, v := range expectedIntSlice {
			if v != actual.intSlice[i] {
				t.Fatalf("expected: %v, actual: %v", expectedIntSlice, actual.intSlice)
			}
		}
	}
}
