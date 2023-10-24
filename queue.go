package main

// Enum for queue dispatch mode
const (
	Default    int = 0
	RoundRobin     = 1
)

type Queue interface {
	Enqueue(item Patient) error
	Dequeue() (Patient, error)
	ChangeMode(int)
	Len() int
	GetMode() string
}

func NewPatientQueue() Queue {
	return &queue{
		items:   make([]Patient, 0),
		hashmap: make(map[string]bool),
	}
}

type queue struct {
	items         []Patient
	mode          int
	hashmap       map[string]bool
	currentGender *Gender
}

func (q *queue) ChangeMode(mode int) {
	q.mode = mode
	if mode == Default {
		q.currentGender = nil
	}
}

func (q *queue) GetMode() string {
	if q.mode == Default {
		return "Default Mode"
	}

	if q.mode == RoundRobin {
		return "Round Robin Mode"
	}

	return "Mode not available"

}

func (q *queue) Len() int {
	return len(q.items)
}

func (q *queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *queue) Enqueue(item Patient) error {
	if q.hashmap[item.Number] {
		return ErrItemInQueue
	}
	q.items = append(q.items, item)
	q.hashmap[item.Number] = true

	return nil
}

func (q *queue) Dequeue() (item Patient, err error) {

	switch q.mode {
	case Default:
		item, err = q.dequeueDefault()
	case RoundRobin:
		item, err = q.dequeueRoundRobin()
	default:
		err = ErrModeNotImplemented
	}

	return
}

func (q *queue) dequeueDefault() (item Patient, err error) {
	if len(q.items) == 0 {
		err = ErrEmptyQueue
		return
	}
	item = q.items[0]
	q.items = q.items[1:]
	delete(q.hashmap, item.Number)
	return item, nil
}

func (q *queue) dequeueRoundRobin() (item Patient, err error) {
	idx := 0

	if len(q.items) == 0 {
		err = ErrEmptyQueue
		return
	}

	currentItem := q.items[idx]

	if q.currentGender == nil || *q.currentGender != currentItem.Gender {
		item = currentItem
		q.currentGender = &currentItem.Gender
		q.items = q.items[1:]
		delete(q.hashmap, currentItem.Number)
		return
	}

	idx++
	found := false
	for !found && idx < len(q.items) {
		currentItem = q.items[idx]
		if *q.currentGender != currentItem.Gender {
			found = true
			item = currentItem
			q.currentGender = &currentItem.Gender
			q.items = remove(q.items, idx)
			delete(q.hashmap, currentItem.Number)
			return
		}

		idx++
	}

	return item, ErrRoundRobinFailed
}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
