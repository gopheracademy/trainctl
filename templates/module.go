package templates

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Module struct {
	ShortName        string `json:"short_name"`
	Description      string `json:"description"`
	Author           string `json:"author"`
	AuthorEmail      string `json:"author_email"`
	AuthorTwitter    string `json:"author_twitter"`
	SourceRepository string `json:"source_repository"`
}

func NewModule(cmd *cobra.Command, description string, importPath string) Module {
	return Module{
		ShortName:        cmd.Flag("name").Value.String(),
		Description:      cmd.Flag("description").Value.String(),
		Author:           viper.GetString("author"),
		AuthorEmail:      viper.GetString("email"),
		AuthorTwitter:    viper.GetString("twitter"),
		SourceRepository: importPath,
	}

}

func (m Module) String() string {
	b := make([]byte, 0, 40)
	b = append(b, []byte("Short Name: ")...)
	b = append(b, m.ShortName...)
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
