# TransactionPartnerChat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the transaction partner, always “chat” | [default to "chat"]
**Chat** | [**Chat**](Chat.md) |  | 
**Gift** | Pointer to [**Gift**](Gift.md) |  | [optional] 

## Methods

### NewTransactionPartnerChat

`func NewTransactionPartnerChat(type_ string, chat Chat, ) *TransactionPartnerChat`

NewTransactionPartnerChat instantiates a new TransactionPartnerChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPartnerChatWithDefaults

`func NewTransactionPartnerChatWithDefaults() *TransactionPartnerChat`

NewTransactionPartnerChatWithDefaults instantiates a new TransactionPartnerChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionPartnerChat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionPartnerChat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionPartnerChat) SetType(v string)`

SetType sets Type field to given value.


### GetChat

`func (o *TransactionPartnerChat) GetChat() Chat`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *TransactionPartnerChat) GetChatOk() (*Chat, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *TransactionPartnerChat) SetChat(v Chat)`

SetChat sets Chat field to given value.


### GetGift

`func (o *TransactionPartnerChat) GetGift() Gift`

GetGift returns the Gift field if non-nil, zero value otherwise.

### GetGiftOk

`func (o *TransactionPartnerChat) GetGiftOk() (*Gift, bool)`

GetGiftOk returns a tuple with the Gift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGift

`func (o *TransactionPartnerChat) SetGift(v Gift)`

SetGift sets Gift field to given value.

### HasGift

`func (o *TransactionPartnerChat) HasGift() bool`

HasGift returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


