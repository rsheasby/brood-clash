# BroodClashApi.DefaultApi

All URIs are relative to *https://localhost/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**deleteQuestion**](DefaultApi.md#deleteQuestion) | **DELETE** /questions/{id} | Delete Question
[**getAllQuestions**](DefaultApi.md#getAllQuestions) | **GET** /questions | Get All Questions
[**getCurrentQuestion**](DefaultApi.md#getCurrentQuestion) | **GET** /currentQuestion | Get Current Question
[**incorrectAnswer**](DefaultApi.md#incorrectAnswer) | **POST** /incorrectAnswer | Incorrect Answer
[**postQuestions**](DefaultApi.md#postQuestions) | **POST** /questions | Post Questions
[**resetGameState**](DefaultApi.md#resetGameState) | **POST** /reset | Reset Game State
[**revealAnswer**](DefaultApi.md#revealAnswer) | **POST** /answers/{id}/reveal | Reveal answer
[**selectQuestion**](DefaultApi.md#selectQuestion) | **POST** /questions/{id}/select | Select Question


<a name="deleteQuestion"></a>
# **deleteQuestion**
> deleteQuestion(id)

Delete Question

### Example
```javascript
import {BroodClashApi} from 'brood_clash_api';
let defaultClient = BroodClashApi.ApiClient.instance;

// Configure API key authorization: CodeAuth
let CodeAuth = defaultClient.authentications['CodeAuth'];
CodeAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//CodeAuth.apiKeyPrefix = 'Token';

let apiInstance = new BroodClashApi.DefaultApi();

let id = "id_example"; // String | Question ID, must be UUID

apiInstance.deleteQuestion(id).then(() => {
  console.log('API called successfully.');
}, (error) => {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**String**](.md)| Question ID, must be UUID | 

### Return type

null (empty response body)

### Authorization

[CodeAuth](../README.md#CodeAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="getAllQuestions"></a>
# **getAllQuestions**
> [ModelsQuestion] getAllQuestions()

Get All Questions

### Example
```javascript
import {BroodClashApi} from 'brood_clash_api';
let defaultClient = BroodClashApi.ApiClient.instance;

// Configure API key authorization: CodeAuth
let CodeAuth = defaultClient.authentications['CodeAuth'];
CodeAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//CodeAuth.apiKeyPrefix = 'Token';

let apiInstance = new BroodClashApi.DefaultApi();
apiInstance.getAllQuestions().then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters
This endpoint does not need any parameter.

### Return type

[**[ModelsQuestion]**](ModelsQuestion.md)

### Authorization

[CodeAuth](../README.md#CodeAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="getCurrentQuestion"></a>
# **getCurrentQuestion**
> ModelsQuestion getCurrentQuestion()

Get Current Question

### Example
```javascript
import {BroodClashApi} from 'brood_clash_api';
let defaultClient = BroodClashApi.ApiClient.instance;

// Configure API key authorization: CodeAuth
let CodeAuth = defaultClient.authentications['CodeAuth'];
CodeAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//CodeAuth.apiKeyPrefix = 'Token';

let apiInstance = new BroodClashApi.DefaultApi();
apiInstance.getCurrentQuestion().then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters
This endpoint does not need any parameter.

### Return type

[**ModelsQuestion**](ModelsQuestion.md)

### Authorization

[CodeAuth](../README.md#CodeAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

<a name="incorrectAnswer"></a>
# **incorrectAnswer**
> incorrectAnswer()

Incorrect Answer

### Example
```javascript
import {BroodClashApi} from 'brood_clash_api';
let defaultClient = BroodClashApi.ApiClient.instance;

// Configure API key authorization: CodeAuth
let CodeAuth = defaultClient.authentications['CodeAuth'];
CodeAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//CodeAuth.apiKeyPrefix = 'Token';

let apiInstance = new BroodClashApi.DefaultApi();
apiInstance.incorrectAnswer().then(() => {
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

[CodeAuth](../README.md#CodeAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="postQuestions"></a>
# **postQuestions**
> postQuestions(questions)

Post Questions

### Example
```javascript
import {BroodClashApi} from 'brood_clash_api';
let defaultClient = BroodClashApi.ApiClient.instance;

// Configure API key authorization: CodeAuth
let CodeAuth = defaultClient.authentications['CodeAuth'];
CodeAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//CodeAuth.apiKeyPrefix = 'Token';

let apiInstance = new BroodClashApi.DefaultApi();

let questions = [new BroodClashApi.ModelsQuestion()]; // [ModelsQuestion] | Questions to be created

apiInstance.postQuestions(questions).then(() => {
  console.log('API called successfully.');
}, (error) => {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **questions** | [**[ModelsQuestion]**](ModelsQuestion.md)| Questions to be created | 

### Return type

null (empty response body)

### Authorization

[CodeAuth](../README.md#CodeAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

<a name="resetGameState"></a>
# **resetGameState**
> resetGameState()

Reset Game State

### Example
```javascript
import {BroodClashApi} from 'brood_clash_api';
let defaultClient = BroodClashApi.ApiClient.instance;

// Configure API key authorization: CodeAuth
let CodeAuth = defaultClient.authentications['CodeAuth'];
CodeAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//CodeAuth.apiKeyPrefix = 'Token';

let apiInstance = new BroodClashApi.DefaultApi();
apiInstance.resetGameState().then(() => {
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

[CodeAuth](../README.md#CodeAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="revealAnswer"></a>
# **revealAnswer**
> revealAnswer(id)

Reveal answer

### Example
```javascript
import {BroodClashApi} from 'brood_clash_api';
let defaultClient = BroodClashApi.ApiClient.instance;

// Configure API key authorization: CodeAuth
let CodeAuth = defaultClient.authentications['CodeAuth'];
CodeAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//CodeAuth.apiKeyPrefix = 'Token';

let apiInstance = new BroodClashApi.DefaultApi();

let id = "id_example"; // String | Answer ID, must be UUID

apiInstance.revealAnswer(id).then(() => {
  console.log('API called successfully.');
}, (error) => {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**String**](.md)| Answer ID, must be UUID | 

### Return type

null (empty response body)

### Authorization

[CodeAuth](../README.md#CodeAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

<a name="selectQuestion"></a>
# **selectQuestion**
> ModelsQuestion selectQuestion(id)

Select Question

### Example
```javascript
import {BroodClashApi} from 'brood_clash_api';
let defaultClient = BroodClashApi.ApiClient.instance;

// Configure API key authorization: CodeAuth
let CodeAuth = defaultClient.authentications['CodeAuth'];
CodeAuth.apiKey = 'YOUR API KEY';
// Uncomment the following line to set a prefix for the API key, e.g. "Token" (defaults to null)
//CodeAuth.apiKeyPrefix = 'Token';

let apiInstance = new BroodClashApi.DefaultApi();

let id = "id_example"; // String | Question ID, must be UUID

apiInstance.selectQuestion(id).then((data) => {
  console.log('API called successfully. Returned data: ' + data);
}, (error) => {
  console.error(error);
});

```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | [**String**](.md)| Question ID, must be UUID | 

### Return type

[**ModelsQuestion**](ModelsQuestion.md)

### Authorization

[CodeAuth](../README.md#CodeAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

