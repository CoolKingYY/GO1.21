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

// GetInstanceResource invokes the foas.GetInstanceResource API synchronously
func (client *Client) GetInstanceResource(request *GetInstanceResourceRequest) (response *GetInstanceResourceResponse, err error) {
	response = CreateGetInstanceResourceResponse()
	err = client.DoAction(request, response)
	return
}

// GetInstanceResourceWithChan invokes the foas.GetInstanceResource API asynchronously
func (client *Client) GetInstanceResourceWithChan(request *GetInstanceResourceRequest) (<-chan *GetInstanceResourceResponse, <-chan error) {
	responseChan := make(chan *GetInstanceResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetInstanceResource(request)
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

// GetInstanceResourceWithCallback invokes the foas.GetInstanceResource API asynchronously
func (client *Client) GetInstanceResourceWithCallback(request *GetInstanceResourceRequest, callback func(response *GetInstanceResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetInstanceResourceResponse
		var err error
		defer close(result)
		response, err = client.GetInstanceResource(request)
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

// GetInstanceResourceRequest is the request struct for api GetInstanceResource
type GetInstanceResourceRequest struct {
	*requests.RoaRequest
	ProjectName string           `position:"Path" name:"projectName"`
	InstanceId  requests.Integer `position:"Path" name:"instanceId"`
	JobName     string           `position:"Path" name:"jobName"`
}

// GetInstanceResourceResponse is the response struct for api GetInstanceResource
type GetInstanceResourceResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Resource  Resource `json:"Resource" xml:"Resource"`
}

// CreateGetInstanceResourceRequest creates a request to invoke GetInstanceResource API
func CreateGetInstanceResourceRequest() (request *GetInstanceResourceRequest) {
	request = &GetInstanceResourceRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "GetInstanceResource", "/api/v2/projects/[projectName]/jobs/[jobName]/instances/[instanceId]/resource", "foas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetInstanceResourceResponse creates a response to parse from GetInstanceResource response
func CreateGetInstanceResourceResponse() (response *GetInstanceResourceResponse) {
	response = &GetInstanceResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}