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

func Test3Split(t *testing.T) {
	ret := Split("abcef", "bc")
	want := []string{"a", "ef"}
	if !reflect.DeepEqual(ret, want) {
		t.Fatalf("want: %v but got:%v\n", want, ret)
	}
}

func TestS(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}
	testGroup := []testCase{
		{"abc", "b", []string{"a", "c"}},
		{"abbc", "bb", []string{"a", "c"}},
	}

	for _, item := range testGroup {
		got := Split(item.str, item.sep)
		if !reflect.DeepEqual(got, item.want) {
			t.Fatalf("eant: %#v got: %#v", item.want, got)
		}
	}
}

func TestMapSplit(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("name:%s excepted:%#v, got:%#v", name, tc.want, got) // 将测试用例的name格式化输出
		}
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d:e", ":")
	}
}
