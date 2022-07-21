// Package balancer provide easy round-robin balancer.
package balancer

import (
	"strings"
	"sync/atomic"
)

func New(hosts []string) *Balancer {
	return &Balancer{
		hosts: hosts,
	}
}

func NewSplit(hosts, delimiter string) *Balancer {
	return New(strings.Split(hosts, delimiter))
}

func (b *Balancer) GetHost() string {
	idx := b.nextIndex()
	return b.hosts[idx]
}

func (b *Balancer) nextIndex() int {
	return int(atomic.AddUint64(&b.current, uint64(1)) % uint64(len(b.hosts)))
}
