package cert

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type Cert struct {
	Course, Name string
	Date         time.Time

	LabelTitle, LabelCompletion, LabelPresented, LabelParticipation, LabelDate string
}

type Saver interface {
	Save(c Cert) error
}

// Maximum characters authorized for a course name
const COURSE_MAX_LEN int = 20
const NAME_MAX_LEN int = 20

func Save(c Cert) error {
	return nil
}

// New creates a new instance of Cert.
// It returns a pointer to a Cert and an error if the creation failed.
func New(course, name, date string) (*Cert, error) {

	// Check values
	if !isValidString(course) {
		return nil, errors.New("invalid course name (empty course is not valid)")
	} else if !isValidString(name) {
		return nil, errors.New("invalid name (empty name is not valid")
	} else if !isValidString(date) {
		return nil, errors.New("invalid date (empty date is not valid")
	} else if len(course) > COURSE_MAX_LEN {
		return nil, fmt.Errorf("invalid course name (too long course name is not valid (%d characters max))", COURSE_MAX_LEN)
	} else if len(name) > NAME_MAX_LEN {
		return nil, fmt.Errorf("invalid name (too long name is not valid (%d characters max))", NAME_MAX_LEN)
	}

	formatedDate, err := time.Parse("2006-01-01", date)
	if err != nil {
		fmt.Printf("%s", err.Error())
		return nil, errors.New("invalid date format, please use yyyy-mm-dd format")
	}

	formatedCourse := formatCourseName(course)
	formatedName := strings.Title(name)

	cert := &Cert{
		Course:             formatedCourse,
		Name:               formatedName,
		Date:               formatedDate,
		LabelTitle:         fmt.Sprintf("%s certificate - %s", formatedCourse, formatedName),
		LabelCompletion:    "Certificate of completion",
		LabelPresented:     "This certificate is presented to :",
		LabelParticipation: fmt.Sprintf("For his/her participation in the %s", formatedCourse),
		LabelDate:          fmt.Sprintf("Date : %s", formatedDate),
	}

	return cert, nil
}

func formatCourseName(course string) string {
	if !strings.HasSuffix(course, " course") {
		course = course + " course"
	}

	return strings.ToTitle(course)
}

func isValidString(str string) bool {
	return len(strings.TrimSpace(str)) > 0
}
