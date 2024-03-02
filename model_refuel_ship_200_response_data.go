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

// checks if the RefuelShip200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefuelShip200ResponseData{}

// RefuelShip200ResponseData struct for RefuelShip200ResponseData
type RefuelShip200ResponseData struct {
	Agent Agent `json:"agent"`
	Fuel ShipFuel `json:"fuel"`
	Transaction MarketTransaction `json:"transaction"`
}

type _RefuelShip200ResponseData RefuelShip200ResponseData

// NewRefuelShip200ResponseData instantiates a new RefuelShip200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefuelShip200ResponseData(agent Agent, fuel ShipFuel, transaction MarketTransaction) *RefuelShip200ResponseData {
	this := RefuelShip200ResponseData{}
	this.Agent = agent
	this.Fuel = fuel
	this.Transaction = transaction
	return &this
}

// NewRefuelShip200ResponseDataWithDefaults instantiates a new RefuelShip200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefuelShip200ResponseDataWithDefaults() *RefuelShip200ResponseData {
	this := RefuelShip200ResponseData{}
	return &this
}

// GetAgent returns the Agent field value
func (o *RefuelShip200ResponseData) GetAgent() Agent {
	if o == nil {
		var ret Agent
		return ret
	}

	return o.Agent
}

// GetAgentOk returns a tuple with the Agent field value
// and a boolean to check if the value has been set.
func (o *RefuelShip200ResponseData) GetAgentOk() (*Agent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agent, true
}

// SetAgent sets field value
func (o *RefuelShip200ResponseData) SetAgent(v Agent) {
	o.Agent = v
}

// GetFuel returns the Fuel field value
func (o *RefuelShip200ResponseData) GetFuel() ShipFuel {
	if o == nil {
		var ret ShipFuel
		return ret
	}

	return o.Fuel
}

// GetFuelOk returns a tuple with the Fuel field value
// and a boolean to check if the value has been set.
func (o *RefuelShip200ResponseData) GetFuelOk() (*ShipFuel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fuel, true
}

// SetFuel sets field value
func (o *RefuelShip200ResponseData) SetFuel(v ShipFuel) {
	o.Fuel = v
}

// GetTransaction returns the Transaction field value
func (o *RefuelShip200ResponseData) GetTransaction() MarketTransaction {
	if o == nil {
		var ret MarketTransaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *RefuelShip200ResponseData) GetTransactionOk() (*MarketTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *RefuelShip200ResponseData) SetTransaction(v MarketTransaction) {
	o.Transaction = v
}

func (o RefuelShip200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefuelShip200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent"] = o.Agent
	toSerialize["fuel"] = o.Fuel
	toSerialize["transaction"] = o.Transaction
	return toSerialize, nil
}

func (o *RefuelShip200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agent",
		"fuel",
		"transaction",
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

	varRefuelShip200ResponseData := _RefuelShip200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRefuelShip200ResponseData)

	if err != nil {
		return err
	}

	*o = RefuelShip200ResponseData(varRefuelShip200ResponseData)

	return err
}

type NullableRefuelShip200ResponseData struct {
	value *RefuelShip200ResponseData
	isSet bool
}

func (v NullableRefuelShip200ResponseData) Get() *RefuelShip200ResponseData {
	return v.value
}

func (v *NullableRefuelShip200ResponseData) Set(val *RefuelShip200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableRefuelShip200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableRefuelShip200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefuelShip200ResponseData(val *RefuelShip200ResponseData) *NullableRefuelShip200ResponseData {
	return &NullableRefuelShip200ResponseData{value: val, isSet: true}
}

func (v NullableRefuelShip200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefuelShip200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


