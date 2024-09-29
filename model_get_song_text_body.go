/*
Music Collection

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the GetSongTextBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSongTextBody{}

// GetSongTextBody struct for GetSongTextBody
type GetSongTextBody struct {
	Group string `json:"group"`
	Song  string `json:"song"`
	// Limit the number of verses returned
	Limit *int32 `json:"limit,omitempty"`
	// Offset for pagination
	Offset *int32 `json:"offset,omitempty"`
}

type _GetSongTextBody GetSongTextBody

// NewGetSongTextBody instantiates a new GetSongTextBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSongTextBody(group string, song string) *GetSongTextBody {
	this := GetSongTextBody{}
	this.Group = group
	this.Song = song
	var limit int32 = 5
	this.Limit = &limit
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// NewGetSongTextBodyWithDefaults instantiates a new GetSongTextBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSongTextBodyWithDefaults() *GetSongTextBody {
	this := GetSongTextBody{}
	var limit int32 = 5
	this.Limit = &limit
	var offset int32 = 0
	this.Offset = &offset
	return &this
}

// GetGroup returns the Group field value
func (o *GetSongTextBody) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *GetSongTextBody) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *GetSongTextBody) SetGroup(v string) {
	o.Group = v
}

// GetSong returns the Song field value
func (o *GetSongTextBody) GetSong() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Song
}

// GetSongOk returns a tuple with the Song field value
// and a boolean to check if the value has been set.
func (o *GetSongTextBody) GetSongOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Song, true
}

// SetSong sets field value
func (o *GetSongTextBody) SetSong(v string) {
	o.Song = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *GetSongTextBody) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSongTextBody) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *GetSongTextBody) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *GetSongTextBody) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *GetSongTextBody) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSongTextBody) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *GetSongTextBody) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *GetSongTextBody) SetOffset(v int32) {
	o.Offset = &v
}

func (o GetSongTextBody) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSongTextBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["group"] = o.Group
	toSerialize["song"] = o.Song
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

func (o *GetSongTextBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"group",
		"song",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetSongTextBody := _GetSongTextBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetSongTextBody)

	if err != nil {
		return err
	}

	*o = GetSongTextBody(varGetSongTextBody)

	return err
}

type NullableGetSongTextBody struct {
	value *GetSongTextBody
	isSet bool
}

func (v NullableGetSongTextBody) Get() *GetSongTextBody {
	return v.value
}

func (v *NullableGetSongTextBody) Set(val *GetSongTextBody) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSongTextBody) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSongTextBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSongTextBody(val *GetSongTextBody) *NullableGetSongTextBody {
	return &NullableGetSongTextBody{value: val, isSet: true}
}

func (v NullableGetSongTextBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSongTextBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
