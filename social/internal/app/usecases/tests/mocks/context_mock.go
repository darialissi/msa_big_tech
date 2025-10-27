package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type MockContext struct {
	mock.Mock
}

func (m *MockContext) Deadline() (time.Time, bool) {
	args := m.Called()
	return args.Get(0).(time.Time), args.Bool(1)
}

func (m *MockContext) Done() <-chan struct{} {
	args := m.Called()
	if ch := args.Get(0); ch != nil {
		return ch.(<-chan struct{})
	}
	return nil
}

func (m *MockContext) Err() error {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (m *MockContext) Value(key interface{}) interface{} {
	args := m.Called(key)
	return args.Get(0)
}
