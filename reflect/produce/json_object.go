package gojson

type JsonObjectType int

const (
	JsonObjectTypeBool JsonObjectType = iota
	JsonObjectTypeString
	JsonObjectTypeInt32
	JsonObjectTypeInt64
	JsonObjectTypeObject
	JsonObjectTypeArray
)

type JsonObject struct {
	data map[string]interface{}
	//dataType

}

func (j *JsonObject) GetString(key string) string {
	return ""
}
func FromBytes(data []byte) (*JsonObject, error) {
	return nil, nil
}
