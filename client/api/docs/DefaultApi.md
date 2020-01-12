# BroodClash.DefaultApi

All URIs are relative to *http://localhost:3000/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**authTest**](DefaultApi.md#authTest) | **GET** /authTest | Test the authorization code
[**createQuestions**](DefaultApi.md#createQuestions) | **POST** /questions | Create a new question
[**getQuestion**](DefaultApi.md#getQuestion) | **GET** /questions/{id} | Get a question by ID
[**getQuestions**](DefaultApi.md#getQuestions) | **GET** /questions | Get all questions
[**revealAnswer**](DefaultApi.md#revealAnswer) | **POST** /questions/{questionId}/answers/{answerId}/reveal | Marks an answer as revealed
[**websocket**](DefaultApi.md#websocket) | **GET** /presenter/websocket | Establish a presenter websocket


<a name="authTest"></a>
# **authTest**
> authTest()

Test the authorization code

### Example
```javascript
import {BroodClash} from 'brood_clash';
let defaultClient = BroodClash.ApiClient.instance;

// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new BroodClash.DefaultApi();
apiInstance.authTest().then(() => {
  console.log('API called successfully.');
}, (error) => {
  console.error(error);
});

```

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="createQuestions"></a>
# **createQuestions**
> createQuestions(questions)

Create a new question

### Example
```javascript
import {BroodClash} from 'brood_clash';
let defaultClient = BroodClash.ApiClient.instance;

// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new BroodClash.DefaultApi();

let questions = [new BroodClash.Question()]; // [Question] | 

apiInstance.createQuestions(questions).then(() => {
  console.log('API called successfully.');
}, (error) => {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questions** | [**[Question]**](Question.md)|  | 

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getQuestion"></a>
# **getQuestion**
> Question getQuestion(id)

Get a question by ID

### Example
```javascript
import {BroodClash} from 'brood_clash';
let defaultClient = BroodClash.ApiClient.instance;

// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new BroodClash.DefaultApi();

let id = 789; // Number | 

apiInstance.getQuestion(id).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **Number**|  | 

### Return type

[**Question**](Question.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="getQuestions"></a>
# **getQuestions**
> [Question] getQuestions()

Get all questions

### Example
```javascript
import {BroodClash} from 'brood_clash';
let defaultClient = BroodClash.ApiClient.instance;

// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new BroodClash.DefaultApi();
apiInstance.getQuestions().then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[Question]**](Question.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="revealAnswer"></a>
# **revealAnswer**
> revealAnswer(questionId, answerId)

Marks an answer as revealed

### Example
```javascript
import {BroodClash} from 'brood_clash';
let defaultClient = BroodClash.ApiClient.instance;

// Configure API key authorization: ApiKey
let ApiKey = defaultClient.authentications['ApiKey'];
ApiKey.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//ApiKey.apiKeyPrefix = 'Token';

let apiInstance = new BroodClash.DefaultApi();

let questionId = 789; // Number | 

let answerId = 789; // Number | 

apiInstance.revealAnswer(questionId, answerId).then(() => {
  console.log('API called successfully.');
}, (error) => {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questionId** | **Number**|  | 
 **answerId** | **Number**|  | 

### Return type

null (empty response body)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

<a name="websocket"></a>
# **websocket**
> websocket()

Establish a presenter websocket

### Example
```javascript
import {BroodClash} from 'brood_clash';

let apiInstance = new BroodClash.DefaultApi();
apiInstance.websocket().then(() => {
  console.log('API called successfully.');
}, (error) => {
  console.error(error);
});

```

### Parameters
This endpoint does not need any parameter.

### Return type

null (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

