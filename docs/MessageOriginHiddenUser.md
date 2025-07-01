# MessageOriginHiddenUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the message origin, always “hidden\\_user” | [default to "hidden_user"]
**Date** | **int32** | Date the message was sent originally in Unix time | 
**SenderUserName** | **string** | Name of the user that sent the message originally | 

## Methods

### NewMessageOriginHiddenUser

`func NewMessageOriginHiddenUser(type_ string, date int32, senderUserName string, ) *MessageOriginHiddenUser`

NewMessageOriginHiddenUser instantiates a new MessageOriginHiddenUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageOriginHiddenUserWithDefaults

`func NewMessageOriginHiddenUserWithDefaults() *MessageOriginHiddenUser`

NewMessageOriginHiddenUserWithDefaults instantiates a new MessageOriginHiddenUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MessageOriginHiddenUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageOriginHiddenUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageOriginHiddenUser) SetType(v string)`

SetType sets Type field to given value.


### GetDate

`func (o *MessageOriginHiddenUser) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *MessageOriginHiddenUser) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *MessageOriginHiddenUser) SetDate(v int32)`

SetDate sets Date field to given value.


### GetSenderUserName

`func (o *MessageOriginHiddenUser) GetSenderUserName() string`

GetSenderUserName returns the SenderUserName field if non-nil, zero value otherwise.

### GetSenderUserNameOk

`func (o *MessageOriginHiddenUser) GetSenderUserNameOk() (*string, bool)`

GetSenderUserNameOk returns a tuple with the SenderUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderUserName

`func (o *MessageOriginHiddenUser) SetSenderUserName(v string)`

SetSenderUserName sets SenderUserName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


