# KeyboardButtonPollType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | *Optional*. If *quiz* is passed, the user will be allowed to create only polls in the quiz mode. If *regular* is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type. | [optional] 

## Methods

### NewKeyboardButtonPollType

`func NewKeyboardButtonPollType() *KeyboardButtonPollType`

NewKeyboardButtonPollType instantiates a new KeyboardButtonPollType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyboardButtonPollTypeWithDefaults

`func NewKeyboardButtonPollTypeWithDefaults() *KeyboardButtonPollType`

NewKeyboardButtonPollTypeWithDefaults instantiates a new KeyboardButtonPollType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *KeyboardButtonPollType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KeyboardButtonPollType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KeyboardButtonPollType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *KeyboardButtonPollType) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


