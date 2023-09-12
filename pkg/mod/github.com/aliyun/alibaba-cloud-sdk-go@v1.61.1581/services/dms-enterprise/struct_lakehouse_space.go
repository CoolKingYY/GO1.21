package dms_enterprise

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

// LakehouseSpace is a nested struct in dms_enterprise response
type LakehouseSpace struct {
	Id          int64  `json:"Id" xml:"Id"`
	SpaceName   string `json:"SpaceName" xml:"SpaceName"`
	CreatorId   string `json:"CreatorId" xml:"CreatorId"`
	TenantId    string `json:"TenantId" xml:"TenantId"`
	Description string `json:"Description" xml:"Description"`
	Mode        int    `json:"Mode" xml:"Mode"`
	DwDbType    string `json:"DwDbType" xml:"DwDbType"`
	SpaceConfig string `json:"SpaceConfig" xml:"SpaceConfig"`
	DevDbId     int    `json:"DevDbId" xml:"DevDbId"`
	ProdDbId    int    `json:"ProdDbId" xml:"ProdDbId"`
	IsDeleted   bool   `json:"IsDeleted" xml:"IsDeleted"`
}