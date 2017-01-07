package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var metadataGetRequestSchema *gojsonschema.Schema

func init() {
	var err error
	metadataGetRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "registry": {
      "type": "string"
    },
    "type": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *MetadataGetRequest) IsValid() (*gojsonschema.Result, error) {
	return metadataGetRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *MetadataGetRequest) IsRequest() {}

func (m *MetadataGetResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
