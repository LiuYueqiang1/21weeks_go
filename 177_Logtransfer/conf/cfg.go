package conf

type LogTransfer struct {
	Kafka `ini:"kafka"`
	EScfg `ini:"es"`
}
type Kafka struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}
type EScfg struct {
	Address string `ini:"address"`
}
