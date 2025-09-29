package dto                                 


type UserID uint64
type Url string

type CreateUser struct {
	Email string // middleware auth ctx
	Nickname string
	Bio string
	Avatar Url
}

type SaveUser struct {
	Email string
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