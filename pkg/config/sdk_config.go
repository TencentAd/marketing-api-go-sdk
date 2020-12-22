package config

import "net/http"

// SDKConfig ...
type SDKConfig struct {
	Configuration
	AccessToken  string
	IsDebug      bool
	DebugFile    string
	SkipMonitor  bool
	IsStrictMode bool
	GlobalConfig GlobalConfig
}

// GlobalConfig ...
type GlobalConfig struct {
	ServiceName ServiceName
	HttpOption  HttpOption
}

// ServiceName ...
type ServiceName struct {
	Name   string
	Schema string
}

// HttpOption ...
type HttpOption struct {
	Header http.Header
}
