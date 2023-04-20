package conf

type SumConf struct {
	KafkaConf `ini:"kafka"`
	TailConf  `ini:"taillog"`
}
type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}
type TailConf struct {
	Filename string `ini:"filename"`
}
