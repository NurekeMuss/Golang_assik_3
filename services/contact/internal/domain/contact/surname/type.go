package surname

type Surname string

func (s Surname) String() string {
	return string(s)
}

func New(surname string) (*Surname, error) {
	s := Surname(surname)
	return &s, nil
}
