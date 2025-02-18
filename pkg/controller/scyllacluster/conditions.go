package scyllacluster

const (
	serviceAccountControllerProgressingCondition = "ServiceAccountControllerProgressing"
	serviceAccountControllerDegradedCondition    = "ServiceAccountControllerDegraded"
	roleBindingControllerProgressingCondition    = "RoleBindingControllerProgressing"
	roleBindingControllerDegradedCondition       = "RoleBindingControllerDegraded"
	agentTokenControllerProgressingCondition     = "AgentTokenControllerProgressing"
	agentTokenControllerDegradedCondition        = "AgentTokenControllerDegraded"
	certControllerProgressingCondition           = "CertControllerProgressing"
	certControllerDegradedCondition              = "CertControllerDegraded"
	statefulSetControllerAvailableCondition      = "StatefulSetControllerAvailable"
	statefulSetControllerProgressingCondition    = "StatefulSetControllerProgressing"
	statefulSetControllerDegradedCondition       = "StatefulSetControllerDegraded"
	serviceControllerProgressingCondition        = "ServiceControllerProgressing"
	serviceControllerDegradedCondition           = "ServiceControllerDegraded"
	pdbControllerProgressingCondition            = "PDBControllerProgressing"
	pdbControllerDegradedCondition               = "PDBControllerDegraded"
	ingressControllerProgressingCondition        = "IngressControllerProgressing"
	ingressControllerDegradedCondition           = "IngressControllerDegraded"
	jobControllerProgressingCondition            = "JobControllerProgressing"
	jobControllerDegradedCondition               = "JobControllerDegraded"
)
