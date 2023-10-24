package hospital_queue

type gender int

const (
	Male   gender = 0
	Female gender = 1
)

type Patient struct {
	Number string
	Gender gender
}

func (p *Patient) GetNumber() string {
	return p.Number
}

func (p *Patient) GetGender() gender {
	return p.Gender
}
