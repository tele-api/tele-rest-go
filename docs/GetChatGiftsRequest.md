# GetChatGiftsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | [**SendMessageRequestChatId**](SendMessageRequestChatId.md) |  | 
**ExcludeUnsaved** | Pointer to **bool** | Pass *True* to exclude gifts that aren&#39;t saved to the chat&#39;s profile page. Always *True*, unless the bot has the *can\\_post\\_messages* administrator right in the channel. | [optional] 
**ExcludeSaved** | Pointer to **bool** | Pass *True* to exclude gifts that are saved to the chat&#39;s profile page. Always *False*, unless the bot has the *can\\_post\\_messages* administrator right in the channel. | [optional] 
**ExcludeUnlimited** | Pointer to **bool** | Pass *True* to exclude gifts that can be purchased an unlimited number of times | [optional] 
**ExcludeLimitedUpgradable** | Pointer to **bool** | Pass *True* to exclude gifts that can be purchased a limited number of times and can be upgraded to unique | [optional] 
**ExcludeLimitedNonUpgradable** | Pointer to **bool** | Pass *True* to exclude gifts that can be purchased a limited number of times and can&#39;t be upgraded to unique | [optional] 
**ExcludeFromBlockchain** | Pointer to **bool** | Pass *True* to exclude gifts that were assigned from the TON blockchain and can&#39;t be resold or transferred in Telegram | [optional] 
**ExcludeUnique** | Pointer to **bool** | Pass *True* to exclude unique gifts | [optional] 
**SortByPrice** | Pointer to **bool** | Pass *True* to sort results by gift price instead of send date. Sorting is applied before pagination. | [optional] 
**Offset** | Pointer to **string** | Offset of the first entry to return as received from the previous request; use an empty string to get the first chunk of results | [optional] 
**Limit** | Pointer to **int32** | The maximum number of gifts to be returned; 1-100. Defaults to 100 | [optional] [default to 100]

## Methods

### NewGetChatGiftsRequest

`func NewGetChatGiftsRequest(chatId SendMessageRequestChatId, ) *GetChatGiftsRequest`

NewGetChatGiftsRequest instantiates a new GetChatGiftsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChatGiftsRequestWithDefaults

`func NewGetChatGiftsRequestWithDefaults() *GetChatGiftsRequest`

