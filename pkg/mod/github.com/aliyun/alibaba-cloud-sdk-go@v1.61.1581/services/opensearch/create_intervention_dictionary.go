package opensearch

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// CreateInterventionDictionary invokes the opensearch.CreateInterventionDictionary API synchronously
func (client *Client) CreateInterventionDictionary(request *CreateInterventionDictionaryRequest) (response *CreateInterventionDictionaryResponse, err error) {
	response = CreateCreateInterventionDictionaryResponse()
	err = client.DoAction(request, response)
	return
}

// CreateInterventionDictionaryWithChan invokes the opensearch.CreateInterventionDictionary API asynchronously
func (client *Client) CreateInterventionDictionaryWithChan(request *CreateInterventionDictionaryRequest) (<-chan *CreateInterventionDictionaryResponse, <-chan error) {
	responseChan := make(chan *CreateInterventionDictionaryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateInterventionDictionary(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// CreateInterventionDictionaryWithCallback invokes the opensearch.CreateInterventionDictionary API asynchronously
func (client *Client) CreateInterventionDictionaryWithCallback(request *CreateInterventionDictionaryRequest, callback func(response *CreateInterventionDictionaryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateInterventionDictionaryResponse
		var err error
		defer close(result)
		response, err = client.CreateInterventionDictionary(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// CreateInterventionDictionaryRequest is the request struct for api CreateInterventionDictionary
type CreateInterventionDictionaryRequest struct {
	*requests.RoaRequest
}

// CreateInterventionDictionaryResponse is the response struct for api CreateInterventionDictionary
type CreateInterventionDictionaryResponse struct {
	*responses.BaseResponse
	RequestId string                               `json:"requestId" xml:"requestId"`
	Result    ResultInCreateInterventionDictionary `json:"result" xml:"result"`
}

// CreateCreateInterventionDictionaryRequest creates a request to invoke CreateInterventionDictionary API
func CreateCreateInterventionDictionaryRequest() (request *CreateInterventionDictionaryRequest) {
	request = &CreateInterventionDictionaryRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "CreateInterventionDictionary", "/v4/openapi/intervention-dictionaries", "opensearch", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateInterventionDictionaryResponse creates a response to parse from CreateInterventionDictionary response
func CreateCreateInterventionDictionaryResponse() (response *CreateInterventionDictionaryResponse) {
	response = &CreateInterventionDictionaryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
