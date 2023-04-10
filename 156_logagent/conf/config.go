package conf

type Sumconfig struct {
	KafkaConf   `ini:"kafka"`
	TaillogConf `ini:"taillog"`
}
type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}
type TaillogConf struct {
	Filename string `ini:"filename"`
}