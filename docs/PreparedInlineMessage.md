# PreparedInlineMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the prepared message | 
**ExpirationDate** | **int32** | Expiration date of the prepared message, in Unix time. Expired prepared messages can no longer be used | 

## Methods

### NewPreparedInlineMessage

`func NewPreparedInlineMessage(id string, expirationDate int32, ) *PreparedInlineMessage`

NewPreparedInlineMessage instantiates a new PreparedInlineMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreparedInlineMessageWithDefaults

`func NewPreparedInlineMessageWithDefaults() *PreparedInlineMessage`

NewPreparedInlineMessageWithDefaults instantiates a new PreparedInlineMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PreparedInlineMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PreparedInlineMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PreparedInlineMessage) SetId(v string)`

SetId sets Id field to given value.


### GetExpirationDate

`func (o *PreparedInlineMessage) GetExpirationDate() int32`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *PreparedInlineMessage) GetExpirationDateOk() (*int32, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *PreparedInlineMessage) SetExpirationDate(v int32)`

SetExpirationDate sets ExpirationDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


