package conf

type AppConf struct {
	KafkaConf   `ini:"kafka"`
	EtcdConf    `ini:"etcd"`
	TailLogConf `ini:"tailLog"`
}

type KafkaConf struct {
	Address string `ini:"address"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Key     string `ini:"key"`
	Timeout int    `ini:"timeout"`
}

type TailLogConf struct {
	Filename string `ini:"filename"`
	// ``表示不转义
}
