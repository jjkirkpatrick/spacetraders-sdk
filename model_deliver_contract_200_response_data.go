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

// checks if the DeliverContract200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliverContract200ResponseData{}

// DeliverContract200ResponseData struct for DeliverContract200ResponseData
type DeliverContract200ResponseData struct {
	Contract Contract `json:"contract"`
	Cargo ShipCargo `json:"cargo"`
}

type _DeliverContract200ResponseData DeliverContract200ResponseData

// NewDeliverContract200ResponseData instantiates a new DeliverContract200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliverContract200ResponseData(contract Contract, cargo ShipCargo) *DeliverContract200ResponseData {
	this := DeliverContract200ResponseData{}
	this.Contract = contract
	this.Cargo = cargo
	return &this
}

// NewDeliverContract200ResponseDataWithDefaults instantiates a new DeliverContract200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliverContract200ResponseDataWithDefaults() *DeliverContract200ResponseData {
	this := DeliverContract200ResponseData{}
	return &this
}

// GetContract returns the Contract field value
func (o *DeliverContract200ResponseData) GetContract() Contract {
	if o == nil {
		var ret Contract
		return ret
	}

	return o.Contract
}

// GetContractOk returns a tuple with the Contract field value
// and a boolean to check if the value has been set.
func (o *DeliverContract200ResponseData) GetContractOk() (*Contract, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contract, true
}

// SetContract sets field value
func (o *DeliverContract200ResponseData) SetContract(v Contract) {
	o.Contract = v
}

// GetCargo returns the Cargo field value
func (o *DeliverContract200ResponseData) GetCargo() ShipCargo {
	if o == nil {
		var ret ShipCargo
		return ret
	}

	return o.Cargo
}

// GetCargoOk returns a tuple with the Cargo field value
// and a boolean to check if the value has been set.
func (o *DeliverContract200ResponseData) GetCargoOk() (*ShipCargo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cargo, true
}

// SetCargo sets field value
func (o *DeliverContract200ResponseData) SetCargo(v ShipCargo) {
	o.Cargo = v
}

func (o DeliverContract200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeliverContract200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contract"] = o.Contract
	toSerialize["cargo"] = o.Cargo
	return toSerialize, nil
}

func (o *DeliverContract200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"contract",
		"cargo",
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

	varDeliverContract200ResponseData := _DeliverContract200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeliverContract200ResponseData)

	if err != nil {
		return err
	}

	*o = DeliverContract200ResponseData(varDeliverContract200ResponseData)

	return err
}

type NullableDeliverContract200ResponseData struct {
	value *DeliverContract200ResponseData
	isSet bool
}

func (v NullableDeliverContract200ResponseData) Get() *DeliverContract200ResponseData {
	return v.value
}

func (v *NullableDeliverContract200ResponseData) Set(val *DeliverContract200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliverContract200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliverContract200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliverContract200ResponseData(val *DeliverContract200ResponseData) *NullableDeliverContract200ResponseData {
	return &NullableDeliverContract200ResponseData{value: val, isSet: true}
}

func (v NullableDeliverContract200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeliverContract200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


