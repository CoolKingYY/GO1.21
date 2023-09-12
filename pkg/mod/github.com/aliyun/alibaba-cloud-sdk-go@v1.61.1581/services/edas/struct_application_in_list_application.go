package edas

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

// ApplicationInListApplication is a nested struct in edas response
type ApplicationInListApplication struct {
	ClusterType          int    `json:"ClusterType" xml:"ClusterType"`
	AppId                string `json:"AppId" xml:"AppId"`
	ResourceGroupId      string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ApplicationType      string `json:"ApplicationType" xml:"ApplicationType"`
	Name                 string `json:"Name" xml:"Name"`
	RunningInstanceCount int    `json:"RunningInstanceCount" xml:"RunningInstanceCount"`
	BuildPackageId       int64  `json:"BuildPackageId" xml:"BuildPackageId"`
	ClusterId            string `json:"ClusterId" xml:"ClusterId"`
	RegionId             string `json:"RegionId" xml:"RegionId"`
}
