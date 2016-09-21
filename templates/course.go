package templates

import "github.com/spf13/cobra"

type Course struct {
	Name              string   `json:"name"`
	Instructor        string   `json:"instructor"`
	InstructorEmail   string   `json:"instructor_email"`
	InstructorTwitter string   `json:"instructor_twitter"`
	Modules           []Module `json:"modules"`
}

func NewCourse(cmd *cobra.Command) Course {

	return Course{}
}
