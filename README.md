# image-splitter
A CLI tool which splits an image into Red-Green-Blue layers

## How to Use
`image-splitter /file/path/to/source-image.png`

The command will create the following files in the same directory

- `source-image-red.png`
- `source-image-green.png`
- `source-image-blue.png`

### Options
TODO

## Development Environment
Go 1.17 is required

## Development Operations
- `go build image-splitter.go` to build executable, named `image-splitter`
- `go run image-splitter.go` to execute
