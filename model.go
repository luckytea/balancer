package balancer

type Balancer struct {
	hosts   []string
	current uint64
}
