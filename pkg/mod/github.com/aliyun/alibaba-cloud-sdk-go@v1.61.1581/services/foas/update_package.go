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

// UpdatePackage invokes the foas.UpdatePackage API synchronously
func (client *Client) UpdatePackage(request *UpdatePackageRequest) (response *UpdatePackageResponse, err error) {
	response = CreateUpdatePackageResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePackageWithChan invokes the foas.UpdatePackage API asynchronously
func (client *Client) UpdatePackageWithChan(request *UpdatePackageRequest) (<-chan *UpdatePackageResponse, <-chan error) {
	responseChan := make(chan *UpdatePackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePackage(request)
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

// UpdatePackageWithCallback invokes the foas.UpdatePackage API asynchronously
func (client *Client) UpdatePackageWithCallback(request *UpdatePackageRequest, callback func(response *UpdatePackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePackageResponse
		var err error
		defer close(result)
		response, err = client.UpdatePackage(request)
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

// UpdatePackageRequest is the request struct for api UpdatePackage
type UpdatePackageRequest struct {
	*requests.RoaRequest
	ProjectName string `position:"Path" name:"projectName"`
	OssBucket   string `position:"Body" name:"ossBucket"`
	OssOwner    string `position:"Body" name:"ossOwner"`
	PackageName string `position:"Path" name:"packageName"`
	OssEndpoint string `position:"Body" name:"ossEndpoint"`
	Description string `position:"Body" name:"description"`
	Tag         string `position:"Body" name:"tag"`
	OriginName  string `position:"Body" name:"originName"`
	OssPath     string `position:"Body" name:"ossPath"`
	Md5         string `position:"Body" name:"md5"`
}

// UpdatePackageResponse is the response struct for api UpdatePackage
type UpdatePackageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUpdatePackageRequest creates a request to invoke UpdatePackage API
func CreateUpdatePackageRequest() (request *UpdatePackageRequest) {
	request = &UpdatePackageRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("foas", "2018-11-11", "UpdatePackage", "/api/v2/projects/[projectName]/packages/[packageName]", "foas", "openAPI")
	request.Method = requests.PUT
	return
}

// CreateUpdatePackageResponse creates a response to parse from UpdatePackage response
func CreateUpdatePackageResponse() (response *UpdatePackageResponse) {
	response = &UpdatePackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
