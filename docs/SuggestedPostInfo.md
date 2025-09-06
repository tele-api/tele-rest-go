# SuggestedPostInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | State of the suggested post. Currently, it can be one of “pending”, “approved”, “declined”. | 
**Price** | Pointer to [**SuggestedPostPrice**](SuggestedPostPrice.md) |  | [optional] 
**SendDate** | Pointer to **int32** | *Optional*. Proposed send date of the post. If the field is omitted, then the post can be published at any time within 30 days at the sole discretion of the user or administrator who approves it. | [optional] 

## Methods

### NewSuggestedPostInfo

`func NewSuggestedPostInfo(state string, ) *SuggestedPostInfo`

NewSuggestedPostInfo instantiates a new SuggestedPostInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuggestedPostInfoWithDefaults

`func NewSuggestedPostInfoWithDefaults() *SuggestedPostInfo`

NewSuggestedPostInfoWithDefaults instantiates a new SuggestedPostInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SuggestedPostInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SuggestedPostInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SuggestedPostInfo) SetState(v string)`

SetState sets State field to given value.


### GetPrice

`func (o *SuggestedPostInfo) GetPrice() SuggestedPostPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *SuggestedPostInfo) GetPriceOk() (*SuggestedPostPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *SuggestedPostInfo) SetPrice(v SuggestedPostPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *SuggestedPostInfo) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSendDate

`func (o *SuggestedPostInfo) GetSendDate() int32`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *SuggestedPostInfo) GetSendDateOk() (*int32, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *SuggestedPostInfo) SetSendDate(v int32)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *SuggestedPostInfo) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


