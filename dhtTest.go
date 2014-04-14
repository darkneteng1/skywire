package main

import (
	"fmt"
	"github.com/skycoin/skywire/src/daemon/dht"
	"log"
	"time"
)

func PeerCallback(infoHash string, peerAddress string) {

	fmt.Printf("PeerCallback: infoHash= %s, peerAddres= %s \n", infoHash, peerAddress)
}

func main() {

	config := daemon_dht.NewDHTConfig()
	config.AddPeerCallback = PeerCallback

	dht := daemon_dht.NewDHT(config)

	err := dht.Init()

	if err != nil {
		log.Panic()
	}

	log.Printf("init")
	go dht.Start()
	log.Printf("start")

	go dht.Listen()

	for i := 0; i < 10; i++ {
		dht.RequestPeers("skycoin-skycoin-skycoin-skycoin-skycoin-skycoin-skycoin")
		time.Sleep(time.Second * 1)
	}

	time.Sleep(time.Second * 60)
	dht.Shutdown()
}