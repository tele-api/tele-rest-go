# UpgradeGiftPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection | 
**OwnedGiftId** | **string** | Unique identifier of the regular gift that should be upgraded to a unique one | 
**KeepOriginalDetails** | Pointer to **bool** | Pass True to keep the original gift text, sender and receiver in the upgraded gift | [optional] 
**StarCount** | Pointer to **int32** | The amount of Telegram Stars that will be paid for the upgrade from the business account balance. If &#x60;gift.prepaid_upgrade_star_count &gt; 0&#x60;, then pass 0, otherwise, the *can\\_transfer\\_stars* business bot right is required and &#x60;gift.upgrade_star_count&#x60; must be passed. | [optional] 

## Methods

### NewUpgradeGiftPostRequest

`func NewUpgradeGiftPostRequest(businessConnectionId string, ownedGiftId string, ) *UpgradeGiftPostRequest`

NewUpgradeGiftPostRequest instantiates a new UpgradeGiftPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeGiftPostRequestWithDefaults

`func NewUpgradeGiftPostRequestWithDefaults() *UpgradeGiftPostRequest`

NewUpgradeGiftPostRequestWithDefaults instantiates a new UpgradeGiftPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *UpgradeGiftPostRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *UpgradeGiftPostRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *UpgradeGiftPostRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetOwnedGiftId

`func (o *UpgradeGiftPostRequest) GetOwnedGiftId() string`

GetOwnedGiftId returns the OwnedGiftId field if non-nil, zero value otherwise.

### GetOwnedGiftIdOk

`func (o *UpgradeGiftPostRequest) GetOwnedGiftIdOk() (*string, bool)`

GetOwnedGiftIdOk returns a tuple with the OwnedGiftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedGiftId

`func (o *UpgradeGiftPostRequest) SetOwnedGiftId(v string)`

SetOwnedGiftId sets OwnedGiftId field to given value.


### GetKeepOriginalDetails

`func (o *UpgradeGiftPostRequest) GetKeepOriginalDetails() bool`

GetKeepOriginalDetails returns the KeepOriginalDetails field if non-nil, zero value otherwise.

### GetKeepOriginalDetailsOk

`func (o *UpgradeGiftPostRequest) GetKeepOriginalDetailsOk() (*bool, bool)`

GetKeepOriginalDetailsOk returns a tuple with the KeepOriginalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepOriginalDetails

`func (o *UpgradeGiftPostRequest) SetKeepOriginalDetails(v bool)`

SetKeepOriginalDetails sets KeepOriginalDetails field to given value.

### HasKeepOriginalDetails

`func (o *UpgradeGiftPostRequest) HasKeepOriginalDetails() bool`

HasKeepOriginalDetails returns a boolean if a field has been set.

### GetStarCount

`func (o *UpgradeGiftPostRequest) GetStarCount() int32`

GetStarCount returns the StarCount field if non-nil, zero value otherwise.

### GetStarCountOk

`func (o *UpgradeGiftPostRequest) GetStarCountOk() (*int32, bool)`

GetStarCountOk returns a tuple with the StarCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarCount

`func (o *UpgradeGiftPostRequest) SetStarCount(v int32)`

SetStarCount sets StarCount field to given value.

### HasStarCount

`func (o *UpgradeGiftPostRequest) HasStarCount() bool`

HasStarCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


