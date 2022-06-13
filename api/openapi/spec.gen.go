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

	"H4sIAAAAAAAC/8xXX2/cRBD/KquFx+v5kgAPfuq1SasrVRslQTxUUbXxTc5bbK+7u256iizRO5CKoAJe",
	"ihAvCCGEQKL8ExJqxZexUsq3QLNrn32273JpK8rTxbs7M7+d3/5mJsfUE2EsIoi0ou4xlaBiESkwH4NI",
	"g4xYsAvyDsgtKYXEZU9EGiKNf7I4DrjHNBeRc0uJCNeU50PI8K/XJRxSl77mlDEcu6ucfsytwzRNO3QI",
	"ypM8Rj/UncUlNjApDnZy3wbczIF7TGMpYpCaW9hQLC+LX/qUcDvhEobUvZGb7neoHsdAXSoOboGnadqh",
	"/UT7EOn8sjtwOwGlm7EjFgL+5vZKSx6N0D5mSh0JOWzZrGEwLioGq6CxpDXhMM8DpW5q8R5ErbAkHEpQ",
	"/pITZuemXT4N+ly4uvM5V22XWkCnJ4btKTVkmSNcQ6hOY3w3OchJn8VmUrIxfoegFButcEMDpjzfdo1Z",
	"nMZNFkdBuiULQYM8HUN5dDmQvSOAlifKzRs8FDJkmrqUR/qtN+jMnEcaRmCTBHd1K9ZEgby5opsadt7+",
	"nt9RIF8Aach40Ip0gRxXB7XNtOcvVPtZ455WBtohCKX/MwQthciEWFqPsDCDl0iux7soNgvtAjAJEisV",
	"fh2Yr0sFl1fe3aN5OUdPdrfk1tc6tp0B7tpesCk843W+U1zi0ZCEQgLh0aEgPkhEnMgg96Fcx4G7LIwD",
	"6HoixPvjwaKLMc+kFGwWqUriWEh9vmpRZJL2twckPzAXwnWco6OjbsXGKY41GlufKF9ITSqrRByS/vag",
	"Szs04B7kVbyIGTPPB7Le7bWFZGa3K+TIyU2Vc3Vwceva7ta59W6v6+swsDqWobp+iM2Ue1D6qEI2Zxwk",
	"gOsAz+yaPYRGO/QOSGVvsNbtdXvoVMQQsZhTl250e90N80C0byhyBEu0v+40esvIlqP5nGSTb7LJo2zy",
	"OJv8mU3vZ9Mn+Dn9Lbv309Ovfn/68GdMDT550+oGQ+rSHeu4b7rNXt5cZlVRUfdGPUgOhRSdiOPi7QTk",
	"uCS43q1KKWiZQKcy19Rls9+Zn5vWe72XNye1d/qWqen621aJSRgyOS7TRGyeiEkUJlOzkTK9LOAIbx+t",
	"Cs5mXMVCtZE1/c4Q9Ec2fXLy2YPs/XvZ9HND3CNcn36U3fsxm9zPJh8/++HBs++fNLmr3AfyJIPSF8Rw",
	"XEtZmASax0xqBxvAuSHT7PmzZqtnmtoS9/+jqpqVhRRp7OfKOTa/g2G6WFLTD7LJt4aQ+0+nH558/cvJ",
	"pw9P/vqiQcZl0GZIuDAebJ4mIROV8GGhHpR7KZ4c0yqyKUeDF9VNXsCKUWGtGFnoFeFHZFOYXlDMKmtY",
	"+1eZE+3Y1OzGaVs559EoAGJMSCiGYArum/YebUFm93Xa/rmafxSXQROb9YMxGWxW34V9Cvm7wCuq1TRr",
	"X8PfXz7+55NfG68BBw0ziC2TZRsjFR7yTnpL+NH5EX7MN9EKMeUcQvXute11Fl6iaWdFqdXnohZyMCtE",
	"C+JJwFLTFP5aM1UXd7b6e1ubcxONkUJ1lrmxjw+3pOmiCUAYieCIILAqUZabCk/OMf4sV2+dr8XqxXCr",
	"iNckY5F2LaBXKt2NGZgd7vlMDsmOMJytpFjzaM8kWLQo9HoWrlGSxrahyIJoMwd5/iq02vGmKUO0z3X4",
	"Ckh9TuHXBX42KVf/y3oeLfdekpYNkMUiNq6wXLfRcVV4LCB2vz6uB7jnC6XdjV6v59xZoxg3999Abpv+",
	"jMp8CEg79YP5YKeFQazmyVdLDfZsA5nv34qm++m/AQAA///lGAMLEhQAAA==",
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
