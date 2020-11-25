package grabx

import (
	"context"
	"fmt"
)

var (
	Client GrabX
)

// GrabX use to get varibale for staging & production
type GrabX interface {
	Get(ctx context.Context, name string) string
}

type grabxImpl struct{}

func (x *grabxImpl) Get(ctx context.Context, name string) string {
	return "some real value"
}

func Initialize() {
	Client = &grabxImpl{}
}

func SyncGetX() {
	fmt.Println(Client.Get(context.Background(), "hello"))
}

func AsyncGetX() {
	go func() {
		fmt.Println(Client.Get(context.Background(), "hello"))
	}()
}

//go:generate mockery -name=GrabX -case underscore -testonly -inpkg
