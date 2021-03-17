// Code generated by swagger-doc. DO NOT EDIT.

package v1beta1

func (DataVolume) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataVolume is an abstraction on top of PersistentVolumeClaims to allow easy population of those PersistentVolumeClaims with relation to VirtualMachines\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion\n+kubebuilder:resource:shortName=dv;dvs,categories=all\n+kubebuilder:printcolumn:name=\"Phase\",type=\"string\",JSONPath=\".status.phase\",description=\"The phase the data volume is in\"\n+kubebuilder:printcolumn:name=\"Progress\",type=\"string\",JSONPath=\".status.progress\",description=\"Transfer progress in percentage if known, N/A otherwise\"\n+kubebuilder:printcolumn:name=\"Restarts\",type=\"integer\",JSONPath=\".status.restartCount\",description=\"The number of times the transfer has been restarted.\"\n+kubebuilder:printcolumn:name=\"Age\",type=\"date\",JSONPath=\".metadata.creationTimestamp\"",
	}
}

func (DataVolumeSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                "DataVolumeSpec defines the DataVolume type specification",
		"source":          "Source is the src of the data for the requested DataVolume",
		"pvc":             "PVC is the PVC specification",
		"pod":             "Pod is the Pod specification for Importer or Uploader pod",
		"contentType":     "DataVolumeContentType options: \"kubevirt\", \"archive\"\n+kubebuilder:validation:Enum=\"kubevirt\";\"archive\"",
		"checkpoints":     "Checkpoints is a list of DataVolumeCheckpoints, representing stages in a multistage import.",
		"finalCheckpoint": "FinalCheckpoint indicates whether the current DataVolumeCheckpoint is the final checkpoint.",
		"preallocation":   "Preallocation controls whether storage for DataVolumes should be allocated in advance.",
	}
}

func (DataVolumeCheckpoint) SwaggerDoc() map[string]string {
	return map[string]string{
		"":         "DataVolumeCheckpoint defines a stage in a warm migration.",
		"previous": "Previous is the identifier of the snapshot from the previous checkpoint.",
		"current":  "Current is the identifier of the snapshot created for this checkpoint.",
	}
}

func (DataVolumeSource) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "DataVolumeSource represents the source for our Data Volume, this can be HTTP, Imageio, S3, Registry or an existing PVC",
	}
}

func (PodSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "PodSpec represents Pod specification for importer/uploader pod",
	}
}

func (DataVolumeSourcePVC) SwaggerDoc() map[string]string {
	return map[string]string{
		"":          "DataVolumeSourcePVC provides the parameters to create a Data Volume from an existing PVC",
		"namespace": "The namespace of the source PVC",
		"name":      "The name of the source PVC",
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
		"":              "DataVolumeSourceS3 provides the parameters to create a Data Volume from an S3 source",
		"url":           "URL is the url of the S3 source",
		"secretRef":     "SecretRef provides the secret reference needed to access the S3 source",
		"certConfigMap": "CertConfigMap is a configmap reference, containing a Certificate Authority(CA) public key, and a base64 encoded pem certificate\n+optional",
	}
}

func (DataVolumeSourceRegistry) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "DataVolumeSourceRegistry provides the parameters to create a Data Volume from an registry source",
		"url":           "URL is the url of the Docker registry source",
		"secretRef":     "SecretRef provides the secret reference needed to access the Registry source",
		"certConfigMap": "CertConfigMap provides a reference to the Registry certs",
	}
}

