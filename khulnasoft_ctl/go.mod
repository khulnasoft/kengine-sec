module github.com/khulnasoft/kengine/khulnasoft_ctl

go 1.21

toolchain go1.21.5

replace github.com/khulnasoft/kengine/khulnasoft_utils => ../khulnasoft_utils/

replace github.com/khulnasoft-lab/golang_sdk/client => ../golang_sdk/client/

replace github.com/khulnasoft-lab/golang_sdk/utils => ../golang_sdk/utils/

require (
	github.com/khulnasoft-lab/golang_sdk/client v0.0.0-00010101000000-000000000000
	github.com/khulnasoft-lab/golang_sdk/utils v0.0.0-00010101000000-000000000000
	github.com/khulnasoft/kengine/khulnasoft_utils v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.8.0

)

require (
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.5 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/rs/zerolog v1.32.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.20.0 // indirect
)
