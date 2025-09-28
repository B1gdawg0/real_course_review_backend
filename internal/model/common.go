package model

type Role string

const (
	UserRole  Role = "USER"
	AdminRole Role = "ADMIN"
)

var ALL_SCHEMA = []interface{}{
	&User{},
	&Review{},
	// &Rating{},
	&Course{},
	&Professor{},
	&Vote{},
	// &Thread{},
	// &Comment{},
	&SystemConfig{},
	&Tag{},
}

type Category int

const (
	Happiness Category = iota
	Easiness
	Quality
)

func (c Category) String() string {
	return [...]string{"happiness", "easiness", "quality"}[c]
}