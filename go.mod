module github.com/fictionbase/agent

go 1.12

replace (
	github.com/fictionbase/fictionbase => ../fictionbase
	github.com/fictionbase/monitor => ../monitor
	github.com/fictionbase/router => ../router
)

require github.com/fictionbase/fictionbase v0.0.0-00010101000000-000000000000
