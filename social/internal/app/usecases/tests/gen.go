package tests

// go install github.com/vektra/mockery/v2@latest

//go:generate mockery --disable-version-string --with-expecter --name FriendRequestRepository --filename friend_request_repository_mock.go --output tests/mocks
//go:generate mockery --disable-version-string --with-expecter --name FriendRepository --filename friend_repository_mock.go --output tests/mocks
//go:generate mockery --disable-version-string --with-expecter --name OutboxRepository --filename outbox_repository_mock.go --output tests/mocks
//go:generate mockery --disable-version-string --with-expecter --name TxManager --filename tx_manager_mock.go --output tests/mocks

// go install github.com/gojuno/minimock/v3/cmd/minimock@latest

//go:generate minimock -i FriendRequestRepository -n FriendRequestRepositoryMock -o ./tests/mocks/friend_request_minimock_mock.go
//go:generate minimock -i FriendRepository -n FriendRepositoryMock -o ./tests/mocks/FriendRepository_minimock_mock.go
