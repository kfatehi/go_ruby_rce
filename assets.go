package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets15e59ec61440888368f15c5f035ea96a0af5db02 = "require 'ripper'\nrequire 'json'\nsexp = Ripper.sexp(ARGF)\nputs JSON.generate sexp"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"assets"}, "/assets": []string{"validator.rb"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1659914054, 1659914054329749501),
		Data:     nil,
	}, "/assets": &assets.File{
		Path:     "/assets",
		FileMode: 0x800001fd,
		Mtime:    time.Unix(1659913006, 1659913006934972516),
		Data:     nil,
	}, "/assets/validator.rb": &assets.File{
		Path:     "/assets/validator.rb",
		FileMode: 0x1b4,
		Mtime:    time.Unix(1659913008, 1659913008574897982),
		Data:     []byte(_Assets15e59ec61440888368f15c5f035ea96a0af5db02),
	}}, "")
