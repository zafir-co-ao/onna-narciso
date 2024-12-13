package auth


type UserCreator interface {
	Create() error
}

type creatorImpl struct {}

func NewUserCreator() UserCreator {
	return &creatorImpl{}
}

func (u *creatorImpl) Create() error {
	return nil
}
