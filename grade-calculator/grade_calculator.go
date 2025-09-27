package esepunittests

// Note: I rearranged the code layout/order to better understand it because it seemed quite scattered before.

type GradeType int

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

// change to have single categories instead of 3 categories
type GradeCalculator struct {
	all_grades []Grade
}

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

// single list for all types of grades as Grade stores type already
func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		all_grades: make([]Grade, 0),
	}
}

// Step 1. Add Grade to GradeCalculator
func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) {
	gc.all_grades = append(gc.all_grades, Grade{
		Name:  name,
		Grade: grade,
		Type:  gradeType,
	})
}

// Step 2. Compute Average of the Grade by its category
func computeAverage(grades []Grade) (int, int, int) {
	assignment_sum := 0
	assignment_num := 0
	assignment_average := 0

	exam_sum := 0
	exam_num := 0
	exam_average := 0

	essay_sum := 0
	essay_num := 0
	essay_average := 0

	for _, grade := range grades {
		switch grade.Type {
		case Assignment:
			assignment_sum += grade.Grade
			assignment_num += 1
		case Exam:
			exam_sum += grade.Grade
			exam_num += 1
		case Essay:
			essay_sum += grade.Grade
			essay_num += 1
		}
	}

	if assignment_num > 0 {
		assignment_average = assignment_sum / assignment_num
	}

	if exam_num > 0 {
		exam_average = exam_sum / exam_num
	}

	if essay_num > 0 {
		essay_average = essay_sum / essay_num
	}

	return assignment_average, exam_average, essay_average
}

// Step 3. Use the averages to find the numerical grade based on weights
func (gc *GradeCalculator) calculateNumericalGrade() int {
	assignment_average, exam_average, essay_average := computeAverage(gc.all_grades)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

// Step 4. Get the final grade
func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}

	return "F"
}
