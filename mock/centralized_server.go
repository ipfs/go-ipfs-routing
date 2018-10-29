package mockrouting

import (
	"context"
	"math/rand"
	"sync"
	"time"

	ds "github.com/ipfs/go-datastore"
	dssync "github.com/ipfs/go-datastore/sync"
	peer "github.com/libp2p/go-libp2p-peer"
	pstore "github.com/libp2p/go-libp2p-peerstore"
	"github.com/libp2p/go-testutil"
	mh "github.com/multiformats/go-multihash"

	offline "github.com/ipfs/go-ipfs-routing/offline"
)

// server is the mockrouting.Client's private interface to the routing server
type server interface {
	Announce(pstore.PeerInfo, mh.Multihash) error
	Providers(mh.Multihash) []pstore.PeerInfo

	Server
}

// s is an implementation of the private server interface
type s struct {
	delayConf DelayConfig

	lock      sync.RWMutex
	providers map[string]map[peer.ID]providerRecord
}

type providerRecord struct {
	Peer    pstore.PeerInfo
	Created time.Time
}

func (rs *s) Announce(p pstore.PeerInfo, c mh.Multihash) error {
	rs.lock.Lock()
	defer rs.lock.Unlock()

	k := string(c)

	_, ok := rs.providers[k]
	if !ok {
		rs.providers[k] = make(map[peer.ID]providerRecord)
	}
	rs.providers[k][p.ID] = providerRecord{
		Created: time.Now(),
		Peer:    p,
	}
	return nil
}

func (rs *s) Providers(c mh.Multihash) []pstore.PeerInfo {
	rs.delayConf.Query.Wait() // before locking

	rs.lock.RLock()
	defer rs.lock.RUnlock()
	k := string(c)

	var ret []pstore.PeerInfo
	records, ok := rs.providers[k]
	if !ok {
		return ret
	}
	for _, r := range records {
		if time.Since(r.Created) > rs.delayConf.ValueVisibility.Get() {
			ret = append(ret, r.Peer)
		}
	}

	for i := range ret {
		j := rand.Intn(i + 1)
		ret[i], ret[j] = ret[j], ret[i]
	}

	return ret
}

func (rs *s) Client(p testutil.Identity) Client {
	return rs.ClientWithDatastore(context.Background(), p, dssync.MutexWrap(ds.NewMapDatastore()))
}

func (rs *s) ClientWithDatastore(_ context.Context, p testutil.Identity, datastore ds.Datastore) Client {
	return &client{
		peer:   p,
		vs:     offline.NewOfflineRouter(datastore, MockValidator{}),
		server: rs,
	}
}
