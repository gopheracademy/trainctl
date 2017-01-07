package models

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Lesson struct {
	ShortName     string `json:"short_name"`
	Description   string `json:"description"`
	Topic         string `json:"topic"`
	Author        string `json:"author"`
	AuthorEmail   string `json:"author_email"`
	AuthorTwitter string `json:"author_twitter"`
}

func NewLesson(cmd *cobra.Command, name, description string, importPath string) Lesson {
	return Lesson{
		ShortName:     name,
		Topic:         cmd.Flag("topic").Value.String(),
		Description:   cmd.Flag("description").Value.String(),
		Author:        viper.GetString("author"),
		AuthorEmail:   viper.GetString("email"),
		AuthorTwitter: viper.GetString("twitter"),
	}

}

func (m Lesson) NumberedPath(i int) string {
	i = i + 1
	s := strconv.Itoa(i)
	if len(s) < 2 {
		return "0" + s + "-" + m.ShortName
	}
	return s + "-" + m.ShortName
}

func (m Lesson) String() string {
	b := make([]byte, 0, 40)
	b = append(b, []byte("Short Name: ")...)
	b = append(b, m.ShortName...)
	b = append(b, '\n')
	b = append(b, []byte("Topic: ")...)
	b = append(b, m.Topic...)
	b = append(b, '\n')
	b = append(b, []byte("Description: ")...)
	b = append(b, m.Description...)
	b = append(b, '\n')
	b = append(b, []byte("Author: ")...)
	b = append(b, m.Author...)
	b = append(b, '\n')
	b = append(b, []byte("Author Email: ")...)
	b = append(b, m.AuthorEmail...)
	b = append(b, '\n')
	b = append(b, []byte("Author Twitter: ")...)
	b = append(b, m.AuthorTwitter...)
	b = append(b, '\n')
	return string(b)

}
