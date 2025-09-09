# SuggestedPostApproved

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuggestedPostMessage** | Pointer to [**Message**](Message.md) |  | [optional] 
**Price** | Pointer to [**SuggestedPostPrice**](SuggestedPostPrice.md) |  | [optional] 
**SendDate** | **int32** | Date when the post will be published | 

## Methods

### NewSuggestedPostApproved

`func NewSuggestedPostApproved(sendDate int32, ) *SuggestedPostApproved`

NewSuggestedPostApproved instantiates a new SuggestedPostApproved object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedPostApprovedWithDefaults

`func NewSuggestedPostApprovedWithDefaults() *SuggestedPostApproved`

NewSuggestedPostApprovedWithDefaults instantiates a new SuggestedPostApproved object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestedPostMessage

`func (o *SuggestedPostApproved) GetSuggestedPostMessage() Message`

GetSuggestedPostMessage returns the SuggestedPostMessage field if non-nil, zero value otherwise.

### GetSuggestedPostMessageOk

`func (o *SuggestedPostApproved) GetSuggestedPostMessageOk() (*Message, bool)`

GetSuggestedPostMessageOk returns a tuple with the SuggestedPostMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedPostMessage

`func (o *SuggestedPostApproved) SetSuggestedPostMessage(v Message)`

SetSuggestedPostMessage sets SuggestedPostMessage field to given value.

### HasSuggestedPostMessage

`func (o *SuggestedPostApproved) HasSuggestedPostMessage() bool`

HasSuggestedPostMessage returns a boolean if a field has been set.

### GetPrice

`func (o *SuggestedPostApproved) GetPrice() SuggestedPostPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SuggestedPostApproved) GetPriceOk() (*SuggestedPostPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SuggestedPostApproved) SetPrice(v SuggestedPostPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SuggestedPostApproved) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSendDate

`func (o *SuggestedPostApproved) GetSendDate() int32`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *SuggestedPostApproved) GetSendDateOk() (*int32, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *SuggestedPostApproved) SetSendDate(v int32)`

SetSendDate sets SendDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


