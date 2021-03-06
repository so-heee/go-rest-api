// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"X7Qe2DJzcFaGCJxsJ38GAAD//2ZX2JtSFgAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
