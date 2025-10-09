package dto                                 

type UserID string
type Url string


type CreateUser struct {
	ID UserID
	Nickname string
	Bio string
	AvatarUrl Url
}

type UpdateUser struct {
	ID UserID
	Nickname string
	Bio string
	AvatarUrl Url
}