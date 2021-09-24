package models

//GenericModel is the struct for all the models which we create in the IOSXE
type GenericModel struct {
	JSONPayload string `json:"json_payload"`
}

//ToMap is the function to map the jsonpayload with the map[string]interface
func (gm *GenericModel) ToMap() (map[string]interface{}, error) {
	gmAttrMap := make(map[string]interface{})

	if gm.JSONPayload != "" {
		gmAttrMap["json_payload"] = gm.JSONPayload
	}
	return gmAttrMap, nil
}
