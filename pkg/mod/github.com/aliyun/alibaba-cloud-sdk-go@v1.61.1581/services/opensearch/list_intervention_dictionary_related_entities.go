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

// ListInterventionDictionaryRelatedEntities invokes the opensearch.ListInterventionDictionaryRelatedEntities API synchronously
func (client *Client) ListInterventionDictionaryRelatedEntities(request *ListInterventionDictionaryRelatedEntitiesRequest) (response *ListInterventionDictionaryRelatedEntitiesResponse, err error) {
	response = CreateListInterventionDictionaryRelatedEntitiesResponse()
	err = client.DoAction(request, response)
	return
}

// ListInterventionDictionaryRelatedEntitiesWithChan invokes the opensearch.ListInterventionDictionaryRelatedEntities API asynchronously
func (client *Client) ListInterventionDictionaryRelatedEntitiesWithChan(request *ListInterventionDictionaryRelatedEntitiesRequest) (<-chan *ListInterventionDictionaryRelatedEntitiesResponse, <-chan error) {
	responseChan := make(chan *ListInterventionDictionaryRelatedEntitiesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListInterventionDictionaryRelatedEntities(request)
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

// ListInterventionDictionaryRelatedEntitiesWithCallback invokes the opensearch.ListInterventionDictionaryRelatedEntities API asynchronously
func (client *Client) ListInterventionDictionaryRelatedEntitiesWithCallback(request *ListInterventionDictionaryRelatedEntitiesRequest, callback func(response *ListInterventionDictionaryRelatedEntitiesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListInterventionDictionaryRelatedEntitiesResponse
		var err error
		defer close(result)
		response, err = client.ListInterventionDictionaryRelatedEntities(request)
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

// ListInterventionDictionaryRelatedEntitiesRequest is the request struct for api ListInterventionDictionaryRelatedEntities
type ListInterventionDictionaryRelatedEntitiesRequest struct {
	*requests.RoaRequest
	Name string `position:"Path" name:"name"`
}

// ListInterventionDictionaryRelatedEntitiesResponse is the response struct for api ListInterventionDictionaryRelatedEntities
type ListInterventionDictionaryRelatedEntitiesResponse struct {
	*responses.BaseResponse
	RequestId string                   `json:"requestId" xml:"requestId"`
	Result    []map[string]interface{} `json:"result" xml:"result"`
}

// CreateListInterventionDictionaryRelatedEntitiesRequest creates a request to invoke ListInterventionDictionaryRelatedEntities API
func CreateListInterventionDictionaryRelatedEntitiesRequest() (request *ListInterventionDictionaryRelatedEntitiesRequest) {
	request = &ListInterventionDictionaryRelatedEntitiesRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListInterventionDictionaryRelatedEntities", "/v4/openapi/intervention-dictionaries/[name]/related", "opensearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListInterventionDictionaryRelatedEntitiesResponse creates a response to parse from ListInterventionDictionaryRelatedEntities response
func CreateListInterventionDictionaryRelatedEntitiesResponse() (response *ListInterventionDictionaryRelatedEntitiesResponse) {
	response = &ListInterventionDictionaryRelatedEntitiesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}