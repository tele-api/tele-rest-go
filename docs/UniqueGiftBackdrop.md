# UniqueGiftBackdrop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the backdrop | 
**Colors** | [**UniqueGiftBackdropColors**](UniqueGiftBackdropColors.md) |  | 
**RarityPerMille** | **int32** | The number of unique gifts that receive this backdrop for every 1000 gifts upgraded | 

## Methods

### NewUniqueGiftBackdrop

`func NewUniqueGiftBackdrop(name string, colors UniqueGiftBackdropColors, rarityPerMille int32, ) *UniqueGiftBackdrop`

NewUniqueGiftBackdrop instantiates a new UniqueGiftBackdrop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUniqueGiftBackdropWithDefaults

`func NewUniqueGiftBackdropWithDefaults() *UniqueGiftBackdrop`

NewUniqueGiftBackdropWithDefaults instantiates a new UniqueGiftBackdrop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UniqueGiftBackdrop) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UniqueGiftBackdrop) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UniqueGiftBackdrop) SetName(v string)`

SetName sets Name field to given value.


### GetColors

`func (o *UniqueGiftBackdrop) GetColors() UniqueGiftBackdropColors`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *UniqueGiftBackdrop) GetColorsOk() (*UniqueGiftBackdropColors, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *UniqueGiftBackdrop) SetColors(v UniqueGiftBackdropColors)`

SetColors sets Colors field to given value.


### GetRarityPerMille

`func (o *UniqueGiftBackdrop) GetRarityPerMille() int32`

GetRarityPerMille returns the RarityPerMille field if non-nil, zero value otherwise.

### GetRarityPerMilleOk

`func (o *UniqueGiftBackdrop) GetRarityPerMilleOk() (*int32, bool)`

GetRarityPerMilleOk returns a tuple with the RarityPerMille field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRarityPerMille

`func (o *UniqueGiftBackdrop) SetRarityPerMille(v int32)`

SetRarityPerMille sets RarityPerMille field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


