package proxying

import (
	"context"
)

type Proxy interface {
	PassThrough(ctx context.Context, s string) (string error)
}
