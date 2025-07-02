# SetMyDescriptionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language. | [optional] 
**LanguageCode** | Pointer to **string** | A two-letter ISO 639-1 language code. If empty, the description will be applied to all users for whose language there is no dedicated description. | [optional] 

## Methods

### NewSetMyDescriptionRequest

`func NewSetMyDescriptionRequest() *SetMyDescriptionRequest`

NewSetMyDescriptionRequest instantiates a new SetMyDescriptionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetMyDescriptionRequestWithDefaults

`func NewSetMyDescriptionRequestWithDefaults() *SetMyDescriptionRequest`

NewSetMyDescriptionRequestWithDefaults instantiates a new SetMyDescriptionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SetMyDescriptionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SetMyDescriptionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SetMyDescriptionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SetMyDescriptionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLanguageCode

`func (o *SetMyDescriptionRequest) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *SetMyDescriptionRequest) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *SetMyDescriptionRequest) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *SetMyDescriptionRequest) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


