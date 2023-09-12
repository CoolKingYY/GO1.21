package airec

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

// PushIntervention invokes the airec.PushIntervention API synchronously
func (client *Client) PushIntervention(request *PushInterventionRequest) (response *PushInterventionResponse, err error) {
	response = CreatePushInterventionResponse()
	err = client.DoAction(request, response)
	return
}

// PushInterventionWithChan invokes the airec.PushIntervention API asynchronously
func (client *Client) PushInterventionWithChan(request *PushInterventionRequest) (<-chan *PushInterventionResponse, <-chan error) {
	responseChan := make(chan *PushInterventionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.PushIntervention(request)
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

// PushInterventionWithCallback invokes the airec.PushIntervention API asynchronously
func (client *Client) PushInterventionWithCallback(request *PushInterventionRequest, callback func(response *PushInterventionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *PushInterventionResponse
		var err error
		defer close(result)
		response, err = client.PushIntervention(request)
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

// PushInterventionRequest is the request struct for api PushIntervention
type PushInterventionRequest struct {
	*requests.RoaRequest
	InstanceId string `position:"Path" name:"instanceId"`
}

// PushInterventionResponse is the response struct for api PushIntervention
type PushInterventionResponse struct {
	*responses.BaseResponse
	Code      string `json:"code" xml:"code"`
	Message   string `json:"message" xml:"message"`
	RequestId string `json:"requestId" xml:"requestId"`
	Result    bool   `json:"result" xml:"result"`
}

// CreatePushInterventionRequest creates a request to invoke PushIntervention API
func CreatePushInterventionRequest() (request *PushInterventionRequest) {
	request = &PushInterventionRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Airec", "2020-11-26", "PushIntervention", "/v2/openapi/instances/[instanceId]/actions/intervene", "airec", "openAPI")
	request.Method = requests.POST
	return
}

// CreatePushInterventionResponse creates a response to parse from PushIntervention response
func CreatePushInterventionResponse() (response *PushInterventionResponse) {
	response = &PushInterventionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
