package main

import "encoding/json"

var gitHash string  // populated with -ldflags
var gitDirty string // populated with -ldflags

type BuildInfo struct {
	GitHash  string `json:"gitHash"`
	GitDirty bool   `json:"gitDirty"`
}

func (b *BuildInfo) ToJson() string {
	jsonBytes, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func GetBuildInfo() *BuildInfo {

	info := &BuildInfo{
		GitHash:  "unknown",
		GitDirty: false,
	}

	if gitHash != "" {
		info.GitHash = gitHash
	}

	if gitDirty == "true" {
		info.GitDirty = true
	}

	return info
}
