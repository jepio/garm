package pool

import (
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	runnerErrors "github.com/cloudbase/garm-provider-common/errors"
	"github.com/cloudbase/garm/params"
)

type poolRoundRobin struct {
	pools []params.Pool
	next  uint32
}

func (p *poolRoundRobin) Next() (params.Pool, error) {
	if len(p.pools) == 0 {
		return params.Pool{}, runnerErrors.ErrNoPoolsAvailable
	}

	n := atomic.AddUint32(&p.next, 1)
	return p.pools[(int(n)-1)%len(p.pools)], nil
}

func (p *poolRoundRobin) Len() int {
	return len(p.pools)
}

func (p *poolRoundRobin) Reset() {
	atomic.StoreUint32(&p.next, 0)
}

type poolsForTags struct {
	pools sync.Map
}

func (p *poolsForTags) Get(tags []string) (*poolRoundRobin, bool) {
	sort.Strings(tags)
	key := strings.Join(tags, "^")

	v, ok := p.pools.Load(key)
	if !ok {
		return nil, false
	}

	return v.(*poolRoundRobin), true
}

func (p *poolsForTags) Add(tags []string, pools []params.Pool) *poolRoundRobin {
	sort.Strings(tags)
	key := strings.Join(tags, "^")

	poolRR := &poolRoundRobin{pools: pools}
	v, _ := p.pools.LoadOrStore(key, poolRR)
	return v.(*poolRoundRobin)
}
