package dataworks_public

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

// UpdateFile invokes the dataworks_public.UpdateFile API synchronously
func (client *Client) UpdateFile(request *UpdateFileRequest) (response *UpdateFileResponse, err error) {
	response = CreateUpdateFileResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateFileWithChan invokes the dataworks_public.UpdateFile API asynchronously
func (client *Client) UpdateFileWithChan(request *UpdateFileRequest) (<-chan *UpdateFileResponse, <-chan error) {
	responseChan := make(chan *UpdateFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateFile(request)
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

// UpdateFileWithCallback invokes the dataworks_public.UpdateFile API asynchronously
func (client *Client) UpdateFileWithCallback(request *UpdateFileRequest, callback func(response *UpdateFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateFileResponse
		var err error
		defer close(result)
		response, err = client.UpdateFile(request)
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

// UpdateFileRequest is the request struct for api UpdateFile
type UpdateFileRequest struct {
	*requests.RpcRequest
	OutputList              string           `position:"Body" name:"OutputList"`
	DependentNodeIdList     string           `position:"Body" name:"DependentNodeIdList"`
	Content                 string           `position:"Body" name:"Content"`
	ProjectIdentifier       string           `position:"Body" name:"ProjectIdentifier"`
	StartImmediately        requests.Boolean `position:"Body" name:"StartImmediately"`
	ProjectId               requests.Integer `position:"Body" name:"ProjectId"`
	AdvancedSettings        string           `position:"Body" name:"AdvancedSettings"`
	StartEffectDate         requests.Integer `position:"Body" name:"StartEffectDate"`
	CycleType               string           `position:"Body" name:"CycleType"`
	FileId                  requests.Integer `position:"Body" name:"FileId"`
	AutoRerunIntervalMillis requests.Integer `position:"Body" name:"AutoRerunIntervalMillis"`
	Owner                   string           `position:"Body" name:"Owner"`
	InputList               string           `position:"Body" name:"InputList"`
	RerunMode               string           `position:"Body" name:"RerunMode"`
	ConnectionName          string           `position:"Body" name:"ConnectionName"`
	OutputParameters        string           `position:"Body" name:"OutputParameters"`
	ParaValue               string           `position:"Body" name:"ParaValue"`
	ResourceGroupIdentifier string           `position:"Body" name:"ResourceGroupIdentifier"`
	AutoRerunTimes          requests.Integer `position:"Body" name:"AutoRerunTimes"`
	CronExpress             string           `position:"Body" name:"CronExpress"`
	EndEffectDate           requests.Integer `position:"Body" name:"EndEffectDate"`
	FileName                string           `position:"Body" name:"FileName"`
	InputParameters         string           `position:"Body" name:"InputParameters"`
	Stop                    requests.Boolean `position:"Body" name:"Stop"`
	DependentType           string           `position:"Body" name:"DependentType"`
	FileFolderPath          string           `position:"Body" name:"FileFolderPath"`
	FileDescription         string           `position:"Body" name:"FileDescription"`
	AutoParsing             requests.Boolean `position:"Body" name:"AutoParsing"`
	SchedulerType           string           `position:"Body" name:"SchedulerType"`
}

// UpdateFileResponse is the response struct for api UpdateFile
type UpdateFileResponse struct {
	*responses.BaseResponse
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateUpdateFileRequest creates a request to invoke UpdateFile API
func CreateUpdateFileRequest() (request *UpdateFileRequest) {
	request = &UpdateFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateFile", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateFileResponse creates a response to parse from UpdateFile response
func CreateUpdateFileResponse() (response *UpdateFileResponse) {
	response = &UpdateFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}