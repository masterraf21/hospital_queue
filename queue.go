package hospital_queue

import "errors"

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
}

func NewPatientQueue() Queue {
	return &queue{}
}

type queue struct {
	items         []Patient
	mode          int
	hashmap       map[string]bool
	currentGender *gender
}

func (q *queue) ChangeMode(mode int) {
	q.mode = mode
}

func (q *queue) Len() int {
	return len(q.items)
}

func (q *queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *queue) Enqueue(item Patient) error {
	if q.hashmap[item.Number] {
		return errors.New("item already in queue")
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
		err = errors.New("mode not implemented")
	}

	return
}

func (q *queue) dequeueDefault() (Patient, error) {
	if len(q.items) == 0 {
		var zeroValue Patient
		return zeroValue, errors.New("empty queue")
	}
	item := q.items[0]
	q.items = q.items[1:]
	delete(q.hashmap, item.Number)
	return item, nil
}

func (q *queue) dequeueRoundRobin() (item Patient, err error) {
	if len(q.items) == 0 {
		err = errors.New("empty queue")
		return
	}

	currentItem := q.items[0]

	if q.currentGender == nil || q.currentGender != &currentItem.Gender {
		item = currentItem
		q.currentGender = &currentItem.Gender
		q.items = q.items[1:]
		delete(q.hashmap, currentItem.Number)
		return
	}

	idx := 1
	found := false
	for !found || idx < len(q.items) {
		tempItem := q.items[idx]
		if q.currentGender != &tempItem.Gender {
			found = true
			item = tempItem
			q.currentGender = &tempItem.Gender
			q.items = remove(q.items, idx)
			delete(q.hashmap, tempItem.Number)
			return

		}

		idx++
	}

	return item, errors.New("Round Robin failed getting alternate gender")
}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}
