module github.com/fictionbase/agent

go 1.12

require github.com/fictionbase/fictionbase v0.0.0-20190510061455-5ee57a5d609d

replace (
	github.com/fictionbase/fictionbase => ../fictionbase
	github.com/fictionbase/monitor => ../monitor
	github.com/fictionbase/router => ../router
)
