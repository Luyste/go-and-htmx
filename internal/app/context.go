package app

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

type Contact struct {
	Name  string
	Email string
}

type DisplayData struct {
	Contacts []Contact
}

type Context struct {
	FormData    FormData
	DisplayData DisplayData
	Counter     int
}

// helper functions to seed objects

func NewFormData() FormData {
	return FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

func NewContact(name, email string) Contact {
	return Contact{
		Name:  name,
		Email: email,
	}
}

func (ctx Context) HasEmail(email string) bool {
	for _, contact := range ctx.DisplayData.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}
