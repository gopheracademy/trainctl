package templates

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Level string

const (
	Beginner     = "beginner"
	Intermediate = "intermediate"
	Advanced     = "advanced"
	Expert       = "expert"
)

type Topic struct {
	ShortName        string `json:"short_name"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Subject          string `json:"subject"`
	Level            Level  `json:"level"`
	Author           string `json:"author"`
	AuthorEmail      string `json:"author_email"`
	AuthorTwitter    string `json:"author_twitter"`
	SourceRepository string `json:"source_repository"`
}

func NewTopic(cmd *cobra.Command, description string, importPath string) Topic {
	return Topic{
		ShortName:        cmd.Flag("shortname").Value.String(),
		Name:             cmd.Flag("name").Value.String(),
		Subject:          cmd.Flag("subject").Value.String(),
		Level:            Level(cmd.Flag("level").Value.String()),
		Description:      cmd.Flag("description").Value.String(),
		Author:           viper.GetString("author"),
		AuthorEmail:      viper.GetString("email"),
		AuthorTwitter:    viper.GetString("twitter"),
		SourceRepository: importPath,
	}

}

func (m Topic) NumberedPath(i int) string {
	s := strconv.Itoa(i)
	if len(s) < 2 {
		return "0" + s + "-" + m.ShortName
	}
	return s + "-" + m.ShortName
}

func (m Topic) String() string {
	b := make([]byte, 0, 40)
	b = append(b, []byte("Short Name: ")...)
	b = append(b, m.ShortName...)
	b = append(b, '\n')
	b = append(b, []byte("Name: ")...)
	b = append(b, m.Name...)
	b = append(b, '\n')
	b = append(b, []byte("Subject: ")...)
	b = append(b, m.Subject...)
	b = append(b, '\n')
	b = append(b, []byte("Level: ")...)
	b = append(b, m.Level...)
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
