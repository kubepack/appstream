package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/appscode/api/dtypes"
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var gitRequestSchema *gojsonschema.Schema
var dockerRequestSchema *gojsonschema.Schema

func init() {
	var err error
	gitRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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
	dockerRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
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

func (m *GitRequest) IsValid() (*gojsonschema.Result, error) {
	return gitRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *GitRequest) IsRequest() {}

func (m *DockerRequest) IsValid() (*gojsonschema.Result, error) {
	return dockerRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *DockerRequest) IsRequest() {}

func (m *GitResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
func (m *DockerResponse) SetStatus(s *dtypes.Status) {
	m.Status = s
}
