package sfdc

type PicklistValue struct {
	Attributes *Attributes `json:"attributes"`
	Label      string      `json:"label"`
	ValidFor   []any       `json:"validFor"`
	Value      string      `json:"value"`
}

type PicklistValues struct {
	ControllerValues map[string]any  `json:"controllerValues"`
	DefaultValue     *PicklistValue  `json:"defaultValue"`
	ETag             string          `json:"eTag"`
	URL              string          `json:"url"`
	Values           []PicklistValue `json:"values"`
}

// Map returns a map of k:v pairs with labels as keys and values as values.
func (plv *PicklistValues) Map() map[string]string {
	m := make(map[string]string, len(plv.Values))
	for _, pv := range plv.Values {
		m[pv.Label] = pv.Value
	}
	return m
}

type StandardValueMetadataItem struct {
	Label string `json:"label"`
	Value string `json:"valueName"`
}

type StandardValueMetadata struct {
	Sorted        bool                        `json:"sorted"`
	StandardValue []StandardValueMetadataItem `json:"standardValue"`
	FullName      string                      `json:"FullName"`
	MasterLabel   string                      `json:"Industry"`
}

type StandardValue struct {
	Attributes *Attributes
	ID         string                 `json:"Id"`
	DurableID  string                 `json:"DurableId"`
	Metadata   *StandardValueMetadata `json:"Metadata"`
}

type StandardValueSet struct {
	Done           bool            `json:"done"`
	TotalSize      int             `json:"totalSize"`
	Records        []StandardValue `json:"records"`
	EntityTypeName string          `json:"entityTypeName"`
}

func (svs *StandardValueSet) Items() []map[string]string {
	count := 0
	for _, rec := range svs.Records {
		count += len(rec.Metadata.StandardValue)
	}
	s := make([]map[string]string, 0, count)
	for _, rec := range svs.Records {
		m := make(map[string]string)
		for _, meta := range rec.Metadata.StandardValue {
			m[meta.Label] = meta.Value
		}
		s = append(s, m)
	}
	return s
}
