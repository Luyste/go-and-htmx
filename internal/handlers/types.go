package handlers

type Context struct {
	PageData Contacts
	FormData FormData
}

type Contacts struct {
	Contacts ContactList
}

type ContactInfo struct {
	Name  string
	Email string
}

type ContactList = []ContactInfo

type FormData struct {
	Values map[string]string
	Errors map[string]string
}
