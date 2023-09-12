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

// DataInGetSoundCodeSchedule is a nested struct in iot response
type DataInGetSoundCodeSchedule struct {
	ScheduleCode string `json:"ScheduleCode" xml:"ScheduleCode"`
	Name         string `json:"Name" xml:"Name"`
	Description  string `json:"Description" xml:"Description"`
	GmtCreate    int64  `json:"GmtCreate" xml:"GmtCreate"`
	Status       string `json:"Status" xml:"Status"`
	StartTime    string `json:"StartTime" xml:"StartTime"`
	EndTime      string `json:"EndTime" xml:"EndTime"`
	StartDate    string `json:"StartDate" xml:"StartDate"`
	EndDate      string `json:"EndDate" xml:"EndDate"`
}
