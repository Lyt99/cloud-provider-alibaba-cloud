package sls

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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// SetAlertCenterResource invokes the sls.SetAlertCenterResource API synchronously
func (client *Client) SetAlertCenterResource(request *SetAlertCenterResourceRequest) (response *SetAlertCenterResourceResponse, err error) {
	response = CreateSetAlertCenterResourceResponse()
	err = client.DoAction(request, response)
	return
}

// SetAlertCenterResourceWithChan invokes the sls.SetAlertCenterResource API asynchronously
func (client *Client) SetAlertCenterResourceWithChan(request *SetAlertCenterResourceRequest) (<-chan *SetAlertCenterResourceResponse, <-chan error) {
	responseChan := make(chan *SetAlertCenterResourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetAlertCenterResource(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// SetAlertCenterResourceWithCallback invokes the sls.SetAlertCenterResource API asynchronously
func (client *Client) SetAlertCenterResourceWithCallback(request *SetAlertCenterResourceRequest, callback func(response *SetAlertCenterResourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetAlertCenterResourceResponse
		var err error
		defer close(result)
		response, err = client.SetAlertCenterResource(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// SetAlertCenterResourceRequest is the request struct for api SetAlertCenterResource
type SetAlertCenterResourceRequest struct {
	*requests.RpcRequest
	App      string `position:"Body" name:"App"`
	Language string `position:"Body" name:"Language"`
	Region   string `position:"Body" name:"Region"`
}

// SetAlertCenterResourceResponse is the response struct for api SetAlertCenterResource
type SetAlertCenterResourceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Code      string `json:"Code" xml:"Code"`
}

// CreateSetAlertCenterResourceRequest creates a request to invoke SetAlertCenterResource API
func CreateSetAlertCenterResourceRequest() (request *SetAlertCenterResourceRequest) {
	request = &SetAlertCenterResourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sls", "2019-10-23", "SetAlertCenterResource", "sls", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSetAlertCenterResourceResponse creates a response to parse from SetAlertCenterResource response
func CreateSetAlertCenterResourceResponse() (response *SetAlertCenterResourceResponse) {
	response = &SetAlertCenterResourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
