package imm

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

// ImageQuality is a nested struct in imm response
type ImageQuality struct {
	ContrastScore    float64 `json:"ContrastScore" xml:"ContrastScore"`
	ClarityScore     float64 `json:"ClarityScore" xml:"ClarityScore"`
	ColorScore       float64 `json:"ColorScore" xml:"ColorScore"`
	OverallScore     float64 `json:"OverallScore" xml:"OverallScore"`
	Contrast         float64 `json:"Contrast" xml:"Contrast"`
	Exposure         float64 `json:"Exposure" xml:"Exposure"`
	CompositionScore float64 `json:"CompositionScore" xml:"CompositionScore"`
	Clarity          float64 `json:"Clarity" xml:"Clarity"`
	Color            float64 `json:"Color" xml:"Color"`
	ExposureScore    float64 `json:"ExposureScore" xml:"ExposureScore"`
}