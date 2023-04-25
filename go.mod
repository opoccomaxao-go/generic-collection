module github.com/opoccomaxao-go/generic-collection

go 1.18

require (
	github.com/stretchr/testify v1.8.1
	golang.org/x/exp v0.0.0-20230127193734-31bee513bff7
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.0.0 // Module abandoned. Use github.com/samber/lo instead.
)
