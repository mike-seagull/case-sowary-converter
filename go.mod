module github.com/mike-seagull/csc

go 1.17

require (
	github.com/magefile/mage v1.11.0
	internal/cli v1.0.0
)

require pkg/case_converter v1.0.0 // indirect

require github.com/spf13/cobra v1.2.1 // indirect

require (
	github.com/cweill/gotests v1.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/sys v0.0.0-20211117180635-dee7805ff2e1 // indirect
	golang.org/x/tools v0.1.7 // indirect
)

replace internal/cli => ./internal/cli

replace pkg/case_converter => ./pkg/case_converter
