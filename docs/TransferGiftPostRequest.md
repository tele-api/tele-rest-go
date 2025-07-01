# TransferGiftPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**OwnedGiftId** | **string** | Unique identifier of the regular gift that should be transferred | 
**NewOwnerChatId** | **int32** | Unique identifier of the chat which will own the gift. The chat must be active in the last 24 hours. | 
**StarCount** | Pointer to **int32** | The amount of Telegram Stars that will be paid for the transfer from the business account balance. If positive, then the *can\\_transfer\\_stars* business bot right is required. | [optional] 

## Methods

### NewTransferGiftPostRequest

`func NewTransferGiftPostRequest(businessConnectionId string, ownedGiftId string, newOwnerChatId int32, ) *TransferGiftPostRequest`

NewTransferGiftPostRequest instantiates a new TransferGiftPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferGiftPostRequestWithDefaults

`func NewTransferGiftPostRequestWithDefaults() *TransferGiftPostRequest`

NewTransferGiftPostRequestWithDefaults instantiates a new TransferGiftPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *TransferGiftPostRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *TransferGiftPostRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *TransferGiftPostRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetOwnedGiftId

`func (o *TransferGiftPostRequest) GetOwnedGiftId() string`

GetOwnedGiftId returns the OwnedGiftId field if non-nil, zero value otherwise.

### GetOwnedGiftIdOk

`func (o *TransferGiftPostRequest) GetOwnedGiftIdOk() (*string, bool)`

GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedGiftId

`func (o *TransferGiftPostRequest) SetOwnedGiftId(v string)`

SetOwnedGiftId sets OwnedGiftId field to given value.


### GetNewOwnerChatId

`func (o *TransferGiftPostRequest) GetNewOwnerChatId() int32`

GetNewOwnerChatId returns the NewOwnerChatId field if non-nil, zero value otherwise.

### GetNewOwnerChatIdOk

`func (o *TransferGiftPostRequest) GetNewOwnerChatIdOk() (*int32, bool)`

GetNewOwnerChatIdOk returns a tuple with the NewOwnerChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOwnerChatId

`func (o *TransferGiftPostRequest) SetNewOwnerChatId(v int32)`

SetNewOwnerChatId sets NewOwnerChatId field to given value.


### GetStarCount

`func (o *TransferGiftPostRequest) GetStarCount() int32`

GetStarCount returns the StarCount field if non-nil, zero value otherwise.

### GetStarCountOk

`func (o *TransferGiftPostRequest) GetStarCountOk() (*int32, bool)`

GetStarCountOk returns a tuple with the StarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarCount

`func (o *TransferGiftPostRequest) SetStarCount(v int32)`

SetStarCount sets StarCount field to given value.

### HasStarCount

`func (o *TransferGiftPostRequest) HasStarCount() bool`

HasStarCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


