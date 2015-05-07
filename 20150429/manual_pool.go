// +build ignore

package main

// Pool holds Clients.
type Pool struct {
	pool chan *Client
}

// NewPool creates a new pool of Clients.
func NewPool(max int) *Pool {
	return &Pool{
		pool: make(chan *Client, max),
	}
}

// Borrow a Client from the pool.
func (p *Pool) Borrow() *Client {
	var c *Client
	select {
	case c = <-p.pool:
	default:
		c = newClient()
	}
	return c
}

// Return returns a Client to the pool.
func (p *Pool) Return(c *Client) {
	select {
	case p.pool <- c:
	default:
		// let it go, let it go...
	}
}

func main() {

}
