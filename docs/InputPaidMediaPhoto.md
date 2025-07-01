# InputPaidMediaPhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the media, must be *photo* | [default to "photo"]
**Media** | **string** | File to send. Pass a file\\_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://\\&lt;file\\_attach\\_name\\&gt;” to upload a new one using multipart/form-data under \\&lt;file\\_attach\\_name\\&gt; name. [More information on Sending Files »](https://core.telegram.org/bots/api/#sending-files) | 

## Methods

### NewInputPaidMediaPhoto

`func NewInputPaidMediaPhoto(type_ string, media string, ) *InputPaidMediaPhoto`

NewInputPaidMediaPhoto instantiates a new InputPaidMediaPhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputPaidMediaPhotoWithDefaults

`func NewInputPaidMediaPhotoWithDefaults() *InputPaidMediaPhoto`

NewInputPaidMediaPhotoWithDefaults instantiates a new InputPaidMediaPhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InputPaidMediaPhoto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InputPaidMediaPhoto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InputPaidMediaPhoto) SetType(v string)`

SetType sets Type field to given value.


### GetMedia

`func (o *InputPaidMediaPhoto) GetMedia() string`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *InputPaidMediaPhoto) GetMediaOk() (*string, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *InputPaidMediaPhoto) SetMedia(v string)`

SetMedia sets Media field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


