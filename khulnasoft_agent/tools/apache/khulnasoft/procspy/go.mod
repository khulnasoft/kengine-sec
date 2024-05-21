module github.com/khulnasoft/kengine/khulnasoft_agent/tools/apache/khulnasoft/procspy

go 1.21

replace github.com/khulnasoft/ke-utils => ../ke-utils/

replace github.com/khulnasoft/ke-utils/osrelease => ../ke-utils/osrelease

replace github.com/khulnasoft/kengine/khulnasoft_utils => ../../../../../khulnasoft_utils

require (
	github.com/khulnasoft/ke-utils v0.0.0-00010101000000-000000000000
	github.com/weaveworks/procspy v0.0.0-20150706124340-cb970aa190c3
)

require (
	github.com/khulnasoft/kengine/khulnasoft_utils v0.0.0-00010101000000-000000000000 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rs/zerolog v1.32.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
)
