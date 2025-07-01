# InlineQueryResultsButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** | Label text on the button | 
**WebApp** | Pointer to [**WebAppInfo**](WebAppInfo.md) |  | [optional] 
**StartParameter** | Pointer to **string** | *Optional*. [Deep-linking](https://core.telegram.org/bots/features#deep-linking) parameter for the /start message sent to the bot when a user presses the button. 1-64 characters, only &#x60;A-Z&#x60;, &#x60;a-z&#x60;, &#x60;0-9&#x60;, &#x60;_&#x60; and &#x60;-&#x60; are allowed.    *Example:* An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a &#39;Connect your YouTube account&#39; button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an OAuth link. Once done, the bot can offer a [*switch\\_inline*](https://core.telegram.org/bots/api/#inlinekeyboardmarkup) button so that the user can easily return to the chat where they wanted to use the bot&#39;s inline capabilities. | [optional] 

## Methods

### NewInlineQueryResultsButton

`func NewInlineQueryResultsButton(text string, ) *InlineQueryResultsButton`

NewInlineQueryResultsButton instantiates a new InlineQueryResultsButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineQueryResultsButtonWithDefaults

`func NewInlineQueryResultsButtonWithDefaults() *InlineQueryResultsButton`

NewInlineQueryResultsButtonWithDefaults instantiates a new InlineQueryResultsButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *InlineQueryResultsButton) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *InlineQueryResultsButton) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *InlineQueryResultsButton) SetText(v string)`

SetText sets Text field to given value.


### GetWebApp

`func (o *InlineQueryResultsButton) GetWebApp() WebAppInfo`

GetWebApp returns the WebApp field if non-nil, zero value otherwise.

### GetWebAppOk

`func (o *InlineQueryResultsButton) GetWebAppOk() (*WebAppInfo, bool)`

GetWebAppOk returns a tuple with the WebApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApp

`func (o *InlineQueryResultsButton) SetWebApp(v WebAppInfo)`

SetWebApp sets WebApp field to given value.

### HasWebApp

`func (o *InlineQueryResultsButton) HasWebApp() bool`

HasWebApp returns a boolean if a field has been set.

### GetStartParameter

`func (o *InlineQueryResultsButton) GetStartParameter() string`

GetStartParameter returns the StartParameter field if non-nil, zero value otherwise.

### GetStartParameterOk

`func (o *InlineQueryResultsButton) GetStartParameterOk() (*string, bool)`

GetStartParameterOk returns a tuple with the StartParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartParameter

`func (o *InlineQueryResultsButton) SetStartParameter(v string)`

SetStartParameter sets StartParameter field to given value.

### HasStartParameter

`func (o *InlineQueryResultsButton) HasStartParameter() bool`

HasStartParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


