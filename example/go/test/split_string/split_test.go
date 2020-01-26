package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	ret := Split("babcbef", "b")
	want := []string{"", "a", "c", "ef"}
	if !reflect.DeepEqual(ret, want) {
		// 测试用例失败了
		t.Errorf("want: %v but got: %v", want, ret)
	}
}

func Test2Split(t *testing.T) {
	ret := Split("ac:ef", ":")
	want := []string{"ac", "ef"}
	if !reflect.DeepEqual(ret, want) {
		// 测试用例失败了
		t.Errorf("want: %v but got: %v", want, ret)
	}
}