func (DataVolumeSourceHTTP) SwaggerDoc() map[string]string {
	return map[string]string{
		"":              "DataVolumeSourceHTTP can be either an http or https endpoint, with an optional basic auth user name and password, and an optional configmap containing additional CAs",
		"url":           "URL is the URL of the http(s) endpoint",
		"secretRef":     "SecretRef A Secret reference, the secret should contain accessKeyId (user name) base64 encoded, and secretKey (password) also base64 encoded\n+optional",
		"certConfigMap": "CertConfigMap is a configmap reference, containing a Certificate Authority(CA) public key, and a base64 encoded pem certificate\n+optional",
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

func (DataVolumeSourceVDDK) SwaggerDoc() map[string]string {
	return map[string]string{
		"":            "DataVolumeSourceVDDK provides the parameters to create a Data Volume from a Vmware source",
		"url":         "URL is the URL of the vCenter or ESXi host with the VM to migrate",
		"uuid":        "UUID is the UUID of the virtual machine that the backing file is attached to in vCenter/ESXi",
		"backingFile": "BackingFile is the path to the virtual hard disk to migrate from vCenter/ESXi",
		"thumbprint":  "Thumbprint is the certificate thumbprint of the vCenter or ESXi host",
		"secretRef":   "SecretRef provides a reference to a secret containing the username and password needed to access the vCenter or ESXi host",
	}
}

func (DataVolumeStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":             "DataVolumeStatus contains the current status of the DataVolume",
		"phase":        "Phase is the current phase of the data volume",
		"restartCount": "RestartCount is the number of times the pod populating the DataVolume has restarted",
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
		"": "DataVolumeCondition represents the state of a data volume condition.",
	}
}

func (StorageProfile) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "StorageProfile provides a CDI specific recommendation for storage parameters\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion\n+kubebuilder:resource:scope=Cluster",
	}
}

func (StorageProfileSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                  "StorageProfileSpec defines specification for StorageProfile",
		"claimPropertySets": "ClaimPropertySets is a provided set of properties applicable to PVC",
	}
}

func (StorageProfileStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                  "StorageProfileStatus provides the most recently observed status of the StorageProfile",
		"storageClass":      "The StorageClass name for which capabilities are defined",
		"provisioner":       "The Storage class provisioner plugin name",
		"claimPropertySets": "ClaimPropertySets computed from the spec and detected in the system",
	}
}

func (ClaimPropertySet) SwaggerDoc() map[string]string {
	return map[string]string{
		"":            "ClaimPropertySet is a set of properties applicable to PVC",
		"accessModes": "AccessModes contains the desired access modes the volume should have.\nMore info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1\n+optional",
		"volumeMode":  "volumeMode defines what type of volume is required by the claim.\nValue of Filesystem is implied when not included in claim spec.\n+optional",
	}
}

func (StorageProfileList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "StorageProfileList provides the needed parameters to request a list of StorageProfile from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of StorageProfile",
	}
}

func (CDI) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "CDI is the CDI Operator CRD\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion\n+kubebuilder:resource:shortName=cdi;cdis,scope=Cluster\n+kubebuilder:printcolumn:name=\"Age\",type=\"date\",JSONPath=\".metadata.creationTimestamp\"\n+kubebuilder:printcolumn:name=\"Phase\",type=\"string\",JSONPath=\".status.phase\"",
		"status": "+optional",
	}
}

func (CertConfig) SwaggerDoc() map[string]string {
	return map[string]string{
		"":            "CertConfig contains the tunables for TLS certificates",
		"duration":    "The requested 'duration' (i.e. lifetime) of the Certificate.",
		"renewBefore": "The amount of time before the currently issued certificate's `notAfter`\ntime that we will begin to attempt to renew the certificate.",
	}
}

func (CDICertConfig) SwaggerDoc() map[string]string {
	return map[string]string{
		"":       "CDICertConfig has the CertConfigs for CDI",
		"ca":     "CA configuration\nCA certs are kept in the CA bundle as long as they are valid",
		"server": "Server configuration\nCerts are rotated and discarded",
	}
}

func (CDISpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                      "CDISpec defines our specification for the CDI installation",
		"imagePullPolicy":       "+kubebuilder:validation:Enum=Always;IfNotPresent;Never\nPullPolicy describes a policy for if/when to pull a container image",
		"uninstallStrategy":     "+kubebuilder:validation:Enum=RemoveWorkloads;BlockUninstallIfWorkloadsExist\nCDIUninstallStrategy defines the state to leave CDI on uninstall",
		"infra":                 "Rules on which nodes CDI infrastructure pods will be scheduled",
		"workload":              "Restrict on which nodes CDI workload pods will be scheduled",
		"cloneStrategyOverride": "Clone strategy override: should we use a host-assisted copy even if snapshots are available?\n+kubebuilder:validation:Enum=\"copy\";\"snapshot\"",
		"config":                "CDIConfig at CDI level",
		"certConfig":            "certificate configuration",
	}
}

