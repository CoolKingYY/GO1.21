package domain

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

// SaveRegistrantProfileRealNameVerification invokes the domain.SaveRegistrantProfileRealNameVerification API synchronously
func (client *Client) SaveRegistrantProfileRealNameVerification(request *SaveRegistrantProfileRealNameVerificationRequest) (response *SaveRegistrantProfileRealNameVerificationResponse, err error) {
	response = CreateSaveRegistrantProfileRealNameVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// SaveRegistrantProfileRealNameVerificationWithChan invokes the domain.SaveRegistrantProfileRealNameVerification API asynchronously
func (client *Client) SaveRegistrantProfileRealNameVerificationWithChan(request *SaveRegistrantProfileRealNameVerificationRequest) (<-chan *SaveRegistrantProfileRealNameVerificationResponse, <-chan error) {
	responseChan := make(chan *SaveRegistrantProfileRealNameVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SaveRegistrantProfileRealNameVerification(request)
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

// SaveRegistrantProfileRealNameVerificationWithCallback invokes the domain.SaveRegistrantProfileRealNameVerification API asynchronously
func (client *Client) SaveRegistrantProfileRealNameVerificationWithCallback(request *SaveRegistrantProfileRealNameVerificationRequest, callback func(response *SaveRegistrantProfileRealNameVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SaveRegistrantProfileRealNameVerificationResponse
		var err error
		defer close(result)
		response, err = client.SaveRegistrantProfileRealNameVerification(request)
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

// SaveRegistrantProfileRealNameVerificationRequest is the request struct for api SaveRegistrantProfileRealNameVerification
type SaveRegistrantProfileRealNameVerificationRequest struct {
	*requests.RpcRequest
	Country                  string           `position:"Query" name:"Country"`
	IdentityCredentialType   string           `position:"Query" name:"IdentityCredentialType"`
	City                     string           `position:"Query" name:"City"`
	RegistrantProfileId      requests.Integer `position:"Query" name:"RegistrantProfileId"`
	IdentityCredential       string           `position:"Query" name:"IdentityCredential"`
	ZhCity                   string           `position:"Query" name:"ZhCity"`
	TelExt                   string           `position:"Query" name:"TelExt"`
	Province                 string           `position:"Query" name:"Province"`
	ZhRegistrantName         string           `position:"Query" name:"ZhRegistrantName"`
	PostalCode               string           `position:"Query" name:"PostalCode"`
	Lang                     string           `position:"Query" name:"Lang"`
	Email                    string           `position:"Query" name:"Email"`
	ZhRegistrantOrganization string           `position:"Query" name:"ZhRegistrantOrganization"`
	Address                  string           `position:"Query" name:"Address"`
	TelArea                  string           `position:"Query" name:"TelArea"`
	ZhAddress                string           `position:"Query" name:"ZhAddress"`
	RegistrantType           string           `position:"Query" name:"RegistrantType"`
	RegistrantProfileType    string           `position:"Query" name:"RegistrantProfileType"`
	Telephone                string           `position:"Query" name:"Telephone"`
	ZhProvince               string           `position:"Query" name:"ZhProvince"`
	RegistrantOrganization   string           `position:"Query" name:"RegistrantOrganization"`
	UserClientIp             string           `position:"Query" name:"UserClientIp"`
	IdentityCredentialNo     string           `position:"Query" name:"IdentityCredentialNo"`
	RegistrantName           string           `position:"Query" name:"RegistrantName"`
}

// SaveRegistrantProfileRealNameVerificationResponse is the response struct for api SaveRegistrantProfileRealNameVerification
type SaveRegistrantProfileRealNameVerificationResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId" xml:"RequestId"`
	RegistrantProfileId int64  `json:"RegistrantProfileId" xml:"RegistrantProfileId"`
}

// CreateSaveRegistrantProfileRealNameVerificationRequest creates a request to invoke SaveRegistrantProfileRealNameVerification API
func CreateSaveRegistrantProfileRealNameVerificationRequest() (request *SaveRegistrantProfileRealNameVerificationRequest) {
	request = &SaveRegistrantProfileRealNameVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "SaveRegistrantProfileRealNameVerification", "", "")
	request.Method = requests.POST
	return
}

// CreateSaveRegistrantProfileRealNameVerificationResponse creates a response to parse from SaveRegistrantProfileRealNameVerification response
func CreateSaveRegistrantProfileRealNameVerificationResponse() (response *SaveRegistrantProfileRealNameVerificationResponse) {
	response = &SaveRegistrantProfileRealNameVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}