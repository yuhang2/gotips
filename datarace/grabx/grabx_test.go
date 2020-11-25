package grabx

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_SyncGetX(t *testing.T) {
	m, restore := mockGrabX(Client)
	defer restore()
	m.On("Get", mock.Anything, mock.Anything).Return("mock data")

	tests := []struct {
		name string
	}{
		{
			name: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SyncGetX()
		})
	}
}

func Test_AsyncGetX(t *testing.T) {
	m, restore := mockGrabX(Client)
	defer restore()
	m.On("Get", mock.Anything, mock.Anything).Return("mock data")

	tests := []struct {
		name string
	}{
		{
			name: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AsyncGetX()
		})
	}
}

func mockGrabX(original GrabX) (*MockGrabX, func()) {
	m := &MockGrabX{}
	Client = m
	return m, func() {
		Client = original
	}
}
