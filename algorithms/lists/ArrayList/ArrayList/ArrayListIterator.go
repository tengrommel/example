package ArrayList

import "errors"

type Iterator interface {
	HasNext() bool              // 是否有下一个
	Next() (interface{}, error) // 下一个
	Remove()                    // 删除
	GetIndex() int              // 得到索引
}

type Iterable interface {
	Iteractor() Iterator
}

// 构造指针访问数组
type ArrayListIterator struct {
	list         *ArrayList // 数组指针
	currentIndex int        // 当前索引
}

func (a *ArrayListIterator) HasNext() bool {
	return a.currentIndex < a.list.theSize
}

func (a *ArrayListIterator) Next() (interface{}, error) {
	if !a.HasNext() {
		return nil, errors.New("没有下一个")
	}
	value, err := a.list.Get(a.currentIndex) //抓取当前数据
	a.currentIndex++
	return value, err
}

func (a *ArrayListIterator) Remove() {
	a.currentIndex--
	a.list.Delete(a.currentIndex) // 删除一个元素
}

func (a *ArrayListIterator) GetIndex() int {
	return a.currentIndex
}

func (l *ArrayList) Iterator() Iterator {
	it := new(ArrayListIterator) // 构造
	it.currentIndex = 0
	it.list = l
	return it
}
