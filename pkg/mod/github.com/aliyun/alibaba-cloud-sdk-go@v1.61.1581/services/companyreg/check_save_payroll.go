package companyreg

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

// CheckSavePayroll invokes the companyreg.CheckSavePayroll API synchronously
func (client *Client) CheckSavePayroll(request *CheckSavePayrollRequest) (response *CheckSavePayrollResponse, err error) {
	response = CreateCheckSavePayrollResponse()
	err = client.DoAction(request, response)
	return
}

// CheckSavePayrollWithChan invokes the companyreg.CheckSavePayroll API asynchronously
func (client *Client) CheckSavePayrollWithChan(request *CheckSavePayrollRequest) (<-chan *CheckSavePayrollResponse, <-chan error) {
	responseChan := make(chan *CheckSavePayrollResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckSavePayroll(request)
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

// CheckSavePayrollWithCallback invokes the companyreg.CheckSavePayroll API asynchronously
func (client *Client) CheckSavePayrollWithCallback(request *CheckSavePayrollRequest, callback func(response *CheckSavePayrollResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckSavePayrollResponse
		var err error
		defer close(result)
		response, err = client.CheckSavePayroll(request)
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

// CheckSavePayrollRequest is the request struct for api CheckSavePayroll
type CheckSavePayrollRequest struct {
	*requests.RpcRequest
	Income                           string           `position:"Query" name:"Income"`
	CorporateHousingAccumulationFund string           `position:"Query" name:"CorporateHousingAccumulationFund"`
	Period                           string           `position:"Query" name:"Period"`
	CorporateSocialInsurance         string           `position:"Query" name:"CorporateSocialInsurance"`
	IdNo                             string           `position:"Query" name:"IdNo"`
	EmployeeTime                     string           `position:"Query" name:"EmployeeTime"`
	PersonHousingAccumulationFund    string           `position:"Query" name:"PersonHousingAccumulationFund"`
	UpdateEmployeeFlag               requests.Boolean `position:"Query" name:"UpdateEmployeeFlag"`
	Phone                            string           `position:"Query" name:"Phone"`
	BizId                            string           `position:"Query" name:"BizId"`
	Name                             string           `position:"Query" name:"Name"`
	Id                               requests.Integer `position:"Query" name:"Id"`
	PersonSocialInsurance            string           `position:"Query" name:"PersonSocialInsurance"`
}

// CheckSavePayrollResponse is the response struct for api CheckSavePayroll
type CheckSavePayrollResponse struct {
	*responses.BaseResponse
	IsExists  bool   `json:"IsExists" xml:"IsExists"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    bool   `json:"Result" xml:"Result"`
}

// CreateCheckSavePayrollRequest creates a request to invoke CheckSavePayroll API
func CreateCheckSavePayrollRequest() (request *CheckSavePayrollRequest) {
	request = &CheckSavePayrollRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2020-10-22", "CheckSavePayroll", "", "")
	request.Method = requests.POST
	return
}

// CreateCheckSavePayrollResponse creates a response to parse from CheckSavePayroll response
func CreateCheckSavePayrollResponse() (response *CheckSavePayrollResponse) {
	response = &CheckSavePayrollResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
