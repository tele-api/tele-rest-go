# GetUserGiftsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **int32** | Unique identifier of the user | 
**ExcludeUnlimited** | Pointer to **bool** | Pass *True* to exclude gifts that can be purchased an unlimited number of times | [optional] 
**ExcludeLimitedUpgradable** | Pointer to **bool** | Pass *True* to exclude gifts that can be purchased a limited number of times and can be upgraded to unique | [optional] 
**ExcludeLimitedNonUpgradable** | Pointer to **bool** | Pass *True* to exclude gifts that can be purchased a limited number of times and can&#39;t be upgraded to unique | [optional] 
**ExcludeFromBlockchain** | Pointer to **bool** | Pass *True* to exclude gifts that were assigned from the TON blockchain and can&#39;t be resold or transferred in Telegram | [optional] 
**ExcludeUnique** | Pointer to **bool** | Pass *True* to exclude unique gifts | [optional] 
**SortByPrice** | Pointer to **bool** | Pass *True* to sort results by gift price instead of send date. Sorting is applied before pagination. | [optional] 
**Offset** | Pointer to **string** | Offset of the first entry to return as received from the previous request; use an empty string to get the first chunk of results | [optional] 
**Limit** | Pointer to **int32** | The maximum number of gifts to be returned; 1-100. Defaults to 100 | [optional] [default to 100]

## Methods

### NewGetUserGiftsRequest

`func NewGetUserGiftsRequest(userId int32, ) *GetUserGiftsRequest`

NewGetUserGiftsRequest instantiates a new GetUserGiftsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserGiftsRequestWithDefaults

`func NewGetUserGiftsRequestWithDefaults() *GetUserGiftsRequest`

NewGetUserGiftsRequestWithDefaults instantiates a new GetUserGiftsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *GetUserGiftsRequest) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetUserGiftsRequest) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetUserGiftsRequest) SetUserId(v int32)`

SetUserId sets UserId field to given value.


### GetExcludeUnlimited

`func (o *GetUserGiftsRequest) GetExcludeUnlimited() bool`

GetExcludeUnlimited returns the ExcludeUnlimited field if non-nil, zero value otherwise.

### GetExcludeUnlimitedOk

`func (o *GetUserGiftsRequest) GetExcludeUnlimitedOk() (*bool, bool)`

GetExcludeUnlimitedOk returns a tuple with the ExcludeUnlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUnlimited

`func (o *GetUserGiftsRequest) SetExcludeUnlimited(v bool)`

SetExcludeUnlimited sets ExcludeUnlimited field to given value.

### HasExcludeUnlimited

`func (o *GetUserGiftsRequest) HasExcludeUnlimited() bool`

HasExcludeUnlimited returns a boolean if a field has been set.

### GetExcludeLimitedUpgradable

`func (o *GetUserGiftsRequest) GetExcludeLimitedUpgradable() bool`

GetExcludeLimitedUpgradable returns the ExcludeLimitedUpgradable field if non-nil, zero value otherwise.

### GetExcludeLimitedUpgradableOk

`func (o *GetUserGiftsRequest) GetExcludeLimitedUpgradableOk() (*bool, bool)`

GetExcludeLimitedUpgradableOk returns a tuple with the ExcludeLimitedUpgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLimitedUpgradable

`func (o *GetUserGiftsRequest) SetExcludeLimitedUpgradable(v bool)`

SetExcludeLimitedUpgradable sets ExcludeLimitedUpgradable field to given value.

### HasExcludeLimitedUpgradable

`func (o *GetUserGiftsRequest) HasExcludeLimitedUpgradable() bool`

HasExcludeLimitedUpgradable returns a boolean if a field has been set.

### GetExcludeLimitedNonUpgradable

`func (o *GetUserGiftsRequest) GetExcludeLimitedNonUpgradable() bool`

GetExcludeLimitedNonUpgradable returns the ExcludeLimitedNonUpgradable field if non-nil, zero value otherwise.

### GetExcludeLimitedNonUpgradableOk

`func (o *GetUserGiftsRequest) GetExcludeLimitedNonUpgradableOk() (*bool, bool)`

GetExcludeLimitedNonUpgradableOk returns a tuple with the ExcludeLimitedNonUpgradable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeLimitedNonUpgradable

`func (o *GetUserGiftsRequest) SetExcludeLimitedNonUpgradable(v bool)`

SetExcludeLimitedNonUpgradable sets ExcludeLimitedNonUpgradable field to given value.

### HasExcludeLimitedNonUpgradable

`func (o *GetUserGiftsRequest) HasExcludeLimitedNonUpgradable() bool`

HasExcludeLimitedNonUpgradable returns a boolean if a field has been set.

### GetExcludeFromBlockchain

`func (o *GetUserGiftsRequest) GetExcludeFromBlockchain() bool`

GetExcludeFromBlockchain returns the ExcludeFromBlockchain field if non-nil, zero value otherwise.

### GetExcludeFromBlockchainOk

`func (o *GetUserGiftsRequest) GetExcludeFromBlockchainOk() (*bool, bool)`

GetExcludeFromBlockchainOk returns a tuple with the ExcludeFromBlockchain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromBlockchain

`func (o *GetUserGiftsRequest) SetExcludeFromBlockchain(v bool)`

SetExcludeFromBlockchain sets ExcludeFromBlockchain field to given value.

### HasExcludeFromBlockchain

`func (o *GetUserGiftsRequest) HasExcludeFromBlockchain() bool`

HasExcludeFromBlockchain returns a boolean if a field has been set.

### GetExcludeUnique

`func (o *GetUserGiftsRequest) GetExcludeUnique() bool`

GetExcludeUnique returns the ExcludeUnique field if non-nil, zero value otherwise.

### GetExcludeUniqueOk

`func (o *GetUserGiftsRequest) GetExcludeUniqueOk() (*bool, bool)`

GetExcludeUniqueOk returns a tuple with the ExcludeUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeUnique

`func (o *GetUserGiftsRequest) SetExcludeUnique(v bool)`

SetExcludeUnique sets ExcludeUnique field to given value.

### HasExcludeUnique

`func (o *GetUserGiftsRequest) HasExcludeUnique() bool`

HasExcludeUnique returns a boolean if a field has been set.

### GetSortByPrice

`func (o *GetUserGiftsRequest) GetSortByPrice() bool`

GetSortByPrice returns the SortByPrice field if non-nil, zero value otherwise.

### GetSortByPriceOk

`func (o *GetUserGiftsRequest) GetSortByPriceOk() (*bool, bool)`

GetSortByPriceOk returns a tuple with the SortByPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortByPrice

`func (o *GetUserGiftsRequest) SetSortByPrice(v bool)`

SetSortByPrice sets SortByPrice field to given value.

### HasSortByPrice

`func (o *GetUserGiftsRequest) HasSortByPrice() bool`

HasSortByPrice returns a boolean if a field has been set.

### GetOffset

`func (o *GetUserGiftsRequest) GetOffset() string`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *GetUserGiftsRequest) GetOffsetOk() (*string, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *GetUserGiftsRequest) SetOffset(v string)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *GetUserGiftsRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *GetUserGiftsRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetUserGiftsRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetUserGiftsRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetUserGiftsRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


