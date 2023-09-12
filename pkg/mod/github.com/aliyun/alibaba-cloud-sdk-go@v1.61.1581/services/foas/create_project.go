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

// CreateProject invokes the foas.CreateProject API synchronously
func (client *Client) CreateProject(request *CreateProjectRequest) (response *CreateProjectResponse, err error) {
	response = CreateCreateProjectResponse()
	err = client.DoAction(request, response)
	return
}

// CreateProjectWithChan invokes the foas.CreateProject API asynchronously
func (client *Client) CreateProjectWithChan(request *CreateProjectRequest) (<-chan *CreateProjectResponse, <-chan error) {
	responseChan := make(chan *CreateProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateProject(request)
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

// CreateProjectWithCallback invokes the foas.CreateProject API asynchronously
func (client *Client) CreateProjectWithCallback(request *CreateProjectRequest, callback func(response *CreateProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateProjectResponse
		var err error
		defer close(result)
		response, err = client.CreateProject(request)
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

// CreateProjectRequest is the request struct for api CreateProject
type CreateProjectRequest struct {
	*requests.RoaRequest
	ManagerIds  string `position:"Body" name:"managerIds"`
	OrderId     string `position:"Body" name:"orderId"`
	Name        string `position:"Body" name:"name"`
	Description string `position:"Body" name:"description"`
	ClusterId   string `position:"Body" name:"clusterId"`
	DeployType  string `position:"Body" name:"deployType"`
}

// CreateProjectResponse is the response struct for api CreateProject
type CreateProjectResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateProjectRequest creates a request to invoke CreateProject API
func CreateCreateProjectRequest() (request *CreateProjectRequest) {
	request = &CreateProjectRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "CreateProject", "/api/v2/projects", "foas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateProjectResponse creates a response to parse from CreateProject response
func CreateCreateProjectResponse() (response *CreateProjectResponse) {
	response = &CreateProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
