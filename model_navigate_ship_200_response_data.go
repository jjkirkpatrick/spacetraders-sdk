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

// checks if the NavigateShip200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NavigateShip200ResponseData{}

// NavigateShip200ResponseData struct for NavigateShip200ResponseData
type NavigateShip200ResponseData struct {
	Fuel ShipFuel `json:"fuel"`
	Nav ShipNav `json:"nav"`
}

type _NavigateShip200ResponseData NavigateShip200ResponseData

// NewNavigateShip200ResponseData instantiates a new NavigateShip200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNavigateShip200ResponseData(fuel ShipFuel, nav ShipNav) *NavigateShip200ResponseData {
	this := NavigateShip200ResponseData{}
	this.Fuel = fuel
	this.Nav = nav
	return &this
}

// NewNavigateShip200ResponseDataWithDefaults instantiates a new NavigateShip200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNavigateShip200ResponseDataWithDefaults() *NavigateShip200ResponseData {
	this := NavigateShip200ResponseData{}
	return &this
}

// GetFuel returns the Fuel field value
func (o *NavigateShip200ResponseData) GetFuel() ShipFuel {
	if o == nil {
		var ret ShipFuel
		return ret
	}

	return o.Fuel
}

// GetFuelOk returns a tuple with the Fuel field value
// and a boolean to check if the value has been set.
func (o *NavigateShip200ResponseData) GetFuelOk() (*ShipFuel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fuel, true
}

// SetFuel sets field value
func (o *NavigateShip200ResponseData) SetFuel(v ShipFuel) {
	o.Fuel = v
}

// GetNav returns the Nav field value
func (o *NavigateShip200ResponseData) GetNav() ShipNav {
	if o == nil {
		var ret ShipNav
		return ret
	}

	return o.Nav
}

// GetNavOk returns a tuple with the Nav field value
// and a boolean to check if the value has been set.
func (o *NavigateShip200ResponseData) GetNavOk() (*ShipNav, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nav, true
}

// SetNav sets field value
func (o *NavigateShip200ResponseData) SetNav(v ShipNav) {
	o.Nav = v
}

func (o NavigateShip200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NavigateShip200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fuel"] = o.Fuel
	toSerialize["nav"] = o.Nav
	return toSerialize, nil
}

func (o *NavigateShip200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fuel",
		"nav",
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

	varNavigateShip200ResponseData := _NavigateShip200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNavigateShip200ResponseData)

	if err != nil {
		return err
	}

	*o = NavigateShip200ResponseData(varNavigateShip200ResponseData)

	return err
}

type NullableNavigateShip200ResponseData struct {
	value *NavigateShip200ResponseData
	isSet bool
}

func (v NullableNavigateShip200ResponseData) Get() *NavigateShip200ResponseData {
	return v.value
}

func (v *NullableNavigateShip200ResponseData) Set(val *NavigateShip200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableNavigateShip200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableNavigateShip200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNavigateShip200ResponseData(val *NavigateShip200ResponseData) *NullableNavigateShip200ResponseData {
	return &NullableNavigateShip200ResponseData{value: val, isSet: true}
}

func (v NullableNavigateShip200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNavigateShip200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


