#!/bin/bash

GOOS="linux" go build -ldflags="-s -w" -tags=jsoniter -o ./bin/simple-api ./cmd/main.go

# descomente se tiver o utilitario UPX instalado para habilitar alta compressao
# upx ./bin/simple-api
