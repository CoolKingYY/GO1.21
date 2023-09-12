package iot

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

// DataInQueryJobStatistics is a nested struct in iot response
type DataInQueryJobStatistics struct {
	Total      int `json:"Total" xml:"Total"`
	Queued     int `json:"Queued" xml:"Queued"`
	Sent       int `json:"Sent" xml:"Sent"`
	InProgress int `json:"InProgress" xml:"InProgress"`
	Succeeded  int `json:"Succeeded" xml:"Succeeded"`
	Failed     int `json:"Failed" xml:"Failed"`
	Rejected   int `json:"Rejected" xml:"Rejected"`
	TimeOut    int `json:"TimeOut" xml:"TimeOut"`
	Cancelled  int `json:"Cancelled" xml:"Cancelled"`
}
