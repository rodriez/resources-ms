package repositories

type AuthRepository struct{}

func (repo *AuthRepository) ValidateUser(name, pass string) bool {
	for _, user := range authorizedUsers {
		if name == user.name && pass == user.pass {
			return true
		}
	}

	return false
}

type user struct {
	name, pass string
}

var authorizedUsers []user

func init() {
	authorizedUsers = []user{
		{name: "Jonh", pass: "johnW.67"},
		{name: "jason", pass: "jasonB.85"},
	}
}
