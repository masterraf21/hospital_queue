package main

type Gender int

const (
	Male   Gender = 0
	Female Gender = 1
)

type Patient struct {
	Number string
	Gender Gender
}

func NewPatient(number string, gender string) Patient {
	var g Gender
	switch gender {
	case "M":
		g = Male
	case "F":
		g = Female
	}

	return Patient{
		Number: number,
		Gender: g,
	}
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
