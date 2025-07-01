# ShippingOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Shipping option identifier | 
**Title** | **string** | Option title | 
**Prices** | [**[]LabeledPrice**](LabeledPrice.md) | List of price portions | 

## Methods

### NewShippingOption

`func NewShippingOption(id string, title string, prices []LabeledPrice, ) *ShippingOption`

NewShippingOption instantiates a new ShippingOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShippingOptionWithDefaults

`func NewShippingOptionWithDefaults() *ShippingOption`

NewShippingOptionWithDefaults instantiates a new ShippingOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShippingOption) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShippingOption) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShippingOption) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ShippingOption) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ShippingOption) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ShippingOption) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPrices

`func (o *ShippingOption) GetPrices() []LabeledPrice`

GetPrices returns the Prices field if non-nil, zero value otherwise.

### GetPricesOk

`func (o *ShippingOption) GetPricesOk() (*[]LabeledPrice, bool)`

GetPricesOk returns a tuple with the Prices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrices

`func (o *ShippingOption) SetPrices(v []LabeledPrice)`

SetPrices sets Prices field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


