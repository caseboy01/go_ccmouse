package queue

//An FIFO queue.
type Queue []interface{}

//Pushes the element info the queue.
//		e.g. q.Push(123)
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

//Pops the element from head.
func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//Returns if the queue is empty or not.
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
