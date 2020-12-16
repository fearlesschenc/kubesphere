/*
Copyright 2020 The KubeSphere Authors.

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

package v1alpha2

import (
	"github.com/fearlesschenc/kubesphere/pkg/client/clientset/versioned"
	"github.com/fearlesschenc/kubesphere/pkg/client/informers/externalversions"
	"github.com/fearlesschenc/kubesphere/pkg/models/devops"
	devopsClient "github.com/fearlesschenc/kubesphere/pkg/simple/client/devops"
	"github.com/fearlesschenc/kubesphere/pkg/simple/client/s3"
	"github.com/fearlesschenc/kubesphere/pkg/simple/client/sonarqube"
)

type ProjectPipelineHandler struct {
	devopsOperator          devops.DevopsOperator
	projectCredentialGetter devops.ProjectCredentialGetter
}

type PipelineSonarHandler struct {
	pipelineSonarGetter devops.PipelineSonarGetter
}

func NewProjectPipelineHandler(devopsClient devopsClient.Interface) ProjectPipelineHandler {
	return ProjectPipelineHandler{
		devopsOperator:          devops.NewDevopsOperator(devopsClient, nil, nil, nil, nil),
		projectCredentialGetter: devops.NewProjectCredentialOperator(devopsClient),
	}
}

func NewPipelineSonarHandler(devopsClient devopsClient.Interface, sonarClient sonarqube.SonarInterface) PipelineSonarHandler {
	return PipelineSonarHandler{
		pipelineSonarGetter: devops.NewPipelineSonarGetter(devopsClient, sonarClient),
	}
}

func NewS2iBinaryHandler(client versioned.Interface, informers externalversions.SharedInformerFactory, s3Client s3.Interface) S2iBinaryHandler {
	return S2iBinaryHandler{devops.NewS2iBinaryUploader(client, informers, s3Client)}
}
