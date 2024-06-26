/*
Copyright 2024 KDP(Kubernetes Data Platform).

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package exception

var (
	ErrDefinitionNotFound = NewExceptCode(404, ErrDetail{
		ExceptionLevel: "error",
		ErrInfo: ErrInfo{
			Code:        600404,
			Description: "referenced definition not found, please check your specified def type",
			Solution:    "",
			ManualURL:   "",
		},
		AppName: "",
	})
)