func (CDIStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"": "CDIStatus defines the status of the installation",
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
		"": "CDIConfig provides a user configuration for CDI\n+genclient\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object\n+kubebuilder:object:root=true\n+kubebuilder:storageversion\n+kubebuilder:resource:scope=Cluster",
	}
}

func (FilesystemOverhead) SwaggerDoc() map[string]string {
	return map[string]string{
		"":             "FilesystemOverhead defines the reserved size for PVCs with VolumeMode: Filesystem",
		"global":       "Global is how much space of a Filesystem volume should be reserved for overhead. This value is used unless overridden by a more specific value (per storageClass)",
		"storageClass": "StorageClass specifies how much space of a Filesystem volume should be reserved for safety. The keys are the storageClass and the values are the overhead. This value overrides the global value",
	}
}

func (CDIConfigSpec) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                         "CDIConfigSpec defines specification for user configuration",
		"uploadProxyURLOverride":   "Override the URL used when uploading to a DataVolume",
		"importProxy":              "ImportProxy contains importer pod proxy configuration.\n+optional",
		"scratchSpaceStorageClass": "Override the storage class to used for scratch space during transfer operations. The scratch space storage class is determined in the following order: 1. value of scratchSpaceStorageClass, if that doesn't exist, use the default storage class, if there is no default storage class, use the storage class of the DataVolume, if no storage class specified, use no storage class for scratch space",
		"podResourceRequirements":  "ResourceRequirements describes the compute resource requirements.",
		"featureGates":             "FeatureGates are a list of specific enabled feature gates",
		"filesystemOverhead":       "FilesystemOverhead describes the space reserved for overhead when using Filesystem volumes. A value is between 0 and 1, if not defined it is 0.055 (5.5% overhead)",
		"preallocation":            "Preallocation controls whether storage for DataVolumes should be allocated in advance.",
	}
}

func (CDIConfigStatus) SwaggerDoc() map[string]string {
	return map[string]string{
		"":                               "CDIConfigStatus provides the most recently observed status of the CDI Config resource",
		"uploadProxyURL":                 "The calculated upload proxy URL",
		"importProxy":                    "ImportProxy contains importer pod proxy configuration.\n+optional",
		"scratchSpaceStorageClass":       "The calculated storage class to be used for scratch space",
		"defaultPodResourceRequirements": "ResourceRequirements describes the compute resource requirements.",
		"filesystemOverhead":             "FilesystemOverhead describes the space reserved for overhead when using Filesystem volumes. A percentage value is between 0 and 1",
		"preallocation":                  "Preallocation controls whether storage for DataVolumes should be allocated in advance.",
	}
}

func (CDIConfigList) SwaggerDoc() map[string]string {
	return map[string]string{
		"":      "CDIConfigList provides the needed parameters to do request a list of CDIConfigs from the system\n+k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object",
		"items": "Items provides a list of CDIConfigs",
	}
}

func (ImportProxy) SwaggerDoc() map[string]string {
	return map[string]string{
		"":               "ImportProxy provides the information on how to configure the importer pod proxy.",
		"HTTPProxy":      "HTTPProxy is the URL http://<username>:<pswd>@<ip>:<port> of the import proxy for HTTP requests.  Empty means unset and will not result in the import pod env var.\n+optional",
		"HTTPSProxy":     "HTTPSProxy is the URL https://<username>:<pswd>@<ip>:<port> of the import proxy for HTTPS requests.  Empty means unset and will not result in the import pod env var.\n+optional",
		"noProxy":        "NoProxy is a comma-separated list of hostnames and/or CIDRs for which the proxy should not be used. Empty means unset and will not result in the import pod env var.\n+optional",
		"trustedCAProxy": "TrustedCAProxy is the name of a ConfigMap in the cdi namespace that contains a user-provided trusted certificate authority (CA) bundle.\nThe TrustedCAProxy field is consumed by the import controller that is resposible for coping it to a config map named trusted-ca-proxy-bundle-cm in the cdi namespace.\nHere is an example of the ConfigMap (in yaml):\n\napiVersion: v1\nkind: ConfigMap\nmetadata:\n  name: trusted-ca-proxy-bundle-cm\n  namespace: cdi\ndata:\n  ca.pem: |",
	}
}
