package auth

import "errors"

type Certificate interface {
	Auth(username string, password string) error
}

var UserData map[string]string

func init(){
	UserData = map[string]string{
		"test": "test",
		"jane": "1005",
		"andy": "0521",
	}
}

func CheckUserIsExist(user string) bool{
	_, isExist := UserData[user]
	return isExist
}

func CheckPassword(p1 string, p2 string) error{
	if p1 == p2 {
		return nil
	} else {
		return errors.New("p1!=p2")
	}
}

func Auth(username string, password string) error{
	if isExist := CheckUserIsExist(username); isExist{
		return CheckPassword(UserData[username],password)
	} else {
		return errors.New("user is not exist")
	}
}


