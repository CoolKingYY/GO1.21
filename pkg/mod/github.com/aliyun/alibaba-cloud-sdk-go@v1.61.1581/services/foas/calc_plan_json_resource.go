package foas

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

// CalcPlanJsonResource invokes the foas.CalcPlanJsonResource API synchronously
func (client *Client) CalcPlanJsonResource(request *CalcPlanJsonResourceRequest) (response *CalcPlanJsonResourceResponse, err error) {
	response = CreateCalcPlanJsonResourceResponse()
	err = client.DoAction(request, response)
	return
}

// CalcPlanJsonResourceWithChan invokes the foas.CalcPlanJsonResource API asynchronously
func (client *Client) CalcPlanJsonResourceWithChan(request *CalcPlanJsonResourceRequest) (<-chan *CalcPlanJsonResourceResponse, <-chan error) {
	responseChan := make(chan *CalcPlanJsonResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CalcPlanJsonResource(request)
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

// CalcPlanJsonResourceWithCallback invokes the foas.CalcPlanJsonResource API asynchronously
func (client *Client) CalcPlanJsonResourceWithCallback(request *CalcPlanJsonResourceRequest, callback func(response *CalcPlanJsonResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CalcPlanJsonResourceResponse
		var err error
		defer close(result)
		response, err = client.CalcPlanJsonResource(request)
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

// CalcPlanJsonResourceRequest is the request struct for api CalcPlanJsonResource
type CalcPlanJsonResourceRequest struct {
	*requests.RoaRequest
	ProjectName string `position:"Path" name:"projectName"`
	JobName     string `position:"Path" name:"jobName"`
}

// CalcPlanJsonResourceResponse is the response struct for api CalcPlanJsonResource
type CalcPlanJsonResourceResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	PlanJsonResource PlanJsonResource `json:"PlanJsonResource" xml:"PlanJsonResource"`
}

// CreateCalcPlanJsonResourceRequest creates a request to invoke CalcPlanJsonResource API
func CreateCalcPlanJsonResourceRequest() (request *CalcPlanJsonResourceRequest) {
	request = &CalcPlanJsonResourceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "CalcPlanJsonResource", "/api/v2/projects/[projectName]/jobs/[jobName]/planjson-resource", "foas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateCalcPlanJsonResourceResponse creates a response to parse from CalcPlanJsonResource response
func CreateCalcPlanJsonResourceResponse() (response *CalcPlanJsonResourceResponse) {
	response = &CalcPlanJsonResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
