package cache

import (
	"context"
	"github.com/Jadepypy/distributed-social-media/application/internal/domain"
)

type OtpCache interface {
	Set(ctx context.Context, user domain.User) error
	Get(ctx context.Context, user domain.User) error
}
