package customModel

import (
	"encoding/json"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"io"
	"strconv"
)

// Coordinate 型は -10 から 10 の範囲内の整数を保証します
type Coordinate int

// UnmarshalCoordinate は、入力が範囲内かどうかをチェックする関数
func UnmarshalCoordinate(v interface{}) (Coordinate, error) {
	switch v := v.(type) {
	case int:
		if v < -10 || v > 10 {
			return 0, errors.New("value out of range (-10 to 10)")
		}
		return Coordinate(v), nil
	case float64: // JSONから直接数値が送られてくる場合の対応
		iv := int(v)
		if iv < -10 || iv > 10 {
			return 0, errors.New("value out of range (-10 to 10)")
		}
		return Coordinate(iv), nil
	case string: // 文字列からパースする場合の対応
		iv, err := strconv.Atoi(v)
		if err != nil || iv < -10 || iv > 10 {
			return 0, errors.New("value out of range (-10 to 10)")
		}
		return Coordinate(iv), nil
	case json.Number: // JSONから直接json.Numberが渡された場合の対応
		iv, err := v.Int64()
		if err != nil || iv < -10 || iv > 10 {
			return 0, errors.New("value out of range (-10 to 10)")
		}
		return Coordinate(int(iv)), nil
	default:
		return 0, errors.New("invalid type for Coordinate")
	}
}

// MarshalCoordinate は、Coordinate 型を出力に変換する関数
func MarshalCoordinate(v Coordinate) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Itoa(int(v)))
	})
}
