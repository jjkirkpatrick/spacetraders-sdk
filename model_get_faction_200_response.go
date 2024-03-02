/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spacetraders_sdk

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GetFaction200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFaction200Response{}

// GetFaction200Response struct for GetFaction200Response
type GetFaction200Response struct {
	Data Faction `json:"data"`
}

type _GetFaction200Response GetFaction200Response

// NewGetFaction200Response instantiates a new GetFaction200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFaction200Response(data Faction) *GetFaction200Response {
	this := GetFaction200Response{}
	this.Data = data
	return &this
}

// NewGetFaction200ResponseWithDefaults instantiates a new GetFaction200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFaction200ResponseWithDefaults() *GetFaction200Response {
	this := GetFaction200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetFaction200Response) GetData() Faction {
	if o == nil {
		var ret Faction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetFaction200Response) GetDataOk() (*Faction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetFaction200Response) SetData(v Faction) {
	o.Data = v
}

func (o GetFaction200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFaction200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetFaction200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetFaction200Response := _GetFaction200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetFaction200Response)

	if err != nil {
		return err
	}

	*o = GetFaction200Response(varGetFaction200Response)

	return err
}

type NullableGetFaction200Response struct {
	value *GetFaction200Response
	isSet bool
}

func (v NullableGetFaction200Response) Get() *GetFaction200Response {
	return v.value
}

func (v *NullableGetFaction200Response) Set(val *GetFaction200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFaction200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFaction200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFaction200Response(val *GetFaction200Response) *NullableGetFaction200Response {
	return &NullableGetFaction200Response{value: val, isSet: true}
}

func (v NullableGetFaction200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFaction200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


