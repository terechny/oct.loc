package user

type User struct {
	id         uint32
	firstname  string
	secondname string
	email      string
	phone      string
	password   string
}

func (u *User) SetId(id uint32) {
	u.id = id
}

func (u *User) SetFirstname(firstname string) {
	u.firstname = firstname
}

func (u *User) SetSecondname(secondname string) {
	u.secondname = secondname
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPhone(phone string) {
	u.phone = phone
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) Id() uint32 {
	return u.id
}

func (u *User) Firstname() string {
	return u.firstname
}

func (u *User) Secondname() string {
	return u.secondname
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Phone() string {
	return u.phone
}

func (u *User) Password() string {
	return u.password
}
