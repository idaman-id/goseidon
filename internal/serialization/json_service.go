package serialization

import "encoding/json"

type jsonService struct {
}

func (s *jsonService) Encode(i interface{}) ([]byte, error) {
	data, err := json.Marshal(i)
	return data, err
}

func (s *jsonService) Decode(i []byte, o interface{}) error {
	return json.Unmarshal(i, o)
}

func NewJsonSerialization() *jsonService {
	return &jsonService{}
}
