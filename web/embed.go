package web

import "embed"

//go:embed www/*
var WWWFiles embed.FS
