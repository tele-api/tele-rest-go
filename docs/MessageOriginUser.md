# MessageOriginUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the message origin, always “user” | [default to "user"]
**Date** | **int32** | Date the message was sent originally in Unix time | 
**SenderUser** | [**User**](User.md) |  | 

## Methods

### NewMessageOriginUser

`func NewMessageOriginUser(type_ string, date int32, senderUser User, ) *MessageOriginUser`

NewMessageOriginUser instantiates a new MessageOriginUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageOriginUserWithDefaults

`func NewMessageOriginUserWithDefaults() *MessageOriginUser`

NewMessageOriginUserWithDefaults instantiates a new MessageOriginUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageOriginUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageOriginUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageOriginUser) SetType(v string)`

SetType sets Type field to given value.


### GetDate

`func (o *MessageOriginUser) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MessageOriginUser) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MessageOriginUser) SetDate(v int32)`

SetDate sets Date field to given value.


### GetSenderUser

`func (o *MessageOriginUser) GetSenderUser() User`

GetSenderUser returns the SenderUser field if non-nil, zero value otherwise.

### GetSenderUserOk

`func (o *MessageOriginUser) GetSenderUserOk() (*User, bool)`

GetSenderUserOk returns a tuple with the SenderUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUser

`func (o *MessageOriginUser) SetSenderUser(v User)`

SetSenderUser sets SenderUser field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


