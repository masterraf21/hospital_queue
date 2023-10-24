package main

type gender int

const (
	Male   gender = 0
	Female gender = 1
)

type Patient struct {
	Number string
	Gender gender
}

func (p *Patient) GetGender() string {
	var res string
	switch p.Gender {
	case Male:
		res = "M"
	case Female:
		res = "F"
	}

	return res
}
