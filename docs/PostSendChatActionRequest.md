# PostSendChatActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessConnectionId** | Pointer to **string** | Unique identifier of the business connection on behalf of which the action will be sent | [optional] 
**ChatId** | [**PostSendMessageRequestChatId**](PostSendMessageRequestChatId.md) |  | 
**MessageThreadId** | Pointer to **int32** | Unique identifier for the target message thread; for supergroups only | [optional] 
**Action** | **string** | Type of action to broadcast. Choose one, depending on what the user is about to receive: *typing* for [text messages](https://core.telegram.org/bots/api/#sendmessage), *upload\\_photo* for [photos](https://core.telegram.org/bots/api/#sendphoto), *record\\_video* or *upload\\_video* for [videos](https://core.telegram.org/bots/api/#sendvideo), *record\\_voice* or *upload\\_voice* for [voice notes](https://core.telegram.org/bots/api/#sendvoice), *upload\\_document* for [general files](https://core.telegram.org/bots/api/#senddocument), *choose\\_sticker* for [stickers](https://core.telegram.org/bots/api/#sendsticker), *find\\_location* for [location data](https://core.telegram.org/bots/api/#sendlocation), *record\\_video\\_note* or *upload\\_video\\_note* for [video notes](https://core.telegram.org/bots/api/#sendvideonote). | 

## Methods

### NewPostSendChatActionRequest

`func NewPostSendChatActionRequest(chatId PostSendMessageRequestChatId, action string, ) *PostSendChatActionRequest`

NewPostSendChatActionRequest instantiates a new PostSendChatActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSendChatActionRequestWithDefaults

`func NewPostSendChatActionRequestWithDefaults() *PostSendChatActionRequest`

NewPostSendChatActionRequestWithDefaults instantiates a new PostSendChatActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessConnectionId

`func (o *PostSendChatActionRequest) GetBusinessConnectionId() string`

GetBusinessConnectionId returns the BusinessConnectionId field if non-nil, zero value otherwise.

### GetBusinessConnectionIdOk

`func (o *PostSendChatActionRequest) GetBusinessConnectionIdOk() (*string, bool)`

GetBusinessConnectionIdOk returns a tuple with the BusinessConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessConnectionId

`func (o *PostSendChatActionRequest) SetBusinessConnectionId(v string)`

SetBusinessConnectionId sets BusinessConnectionId field to given value.

### HasBusinessConnectionId

`func (o *PostSendChatActionRequest) HasBusinessConnectionId() bool`

HasBusinessConnectionId returns a boolean if a field has been set.

### GetChatId

`func (o *PostSendChatActionRequest) GetChatId() PostSendMessageRequestChatId`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *PostSendChatActionRequest) GetChatIdOk() (*PostSendMessageRequestChatId, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *PostSendChatActionRequest) SetChatId(v PostSendMessageRequestChatId)`

SetChatId sets ChatId field to given value.


### GetMessageThreadId

`func (o *PostSendChatActionRequest) GetMessageThreadId() int32`

GetMessageThreadId returns the MessageThreadId field if non-nil, zero value otherwise.

### GetMessageThreadIdOk

`func (o *PostSendChatActionRequest) GetMessageThreadIdOk() (*int32, bool)`

GetMessageThreadIdOk returns a tuple with the MessageThreadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageThreadId

`func (o *PostSendChatActionRequest) SetMessageThreadId(v int32)`

SetMessageThreadId sets MessageThreadId field to given value.

### HasMessageThreadId

`func (o *PostSendChatActionRequest) HasMessageThreadId() bool`

HasMessageThreadId returns a boolean if a field has been set.

### GetAction

`func (o *PostSendChatActionRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PostSendChatActionRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PostSendChatActionRequest) SetAction(v string)`

SetAction sets Action field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


