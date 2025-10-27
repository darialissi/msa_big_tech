package tests

// go install github.com/vektra/mockery/v2@latest

//go:generate mockery --disable-version-string --with-expecter --name FriendRequestRepository --filename friend_request_repository_mock.go
//go:generate mockery --disable-version-string --with-expecter --name FriendRepository --filename friend_repository_mock.go
//go:generate mockery --disable-version-string --with-expecter --name TxManager --filename tx_manager_mock.go

// go install github.com/gojuno/minimock/v3/cmd/minimock@latest

//go:generate minimock -i FriendRequestRepository -n FriendRequestRepositoryMock -o ./mocks/friend_request_minimock_mock.go
//go:generate minimock -i FriendRepository -n FriendRepositoryMock -o ./mocks/FriendRepository_minimock_mock.go
