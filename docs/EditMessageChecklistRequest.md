# EditMessageChecklistRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | **string** | Unique identifier of the business connection on behalf of which the message will be sent | 
**ChatId** | **int32** | Unique identifier for the target chat | 
**MessageId** | **int32** | Unique identifier for the target message | 
**Checklist** | [**InputChecklist**](InputChecklist.md) |  | 
**ReplyMarkup** | Pointer to [**InlineKeyboardMarkup**](InlineKeyboardMarkup.md) |  | [optional] 

## Methods

### NewEditMessageChecklistRequest

`func NewEditMessageChecklistRequest(businessConnectionId string, chatId int32, messageId int32, checklist InputChecklist, ) *EditMessageChecklistRequest`

NewEditMessageChecklistRequest instantiates a new EditMessageChecklistRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditMessageChecklistRequestWithDefaults

`func NewEditMessageChecklistRequestWithDefaults() *EditMessageChecklistRequest`

NewEditMessageChecklistRequestWithDefaults instantiates a new EditMessageChecklistRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *EditMessageChecklistRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *EditMessageChecklistRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *EditMessageChecklistRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.


### GetChatId

`func (o *EditMessageChecklistRequest) GetChatId() int32`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *EditMessageChecklistRequest) GetChatIdOk() (*int32, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *EditMessageChecklistRequest) SetChatId(v int32)`

SetChatId sets ChatId field to given value.


### GetMessageId

`func (o *EditMessageChecklistRequest) GetMessageId() int32`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *EditMessageChecklistRequest) GetMessageIdOk() (*int32, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *EditMessageChecklistRequest) SetMessageId(v int32)`

SetMessageId sets MessageId field to given value.


### GetChecklist

`func (o *EditMessageChecklistRequest) GetChecklist() InputChecklist`

GetChecklist returns the Checklist field if non-nil, zero value otherwise.

### GetChecklistOk

`func (o *EditMessageChecklistRequest) GetChecklistOk() (*InputChecklist, bool)`

GetChecklistOk returns a tuple with the Checklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklist

`func (o *EditMessageChecklistRequest) SetChecklist(v InputChecklist)`

SetChecklist sets Checklist field to given value.


### GetReplyMarkup

`func (o *EditMessageChecklistRequest) GetReplyMarkup() InlineKeyboardMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *EditMessageChecklistRequest) GetReplyMarkupOk() (*InlineKeyboardMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *EditMessageChecklistRequest) SetReplyMarkup(v InlineKeyboardMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *EditMessageChecklistRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


