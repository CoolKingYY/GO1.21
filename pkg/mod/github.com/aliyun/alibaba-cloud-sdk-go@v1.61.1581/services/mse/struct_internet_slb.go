package mse

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

// InternetSlb is a nested struct in mse response
type InternetSlb struct {
	SlbSpec             string `json:"SlbSpec" xml:"SlbSpec"`
	SlbPort             string `json:"SlbPort" xml:"SlbPort"`
	GatewaySlbStatus    string `json:"GatewaySlbStatus" xml:"GatewaySlbStatus"`
	GatewaySlbMode      string `json:"GatewaySlbMode" xml:"GatewaySlbMode"`
	InternetNetworkFlow string `json:"InternetNetworkFlow" xml:"InternetNetworkFlow"`
	StatusDesc          string `json:"StatusDesc" xml:"StatusDesc"`
	SlbId               string `json:"SlbId" xml:"SlbId"`
	SlbIp               string `json:"SlbIp" xml:"SlbIp"`
	Type                string `json:"Type" xml:"Type"`
}
