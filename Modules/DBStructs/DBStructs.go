package dbstructs

type User struct {
	Username string
	Email    string
	Pass     string
}

func NewUser(username, email, pass string) User {
	return User{Username: username, Email: email, Pass: pass}
}

type Task struct {
	User_name string
	Name      string
	Context   string
	Type      string
	Status    bool
}

func NewTaskItem(name string, cont string, typ string, username string, status bool) Task {
	return Task{User_name: username, Name: name, Context: cont, Type: typ, Status: status}
}