NewGetChatGiftsRequestWithDefaults instantiates a new GetChatGiftsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *GetChatGiftsRequest) GetChatId() SendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *GetChatGiftsRequest) GetChatIdOk() (*SendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *GetChatGiftsRequest) SetChatId(v SendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetExcludeUnsaved

`func (o *GetChatGiftsRequest) GetExcludeUnsaved() bool`

GetExcludeUnsaved returns the ExcludeUnsaved field if non-nil, zero value otherwise.

### GetExcludeUnsavedOk

`func (o *GetChatGiftsRequest) GetExcludeUnsavedOk() (*bool, bool)`

GetExcludeUnsavedOk returns a tuple with the ExcludeUnsaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUnsaved

`func (o *GetChatGiftsRequest) SetExcludeUnsaved(v bool)`

SetExcludeUnsaved sets ExcludeUnsaved field to given value.

### HasExcludeUnsaved

`func (o *GetChatGiftsRequest) HasExcludeUnsaved() bool`

HasExcludeUnsaved returns a boolean if a field has been set.

### GetExcludeSaved

`func (o *GetChatGiftsRequest) GetExcludeSaved() bool`

GetExcludeSaved returns the ExcludeSaved field if non-nil, zero value otherwise.

### GetExcludeSavedOk

`func (o *GetChatGiftsRequest) GetExcludeSavedOk() (*bool, bool)`

GetExcludeSavedOk returns a tuple with the ExcludeSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSaved

`func (o *GetChatGiftsRequest) SetExcludeSaved(v bool)`

SetExcludeSaved sets ExcludeSaved field to given value.

### HasExcludeSaved

`func (o *GetChatGiftsRequest) HasExcludeSaved() bool`

HasExcludeSaved returns a boolean if a field has been set.

### GetExcludeUnlimited

`func (o *GetChatGiftsRequest) GetExcludeUnlimited() bool`

GetExcludeUnlimited returns the ExcludeUnlimited field if non-nil, zero value otherwise.

### GetExcludeUnlimitedOk

`func (o *GetChatGiftsRequest) GetExcludeUnlimitedOk() (*bool, bool)`

GetExcludeUnlimitedOk returns a tuple with the ExcludeUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUnlimited

`func (o *GetChatGiftsRequest) SetExcludeUnlimited(v bool)`

SetExcludeUnlimited sets ExcludeUnlimited field to given value.

### HasExcludeUnlimited

`func (o *GetChatGiftsRequest) HasExcludeUnlimited() bool`

HasExcludeUnlimited returns a boolean if a field has been set.

### GetExcludeLimitedUpgradable

`func (o *GetChatGiftsRequest) GetExcludeLimitedUpgradable() bool`

GetExcludeLimitedUpgradable returns the ExcludeLimitedUpgradable field if non-nil, zero value otherwise.

### GetExcludeLimitedUpgradableOk

`func (o *GetChatGiftsRequest) GetExcludeLimitedUpgradableOk() (*bool, bool)`

GetExcludeLimitedUpgradableOk returns a tuple with the ExcludeLimitedUpgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLimitedUpgradable

`func (o *GetChatGiftsRequest) SetExcludeLimitedUpgradable(v bool)`

SetExcludeLimitedUpgradable sets ExcludeLimitedUpgradable field to given value.

### HasExcludeLimitedUpgradable

`func (o *GetChatGiftsRequest) HasExcludeLimitedUpgradable() bool`

HasExcludeLimitedUpgradable returns a boolean if a field has been set.

### GetExcludeLimitedNonUpgradable

`func (o *GetChatGiftsRequest) GetExcludeLimitedNonUpgradable() bool`

GetExcludeLimitedNonUpgradable returns the ExcludeLimitedNonUpgradable field if non-nil, zero value otherwise.

### GetExcludeLimitedNonUpgradableOk

`func (o *GetChatGiftsRequest) GetExcludeLimitedNonUpgradableOk() (*bool, bool)`

GetExcludeLimitedNonUpgradableOk returns a tuple with the ExcludeLimitedNonUpgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLimitedNonUpgradable

`func (o *GetChatGiftsRequest) SetExcludeLimitedNonUpgradable(v bool)`

SetExcludeLimitedNonUpgradable sets ExcludeLimitedNonUpgradable field to given value.

### HasExcludeLimitedNonUpgradable

`func (o *GetChatGiftsRequest) HasExcludeLimitedNonUpgradable() bool`

HasExcludeLimitedNonUpgradable returns a boolean if a field has been set.

### GetExcludeFromBlockchain

`func (o *GetChatGiftsRequest) GetExcludeFromBlockchain() bool`

GetExcludeFromBlockchain returns the ExcludeFromBlockchain field if non-nil, zero value otherwise.

### GetExcludeFromBlockchainOk

`func (o *GetChatGiftsRequest) GetExcludeFromBlockchainOk() (*bool, bool)`

GetExcludeFromBlockchainOk returns a tuple with the ExcludeFromBlockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromBlockchain

`func (o *GetChatGiftsRequest) SetExcludeFromBlockchain(v bool)`

SetExcludeFromBlockchain sets ExcludeFromBlockchain field to given value.

### HasExcludeFromBlockchain

`func (o *GetChatGiftsRequest) HasExcludeFromBlockchain() bool`

HasExcludeFromBlockchain returns a boolean if a field has been set.

### GetExcludeUnique

`func (o *GetChatGiftsRequest) GetExcludeUnique() bool`

GetExcludeUnique returns the ExcludeUnique field if non-nil, zero value otherwise.

### GetExcludeUniqueOk

`func (o *GetChatGiftsRequest) GetExcludeUniqueOk() (*bool, bool)`

GetExcludeUniqueOk returns a tuple with the ExcludeUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUnique

`func (o *GetChatGiftsRequest) SetExcludeUnique(v bool)`

SetExcludeUnique sets ExcludeUnique field to given value.

### HasExcludeUnique

`func (o *GetChatGiftsRequest) HasExcludeUnique() bool`

HasExcludeUnique returns a boolean if a field has been set.

### GetSortByPrice

`func (o *GetChatGiftsRequest) GetSortByPrice() bool`

GetSortByPrice returns the SortByPrice field if non-nil, zero value otherwise.

### GetSortByPriceOk

`func (o *GetChatGiftsRequest) GetSortByPriceOk() (*bool, bool)`

GetSortByPriceOk returns a tuple with the SortByPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortByPrice

`func (o *GetChatGiftsRequest) SetSortByPrice(v bool)`

SetSortByPrice sets SortByPrice field to given value.

### HasSortByPrice

`func (o *GetChatGiftsRequest) HasSortByPrice() bool`

HasSortByPrice returns a boolean if a field has been set.

### GetOffset

`func (o *GetChatGiftsRequest) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetChatGiftsRequest) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetChatGiftsRequest) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetChatGiftsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *GetChatGiftsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetChatGiftsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetChatGiftsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetChatGiftsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


