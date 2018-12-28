package concensus

import "github.com/hashicorp/raft"

var rafts map[raft.ServerAddress]*raft.Raft

func init() {
	rafts = make(map[raft.ServerAddress]*raft.Raft)
}

func Config(num int) {
	conf := raft.DefaultConfig()
	snapshotStore := raft.NewDiscardSnapshotStore()

	addrs := []raft.ServerAddress{}
	transports := []*raft.InmemTransport{}
	for i := 0; i < num; i++ {
		addr, transport := raft.NewInmemTransport("127.0.0.1")
		addrs = append(addrs, addr)
		transports = append(transports, transport)
	}
	stableStore := raft.NewInmemStore()
	memstore := raft.NewInmemStore()

	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			if i != j {
				transports[i].Connect(addrs[j], transports[j])
			}
		}
		conf.LocalID = "127.0.0.1:2133"
		r, err := raft.NewRaft(conf, NewFSM(), memstore, stableStore, snapshotStore, transports[i])
		if err != nil {
			panic(err)
		}
		//TODO: add peer to raft
		rafts[addrs[i]] = r
	}
}
