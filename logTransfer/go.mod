module logTransfer

go 1.15

replace logTransfer/conf => ./conf

require (
	github.com/Shopify/sarama v1.19.0
	github.com/Shopify/toxiproxy v2.1.4+incompatible // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/frankban/quicktest v1.13.0 // indirect
	github.com/golang/snappy v0.0.3 // indirect
	github.com/olivere/elastic/v7 v7.0.24
	github.com/pierrec/lz4 v2.6.0+incompatible // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	gopkg.in/ini.v1 v1.62.0
)
