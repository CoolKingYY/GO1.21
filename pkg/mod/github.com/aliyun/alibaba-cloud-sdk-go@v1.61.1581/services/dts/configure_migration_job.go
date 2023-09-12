package dts

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

// ConfigureMigrationJob invokes the dts.ConfigureMigrationJob API synchronously
func (client *Client) ConfigureMigrationJob(request *ConfigureMigrationJobRequest) (response *ConfigureMigrationJobResponse, err error) {
	response = CreateConfigureMigrationJobResponse()
	err = client.DoAction(request, response)
	return
}

// ConfigureMigrationJobWithChan invokes the dts.ConfigureMigrationJob API asynchronously
func (client *Client) ConfigureMigrationJobWithChan(request *ConfigureMigrationJobRequest) (<-chan *ConfigureMigrationJobResponse, <-chan error) {
	responseChan := make(chan *ConfigureMigrationJobResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ConfigureMigrationJob(request)
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

// ConfigureMigrationJobWithCallback invokes the dts.ConfigureMigrationJob API asynchronously
func (client *Client) ConfigureMigrationJobWithCallback(request *ConfigureMigrationJobRequest, callback func(response *ConfigureMigrationJobResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ConfigureMigrationJobResponse
		var err error
		defer close(result)
		response, err = client.ConfigureMigrationJob(request)
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

// ConfigureMigrationJobRequest is the request struct for api ConfigureMigrationJob
type ConfigureMigrationJobRequest struct {
	*requests.RpcRequest
	SourceEndpointInstanceID            string           `position:"Query" name:"SourceEndpoint.InstanceID"`
	Checkpoint                          string           `position:"Query" name:"Checkpoint"`
	SourceEndpointEngineName            string           `position:"Query" name:"SourceEndpoint.EngineName"`
	SourceEndpointOracleSID             string           `position:"Query" name:"SourceEndpoint.OracleSID"`
	DestinationEndpointInstanceID       string           `position:"Query" name:"DestinationEndpoint.InstanceID"`
	SourceEndpointIP                    string           `position:"Query" name:"SourceEndpoint.IP"`
	DestinationEndpointPassword         string           `position:"Query" name:"DestinationEndpoint.Password"`
	MigrationObject                     string           `position:"Body" name:"MigrationObject"`
	MigrationModeDataIntialization      requests.Boolean `position:"Query" name:"MigrationMode.DataIntialization"`
	MigrationJobId                      string           `position:"Query" name:"MigrationJobId"`
	SourceEndpointInstanceType          string           `position:"Query" name:"SourceEndpoint.InstanceType"`
	DestinationEndpointEngineName       string           `position:"Query" name:"DestinationEndpoint.EngineName"`
	AccountId                           string           `position:"Query" name:"AccountId"`
	MigrationModeStructureIntialization requests.Boolean `position:"Query" name:"MigrationMode.StructureIntialization"`
	MigrationModeDataSynchronization    requests.Boolean `position:"Query" name:"MigrationMode.DataSynchronization"`
	DestinationEndpointRegion           string           `position:"Query" name:"DestinationEndpoint.Region"`
	SourceEndpointUserName              string           `position:"Query" name:"SourceEndpoint.UserName"`
	SourceEndpointDatabaseName          string           `position:"Query" name:"SourceEndpoint.DatabaseName"`
	SourceEndpointPort                  string           `position:"Query" name:"SourceEndpoint.Port"`
	SourceEndpointOwnerID               string           `position:"Query" name:"SourceEndpoint.OwnerID"`
	DestinationEndpointUserName         string           `position:"Query" name:"DestinationEndpoint.UserName"`
	DestinationEndpointOracleSID        string           `position:"Query" name:"DestinationEndpoint.OracleSID"`
	DestinationEndpointPort             string           `position:"Query" name:"DestinationEndpoint.Port"`
	SourceEndpointRegion                string           `position:"Query" name:"SourceEndpoint.Region"`
	SourceEndpointRole                  string           `position:"Query" name:"SourceEndpoint.Role"`
	OwnerId                             string           `position:"Query" name:"OwnerId"`
	DestinationEndpointDataBaseName     string           `position:"Query" name:"DestinationEndpoint.DataBaseName"`
	SourceEndpointPassword              string           `position:"Query" name:"SourceEndpoint.Password"`
	MigrationReserved                   string           `position:"Query" name:"MigrationReserved"`
	DestinationEndpointIP               string           `position:"Query" name:"DestinationEndpoint.IP"`
	MigrationJobName                    string           `position:"Query" name:"MigrationJobName"`
	DestinationEndpointInstanceType     string           `position:"Query" name:"DestinationEndpoint.InstanceType"`
}

// ConfigureMigrationJobResponse is the response struct for api ConfigureMigrationJob
type ConfigureMigrationJobResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	ErrCode    string `json:"ErrCode" xml:"ErrCode"`
	Success    string `json:"Success" xml:"Success"`
	ErrMessage string `json:"ErrMessage" xml:"ErrMessage"`
}

// CreateConfigureMigrationJobRequest creates a request to invoke ConfigureMigrationJob API
func CreateConfigureMigrationJobRequest() (request *ConfigureMigrationJobRequest) {
	request = &ConfigureMigrationJobRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dts", "2020-01-01", "ConfigureMigrationJob", "dts", "openAPI")
	request.Method = requests.POST
	return
}

// CreateConfigureMigrationJobResponse creates a response to parse from ConfigureMigrationJob response
func CreateConfigureMigrationJobResponse() (response *ConfigureMigrationJobResponse) {
	response = &ConfigureMigrationJobResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
