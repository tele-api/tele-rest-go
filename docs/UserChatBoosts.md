# UserChatBoosts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boosts** | [**[]ChatBoost**](ChatBoost.md) | The list of boosts added to the chat by the user | 

## Methods

### NewUserChatBoosts

`func NewUserChatBoosts(boosts []ChatBoost, ) *UserChatBoosts`

NewUserChatBoosts instantiates a new UserChatBoosts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserChatBoostsWithDefaults

`func NewUserChatBoostsWithDefaults() *UserChatBoosts`

NewUserChatBoostsWithDefaults instantiates a new UserChatBoosts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoosts

`func (o *UserChatBoosts) GetBoosts() []ChatBoost`

GetBoosts returns the Boosts field if non-nil, zero value otherwise.

### GetBoostsOk

`func (o *UserChatBoosts) GetBoostsOk() (*[]ChatBoost, bool)`

GetBoostsOk returns a tuple with the Boosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoosts

`func (o *UserChatBoosts) SetBoosts(v []ChatBoost)`

SetBoosts sets Boosts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


