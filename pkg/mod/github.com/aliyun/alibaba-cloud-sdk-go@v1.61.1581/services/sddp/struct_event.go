package sddp

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

// Event is a nested struct in sddp response
type Event struct {
	EventTime         int64        `json:"EventTime" xml:"EventTime"`
	DealUserIdValue   string       `json:"dealUserIdValue" xml:"dealUserIdValue"`
	TypeName          string       `json:"TypeName" xml:"TypeName"`
	SubTypeCode       string       `json:"SubTypeCode" xml:"SubTypeCode"`
	AlertTime         int64        `json:"AlertTime" xml:"AlertTime"`
	DealReason        string       `json:"DealReason" xml:"DealReason"`
	TargetProductCode string       `json:"TargetProductCode" xml:"TargetProductCode"`
	DealDisplayName   string       `json:"DealDisplayName" xml:"DealDisplayName"`
	Backed            bool         `json:"Backed" xml:"Backed"`
	UserIdValue       string       `json:"UserIdValue" xml:"UserIdValue"`
	DealLoginName     string       `json:"DealLoginName" xml:"DealLoginName"`
	DataInstance      string       `json:"DataInstance" xml:"DataInstance"`
	LogDetail         string       `json:"LogDetail" xml:"LogDetail"`
	TypeCode          string       `json:"TypeCode" xml:"TypeCode"`
	Status            int          `json:"Status" xml:"Status"`
	DealTime          int64        `json:"DealTime" xml:"DealTime"`
	DealUserId        int64        `json:"DealUserId" xml:"DealUserId"`
	DepartName        string       `json:"DepartName" xml:"DepartName"`
	InstanceName      string       `json:"InstanceName" xml:"InstanceName"`
	Id                int64        `json:"Id" xml:"Id"`
	ProductCode       string       `json:"ProductCode" xml:"ProductCode"`
	DisplayName       string       `json:"DisplayName" xml:"DisplayName"`
	StatusName        string       `json:"StatusName" xml:"StatusName"`
	WarnLevel         int          `json:"WarnLevel" xml:"WarnLevel"`
	LoginName         string       `json:"LoginName" xml:"LoginName"`
	SubTypeName       string       `json:"SubTypeName" xml:"SubTypeName"`
	UserId            int64        `json:"UserId" xml:"UserId"`
	Detail            Detail       `json:"Detail" xml:"Detail"`
	HandleInfoList    []HandleInfo `json:"HandleInfoList" xml:"HandleInfoList"`
}
