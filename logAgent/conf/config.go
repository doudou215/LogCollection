package conf

type AppConf struct {
	KafkaConf   `ini:"kafka"` // 在使用了ini包之后， kafka会作为标签，去寻找配置文件中对应的值
	EtcdConf    `ini:"etcd"`  // 使用etcd作为标签
	TailLogConf `ini:"tailLog"`
}

type KafkaConf struct {
	Address     string `ini:"address"`
	MaxChanSize int    `ini:"max_chan_size"`
}

type EtcdConf struct {
	Address string `ini:"address"`
	Key     string `ini:"collect_log_name"`
	Timeout int    `ini:"timeout"`
}

type TailLogConf struct {
	Filename string `ini:"filename"`
	// ``表示不转义， 不对，是要和下面的文件里面的名字保持一致
}

/*
[kafka]
address = 127.0.0.1:9092
max_chan_size = 100000

[etcd]
address = 127.0.0.1:2379
timeout = 5
collect_log_name = loges
*/
