# BotCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | **string** | Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and underscores. | 
**Description** | **string** | Description of the command; 1-256 characters. | 

## Methods

### NewBotCommand

`func NewBotCommand(command string, description string, ) *BotCommand`

NewBotCommand instantiates a new BotCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBotCommandWithDefaults

`func NewBotCommandWithDefaults() *BotCommand`

NewBotCommandWithDefaults instantiates a new BotCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *BotCommand) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *BotCommand) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *BotCommand) SetCommand(v string)`

SetCommand sets Command field to given value.


### GetDescription

`func (o *BotCommand) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BotCommand) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BotCommand) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


