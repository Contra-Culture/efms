package efms

import "time"

type Account struct {
	ID string
}
type Family struct {
	Name string
	Account
}
type Person struct {
	Account
}
type Task struct {
	Title       string
	Description string
	Deadline    time.Time
}
type Project struct {
	Name  string
	Tasks []Task
}
type Artefact struct {
}
type Income struct {
}
type Expense struct {
}
type Ritual struct {
}
type Value struct {
}
type Asset struct {
}
