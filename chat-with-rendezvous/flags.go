package main

import (
	"flag"
	"strings"

	maddr "github.com/multiformats/go-multiaddr"
	multiaddr "github.com/multiformats/go-multiaddr"
)

// A new type we need for writing a custom flag parser
type addrList []maddr.Multiaddr

func (al *addrList) String() string {
	strs := make([]string, len(*al))
	for i, addr := range *al {
		strs[i] = addr.String()
	}
	return strings.Join(strs, ",")
}

func (al *addrList) Set(value string) error {
	addr, err := maddr.NewMultiaddr(value)
	if err != nil {
		return err
	}
	*al = append(*al, addr)
	return nil
}

func StringsToAddrs(addrStrings []string) (maddrs []maddr.Multiaddr, err error) {
	for _, addrString := range addrStrings {
		addr, err := maddr.NewMultiaddr(addrString)
		if err != nil {
			return maddrs, err
		}
		maddrs = append(maddrs, addr)
	}
	return
}

type Config struct {
	RendezvousString string
	BootstrapPeers   addrList
	ListenAddresses  addrList
	ProtocolID       string
}

func ParseFlags() (Config, error) {
	config := Config{}
	flag.StringVar(&config.RendezvousString, "rendezvous", "meet me here",
		"Unique string to identify group of nodes. Share this with your friends to let them connect with you")
	flag.Var(&config.BootstrapPeers, "peer", "Adds a peer multiaddress to the bootstrap list")
	flag.Var(&config.ListenAddresses, "listen", "Adds a multiaddress to the listen list")
	flag.StringVar(&config.ProtocolID, "pid", "/chat/1.1.0", "Sets a protocol id for stream headers")
	flag.Parse()

	if len(config.BootstrapPeers) == 0 {
		//config.BootstrapPeers = dht.DefaultBootstrapPeers
		// ma, err := maddr.NewMultiaddr("/ip4/220.194.157.80/tcp/4001/p2p/QmP2C45o2vZfy1JXWFZDUEzrQCigMtd4r3nesvArV8dFKd")
		// ma, err := maddr.NewMultiaddr("/ip4/192.168.1.175/tcp/4001/p2p/QmdSyhb8eR9dDSR5jjnRoTDBwpBCSAjT7WueKJ9cQArYoA")
		// if err != nil {
		// 	panic(err)
		// }

		// config.BootstrapPeers = append(config.BootstrapPeers, ma)

		for _, s := range []string{
			// "/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
			// "/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
			// "/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
			// "/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
			// "/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ", // mars.i.ipfs.io
			"/ip4/220.194.157.80/tcp/4001/p2p/QmP2C45o2vZfy1JXWFZDUEzrQCigMtd4r3nesvArV8dFKd",
			"/ip4/192.168.1.175/tcp/4001/p2p/QmdSyhb8eR9dDSR5jjnRoTDBwpBCSAjT7WueKJ9cQArYoA",
		} {
			ma, err := multiaddr.NewMultiaddr(s)
			if err != nil {
				panic(err)
			}
			// DefaultBootstrapPeers = append(DefaultBootstrapPeers, ma)
			config.BootstrapPeers = append(config.BootstrapPeers, ma)
		}
	}

	return config, nil
}
