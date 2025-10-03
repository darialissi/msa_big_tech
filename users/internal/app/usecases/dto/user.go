package dto                                 

type UserID uint64
type Url string


type CreateUser struct {
	ID UserID
	Nickname string
	Bio string
	Avatar Url
}

type UpdateUser struct {
	ID UserID
	Nickname string
	Bio string
	Avatar Url
}