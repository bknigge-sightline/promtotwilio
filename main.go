package main

import (
	"fmt"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type options struct {
	Host       string
	Port       int
	AccountSid string
	AuthToken  string
	Receiver   string
	Sender     string
}

func main() {

	k := koanf.New(".")
	err := k.Load(file.Provider("conf.toml"), toml.Parser())
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	opts := options{
		Host:       k.String("host"),
		Port:       k.Int("port"),
		AccountSid: k.String("accountSid"),
		AuthToken:  k.String("authToken"),
		Receiver:   k.String("receiver"),
		Sender:     k.String("accountSid"),
	}

	if opts.AccountSid == "" || opts.AuthToken == "" || opts.Sender == "" {
		log.Fatal("'SID', 'TOKEN' and 'SENDER' environment variables need to be set")
	}

	o := NewMOptionsWithHandler(&opts)
	err = fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", opts.Host, opts.Port), o.HandleFastHTTP)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
