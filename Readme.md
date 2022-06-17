# flatfile-go

[![Documentation](https://godoc.org/github.com/be-streamlined/flatfile-go?status.svg)](http://godoc.org/github.com/be-streamlined/flatfile-go)
[![CI](https://github.com/Be-Streamlined/flatfile-go/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/Be-Streamlined/flatfile-go/actions/workflows/ci.yaml)

A package to interface with the [Flatfile](https://flatfile.com) APIs.


## Install

```
go get -u github.com/be-streamlined/flatfile-go
```

## Usage 

### Basic usage

#### From Environment

```golang
flatfile, err := flatfile.NewFromEnv()
```
#### With Config Variables

```golang
apiKey := ""
secretKey := ""
licenseKey := ""

flatfile, err := flatfile.New(apiKey, secretKey, licenseKey)
```
## Environment Variables
- FLATFILE_API_KEY
    - The API Key for Flatfile
- FLATFILE_SECRET_API_KEY
    - The Flatfile Secret API Key
- FLATFILE_LICENSE_KEY
    - The License Key for Flatfile
- FLATFILE_BASE_URL
    - The base url for Flatfile

## Examples

### Get a batch
```golang
flatfile, err := flatfile.NewFromEnv()

batch, err := flatfile.GetBatch("1234", 0)
```


## Contributing

Pull requests are welcome. Feel free to...

- Revise documentation
- Add new features
- Fix bugs
- Suggest improvements

## License

MIT
