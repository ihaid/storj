module storj.io/storj

go 1.13

// force specific versions for minio
require (
	github.com/btcsuite/btcutil v0.0.0-20180706230648-ab6388e0c60a
	github.com/graphql-go/graphql v0.7.9-0.20190403165646-199d20bbfed7
	github.com/segmentio/go-prompt v1.2.1-0.20161017233205-f0d19b6901ad
)

exclude gopkg.in/olivere/elastic.v5 v5.0.72 // buggy import, see https://github.com/olivere/elastic/pull/869

replace google.golang.org/grpc => github.com/storj/grpc-go v1.23.1-0.20190918084400-1c4561bf5127

require (
	github.com/Shopify/go-lua v0.0.0-20181106184032-48449c60c0a9
	github.com/alessio/shellescape v0.0.0-20190409004728-b115ca0f9053
	github.com/alicebob/gopher-json v0.0.0-20180125190556-5a6b3ba71ee6 // indirect
	github.com/alicebob/miniredis v0.0.0-20180911162847-3657542c8629
	github.com/blang/semver v3.5.1+incompatible
	github.com/boltdb/bolt v1.3.1
	github.com/cheggaaa/pb/v3 v3.0.1
	github.com/cockroachdb/cockroach-go v0.0.0-20181001143604-e0a95dfd547c
	github.com/fatih/color v1.7.0
	github.com/go-ini/ini v1.38.2 // indirect
	github.com/go-redis/redis v6.14.1+incompatible
	github.com/gogo/protobuf v1.2.1
	github.com/golang-migrate/migrate/v4 v4.7.0
	github.com/golang/protobuf v1.3.2
	github.com/gomodule/redigo v2.0.0+incompatible // indirect
	github.com/google/go-cmp v0.3.0
	github.com/gopherjs/gopherjs v0.0.0-20181103185306-d547d1d9531e // indirect
	github.com/gorilla/mux v1.7.1
	github.com/gorilla/schema v1.1.0
	github.com/howeyc/gopass v0.0.0-20170109162249-bf9dde6d0d2c // indirect
	github.com/jtolds/gls v4.2.1+incompatible // indirect
	github.com/jtolds/go-luar v0.0.0-20170419063437-0786921db8c0
	github.com/jtolds/monkit-hw v0.0.0-20190108155550-0f753668cf20
	github.com/lib/pq v1.0.0
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/minio/minio-go v6.0.3+incompatible
	github.com/minio/sha256-simd v0.0.0-20190328051042-05b4dd3047e5
	github.com/nsf/jsondiff v0.0.0-20160203110537-7de28ed2b6e3
	github.com/nsf/termbox-go v0.0.0-20190121233118-02980233997d
	github.com/skyrings/skyring-common v0.0.0-20160929130248-d1c0bb1cbd5e
	github.com/smartystreets/assertions v0.0.0-20180820201707-7c9eb446e3cf // indirect
	github.com/smartystreets/goconvey v0.0.0-20180222194500-ef6db91d284a // indirect
	github.com/spacemonkeygo/errors v0.0.0-20171212215202-9064522e9fd1 // indirect
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
	github.com/stripe/stripe-go v63.1.1+incompatible
	github.com/vivint/infectious v0.0.0-20190108171102-2455b059135b
	github.com/yuin/gopher-lua v0.0.0-20180918061612-799fa34954fb // indirect
	github.com/zeebo/admission v0.0.0-20180821192747-f24f2a94a40c
	github.com/zeebo/errs v1.2.2
	github.com/zeebo/float16 v0.1.0 // indirect
	github.com/zeebo/incenc v0.0.0-20180505221441-0d92902eec54 // indirect
	github.com/zeebo/structs v1.0.2
	go.uber.org/zap v1.10.0
	golang.org/x/crypto v0.0.0-20190911031432-227b76d455e7
	golang.org/x/net v0.0.0-20190916140828-c8589233b77d // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys v0.0.0-20190916202348-b4ddaad3f8a3
	golang.org/x/tools v0.0.0-20190917215024-905c8ffbfa41
	google.golang.org/genproto v0.0.0-20190716160619-c506a9f90610 // indirect
	google.golang.org/grpc v1.23.1
	gopkg.in/ini.v1 v1.38.2 // indirect
	gopkg.in/spacemonkeygo/monkit.v2 v2.0.0-20190612171030-cf5a9e6f8fd2
	gopkg.in/yaml.v2 v2.2.2
	storj.io/drpc v0.0.7-0.20191115031725-2171c57838d2
)
