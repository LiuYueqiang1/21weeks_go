package conf

type SumConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}
type KafkaConf struct {
	Address string `ini:"address"`
}
type EtcdConf struct {
	Address string `ini:"address"`
}

// -----------
type TailConf struct {
	Filename string `ini:"filename"`
}
