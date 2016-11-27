package templates

type Conf struct {
	Twitter   string
	Author    string
	Email     string
	TopicDir  string
	CourseDir string
}

const Config = `twitter: {{.Twitter}}
author: {{.Author}}
email: {{.Email}}
topicdir: {{.TopicDir}}
coursedir: {{.CourseDir}}
`
