package phones_gateway

import (
	"context"
	"fmt"
	"github.com/exitialis/workshop/homework/complex/internal/cache"
	"golang.org/x/time/rate"
	"math/rand"
	"time"
)

func (p *PhoneService) fillCache(ctx context.Context) {
	limiter := rate.NewLimiter(rate.Every(time.Second * 5), 1)
	for err := limiter.Wait(ctx); err == nil; err = limiter.Wait(ctx) {
		p.refillCache()
	}
}

func (p *PhoneService) refillCache() {
	rand.Seed(time.Now().Unix())
	for i := 1; i < 5_000_000; i++ {
		phone := rand.Intn(9_999_999_999 - 9_000_000_000) + 9_000_000_000
		err := p.lru.SetPhone(int64(i), lru.RealPhone(fmt.Sprintf("+7%d", phone)))
		if err != nil {
			panic(err)
		}
	}
}
