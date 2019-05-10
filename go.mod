module github.com/fictionbase/agent

go 1.12

require github.com/fictionbase/fictionbase v0.0.0-20190510051858-251968169696

replace (
	github.com/fictionbase/agent => ../agent
	github.com/fictionbase/fictionbase => ../fictionbase
	github.com/fictionbase/monitor => ../monitor
	github.com/fictionbase/router => ../router
)
