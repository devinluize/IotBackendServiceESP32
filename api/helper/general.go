package helper

func Paniciferror(err error) {
	if err != nil {
		panic(err)
	} else {
		return
	}
}

type UserContext struct {
	UserName string
	UserId   int
}
