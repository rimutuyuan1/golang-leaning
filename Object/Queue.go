package Queue

type Queue []interface{}

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	//如果限制类型，需要强转
	return head.(int)
	//return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
