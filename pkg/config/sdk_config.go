package config

import "net/http"

type SDKConfig struct {
	Configuration
	AccessToken  string
	IsDebug      bool
	DebugFile    string
	SkipMonitor  bool
	GlobalConfig GlobalConfig
}

type GlobalConfig struct {
	ServiceName ServiceName
	HttpOption  HttpOption
}

type ServiceName struct {
	Name   string
	Schema string
}

type HttpOption struct {
	Header http.Header
}
