package init

import (
	"fmt"
	initAdapter "game-go/core/adapter/init"
)

type controller struct {
	adapter initAdapter.Adapter
}

func New(adapter initAdapter.Adapter) Controller {
	return &controller{adapter: adapter}
}

func (c *controller) InitBetAreaCache() {
	if err := c.adapter.InitBetAreaCache(); err != nil {
		fmt.Println(err)
	}
}
