# SetMyCommandsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commands** | [**[]BotCommand**](BotCommand.md) | A JSON-serialized list of bot commands to be set as the list of the bot&#39;s commands. At most 100 commands can be specified. | 
**Scope** | Pointer to [**BotCommandScope**](BotCommandScope.md) |  | [optional] 
**LanguageCode** | Pointer to **string** | A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands | [optional] 

## Methods

### NewSetMyCommandsRequest

`func NewSetMyCommandsRequest(commands []BotCommand, ) *SetMyCommandsRequest`

NewSetMyCommandsRequest instantiates a new SetMyCommandsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetMyCommandsRequestWithDefaults

`func NewSetMyCommandsRequestWithDefaults() *SetMyCommandsRequest`

NewSetMyCommandsRequestWithDefaults instantiates a new SetMyCommandsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommands

`func (o *SetMyCommandsRequest) GetCommands() []BotCommand`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *SetMyCommandsRequest) GetCommandsOk() (*[]BotCommand, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *SetMyCommandsRequest) SetCommands(v []BotCommand)`

SetCommands sets Commands field to given value.


### GetScope

`func (o *SetMyCommandsRequest) GetScope() BotCommandScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SetMyCommandsRequest) GetScopeOk() (*BotCommandScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SetMyCommandsRequest) SetScope(v BotCommandScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SetMyCommandsRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetLanguageCode

`func (o *SetMyCommandsRequest) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *SetMyCommandsRequest) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *SetMyCommandsRequest) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *SetMyCommandsRequest) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


