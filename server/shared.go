package server

import "context"

type ServerShutdownFunc = func(ctx context.Context) error
