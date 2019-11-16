package ratelimit

import (
	//"fmt"
	//"math"
	"time"
)

type Limit float64

type Limiter struct {
	limit		Limit
	last		time.Time
	lastEvent	time.Time
}

func NewRateLimiter(r Limit) *Limiter {
	return &Limiter{
		limit: r,
		last: time.Now(),
	}
}

// AddLimit set a new limit for the limiter
// if limit over the burst, return error and wait 1 min
func (lim *Limiter) AddLimit() {
	now, newLimit := lim.queue(time.Now())

	lim.last = now
	lim.limit = newLimit

	//fmt.Println(lim)
}

func (lim *Limiter) GetLimit() Limit {
	return lim.limit
}

func (lim *Limiter) queue(now time.Time) (newNow time.Time, newLimit Limit) {
	t := now.Sub(lim.last)

	newNow = lim.last
	newLimit = lim.limit + Limit(1)

	// over 1 min and reset limit
	if t.Minutes() > float64(1) {
		newNow = time.Now()
		newLimit = Limit(1)
	}

	return newNow, newLimit
}
