# PostSetMyNameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New bot name; 0-64 characters. Pass an empty string to remove the dedicated name for the given language. | [optional] 
**LanguageCode** | Pointer to **string** | A two-letter ISO 639-1 language code. If empty, the name will be shown to all users for whose language there is no dedicated name. | [optional] 

## Methods

### NewPostSetMyNameRequest

`func NewPostSetMyNameRequest() *PostSetMyNameRequest`

NewPostSetMyNameRequest instantiates a new PostSetMyNameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSetMyNameRequestWithDefaults

`func NewPostSetMyNameRequestWithDefaults() *PostSetMyNameRequest`

NewPostSetMyNameRequestWithDefaults instantiates a new PostSetMyNameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostSetMyNameRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostSetMyNameRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostSetMyNameRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostSetMyNameRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanguageCode

`func (o *PostSetMyNameRequest) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *PostSetMyNameRequest) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *PostSetMyNameRequest) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *PostSetMyNameRequest) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


