# MenuButtonWebApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the button, must be *web\\_app* | [default to "web_app"]
**Text** | **string** | Text on the button | 
**WebApp** | [**WebAppInfo**](WebAppInfo.md) |  | 

## Methods

### NewMenuButtonWebApp

`func NewMenuButtonWebApp(type_ string, text string, webApp WebAppInfo, ) *MenuButtonWebApp`

NewMenuButtonWebApp instantiates a new MenuButtonWebApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuButtonWebAppWithDefaults

`func NewMenuButtonWebAppWithDefaults() *MenuButtonWebApp`

NewMenuButtonWebAppWithDefaults instantiates a new MenuButtonWebApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MenuButtonWebApp) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MenuButtonWebApp) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MenuButtonWebApp) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *MenuButtonWebApp) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MenuButtonWebApp) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MenuButtonWebApp) SetText(v string)`

SetText sets Text field to given value.


### GetWebApp

`func (o *MenuButtonWebApp) GetWebApp() WebAppInfo`

GetWebApp returns the WebApp field if non-nil, zero value otherwise.

### GetWebAppOk

`func (o *MenuButtonWebApp) GetWebAppOk() (*WebAppInfo, bool)`

GetWebAppOk returns a tuple with the WebApp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebApp

`func (o *MenuButtonWebApp) SetWebApp(v WebAppInfo)`

SetWebApp sets WebApp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


