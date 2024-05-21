module github.com/khulnasoft/kengine/khulnasoft_agent/tools/apache/khulnasoft/ke-utils/get_cloud_instance_id

go 1.21

replace github.com/khulnasoft/ke-utils => ../../ke-utils/

replace github.com/khulnasoft/ke-utils/cloud_metadata => ../cloud_metadata

require github.com/khulnasoft/ke-utils/cloud_metadata v0.0.0-00010101000000-000000000000

replace github.com/khulnasoft/kengine/khulnasoft_utils => ../../../../../../khulnasoft_utils

require (
	github.com/khulnasoft/ke-utils v0.0.0-00010101000000-000000000000 // indirect
	github.com/khulnasoft/kengine/khulnasoft_utils v0.0.0-00010101000000-000000000000 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rs/zerolog v1.32.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
)
