module github.com/jeffizhungry/workflows

go 1.14

// Need to version pin thrift for compilation issues
// https://github.com/uber-go/cadence-client/issues/523
replace github.com/apache/thrift => github.com/apache/thrift v0.0.0-20190309152529-a9b748bb0e02

require (
	github.com/codeskyblue/fswatch v0.0.0-20191227065248-65cdcfddf017 // indirect
	github.com/codeskyblue/kexec v0.0.0-20180119015717-5a4bed90d99a // indirect
	github.com/crossdock/crossdock-go v0.0.0-20160816171116-049aabb0122b // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/google/shlex v0.0.0-20191202100458-e7afc7fbc510 // indirect
	github.com/k0kubun/pp v3.0.1+incompatible // indirect
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/echo/v4 v4.1.16
	github.com/pborman/uuid v0.0.0-20160209185913-a97ce2ca70fa // indirect
	github.com/robfig/cron v1.2.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.7.0
	github.com/uber-go/atomic v1.4.0 // indirect
	github.com/uber-go/tally v3.3.17+incompatible
	go.uber.org/cadence v0.12.2
	go.uber.org/yarpc v1.46.0
	go.uber.org/zap v1.15.0
	golang.org/x/sys v0.0.0-20200625212154-ddb9806d33ae // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
)
