package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var helloRequestSchema *gojsonschema.Schema

func init() {
	var err error
	helloRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *HelloRequest) IsValid() (*gojsonschema.Result, error) {
	return helloRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *HelloRequest) IsRequest() {}

func (m *HelloResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
