module unity.service

go 1.18

require (
	github.com/buger/jsonparser v1.1.1
	github.com/chenjiandongx/ginprom v0.0.0-20201217063207-fe11b7f55a35
	github.com/gin-gonic/gin v1.7.1
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.6.0
	github.com/go-redis/redis/v8 v8.11.4
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0
	github.com/jinzhu/gorm v1.9.16
	github.com/prometheus/client_golang v1.13.0
	github.com/prometheus/common v0.37.0 // indirect
	github.com/spaolacci/murmur3 v1.1.0
	github.com/spf13/cobra v1.6.1
	github.com/xxjwxc/ginrpc v0.0.0-20200904081558-8004c9db8189
	golang.org/x/crypto v0.1.0
)

require github.com/swaggo/swag v1.7.0

require (
	entgo.io/ent v0.11.4
	github.com/aws/aws-sdk-go v1.44.154
	github.com/gookit/color v1.2.5
	github.com/tealeg/xlsx v1.0.5
	go.uber.org/automaxprocs v1.5.1
	golang.org/x/text v0.5.0
)

require (
	ariga.io/atlas v0.8.3 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/PuerkitoBio/purell v1.1.1 // indirect
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.3 // indirect
	github.com/go-openapi/jsonreference v0.19.4 // indirect
	github.com/go-openapi/spec v0.19.14 // indirect
	github.com/go-openapi/swag v0.19.11 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/hashicorp/hcl/v2 v2.15.0 // indirect
	github.com/hibiken/asynq v0.23.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/pretty v0.2.1 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.1 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.24.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/sergi/go-diff v1.1.0 // indirect
	github.com/smartystreets/goconvey v1.7.2 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	github.com/ugorji/go/codec v1.1.13 // indirect
	github.com/xxjwxc/public v0.0.0-20200601115915-ab2b4ce31a9c // indirect
	github.com/zclconf/go-cty v1.12.1 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.3.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
	golang.org/x/tools v0.4.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/gin-gonic/gin => ./pkg/gin

replace github.com/xxjwxc/ginrpc => ./pkg/ginrpc

replace github.com/TheAlgorithms/Go v0.0.1 => github.com/tjgurwara99/Go v0.0.4

//replace github.com/TheAlgorithms/Go => ./pkg/algorithm

//replace gopkg.in/ory-am/dockertest.v3 v3.8.1 => github.com/ory/dockertest/v3 v3.8.1

//replace github.com/TheAlgorithms/Go master => github.com/tjgurwara99/Go master
