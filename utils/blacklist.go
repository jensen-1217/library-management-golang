package utils

import "sync"

var (
	blacklist []string
	mu        sync.Mutex
)

// AddTokenToBlacklist 将令牌添加到黑名单中
func AddTokenToBlacklist(token string) {
	mu.Lock()
	defer mu.Unlock()
	blacklist = append(blacklist, token)
}

// IsTokenBlacklisted 检查令牌是否在黑名单中
func IsTokenBlacklisted(token string) bool {
	mu.Lock()
	defer mu.Unlock()
	for _, t := range blacklist {
		if t == token {
			return true
		}
	}
	return false
}
