module logAgent

go 1.15

require (
	github.com/doudou215/LogCollection/logAgent/etcd v0.0.0
	github.com/doudou215/LogCollection/logAgent/kafka v0.0.0
	github.com/doudou215/LogCollection/logAgent/tailLog v0.0.0
	github.com/olivere/elastic/v7 v7.0.24
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.62.0
)

replace (
	github.com/doudou215/LogCollection/logAgent/etcd => ./etcd
	github.com/doudou215/LogCollection/logAgent/kafka => ./kafka
	github.com/doudou215/LogCollection/logAgent/tailLog => ./tailLog

)
