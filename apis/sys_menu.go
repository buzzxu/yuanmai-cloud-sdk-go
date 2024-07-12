package apis

import (
	"github.com/buzzxu/boys/types"
	objects "github.com/buzzxu/yuanmai-objects-go"
)

type (
	SysMenuApi interface {
		List() types.Pager[objects.Menu]
	}
)
