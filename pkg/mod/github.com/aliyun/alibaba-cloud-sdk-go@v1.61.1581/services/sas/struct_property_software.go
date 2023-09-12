package sas

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

// PropertySoftware is a nested struct in sas response
type PropertySoftware struct {
	Name            string `json:"Name" xml:"Name"`
	Path            string `json:"Path" xml:"Path"`
	InstanceName    string `json:"InstanceName" xml:"InstanceName"`
	Ip              string `json:"Ip" xml:"Ip"`
	Create          int64  `json:"Create" xml:"Create"`
	CreateTimestamp int64  `json:"CreateTimestamp" xml:"CreateTimestamp"`
	InstallTime     string `json:"InstallTime" xml:"InstallTime"`
	Version         string `json:"Version" xml:"Version"`
	Uuid            string `json:"Uuid" xml:"Uuid"`
	InstanceId      string `json:"InstanceId" xml:"InstanceId"`
	IntranetIp      string `json:"IntranetIp" xml:"IntranetIp"`
	InternetIp      string `json:"InternetIp" xml:"InternetIp"`
}
