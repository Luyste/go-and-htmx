package app

type Item struct {
	Id   string
	Item string
}

type Validation struct {
	NameError  string
	NameInput  string
	EmailError string
	EmailInput string
	Success    bool
}

type Context struct {
	List       []Item
	Validation Validation
}
