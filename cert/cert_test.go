package cert

import "testing"

// Cert functions testing
func TestValidCertData(t *testing.T) {
	c, err := New("golang", "john doe", "2022-12-12")
	if err != nil {
		t.Errorf("New() returned an error: %v. Cert data should be valid", err)
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("New() returned an invalid course name: %v. Cert data should be valid", c.Course)
	} else if c.Name != "John Doe" {
		t.Errorf("New() returned an invalid name: %v. Cert data should be valid", c.Name)
	}
}

func TestMissingName(t *testing.T) {
	_, err := New("", "John Doe", "2022-12-12")
	if err == nil {
		t.Errorf("New() did not return an error. Cert data should be invalid because of missing course name")
	}
}

func TestMissingCourse(t *testing.T) {
	_, err := New("Golang", "", "2022-12-12")
	if err == nil {
		t.Errorf("New() did not return an error. Cert data should be invalid because of missing name")
	}
}

func TestMissingDate(t *testing.T) {
	_, err := New("Golang", "John Doe", "")
	if err == nil {
		t.Errorf("New() did not return an error. Cert data should be invalid because of missing date")
	}
}

func TestCourseTooLong(t *testing.T) {
	course := "fje,rfejrf,jearfj,eraj,ferj,ferjruefbvhgbevubrueavcnjaecokxz,freifea"
	_, err := New(course, "John Doe", "2022-12-12")
	if err == nil {
		t.Errorf("New() did not return an error. Cert data should be invalid because of a too long course name")
	}
}

func TestNameTooLong(t *testing.T) {
	name := "fje,rfejrf,jearfj,eraj,ferj,ferjruefbvhgbevubrueavcnjaecokxz,freifea"
	_, err := New("Golang", name, "2022-12-12")
	if err == nil {
		t.Errorf("New() did not return an error. Cert data should be invalid because of a too long name")
	}
}

// Utils functions testing
func TestInvalidDateFormat(t *testing.T) {
	_, err := New("Golang", "John Doe", "2022-12-12-12")
	if err == nil {
		t.Errorf("New() did not return an error. Cert data should be invalid because of invalid date format")
	}
}

func TestIsValidString(t *testing.T) {
	if !isValidString("Golang") {
		t.Errorf("isValidString() returned false. String should be valid")
	} else if isValidString("") {
		t.Errorf("isValidString() returned true. String should be invalid")
	}
}

func TestFormatCourseName(t *testing.T) {
	course := formatCourseName("Golang")
	if course != "GOLANG COURSE" {
		t.Errorf("formatCourseName() returned an invalid course name: %v. Course name should be formated", course)
	}
}
