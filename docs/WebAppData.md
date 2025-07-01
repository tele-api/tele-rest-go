# WebAppData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | The data. Be aware that a bad client can send arbitrary data in this field. | 
**ButtonText** | **string** | Text of the *web\\_app* keyboard button from which the Web App was opened. Be aware that a bad client can send arbitrary data in this field. | 

## Methods

### NewWebAppData

`func NewWebAppData(data string, buttonText string, ) *WebAppData`

NewWebAppData instantiates a new WebAppData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebAppDataWithDefaults

`func NewWebAppDataWithDefaults() *WebAppData`

NewWebAppDataWithDefaults instantiates a new WebAppData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WebAppData) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WebAppData) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WebAppData) SetData(v string)`

SetData sets Data field to given value.


### GetButtonText

`func (o *WebAppData) GetButtonText() string`

GetButtonText returns the ButtonText field if non-nil, zero value otherwise.

### GetButtonTextOk

`func (o *WebAppData) GetButtonTextOk() (*string, bool)`

GetButtonTextOk returns a tuple with the ButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonText

`func (o *WebAppData) SetButtonText(v string)`

SetButtonText sets ButtonText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


