package retailadvqa_public

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

// SaveBehaviorImportTask invokes the retailadvqa_public.SaveBehaviorImportTask API synchronously
func (client *Client) SaveBehaviorImportTask(request *SaveBehaviorImportTaskRequest) (response *SaveBehaviorImportTaskResponse, err error) {
	response = CreateSaveBehaviorImportTaskResponse()
	err = client.DoAction(request, response)
	return
}

// SaveBehaviorImportTaskWithChan invokes the retailadvqa_public.SaveBehaviorImportTask API asynchronously
func (client *Client) SaveBehaviorImportTaskWithChan(request *SaveBehaviorImportTaskRequest) (<-chan *SaveBehaviorImportTaskResponse, <-chan error) {
	responseChan := make(chan *SaveBehaviorImportTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveBehaviorImportTask(request)
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

// SaveBehaviorImportTaskWithCallback invokes the retailadvqa_public.SaveBehaviorImportTask API asynchronously
func (client *Client) SaveBehaviorImportTaskWithCallback(request *SaveBehaviorImportTaskRequest, callback func(response *SaveBehaviorImportTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveBehaviorImportTaskResponse
		var err error
		defer close(result)
		response, err = client.SaveBehaviorImportTask(request)
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

// SaveBehaviorImportTaskRequest is the request struct for api SaveBehaviorImportTask
type SaveBehaviorImportTaskRequest struct {
	*requests.RpcRequest
	BehaviorTypeGroupModelList        *[]SaveBehaviorImportTaskBehaviorTypeGroupModelList        `position:"Body" name:"BehaviorTypeGroupModelList"  type:"Json"`
	DatailTable                       requests.Boolean                                           `position:"Body" name:"DatailTable"`
	BehaviorImportTaskColumnModelList *[]SaveBehaviorImportTaskBehaviorImportTaskColumnModelList `position:"Body" name:"BehaviorImportTaskColumnModelList"  type:"Json"`
	AccessId                          string                                                     `position:"Query" name:"AccessId"`
	FilePath                          string                                                     `position:"Body" name:"FilePath"`
	Delimiter                         string                                                     `position:"Body" name:"Delimiter"`
	TenantId                          string                                                     `position:"Query" name:"TenantId"`
	FileFormat                        string                                                     `position:"Body" name:"FileFormat"`
	TableName                         string                                                     `position:"Body" name:"TableName"`
	TaskId                            string                                                     `position:"Body" name:"TaskId"`
	TableDesc                         string                                                     `position:"Body" name:"TableDesc"`
	WorkspaceId                       string                                                     `position:"Query" name:"WorkspaceId"`
}

// SaveBehaviorImportTaskBehaviorTypeGroupModelList is a repeated param struct in SaveBehaviorImportTaskRequest
type SaveBehaviorImportTaskBehaviorTypeGroupModelList struct {
	BehaviorType           string `name:"BehaviorType"`
	BehaviorTypeGroup      string `name:"BehaviorTypeGroup"`
	BehaviorTypeColumnName string `name:"BehaviorTypeColumnName"`
}

// SaveBehaviorImportTaskBehaviorImportTaskColumnModelList is a repeated param struct in SaveBehaviorImportTaskRequest
type SaveBehaviorImportTaskBehaviorImportTaskColumnModelList struct {
	IsAvailable    string `name:"IsAvailable"`
	IdentityType   string `name:"IdentityType"`
	FieldClassify  string `name:"FieldClassify"`
	EncryptionType string `name:"EncryptionType"`
	FieldAlias     string `name:"FieldAlias"`
	FieldType      string `name:"FieldType"`
	ColumnName     string `name:"ColumnName"`
	FieldSeparator string `name:"FieldSeparator"`
	FieldId        string `name:"FieldId"`
	ColumnType     string `name:"ColumnType"`
}

// SaveBehaviorImportTaskResponse is the response struct for api SaveBehaviorImportTask
type SaveBehaviorImportTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ErrorDesc string `json:"ErrorDesc" xml:"ErrorDesc"`
	TraceId   string `json:"TraceId" xml:"TraceId"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSaveBehaviorImportTaskRequest creates a request to invoke SaveBehaviorImportTask API
func CreateSaveBehaviorImportTaskRequest() (request *SaveBehaviorImportTaskRequest) {
	request = &SaveBehaviorImportTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailadvqa-public", "2020-05-15", "SaveBehaviorImportTask", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveBehaviorImportTaskResponse creates a response to parse from SaveBehaviorImportTask response
func CreateSaveBehaviorImportTaskResponse() (response *SaveBehaviorImportTaskResponse) {
	response = &SaveBehaviorImportTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
