package parser

type ParsedData struct {
	Name   string
	Email  string
	Phone  string
	Skills []string
}

func Extract(path string) (string, ParsedData) {
	raw := "Fake resume content..."

	parsed := ParsedData{
		Name:   "John Doe",
		Email:  "john@example.com",
		Phone:  "123-456-7890",
		Skills: []string{"Go", "Docker", "MongoDB"},
	}

	return raw, parsed
}
