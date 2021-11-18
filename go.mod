module github.com/anschwa/giftopotamus

// https://github.com/heroku/heroku-buildpack-go
// +heroku goVersion go1.17
// +heroku install ./cmd/app
go 1.17

require (
	github.com/aws/aws-sdk-go v1.41.14
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect
