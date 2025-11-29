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
	&Report{},
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

type ReportType int

const (
	Inappropriate ReportType = iota + 1
	Misleading
	Spam
	Sensitive
)

type ReportContext string

const (
	ReportInappropriate ReportContext = "Inappropriate or Offensive Content"
	ReportMisleading    ReportContext = "False or Misleading Information"
	ReportSpam          ReportContext = "Spam or Promotional Content"
	ReportSensitive     ReportContext = "Privacy or Sensitive Information"
)

var ReportTypeDescriptions = map[ReportType]ReportContext{
	Inappropriate: ReportInappropriate,
	Misleading:    ReportMisleading,
	Spam:          ReportSpam,
	Sensitive:     ReportSensitive,
}

func (rt ReportType) String() string {
	if desc, ok := ReportTypeDescriptions[rt]; ok {
		return string(desc)
	}
	return "Unknown Report Type"
}