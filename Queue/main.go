package main

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() int {
	front := q.items[0]
	q.items = q.items[1:]
	return front
}

func (q *Queue) Front() int {
	return q.items[0]
}

func main() {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Front()
}
