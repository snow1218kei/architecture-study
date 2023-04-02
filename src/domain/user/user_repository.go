package user

type UserRepository interface {
	Store(*User) error
	FindByName(string) error
}
