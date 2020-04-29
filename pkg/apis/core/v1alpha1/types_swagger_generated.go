// Code generated by swagger-doc. DO NOT EDIT.

package v1alpha1

func (DataVolume) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataVolume provides a representation of our data volume\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}

func (DataVolumeSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":            "DataVolumeSpec defines our specification for a DataVolume type",
		"source":      "Source is the src of the data for the requested DataVolume",
		"pvc":         "PVC is a pointer to the PVC Spec we want to use",
		"contentType": "DataVolumeContentType options: \"kubevirt\", \"archive\"",
	}
}

func (DataVolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataVolumeSource represents the source for our Data Volume, this can be HTTP, Imageio, S3, Registry or an existing PVC",
	}
}

func (DataVolumeSourcePVC) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataVolumeSourcePVC provides the parameters to create a Data Volume from an existing PVC",
	}
}

func (DataVolumeBlankImage) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataVolumeBlankImage provides the parameters to create a new raw blank image for the PVC",
	}
}

func (DataVolumeSourceUpload) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataVolumeSourceUpload provides the parameters to create a Data Volume by uploading the source",
	}
}

func (DataVolumeSourceS3) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "DataVolumeSourceS3 provides the parameters to create a Data Volume from an S3 source",
		"url":       "URL is the url of the S3 source",
		"secretRef": "SecretRef provides the secret reference needed to access the S3 source",
	}
}

func (DataVolumeSourceRegistry) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "DataVolumeSourceRegistry provides the parameters to create a Data Volume from an registry source",
		"url":           "URL is the url of the Registry source",
		"secretRef":     "SecretRef provides the secret reference needed to access the Registry source",
		"certConfigMap": "CertConfigMap provides a reference to the Registry certs",
	}
}

func (DataVolumeSourceHTTP) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "DataVolumeSourceHTTP provides the parameters to create a Data Volume from an HTTP source",
		"url":           "URL is the URL of the http source",
		"secretRef":     "SecretRef provides the secret reference needed to access the HTTP source",
		"certConfigMap": "CertConfigMap provides a reference to the Registry certs",
	}
}

func (DataVolumeSourceImageIO) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "DataVolumeSourceImageIO provides the parameters to create a Data Volume from an imageio source",
		"url":           "URL is the URL of the ovirt-engine",
		"diskId":        "DiskID provides id of a disk to be imported",
		"secretRef":     "SecretRef provides the secret reference needed to access the ovirt-engine",
		"certConfigMap": "CertConfigMap provides a reference to the CA cert",
	}
}

func (DataVolumeStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":           "DataVolumeStatus provides the parameters to store the phase of the Data Volume",
		"phase":      "Phase is the current phase of the data volume",
		"conditions": "+listType=set",
	}
}

func (DataVolumeList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "DataVolumeList provides the needed parameters to do request a list of Data Volumes from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of DataVolumes",
	}
}

func (DataVolumeCondition) SwaggerDoc() map[string]string {
	return map[string]string{
		"":        "DataVolumeCondition represents the state of a data volume condition.",
		"reason":  "+optional",
		"message": "+optional",
	}
}

func (CDI) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "CDI is the CDI Operator CRD\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}

func (CDISpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "CDISpec defines our specification for the CDI installation",
	}
}

func (CDIStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "CDIStatus defines the status of the CDI installation",
	}
}

func (CDIList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "CDIList provides the needed parameters to do request a list of CDIs from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of CDIs",
	}
}

func (CDIConfig) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "CDIConfig provides a user configuration for CDI\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
	}
}

func (CDIConfigSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "CDIConfigSpec defines specification for user configuration",
	}
}

func (CDIConfigStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "CDIConfigStatus provides",
	}
}

func (CDIConfigList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "CDIConfigList provides the needed parameters to do request a list of CDIConfigs from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of CDIConfigs",
	}
}
