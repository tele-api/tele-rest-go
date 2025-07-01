# ChatBoostSourcePremium

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Source of the boost, always “premium” | [default to "premium"]
**User** | [**User**](User.md) |  | 

## Methods

### NewChatBoostSourcePremium

`func NewChatBoostSourcePremium(source string, user User, ) *ChatBoostSourcePremium`

NewChatBoostSourcePremium instantiates a new ChatBoostSourcePremium object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatBoostSourcePremiumWithDefaults

`func NewChatBoostSourcePremiumWithDefaults() *ChatBoostSourcePremium`

NewChatBoostSourcePremiumWithDefaults instantiates a new ChatBoostSourcePremium object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *ChatBoostSourcePremium) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ChatBoostSourcePremium) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ChatBoostSourcePremium) SetSource(v string)`

SetSource sets Source field to given value.


### GetUser

`func (o *ChatBoostSourcePremium) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ChatBoostSourcePremium) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ChatBoostSourcePremium) SetUser(v User)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


