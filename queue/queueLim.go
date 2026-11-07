package queue

type (
	
	elemSet [T any] []T
	
	QueueLim [T any] struct {
		set elemSet[T]
		size,
		front,
		rear int
	}

)

func (q *QueueLim[T]) Init (n int) {
	(*q) = QueueLim[T]{set: make(elemSet[T], n + 1), size: n + 1, front: 0, rear: 0}
} //Init

func NewLim [T any] (n int) *QueueLim[T] {
	q := new(QueueLim[T])
	q.Init(n)
	return q
} //NewLim

func (q QueueLim[T]) IsEmpty () bool {
	return q.rear == q.front
} //IsEmpty

func (q QueueLim[T]) IsFull () bool {
	return (q.rear + 1) % q.size == q.front
} //IsFull

func (q *QueueLim[T]) Push (val T) (ok bool) {
	r1 := (q.rear + 1) % q.size
	ok = r1 != q.front
	if ok {
		q.set[r1] = val
		q .rear = r1
	}
	return
} //Push

func (q *QueueLim[T]) Pull (val *T) (ok bool) {
	ok = q.front != q.rear
	if ok {
	q.front = (q.front + 1) % q.size
	*val = q.set[q.front]
	}
	return
} //Pull
