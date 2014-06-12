package digitalocean

import (
	"errors"
	"sync"

	"github.com/koding/logging"
)

type Factory interface {
	Create() (*Droplet, error)
	Destroy(uint) error
}

type Pool struct {
	mu       sync.Mutex
	droplets chan *Droplet
	log      logging.Logger

	factory Factory
}

// InitializeCaching should be called only once. It creates
func NewPool(initialCap, maxCap int, factory Factory) (*Pool, error) {
	if initialCap <= 0 || maxCap <= 0 || initialCap > maxCap {
		return nil, errors.New("invalid capacity settings")
	}

	if factory == nil {
		return nil, errors.New("factory function is not given.")
	}

	p := &Pool{
		droplets: make(chan *Droplet, maxCap),
		factory:  factory,
	}

	for i := 0; i < initialCap; i++ {
		go func() {
			droplet, err := p.factory.Create()
			if err != nil {
				p.log.Error("filling cache channel: %s", err.Error())
			}

			p.Put(droplet)
		}()
	}

	return p, nil
}

// PutCacheDroplet puts the droplet into the cache pool. If the pool is full,
// it factory is invoked to destroy the passed droplet.
func (p *Pool) Put(droplet *Droplet) error {
	select {
	case p.droplets <- droplet:
		return nil
	default:
		go p.factory.Destroy(uint(droplet.Droplet.Id))
		return errors.New("cache is already full, deleting previous droplet")
	}
}

// Get fetchs a droplet from the cache pool. After a succesfull fetch another
// goroutine is going to create a cached droplet to fill the cache pool. If
// there is no droplet in the cache pool a new droplet is created and returned.
func (p *Pool) Get() (*Droplet, error) {
	select {
	case droplet := <-p.droplets:
		if droplet == nil {
			return nil, errors.New("cache channel is closed")
		}

		// create and put another cached machine into the channel after we get
		// one, let's try that that channel is always be filled.
		go func() {
			cacheDroplet, err := p.factory.Create()
			if err != nil {
				p.log.Error("getCacheDroplet, creating err: %s", err.Error())
				return
			}

			if err := p.Put(cacheDroplet); err != nil {
				p.log.Error("getCacheDroplet, putting err: %s", err.Error())
			}
		}()

		return droplet, nil
	default:
		return p.factory.Create()
	}
}
