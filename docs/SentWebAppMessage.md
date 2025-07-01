# SentWebAppMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InlineMessageId** | Pointer to **string** | *Optional*. Identifier of the sent inline message. Available only if there is an [inline keyboard](https://core.telegram.org/bots/api/#inlinekeyboardmarkup) attached to the message. | [optional] 

## Methods

### NewSentWebAppMessage

`func NewSentWebAppMessage() *SentWebAppMessage`

NewSentWebAppMessage instantiates a new SentWebAppMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSentWebAppMessageWithDefaults

`func NewSentWebAppMessageWithDefaults() *SentWebAppMessage`

NewSentWebAppMessageWithDefaults instantiates a new SentWebAppMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInlineMessageId

`func (o *SentWebAppMessage) GetInlineMessageId() string`

GetInlineMessageId returns the InlineMessageId field if non-nil, zero value otherwise.

### GetInlineMessageIdOk

`func (o *SentWebAppMessage) GetInlineMessageIdOk() (*string, bool)`

GetInlineMessageIdOk returns a tuple with the InlineMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInlineMessageId

`func (o *SentWebAppMessage) SetInlineMessageId(v string)`

SetInlineMessageId sets InlineMessageId field to given value.

### HasInlineMessageId

`func (o *SentWebAppMessage) HasInlineMessageId() bool`

HasInlineMessageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


