module github.com/sogvin/website

go 1.21

toolchain go1.21.3

require (
	github.com/gregoryv/asserter v0.4.2
	github.com/gregoryv/cmdline v0.15.2
	github.com/gregoryv/draw v0.33.0
	github.com/gregoryv/logger v0.2.0
	github.com/gregoryv/navstar v0.3.0
	github.com/gregoryv/qual v0.4.2
	github.com/gregoryv/web v0.25.0
	github.com/gregoryv/workdir v0.2.1
)

require (
	github.com/gregoryv/gocyclo v0.1.1 // indirect
	github.com/gregoryv/nexus v0.6.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
)

replace github.com/gregoryv/navstar => ../navstar
