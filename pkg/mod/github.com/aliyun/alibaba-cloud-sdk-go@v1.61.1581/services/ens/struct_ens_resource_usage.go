package ens

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

// EnsResourceUsage is a nested struct in ens response
type EnsResourceUsage struct {
	InstanceCount        int    `json:"InstanceCount" xml:"InstanceCount"`
	CpuSum               int64  `json:"CpuSum" xml:"CpuSum"`
	GpuSum               int64  `json:"GpuSum" xml:"GpuSum"`
	StorageSum           int64  `json:"StorageSum" xml:"StorageSum"`
	ServiceType          string `json:"ServiceType" xml:"ServiceType"`
	ExpiredCount         int    `json:"ExpiredCount" xml:"ExpiredCount"`
	RunningCount         int    `json:"RunningCount" xml:"RunningCount"`
	DownCount            int    `json:"DownCount" xml:"DownCount"`
	ExpiringCount        int    `json:"ExpiringCount" xml:"ExpiringCount"`
	DiskCount            int    `json:"DiskCount" xml:"DiskCount"`
	ComputeResourceCount int    `json:"ComputeResourceCount" xml:"ComputeResourceCount"`
}