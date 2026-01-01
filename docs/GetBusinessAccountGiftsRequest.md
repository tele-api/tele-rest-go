# GetBusinessAccountGiftsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**ExcludeUnsaved** | Pointer to **bool** | Pass *True* to exclude gifts that aren&#39;t saved to the account&#39;s profile page | [optional] 
**ExcludeSaved** | Pointer to **bool** | Pass *True* to exclude gifts that are saved to the account&#39;s profile page | [optional] 
**ExcludeUnlimited** | Pointer to **bool** | Pass *True* to exclude gifts that can be purchased an unlimited number of times | [optional] 
**ExcludeLimitedUpgradable** | Pointer to **bool** | Pass *True* to exclude gifts that can be purchased a limited number of times and can be upgraded to unique | [optional] 
**ExcludeLimitedNonUpgradable** | Pointer to **bool** | Pass *True* to exclude gifts that can be purchased a limited number of times and can&#39;t be upgraded to unique | [optional] 
**ExcludeUnique** | Pointer to **bool** | Pass *True* to exclude unique gifts | [optional] 
**ExcludeFromBlockchain** | Pointer to **bool** | Pass *True* to exclude gifts that were assigned from the TON blockchain and can&#39;t be resold or transferred in Telegram | [optional] 
**SortByPrice** | Pointer to **bool** | Pass *True* to sort results by gift price instead of send date. Sorting is applied before pagination. | [optional] 
**Offset** | Pointer to **string** | Offset of the first entry to return as received from the previous request; use empty string to get the first chunk of results | [optional] 
**Limit** | Pointer to **int32** | The maximum number of gifts to be returned; 1-100. Defaults to 100 | [optional] [default to 100]

## Methods

### NewGetBusinessAccountGiftsRequest

`func NewGetBusinessAccountGiftsRequest(businessConnectionId string, ) *GetBusinessAccountGiftsRequest`

NewGetBusinessAccountGiftsRequest instantiates a new GetBusinessAccountGiftsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBusinessAccountGiftsRequestWithDefaults

`func NewGetBusinessAccountGiftsRequestWithDefaults() *GetBusinessAccountGiftsRequest`

NewGetBusinessAccountGiftsRequestWithDefaults instantiates a new GetBusinessAccountGiftsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *GetBusinessAccountGiftsRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *GetBusinessAccountGiftsRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *GetBusinessAccountGiftsRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetExcludeUnsaved

`func (o *GetBusinessAccountGiftsRequest) GetExcludeUnsaved() bool`

GetExcludeUnsaved returns the ExcludeUnsaved field if non-nil, zero value otherwise.

### GetExcludeUnsavedOk

`func (o *GetBusinessAccountGiftsRequest) GetExcludeUnsavedOk() (*bool, bool)`

GetExcludeUnsavedOk returns a tuple with the ExcludeUnsaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUnsaved

`func (o *GetBusinessAccountGiftsRequest) SetExcludeUnsaved(v bool)`

SetExcludeUnsaved sets ExcludeUnsaved field to given value.

### HasExcludeUnsaved

`func (o *GetBusinessAccountGiftsRequest) HasExcludeUnsaved() bool`

HasExcludeUnsaved returns a boolean if a field has been set.

### GetExcludeSaved

`func (o *GetBusinessAccountGiftsRequest) GetExcludeSaved() bool`

GetExcludeSaved returns the ExcludeSaved field if non-nil, zero value otherwise.

### GetExcludeSavedOk

`func (o *GetBusinessAccountGiftsRequest) GetExcludeSavedOk() (*bool, bool)`

GetExcludeSavedOk returns a tuple with the ExcludeSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSaved

`func (o *GetBusinessAccountGiftsRequest) SetExcludeSaved(v bool)`

SetExcludeSaved sets ExcludeSaved field to given value.

### HasExcludeSaved

`func (o *GetBusinessAccountGiftsRequest) HasExcludeSaved() bool`

HasExcludeSaved returns a boolean if a field has been set.

### GetExcludeUnlimited

`func (o *GetBusinessAccountGiftsRequest) GetExcludeUnlimited() bool`

GetExcludeUnlimited returns the ExcludeUnlimited field if non-nil, zero value otherwise.

### GetExcludeUnlimitedOk

`func (o *GetBusinessAccountGiftsRequest) GetExcludeUnlimitedOk() (*bool, bool)`

GetExcludeUnlimitedOk returns a tuple with the ExcludeUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUnlimited

`func (o *GetBusinessAccountGiftsRequest) SetExcludeUnlimited(v bool)`

SetExcludeUnlimited sets ExcludeUnlimited field to given value.

### HasExcludeUnlimited

`func (o *GetBusinessAccountGiftsRequest) HasExcludeUnlimited() bool`

HasExcludeUnlimited returns a boolean if a field has been set.

