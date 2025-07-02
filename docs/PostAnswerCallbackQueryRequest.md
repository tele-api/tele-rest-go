# PostAnswerCallbackQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackQueryId** | **string** | Unique identifier for the query to be answered | 
**Text** | Pointer to **string** | Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters | [optional] 
**ShowAlert** | Pointer to **bool** | If *True*, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to *false*. | [optional] [default to false]
**Url** | Pointer to **string** | URL that will be opened by the user&#39;s client. If you have created a [Game](https://core.telegram.org/bots/api/#game) and accepted the conditions via [@BotFather](https://t.me/botfather), specify the URL that opens your game - note that this will only work if the query comes from a [*callback\\_game*](https://core.telegram.org/bots/api/#inlinekeyboardbutton) button.    Otherwise, you may use links like &#x60;t.me/your_bot?start&#x3D;XXXX&#x60; that open your bot with a parameter. | [optional] 
**CacheTime** | Pointer to **int32** | The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0. | [optional] [default to 0]

## Methods

### NewPostAnswerCallbackQueryRequest

`func NewPostAnswerCallbackQueryRequest(callbackQueryId string, ) *PostAnswerCallbackQueryRequest`

NewPostAnswerCallbackQueryRequest instantiates a new PostAnswerCallbackQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAnswerCallbackQueryRequestWithDefaults

`func NewPostAnswerCallbackQueryRequestWithDefaults() *PostAnswerCallbackQueryRequest`

NewPostAnswerCallbackQueryRequestWithDefaults instantiates a new PostAnswerCallbackQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackQueryId

`func (o *PostAnswerCallbackQueryRequest) GetCallbackQueryId() string`

GetCallbackQueryId returns the CallbackQueryId field if non-nil, zero value otherwise.

### GetCallbackQueryIdOk

`func (o *PostAnswerCallbackQueryRequest) GetCallbackQueryIdOk() (*string, bool)`

GetCallbackQueryIdOk returns a tuple with the CallbackQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackQueryId

`func (o *PostAnswerCallbackQueryRequest) SetCallbackQueryId(v string)`

SetCallbackQueryId sets CallbackQueryId field to given value.


### GetText

`func (o *PostAnswerCallbackQueryRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *PostAnswerCallbackQueryRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *PostAnswerCallbackQueryRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *PostAnswerCallbackQueryRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetShowAlert

`func (o *PostAnswerCallbackQueryRequest) GetShowAlert() bool`

GetShowAlert returns the ShowAlert field if non-nil, zero value otherwise.

### GetShowAlertOk

`func (o *PostAnswerCallbackQueryRequest) GetShowAlertOk() (*bool, bool)`

GetShowAlertOk returns a tuple with the ShowAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAlert

`func (o *PostAnswerCallbackQueryRequest) SetShowAlert(v bool)`

SetShowAlert sets ShowAlert field to given value.

### HasShowAlert

`func (o *PostAnswerCallbackQueryRequest) HasShowAlert() bool`

HasShowAlert returns a boolean if a field has been set.

### GetUrl

`func (o *PostAnswerCallbackQueryRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PostAnswerCallbackQueryRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PostAnswerCallbackQueryRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PostAnswerCallbackQueryRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetCacheTime

`func (o *PostAnswerCallbackQueryRequest) GetCacheTime() int32`

GetCacheTime returns the CacheTime field if non-nil, zero value otherwise.

### GetCacheTimeOk

`func (o *PostAnswerCallbackQueryRequest) GetCacheTimeOk() (*int32, bool)`

GetCacheTimeOk returns a tuple with the CacheTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTime

`func (o *PostAnswerCallbackQueryRequest) SetCacheTime(v int32)`

SetCacheTime sets CacheTime field to given value.

### HasCacheTime

`func (o *PostAnswerCallbackQueryRequest) HasCacheTime() bool`

HasCacheTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


