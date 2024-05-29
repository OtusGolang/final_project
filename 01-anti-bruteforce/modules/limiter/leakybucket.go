package limiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu               sync.Mutex
	loginAttempts    map[string]*LeakyBucket
	passwordAttempts map[string]*LeakyBucket
	ipAttempts       map[string]*LeakyBucket
}

type LeakyBucket struct {
	rate         time.Duration
	capacity     int
	tokens       int
	lastLeakTime time.Time
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		loginAttempts:    make(map[string]*LeakyBucket),
		passwordAttempts: make(map[string]*LeakyBucket),
		ipAttempts:       make(map[string]*LeakyBucket),
	}
}

func (l *RateLimiter) ClearBucket(r *RateLimiter, login, ip string) bool {
	r.mu.Lock()

	if _, ok := r.loginAttempts[login]; ok {
		delete(r.loginAttempts, login)
	}
	if _, ok := r.ipAttempts[ip]; ok {
		delete(r.ipAttempts, ip)
	}
	r.mu.Unlock()
	return true
}

func (l *RateLimiter) IsAllowed(login, password, ip string, N, M, K int) bool {
	// Check login bucket
	l.mu.Lock()
	loginBucket, ok := l.loginAttempts[login]
	if !ok {
		loginBucket = &LeakyBucket{
			rate:         time.Minute / 10,
			capacity:     N,
			tokens:       0,
			lastLeakTime: time.Now(),
		}
		l.loginAttempts[login] = loginBucket
	}
	l.mu.Unlock()

	if !loginBucket.consumeToken() {
		return false
	}

	// Check password bucket
	l.mu.Lock()
	passwordBucket, ok := l.passwordAttempts[password]
	if !ok {
		passwordBucket = &LeakyBucket{
			rate:         time.Minute / 100,
			capacity:     M,
			tokens:       0,
			lastLeakTime: time.Now(),
		}
		l.passwordAttempts[password] = passwordBucket
	}
	l.mu.Unlock()

	if !passwordBucket.consumeToken() {
		return false
	}

	// Check IP bucket
	l.mu.Lock()
	ipBucket, ok := l.ipAttempts[ip]
	if !ok {
		ipBucket = &LeakyBucket{
			rate:         time.Minute / 1000,
			capacity:     K,
			tokens:       0,
			lastLeakTime: time.Now(),
		}
		l.ipAttempts[ip] = ipBucket
	}
	l.mu.Unlock()

	if !ipBucket.consumeToken() {
		return false
	}

	return true
}

func (b *LeakyBucket) consumeToken() bool {
	b.leakTokens()

	if b.tokens < b.capacity {
		b.tokens++
		return true
	}

	return false
}

func (b *LeakyBucket) leakTokens() {
	now := time.Now()
	elapsed := now.Sub(b.lastLeakTime)
	numTokens := int(elapsed / b.rate)
	b.lastLeakTime = now
	b.tokens -= numTokens
	if b.tokens < 0 {
		b.tokens = 0
	}
}