### GetExcludeLimitedUpgradable

`func (o *GetBusinessAccountGiftsRequest) GetExcludeLimitedUpgradable() bool`

GetExcludeLimitedUpgradable returns the ExcludeLimitedUpgradable field if non-nil, zero value otherwise.

### GetExcludeLimitedUpgradableOk

`func (o *GetBusinessAccountGiftsRequest) GetExcludeLimitedUpgradableOk() (*bool, bool)`

GetExcludeLimitedUpgradableOk returns a tuple with the ExcludeLimitedUpgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLimitedUpgradable

`func (o *GetBusinessAccountGiftsRequest) SetExcludeLimitedUpgradable(v bool)`

SetExcludeLimitedUpgradable sets ExcludeLimitedUpgradable field to given value.

### HasExcludeLimitedUpgradable

`func (o *GetBusinessAccountGiftsRequest) HasExcludeLimitedUpgradable() bool`

HasExcludeLimitedUpgradable returns a boolean if a field has been set.

### GetExcludeLimitedNonUpgradable

`func (o *GetBusinessAccountGiftsRequest) GetExcludeLimitedNonUpgradable() bool`

GetExcludeLimitedNonUpgradable returns the ExcludeLimitedNonUpgradable field if non-nil, zero value otherwise.

### GetExcludeLimitedNonUpgradableOk

`func (o *GetBusinessAccountGiftsRequest) GetExcludeLimitedNonUpgradableOk() (*bool, bool)`

GetExcludeLimitedNonUpgradableOk returns a tuple with the ExcludeLimitedNonUpgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLimitedNonUpgradable

`func (o *GetBusinessAccountGiftsRequest) SetExcludeLimitedNonUpgradable(v bool)`

SetExcludeLimitedNonUpgradable sets ExcludeLimitedNonUpgradable field to given value.

### HasExcludeLimitedNonUpgradable

`func (o *GetBusinessAccountGiftsRequest) HasExcludeLimitedNonUpgradable() bool`

HasExcludeLimitedNonUpgradable returns a boolean if a field has been set.

### GetExcludeUnique

`func (o *GetBusinessAccountGiftsRequest) GetExcludeUnique() bool`

GetExcludeUnique returns the ExcludeUnique field if non-nil, zero value otherwise.

### GetExcludeUniqueOk

`func (o *GetBusinessAccountGiftsRequest) GetExcludeUniqueOk() (*bool, bool)`

GetExcludeUniqueOk returns a tuple with the ExcludeUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUnique

`func (o *GetBusinessAccountGiftsRequest) SetExcludeUnique(v bool)`

SetExcludeUnique sets ExcludeUnique field to given value.

### HasExcludeUnique

`func (o *GetBusinessAccountGiftsRequest) HasExcludeUnique() bool`

HasExcludeUnique returns a boolean if a field has been set.

### GetExcludeFromBlockchain

`func (o *GetBusinessAccountGiftsRequest) GetExcludeFromBlockchain() bool`

GetExcludeFromBlockchain returns the ExcludeFromBlockchain field if non-nil, zero value otherwise.

### GetExcludeFromBlockchainOk

`func (o *GetBusinessAccountGiftsRequest) GetExcludeFromBlockchainOk() (*bool, bool)`

GetExcludeFromBlockchainOk returns a tuple with the ExcludeFromBlockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromBlockchain

`func (o *GetBusinessAccountGiftsRequest) SetExcludeFromBlockchain(v bool)`

SetExcludeFromBlockchain sets ExcludeFromBlockchain field to given value.

### HasExcludeFromBlockchain

`func (o *GetBusinessAccountGiftsRequest) HasExcludeFromBlockchain() bool`

HasExcludeFromBlockchain returns a boolean if a field has been set.

### GetSortByPrice

`func (o *GetBusinessAccountGiftsRequest) GetSortByPrice() bool`

GetSortByPrice returns the SortByPrice field if non-nil, zero value otherwise.

### GetSortByPriceOk

`func (o *GetBusinessAccountGiftsRequest) GetSortByPriceOk() (*bool, bool)`

GetSortByPriceOk returns a tuple with the SortByPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortByPrice

`func (o *GetBusinessAccountGiftsRequest) SetSortByPrice(v bool)`

SetSortByPrice sets SortByPrice field to given value.

### HasSortByPrice

`func (o *GetBusinessAccountGiftsRequest) HasSortByPrice() bool`

HasSortByPrice returns a boolean if a field has been set.

### GetOffset

`func (o *GetBusinessAccountGiftsRequest) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetBusinessAccountGiftsRequest) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetBusinessAccountGiftsRequest) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetBusinessAccountGiftsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *GetBusinessAccountGiftsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetBusinessAccountGiftsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetBusinessAccountGiftsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetBusinessAccountGiftsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


