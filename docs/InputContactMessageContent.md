# InputContactMessageContent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | Contact&#39;s phone number | 
**FirstName** | **string** | Contact&#39;s first name | 
**LastName** | Pointer to **string** | *Optional*. Contact&#39;s last name | [optional] 
**Vcard** | Pointer to **string** | *Optional*. Additional data about the contact in the form of a [vCard](https://en.wikipedia.org/wiki/VCard), 0-2048 bytes | [optional] 

## Methods

### NewInputContactMessageContent

`func NewInputContactMessageContent(phoneNumber string, firstName string, ) *InputContactMessageContent`

NewInputContactMessageContent instantiates a new InputContactMessageContent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputContactMessageContentWithDefaults

`func NewInputContactMessageContentWithDefaults() *InputContactMessageContent`

NewInputContactMessageContentWithDefaults instantiates a new InputContactMessageContent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *InputContactMessageContent) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *InputContactMessageContent) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *InputContactMessageContent) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetFirstName

`func (o *InputContactMessageContent) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *InputContactMessageContent) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *InputContactMessageContent) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *InputContactMessageContent) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *InputContactMessageContent) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *InputContactMessageContent) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *InputContactMessageContent) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetVcard

`func (o *InputContactMessageContent) GetVcard() string`

GetVcard returns the Vcard field if non-nil, zero value otherwise.

### GetVcardOk

`func (o *InputContactMessageContent) GetVcardOk() (*string, bool)`

GetVcardOk returns a tuple with the Vcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcard

`func (o *InputContactMessageContent) SetVcard(v string)`

SetVcard sets Vcard field to given value.

### HasVcard

`func (o *InputContactMessageContent) HasVcard() bool`

HasVcard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


