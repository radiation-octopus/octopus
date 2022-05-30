package utils

import "container/heap"

// The size of a block of data
const blockSize = 4096

// 优先级队列数据结构。
type Prque struct {
	cont *sstack
}

// 新建创建新的优先级队列。
func New(setIndex SetIndexCallback) *Prque {
	return &Prque{newSstack(setIndex, false)}
}

// newwrapparound使用环绕优先级处理创建新的优先级队列。
func NewWrapAround(setIndex SetIndexCallback) *Prque {
	return &Prque{newSstack(setIndex, true)}
}

// 将具有给定优先级的值推送到队列中，必要时进行扩展。
func (p *Prque) Push(data interface{}, priority int64) {
	heap.Push(p.cont, &item{data, priority})
}

// Peek返回具有greates优先级的值，但不会将其弹出。
func (p *Prque) Peek() (interface{}, int64) {
	item := p.cont.blocks[0][0]
	return item.value, item.priority
}

// 从堆栈中弹出具有greates优先级的值并返回它。目前未进行收缩。
func (p *Prque) Pop() (interface{}, int64) {
	item := heap.Pop(p.cont).(*item)
	return item.value, item.priority
}

// 仅从队列中弹出项目，删除关联的优先级值。
func (p *Prque) PopItem() interface{} {
	return heap.Pop(p.cont).(*item).value
}

// Remove删除具有给定索引的元素。
func (p *Prque) Remove(i int) interface{} {
	if i < 0 {
		return nil
	}
	return heap.Remove(p.cont, i)
}

// 检查优先级队列是否为空。
func (p *Prque) Empty() bool {
	return p.cont.Len() == 0
}

// 返回优先级队列中的元素数。
func (p *Prque) Size() int {
	return p.cont.Len()
}

// 清除优先级队列的内容。
func (p *Prque) Reset() {
	*p = *New(p.cont.setIndex)
}

/**
当元素移动到新索引时，将调用SetIndexCallback。提供SetIndexCallback是可选的，只有当应用程序需要删除除顶部元素以外的其他元素时才需要它。
*/
type SetIndexCallback func(data interface{}, index int)

/**
内部可排序堆栈数据结构。实现堆栈（heap）功能的Push和Pop操作，以及堆可排序性要求的Len、Less和Swap方法。
*/
type sstack struct {
	setIndex   SetIndexCallback
	size       int
	capacity   int
	offset     int
	wrapAround bool

	blocks [][]*item
	active []*item
}

// 创建新的空堆栈。
func newSstack(setIndex SetIndexCallback, wrapAround bool) *sstack {
	result := new(sstack)
	result.setIndex = setIndex
	result.active = make([]*item, blockSize)
	result.blocks = [][]*item{result.active}
	result.capacity = blockSize
	result.wrapAround = wrapAround
	return result
}

//将值推送到堆栈上，必要时进行扩展。堆所需。界面
func (s *sstack) Push(data interface{}) {
	if s.size == s.capacity {
		s.active = make([]*item, blockSize)
		s.blocks = append(s.blocks, s.active)
		s.capacity += blockSize
		s.offset = 0
	} else if s.offset == blockSize {
		s.active = s.blocks[s.size/blockSize]
		s.offset = 0
	}
	if s.setIndex != nil {
		s.setIndex(data.(*item).value, s.size)
	}
	s.active[s.offset] = data.(*item)
	s.offset++
	s.size++
}

// 从堆栈中弹出一个值并返回它。目前未进行收缩。堆所需。界面
func (s *sstack) Pop() (res interface{}) {
	s.size--
	s.offset--
	if s.offset < 0 {
		s.offset = blockSize - 1
		s.active = s.blocks[s.size/blockSize]
	}
	res, s.active[s.offset] = s.active[s.offset], nil
	if s.setIndex != nil {
		s.setIndex(res.(*item).value, -1)
	}
	return
}

// 返回堆栈的长度。排序时需要。界面
func (s *sstack) Len() int {
	return s.size
}

// 比较堆栈中两个元素的优先级（优先级越高）。排序时需要。界面
func (s *sstack) Less(i, j int) bool {
	a, b := s.blocks[i/blockSize][i%blockSize].priority, s.blocks[j/blockSize][j%blockSize].priority
	if s.wrapAround {
		return a-b > 0
	}
	return a > b
}

// 交换堆栈中的两个元素。排序时需要。界面
func (s *sstack) Swap(i, j int) {
	ib, io, jb, jo := i/blockSize, i%blockSize, j/blockSize, j%blockSize
	a, b := s.blocks[jb][jo], s.blocks[ib][io]
	if s.setIndex != nil {
		s.setIndex(a.value, i)
		s.setIndex(b.value, j)
	}
	s.blocks[ib][io], s.blocks[jb][jo] = a, b
}

// 重置堆栈，有效地清除其内容。
func (s *sstack) Reset() {
	*s = *newSstack(s.setIndex, false)
}

/**
排序堆栈中的优先项目。

注意：优先级可以“环绕”int64范围，如果（a.priority-b.priority）>0，则a位于b之前。

队列中任何点的最低优先级和最高优先级之间的差值应小于2^63。
*/
type item struct {
	value    interface{}
	priority int64
}
