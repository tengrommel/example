package ArrayList

import (
	"errors"
	"fmt"
)

// 接口
type List interface {
	Size() int                                  // 数组的大小
	Get(index int) (interface{}, error)         // 抓取第几个元素
	Set(index int, newVal interface{}) error    // 修改数据
	Insert(index int, newVal interface{}) error // 插入数据
	Append(newVal interface{})                  // 追加
	Clear()                                     // 清空
	Delete(index int) error                     // 删除
	String() string                             // 返回字符串
}

// 数据结构，字符串，整数，实数
type ArrayList struct {
	dataStore []interface{} // 数组的存储
	theSize   int           // 数组的大小
}

func NewArrayList() *ArrayList {
	list := new(ArrayList) // 初始化结构体
	list.dataStore = make([]interface{}, 0, 10)
	list.theSize = 0
	return list
}

func (l *ArrayList) checkIsFull() bool {
	if l.theSize == cap(l.dataStore) {
		// 判断内存的使用
		newDataSource := make([]interface{}, 0, 2*l.theSize) // 开辟双倍内存
		copy(newDataSource, l.dataStore)
		l.dataStore = newDataSource
		return true
	}
	return false
}

func (l *ArrayList) Size() int {
	return l.theSize
}

func (l *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || index > l.theSize {
		return nil, errors.New("索引越界")
	}
	return l.dataStore[index], nil
}

func (l *ArrayList) Set(index int, newVal interface{}) error {
	if index < 0 || index > l.theSize {
		return errors.New("索引越界")
	}
	l.dataStore[index] = newVal // 对数据进行设置
	return nil
}

func (l *ArrayList) Insert(index int, newVal interface{}) error {
	if index < 0 || index > l.theSize {
		return errors.New("索引越界")
	}
	l.checkIsFull()
	l.dataStore = l.dataStore[:l.theSize+1] // 插入数据时需要内存移动一位
	for i := l.theSize; i > index; i-- {    // 从后往前移动
		l.dataStore[i] = l.dataStore[i-1]
	}
	l.dataStore[index] = newVal
	l.theSize++
	return nil
}

func (l *ArrayList) Append(newVal interface{}) {
	l.dataStore = append(l.dataStore, newVal)
	l.theSize++
}

func (l *ArrayList) Clear() {
	l.dataStore = make([]interface{}, 0, 10)
	l.theSize = 0
}

func (l *ArrayList) Delete(index int) error {
	l.dataStore = append(l.dataStore[:index], l.dataStore[index+1:]...) // 重新叠加跳过index
	l.theSize--
	return nil
}

func (l *ArrayList) String() string {
	return fmt.Sprint(l.dataStore)
}
