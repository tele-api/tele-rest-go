# TransferGiftRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**OwnedGiftId** | **string** | Unique identifier of the regular gift that should be transferred | 
**NewOwnerChatId** | **int32** | Unique identifier of the chat which will own the gift. The chat must be active in the last 24 hours. | 
**StarCount** | Pointer to **int32** | The amount of Telegram Stars that will be paid for the transfer from the business account balance. If positive, then the *can\\_transfer\\_stars* business bot right is required. | [optional] 

## Methods

### NewTransferGiftRequest

`func NewTransferGiftRequest(businessConnectionId string, ownedGiftId string, newOwnerChatId int32, ) *TransferGiftRequest`

NewTransferGiftRequest instantiates a new TransferGiftRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferGiftRequestWithDefaults

`func NewTransferGiftRequestWithDefaults() *TransferGiftRequest`

NewTransferGiftRequestWithDefaults instantiates a new TransferGiftRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *TransferGiftRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *TransferGiftRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *TransferGiftRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetOwnedGiftId

`func (o *TransferGiftRequest) GetOwnedGiftId() string`

GetOwnedGiftId returns the OwnedGiftId field if non-nil, zero value otherwise.

### GetOwnedGiftIdOk

`func (o *TransferGiftRequest) GetOwnedGiftIdOk() (*string, bool)`

GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedGiftId

`func (o *TransferGiftRequest) SetOwnedGiftId(v string)`

SetOwnedGiftId sets OwnedGiftId field to given value.


### GetNewOwnerChatId

`func (o *TransferGiftRequest) GetNewOwnerChatId() int32`

GetNewOwnerChatId returns the NewOwnerChatId field if non-nil, zero value otherwise.

### GetNewOwnerChatIdOk

`func (o *TransferGiftRequest) GetNewOwnerChatIdOk() (*int32, bool)`

GetNewOwnerChatIdOk returns a tuple with the NewOwnerChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewOwnerChatId

`func (o *TransferGiftRequest) SetNewOwnerChatId(v int32)`

SetNewOwnerChatId sets NewOwnerChatId field to given value.


### GetStarCount

`func (o *TransferGiftRequest) GetStarCount() int32`

GetStarCount returns the StarCount field if non-nil, zero value otherwise.

### GetStarCountOk

`func (o *TransferGiftRequest) GetStarCountOk() (*int32, bool)`

GetStarCountOk returns a tuple with the StarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarCount

`func (o *TransferGiftRequest) SetStarCount(v int32)`

SetStarCount sets StarCount field to given value.

### HasStarCount

`func (o *TransferGiftRequest) HasStarCount() bool`

HasStarCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


