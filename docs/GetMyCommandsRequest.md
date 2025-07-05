# GetMyCommandsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to [**BotCommandScope**](BotCommandScope.md) |  | [optional] 
**LanguageCode** | Pointer to **string** | A two-letter ISO 639-1 language code or an empty string | [optional] 

## Methods

### NewGetMyCommandsRequest

`func NewGetMyCommandsRequest() *GetMyCommandsRequest`

NewGetMyCommandsRequest instantiates a new GetMyCommandsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMyCommandsRequestWithDefaults

`func NewGetMyCommandsRequestWithDefaults() *GetMyCommandsRequest`

NewGetMyCommandsRequestWithDefaults instantiates a new GetMyCommandsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *GetMyCommandsRequest) GetScope() BotCommandScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetMyCommandsRequest) GetScopeOk() (*BotCommandScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetMyCommandsRequest) SetScope(v BotCommandScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GetMyCommandsRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetLanguageCode

`func (o *GetMyCommandsRequest) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *GetMyCommandsRequest) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *GetMyCommandsRequest) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *GetMyCommandsRequest) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


