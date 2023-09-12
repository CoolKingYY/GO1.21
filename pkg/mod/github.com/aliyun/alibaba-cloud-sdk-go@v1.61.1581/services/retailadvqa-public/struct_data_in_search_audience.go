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

// DataInSearchAudience is a nested struct in retailadvqa_public response
type DataInSearchAudience struct {
	DataModelName          string                   `json:"DataModelName" xml:"DataModelName"`
	GmtModified            int64                    `json:"GmtModified" xml:"GmtModified"`
	DbName                 string                   `json:"DbName" xml:"DbName"`
	NumberOfAudiences      int64                    `json:"NumberOfAudiences" xml:"NumberOfAudiences"`
	ErrorMessage           string                   `json:"ErrorMessage" xml:"ErrorMessage"`
	DbType                 string                   `json:"DbType" xml:"DbType"`
	Permission             string                   `json:"Permission" xml:"Permission"`
	Type                   int                      `json:"Type" xml:"Type"`
	GmtCreate              int64                    `json:"GmtCreate" xml:"GmtCreate"`
	Version                string                   `json:"Version" xml:"Version"`
	ParentId               string                   `json:"ParentId" xml:"ParentId"`
	ModifyUser             string                   `json:"ModifyUser" xml:"ModifyUser"`
	ModifyUserName         string                   `json:"ModifyUserName" xml:"ModifyUserName"`
	LatestDataModifyStatus int                      `json:"LatestDataModifyStatus" xml:"LatestDataModifyStatus"`
	Public                 bool                     `json:"Public" xml:"Public"`
	Subtype                int                      `json:"Subtype" xml:"Subtype"`
	Name                   string                   `json:"Name" xml:"Name"`
	AutoUpdateData         bool                     `json:"AutoUpdateData" xml:"AutoUpdateData"`
	CreateUser             string                   `json:"CreateUser" xml:"CreateUser"`
	Id                     string                   `json:"Id" xml:"Id"`
	LatestDataModifyTime   int64                    `json:"LatestDataModifyTime" xml:"LatestDataModifyTime"`
	Desc                   string                   `json:"Desc" xml:"Desc"`
	ExtendMappingTypes     []ExtendMappingTypesItem `json:"ExtendMappingTypes" xml:"ExtendMappingTypes"`
}