package model

type Requirement interface {
	IsRequirement()
}

type ChoiceRequirement struct {
	Choices *CollectionRequirement `json:"choices"`
}

func (ChoiceRequirement) IsRequirement() {}

type CollectionRequirement struct {
	Name     string        `json:"name"`
	Required int           `json:"required"`
	Options  []Requirement `json:"options"`
}

func (CollectionRequirement) IsRequirement() {}

type ConsentRequirement struct {
	Granter string `json:"granter"`
}

func (ConsentRequirement) IsRequirement() {}

type CoreRequirement struct {
	CoreFlag string `json:"core_flag"`
	Hours    int    `json:"hours"`
}

func (CoreRequirement) IsRequirement() {}

/*
type CourseRequirement struct {
	ClassReference *Course `json:"class_reference"`
	MinimumGrade   string  `json:"minimum_grade"`
}

func (CourseRequirement) IsRequirement() {}
*/

// @TODO: Remove, come up with better solution
// -------------------------
type TempCourse struct {
	SubjectPrefix string `json:"subject_prefix"`
	CourseNumber  string `json:"course_number"`
	MinimumGrade  string `json:"minimum_grade"`
}

func (TempCourse) IsRequirement() {}

// -------------------------

/*
type ExamRequirement struct {
	ExamReference Exam `json:"exam_reference"`
	MinimumScore  int  `json:"minimum_score"`
}
func (ExamRequirement) IsRequirement() {}
*/

type GPARequirement struct {
	Minimum float64 `json:"minimum"`
	Subset  string  `json:"subset"`
}

func (GPARequirement) IsRequirement() {}

// @TODO: Replace TempCourse
type HoursRequirement struct {
	Required int           `json:"required"`
	Options  []*TempCourse `json:"options"`
}

func (HoursRequirement) IsRequirement() {}

type LimitRequirement struct {
	MaxHours int `json:"max_hours"`
}

func (LimitRequirement) IsRequirement() {}

type MajorRequirement struct {
	Major string `json:"major"`
}

func (MajorRequirement) IsRequirement() {}

type MinorRequirement struct {
	Minor string `json:"minor"`
}

func (MinorRequirement) IsRequirement() {}

type OtherRequirement struct {
	Description string `json:"description"`
	Condition   string `json:"condition"`
}

func (OtherRequirement) IsRequirement() {}

// type SectionRequirement struct {
// 	SectionReference *Section `json:"section_reference"`
// }

// func (SectionRequirement) IsRequirement() {}