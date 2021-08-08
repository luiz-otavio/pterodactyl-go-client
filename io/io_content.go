package io

import (
	"io"
	"log"

	"github.com/valyala/fastjson"
)

func JSONBody(reader *io.ReadCloser) (*fastjson.Value, error) {
	content, err := io.ReadAll(*reader)

	if err != nil {
		log.Fatal(err)
	}

	return fastjson.ParseBytes(content)
}
