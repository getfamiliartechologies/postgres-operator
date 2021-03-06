package util

/*
 Copyright 2020 Crunchy Data Solutions, Inc.
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

import "fmt"

// exporterSecretFormat is the format of the name of the exporter secret, i.e.
// "<clusterNamer>-exporter-secret"
// #nosec G101
const exporterSecretFormat = "%s-exporter-secret"

// GenerateExporterSecretName returns the name of the secret that contains
// information around a monitoring user
func GenerateExporterSecretName(clusterName string) string {
	return fmt.Sprintf(exporterSecretFormat, clusterName)
}
