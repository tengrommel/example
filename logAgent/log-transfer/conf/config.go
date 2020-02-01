package conf

// LogTransfer全局配置
type LogTransfer struct {
	KafkaCfg `ini:"kafka"`
	ESCfg    `ini:"es"`
}

// KafkaCfg ...
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

type ESCfg struct {
	Address string `ini:"address"`
}
