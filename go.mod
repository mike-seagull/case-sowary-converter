module github.com/mike-seagull/csc

go 1.17

require internal/cli v1.0.0

require pkg/case_converter v1.0.0 // indirect

require github.com/spf13/cobra v1.2.1 // indirect

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace internal/cli => ./internal/cli

replace pkg/case_converter => ./pkg/case_converter
