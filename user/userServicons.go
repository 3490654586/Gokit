package user


type IuserService interface {
	GetName(id int)string
}

type User struct {

}

func (this *User)GetName(id int)string {
	if id == 101 {
		return "zhang"
	}

	return "guest"
}