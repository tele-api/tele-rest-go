# SetMyShortDescriptionPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortDescription** | Pointer to **string** | New short description for the bot; 0-120 characters. Pass an empty string to remove the dedicated short description for the given language. | [optional] 
**LanguageCode** | Pointer to **string** | A two-letter ISO 639-1 language code. If empty, the short description will be applied to all users for whose language there is no dedicated short description. | [optional] 

## Methods

### NewSetMyShortDescriptionPostRequest

`func NewSetMyShortDescriptionPostRequest() *SetMyShortDescriptionPostRequest`

NewSetMyShortDescriptionPostRequest instantiates a new SetMyShortDescriptionPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetMyShortDescriptionPostRequestWithDefaults

`func NewSetMyShortDescriptionPostRequestWithDefaults() *SetMyShortDescriptionPostRequest`

NewSetMyShortDescriptionPostRequestWithDefaults instantiates a new SetMyShortDescriptionPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortDescription

`func (o *SetMyShortDescriptionPostRequest) GetShortDescription() string`

GetShortDescription returns the ShortDescription field if non-nil, zero value otherwise.

### GetShortDescriptionOk

`func (o *SetMyShortDescriptionPostRequest) GetShortDescriptionOk() (*string, bool)`

GetShortDescriptionOk returns a tuple with the ShortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortDescription

`func (o *SetMyShortDescriptionPostRequest) SetShortDescription(v string)`

SetShortDescription sets ShortDescription field to given value.

### HasShortDescription

`func (o *SetMyShortDescriptionPostRequest) HasShortDescription() bool`

HasShortDescription returns a boolean if a field has been set.

### GetLanguageCode

`func (o *SetMyShortDescriptionPostRequest) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *SetMyShortDescriptionPostRequest) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *SetMyShortDescriptionPostRequest) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *SetMyShortDescriptionPostRequest) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


