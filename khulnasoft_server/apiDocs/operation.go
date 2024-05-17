package apiDocs //nolint:stylecheck

import (
	"net/http"

	"github.com/khulnasoft/kengine/khulnasoft_server/diagnosis"
	"github.com/khulnasoft/kengine/khulnasoft_server/ingesters"
	. "github.com/khulnasoft/kengine/khulnasoft_server/model"                //nolint:stylecheck
	. "github.com/khulnasoft/kengine/khulnasoft_server/reporters/completion" //nolint:stylecheck
	. "github.com/khulnasoft/kengine/khulnasoft_server/reporters/graph"      //nolint:stylecheck
	. "github.com/khulnasoft/kengine/khulnasoft_server/reporters/lookup"     //nolint:stylecheck
	. "github.com/khulnasoft/kengine/khulnasoft_server/reporters/search"     //nolint:stylecheck
	"github.com/khulnasoft/kengine/khulnasoft_utils/controls"
	postgresqldb "github.com/khulnasoft/kengine/khulnasoft_utils/postgresql/postgresql-db"
	"github.com/khulnasoft/kengine/khulnasoft_utils/report"
	"github.com/khulnasoft/kengine/khulnasoft_utils/threatintel"
	"github.com/khulnasoft/kengine/khulnasoft_utils/utils"
	ingestersUtil "github.com/khulnasoft/kengine/khulnasoft_utils/utils/ingesters"
)

func (d *OpenAPIDocs) AddUserAuthOperations() {
	d.AddOperation("registerUser", http.MethodPost, "/khulnasoft/user/register",
		"Register User", "First user registration. Further users needs to be invited.",
		http.StatusOK, []string{tagUser}, nil, new(UserRegisterRequest), new(LoginResponse))
	d.AddOperation("authToken", http.MethodPost, "/khulnasoft/auth/token",
		"Get Access Token for API Token", "Get access token for programmatic API access, by providing API Token",
		http.StatusOK, []string{tagAuthentication}, nil, new(APIAuthRequest), new(ResponseAccessToken))
	d.AddOperation("authTokenRefresh", http.MethodPost, "/khulnasoft/auth/token/refresh",
		"Refresh access token", "Reissue access token using refresh token",
		http.StatusOK, []string{tagAuthentication}, bearerToken, nil, new(ResponseAccessToken))
	d.AddOperation("login", http.MethodPost, "/khulnasoft/user/login",
		"Login API", "Login API",
		http.StatusOK, []string{tagAuthentication}, nil, new(LoginRequest), new(LoginResponse))
	d.AddOperation("logout", http.MethodPost, "/khulnasoft/user/logout",
		"Logout API", "Logout API",
		http.StatusNoContent, []string{tagAuthentication}, bearerToken, nil, nil)
}

func (d *OpenAPIDocs) AddUserOperations() {
	d.AddOperation("getCurrentUser", http.MethodGet, "/khulnasoft/user",
		"Get Current User", "Get logged in user information",
		http.StatusOK, []string{tagUser}, bearerToken, nil, new(User))
	d.AddOperation("updateCurrentUser", http.MethodPut, "/khulnasoft/user",
		"Update Current User", "Update logged in user information",
		http.StatusOK, []string{tagUser}, bearerToken, new(UpdateUserRequest), new(User))
	d.AddOperation("updatePassword", http.MethodPut, "/khulnasoft/user/password",
		"Update Password", "Update current user's password",
		http.StatusNoContent, []string{tagUser}, bearerToken, new(UpdateUserPasswordRequest), nil)
	d.AddOperation("deleteCurrentUser", http.MethodDelete, "/khulnasoft/user",
		"Delete Current User", "Delete logged in user",
		http.StatusNoContent, []string{tagUser}, bearerToken, nil, nil)

	d.AddOperation("getUsers", http.MethodGet, "/khulnasoft/users",
		"Get all users", "Get all users",
		http.StatusOK, []string{tagUser}, bearerToken, nil, new([]User))
	d.AddOperation("getUser", http.MethodGet, "/khulnasoft/users/{id}",
		"Get User by User ID", "Get User by User ID",
		http.StatusOK, []string{tagUser}, bearerToken, new(UserIDRequest), new(User))
	d.AddOperation("updateUser", http.MethodPut, "/khulnasoft/users/{id}",
		"Update User by User ID", "Update User by User ID",
		http.StatusOK, []string{tagUser}, bearerToken, new(UpdateUserIDRequest), new(User))
	d.AddOperation("deleteUser", http.MethodDelete, "/khulnasoft/users/{id}",
		"Delete User by User ID", "Delete User by User ID",
		http.StatusNoContent, []string{tagUser}, bearerToken, new(UserIDRequest), nil)

	d.AddOperation("getApiTokens", http.MethodGet, "/khulnasoft/api-token",
		"Get User's API Tokens", "Get logged in user's API Tokens",
		http.StatusOK, []string{tagUser}, bearerToken, nil, new([]APITokenResponse))
	d.AddOperation("resetApiTokens", http.MethodPost, "/khulnasoft/api-token/reset",
		"Reset User's API Tokens", "Reset user's API Tokens",
		http.StatusOK, []string{tagUser}, bearerToken, nil, new([]APITokenResponse))

	d.AddOperation("resetPasswordRequest", http.MethodPost, "/khulnasoft/user/reset-password/request",
		"Reset Password Request", "Request for resetting the password",
		http.StatusOK, []string{tagUser}, nil, new(PasswordResetRequest), new(MessageResponse))
	d.AddOperation("verifyResetPasswordRequest", http.MethodPost, "/khulnasoft/user/reset-password/verify",
		"Verify and Reset Password", "Verify code and reset the password",
		http.StatusNoContent, []string{tagUser}, nil, new(PasswordResetVerifyRequest), nil)

	d.AddOperation("inviteUser", http.MethodPost, "/khulnasoft/user/invite",
		"Invite User", "Invite a user",
		http.StatusOK, []string{tagUser}, bearerToken, new(InviteUserRequest), new(InviteUserResponse))
	d.AddOperation("registerInvitedUser", http.MethodPost, "/khulnasoft/user/invite/register",
		"Register Invited User", "Register invited user",
		http.StatusOK, []string{tagUser}, nil, new(RegisterInvitedUserRequest), new(LoginResponse))

	d.AddOperation("eula", http.MethodGet, "/khulnasoft/end-user-license-agreement",
		"Get End User License Agreement", "Get End User License Agreement",
		http.StatusOK, []string{tagCommon}, nil, nil, new(MessageResponse))
}

func (d *OpenAPIDocs) AddGraphOperations() {
	d.AddOperation("getTopologyGraph", http.MethodPost, "/khulnasoft/graph/topology/",
		"Get Topology Graph", "Retrieve the full topology graph associated with the account",
		http.StatusOK, []string{tagTopology}, bearerToken, new(TopologyFilters), new(GraphResult))

	d.AddOperation("getHostsTopologyGraph", http.MethodPost, "/khulnasoft/graph/topology/hosts",
		"Get Hosts Topology Graph", "Retrieve the full topology graph associated with the account from Hosts",
		http.StatusOK, []string{tagTopology}, bearerToken, new(TopologyFilters), new(GraphResult))

	d.AddOperation("getKubernetesTopologyGraph", http.MethodPost, "/khulnasoft/graph/topology/kubernetes",
		"Get Kubernetes Topology Graph", "Retrieve the full topology graph associated with the account from Kubernetes",
		http.StatusOK, []string{tagTopology}, bearerToken, new(TopologyFilters), new(GraphResult))

	d.AddOperation("getContainersTopologyGraph", http.MethodPost, "/khulnasoft/graph/topology/containers",
		"Get Containers Topology Graph", "Retrieve the full topology graph associated with the account from Containers",
		http.StatusOK, []string{tagTopology}, bearerToken, new(TopologyFilters), new(GraphResult))

	d.AddOperation("getPodsTopologyGraph", http.MethodPost, "/khulnasoft/graph/topology/pods",
		"Get Pods Topology Graph", "Retrieve the full topology graph associated with the account from Pods",
		http.StatusOK, []string{tagTopology}, bearerToken, new(TopologyFilters), new(GraphResult))

	d.AddOperation("getTopologyDelta", http.MethodPost, "/khulnasoft/graph/topology/delta",
		"Get Topology Delta", "Retrieve addition or deletion toplogy deltas",
		http.StatusOK, []string{tagTopology}, bearerToken, new(TopologyDeltaReq), new(TopologyDeltaResponse))

	d.AddOperation("getThreatGraph", http.MethodPost, "/khulnasoft/graph/threat",
		"Get Threat Graph", "Retrieve the full threat graph associated with the account",
		http.StatusOK, []string{tagThreat}, bearerToken, new(ThreatFilters), new(ThreatGraph))

	d.AddOperation("getIndividualThreatGraph", http.MethodPost, "/khulnasoft/graph/threat/individual",
		"Get Vulnerability Threat Graph", "Retrieve threat graph associated with vulnerabilities",
		http.StatusOK, []string{tagThreat}, bearerToken, new(IndividualThreatGraphRequest), new([]IndividualThreatGraph))
}

func (d *OpenAPIDocs) AddLookupOperations() {
	d.AddOperation("getHosts", http.MethodPost, "/khulnasoft/lookup/hosts",
		"Retrieve Hosts data", "Retrieve all the data associated with hosts",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]Host))

	d.AddOperation("getContainers", http.MethodPost, "/khulnasoft/lookup/containers",
		"Retrieve Containers data", "Retrieve all the data associated with containers",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]Container))

	d.AddOperation("getProcesses", http.MethodPost, "/khulnasoft/lookup/processes",
		"Retrieve Processes data", "Retrieve all the data associated with processes",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]Process))

	d.AddOperation("getKubernetesClusters", http.MethodPost, "/khulnasoft/lookup/kubernetesclusters",
		"Retrieve K8s data", "Retrieve all the data associated with k8s clusters",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]KubernetesCluster))

	d.AddOperation("getPods", http.MethodPost, "/khulnasoft/lookup/pods",
		"Retrieve Pods data", "Retrieve all the data associated with pods",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]Pod))

	d.AddOperation("getContainerImages", http.MethodPost, "/khulnasoft/lookup/containerimages",
		"Retrieve Container Images data", "Retrieve all the data associated with images",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]ContainerImage))

	d.AddOperation("getRegistryAccount", http.MethodPost, "/khulnasoft/lookup/registryaccount",
		"Get Images in Registry", "List all the images present in the given registry",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]RegistryAccount))

	d.AddOperation("getCloudResources", http.MethodPost, "/khulnasoft/lookup/cloud-resources",
		"Get Cloud Resources", "Retrieve the cloud resources",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]CloudResource))

	d.AddOperation("getVulnerabilities", http.MethodPost, "/khulnasoft/lookup/vulnerabilities",
		"Retrieve Vulnerabilities data", "Retrieve all the data associated with vulnerabilities",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]Vulnerability))

	d.AddOperation("getSecrets", http.MethodPost, "/khulnasoft/lookup/secrets",
		"Retrieve Secrets data", "Retrieve all the data associated with secrets",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]Secret))

	d.AddOperation("getMalwares", http.MethodPost, "/khulnasoft/lookup/malwares",
		"Retrieve Malwares data", "Retrieve all the data associated with malwares",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]Malware))

	d.AddOperation("getCompliances", http.MethodPost, "/khulnasoft/lookup/compliances",
		"Retrieve Compliances data", "Retrieve all the data associated with compliances",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]Compliance))

	d.AddOperation("getCloudCompliances", http.MethodPost, "/khulnasoft/lookup/cloud-compliances",
		"Retrieve Cloud Compliances data", "Retrieve all the data associated with cloud-compliances",
		http.StatusOK, []string{tagLookup}, bearerToken, new(LookupFilter), new([]CloudCompliance))
}

func (d *OpenAPIDocs) AddSearchOperations() {
	// Search APIs
	d.AddOperation("searchHosts", http.MethodPost, "/khulnasoft/search/hosts",
		"Search hosts", "Search across all data associated with hosts",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]Host))

	d.AddOperation("searchContainers", http.MethodPost, "/khulnasoft/search/containers",
		"Search Containers data", "Search across all data associated with containers",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]Container))

	d.AddOperation("searchContainerImages", http.MethodPost, "/khulnasoft/search/images",
		"Search Container images", "Search across all the data associated with container images",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]ContainerImage))

	d.AddOperation("searchCloudResources", http.MethodPost, "/khulnasoft/search/cloud-resources",
		"Search Cloud Resources", "Search across all data associated with CloudResources",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]CloudResource))

	d.AddOperation("searchKubernetesClusters", http.MethodPost, "/khulnasoft/search/kubernetes-clusters",
		"Search Kuberenetes Clusters", "Search across all data associated with kuberentes clusters",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]KubernetesCluster))

	d.AddOperation("searchVulnerabilities", http.MethodPost, "/khulnasoft/search/vulnerabilities",
		"Search Vulnerabilities", "Search across all the data associated with vulnerabilities",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]Vulnerability))

	d.AddOperation("searchSecrets", http.MethodPost, "/khulnasoft/search/secrets",
		"Search Secrets", "Search across all the data associated with secrets",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]Secret))

	d.AddOperation("searchMalwares", http.MethodPost, "/khulnasoft/search/malwares",
		"Search Malwares", "Search across all the data associated with malwares",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]Malware))

	d.AddOperation("searchCloudCompliances", http.MethodPost, "/khulnasoft/search/cloud-compliances",
		"Search Cloud compliances", "Search across all the data associated with cloud-compliances",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]CloudCompliance))

	d.AddOperation("searchCompliances", http.MethodPost, "/khulnasoft/search/compliances",
		"Search Compliances", "Search across all the data associated with compliances",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]Compliance))

	d.AddOperation("searchVulerabilityRules", http.MethodPost, "/khulnasoft/search/vulnerability-rules",
		"Search Vulnerability Rules", "Search across all the data associated with vulnerability rules",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]VulnerabilityRule))

	d.AddOperation("searchSecretRules", http.MethodPost, "/khulnasoft/search/secret-rules",
		"Search Secret Rules", "Search across all the data associated with secret ruless",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]SecretRule))

	d.AddOperation("searchMalwareRules", http.MethodPost, "/khulnasoft/search/malware-rules",
		"Search Malware Rules", "Search across all the data associated with malware rules",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]MalwareRule))

	d.AddOperation("searchComplianceRules", http.MethodPost, "/khulnasoft/search/compliance-rules",
		"Search Compliance Rules", "Search across all the data associated with compliance rules",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]ComplianceRule))

	d.AddOperation("searchPods", http.MethodPost, "/khulnasoft/search/pods",
		"Search Pods", "Search across all the data associated with pods",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]Pod))

	d.AddOperation("searchVulnerabilityScans", http.MethodPost, "/khulnasoft/search/vulnerability/scans",
		"Search Vulnerability Scan results", "Search across all the data associated with vulnerability scan",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchScanReq), new([]ScanInfo))

	d.AddOperation("searchSecretsScans", http.MethodPost, "/khulnasoft/search/secret/scans",
		"Search Secrets Scan results", "Search across all the data associated with secrets scan",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchScanReq), new([]ScanInfo))

	d.AddOperation("searchMalwareScans", http.MethodPost, "/khulnasoft/search/malware/scans",
		"Search Malware Scan results", "Search across all the data associated with malwares scan",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchScanReq), new([]ScanInfo))

	d.AddOperation("searchComplianceScans", http.MethodPost, "/khulnasoft/search/compliance/scans",
		"Search Compliance Scan results", "Search across all the data associated with compliance scan",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchScanReq), new([]ScanInfo))

	d.AddOperation("searchCloudComplianceScans", http.MethodPost, "/khulnasoft/search/cloud-compliance/scans",
		"Search Cloud Compliance Scan results", "Search across all the data associated with cloud-compliance scan",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchScanReq), new([]ScanInfo))

	d.AddOperation("searchCloudAccounts", http.MethodPost, "/khulnasoft/search/cloud-accounts",
		"Search Cloud Nodes", "Search across all the data associated with cloud nodes",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]CloudNodeAccountInfo))

	d.AddOperation("searchRegistryAccounts", http.MethodPost, "/khulnasoft/search/registry-accounts",
		"Search Registry Accounts", "Search across all the data associated with registry account",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new([]RegistryAccount))

	d.AddOperation("getCloudComplianceFilters", http.MethodPost, "/khulnasoft/filters/cloud-compliance",
		"Get Cloud Compliance Filters", "Get all applicable filter values for cloud compliance",
		http.StatusOK, []string{tagSearch}, bearerToken, new(FiltersReq), new(FiltersResult))

	d.AddOperation("getComplianceFilters", http.MethodPost, "/khulnasoft/filters/compliance",
		"Get Compliance Filters", "Get all applicable filter values for compliance",
		http.StatusOK, []string{tagSearch}, bearerToken, new(FiltersReq), new(FiltersResult))

	// Count APIs
	d.AddOperation("countNodes", http.MethodGet, "/khulnasoft/search/count/nodes",
		"Count nodes", "Count hosts, containers, pods, k8s clusters, images",
		http.StatusOK, []string{tagSearch}, bearerToken, nil, new(NodeCountResp))

	d.AddOperation("countHosts", http.MethodPost, "/khulnasoft/search/count/hosts",
		"Count hosts", "Count across all the data associated with hosts",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countContainers", http.MethodPost, "/khulnasoft/search/count/containers",
		"Count Containers data", "Count across all the data associated with containers",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countContainerImages", http.MethodPost, "/khulnasoft/search/count/images",
		"Count Container images", "Count across all the data associated with container images",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countVulnerabilities", http.MethodPost, "/khulnasoft/search/count/vulnerabilities",
		"Count Vulnerabilities", "Search across all the data associated with vulnerabilities",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countSecrets", http.MethodPost, "/khulnasoft/search/count/secrets",
		"Count Secrets", "Count across all the data associated with secrets",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countMalwares", http.MethodPost, "/khulnasoft/search/count/malwares",
		"Count Malwares", "Count across all the data associated with malwares",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countPods", http.MethodPost, "/khulnasoft/search/count/pods",
		"Count Pods", "Count across all the data associated with pods",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countCloudCompliances", http.MethodPost, "/khulnasoft/search/count/cloud-compliances",
		"Count Cloud compliances", "Count across all the data ssociated with cloud compliances",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countCloudResources", http.MethodPost, "/khulnasoft/search/count/cloud-resources",
		"Count Cloud resources", "Count across all the data ssociated with cloud resources",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countKubernetesClusters", http.MethodPost, "/khulnasoft/search/count/kubernetes-clusters",
		"Count Kubernetes clusters", "Count across all the data ssociated with kubernetes clusters",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countCompliances", http.MethodPost, "/khulnasoft/search/count/compliances",
		"Count Compliances", "Count across all the data associated with compliances",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countVulnerabilityRules", http.MethodPost, "/khulnasoft/search/count/vulnerability-rules",
		"Count Vulnerability Rules", "Count across all the data associated with vulnerability rules",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countSecretRules", http.MethodPost, "/khulnasoft/search/count/secret-rules",
		"Count Secret Rules", "Count across all the data associated with secret rules",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countMalwareRules", http.MethodPost, "/khulnasoft/search/count/malware-rules",
		"Count Malware Rules", "Count across all the data associated with malware rules",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countComplianceRules", http.MethodPost, "/khulnasoft/search/count/compliance-rules",
		"Count Compliance Rules", "Count across all the data associated with compliance rules",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countVulnerabilityScans", http.MethodPost, "/khulnasoft/search/count/vulnerability/scans",
		"Count Vulnerability Scan results", "Count across all the data associated with vulnerability scans",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchScanReq), new(SearchCountResp))

	d.AddOperation("countSecretsScans", http.MethodPost, "/khulnasoft/search/count/secret/scans",
		"Count Secret Scan results", "Count across all the data associated with secret scans",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchScanReq), new(SearchCountResp))

	d.AddOperation("countMalwareScans", http.MethodPost, "/khulnasoft/search/count/malware/scans",
		"Count Malware Scan results", "Count across all the data associated with malware scans",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchScanReq), new(SearchCountResp))

	d.AddOperation("countComplianceScans", http.MethodPost, "/khulnasoft/search/count/compliance/scans",
		"Count Compliance Scan results", "Count across all the data associated with compliance scans",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchScanReq), new(SearchCountResp))

	d.AddOperation("countCloudComplianceScans", http.MethodPost, "/khulnasoft/search/count/cloud-compliance/scans",
		"Count Cloud Compliance Scan results", "Count across all the data associated with cloud-compliance scans",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchScanReq), new(SearchCountResp))

	d.AddOperation("countCloudAccounts", http.MethodPost, "/khulnasoft/search/count/cloud-accounts",
		"Count Cloud Nodes", "Search across all the data associated with cloud nodes",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))

	d.AddOperation("countRegistryAccounts", http.MethodPost, "/khulnasoft/search/count/registry-accounts",
		"Count Registry Accounts", "Search across all the data associated with registry account",
		http.StatusOK, []string{tagSearch}, bearerToken, new(SearchNodeReq), new(SearchCountResp))
}

func (d *OpenAPIDocs) AddControlsOperations() {
	d.AddOperation("getAgentControls", http.MethodPost, "/khulnasoft/controls/agent",
		"Fetch Agent Actions", "Fetch actions for a given agent",
		http.StatusOK, []string{tagControls}, bearerToken, new(AgentID), new(controls.AgentControls))

	d.AddOperation("getKubernetesClusterControls", http.MethodPost, "/khulnasoft/controls/kubernetes-cluster",
		"Fetch Kubernetes Cluster Actions", "Fetch actions for a given Kubernetes Cluster",
		http.StatusOK, []string{tagControls}, bearerToken, new(AgentID), new(controls.AgentControls))

	d.AddOperation("getAgentInitControls", http.MethodPost, "/khulnasoft/controls/agent-init",
		"Fetch Agent Init Actions", "Fetch initial actions for a given agent after it started",
		http.StatusOK, []string{tagControls}, bearerToken, new(InitAgentReq), new(controls.AgentControls))

	d.AddOperation("upgradeAgentVersion", http.MethodPost, "/khulnasoft/controls/agent-upgrade",
		"Schedule new agent version upgrade", "Schedule new agent version upgrade",
		http.StatusOK, []string{tagControls}, bearerToken, new(AgentUpgrade), nil)

	d.AddOperation("enableAgentPlugin", http.MethodPost, "/khulnasoft/controls/agent-plugins/enable",
		"Schedule new agent plugin version enabling", "Schedule agent plugin enable",
		http.StatusOK, []string{tagControls}, bearerToken, new(AgentPluginEnable), nil)

	d.AddOperation("disableAgentPlugin", http.MethodPost, "/khulnasoft/controls/agent-plugins/disable",
		"Schedule new agent plugin version disabling", "Schedule agent plugin disable",
		http.StatusOK, []string{tagControls}, bearerToken, new(AgentPluginDisable), nil)

	d.AddOperation("getCloudNodeControls", http.MethodPost, "/khulnasoft/controls/cloud-node",
		"Fetch Cloud Node Controls", "Fetch controls for a cloud node",
		http.StatusOK, []string{tagControls}, bearerToken, new(CloudNodeControlReq), new(CloudNodeControlResp))

	d.AddOperation("enableCloudNodeControls", http.MethodPost, "/khulnasoft/controls/cloud-node/enable",
		"Enable Cloud Node Controls", "Enable controls for a cloud node",
		http.StatusOK, []string{tagControls}, bearerToken, new(CloudNodeEnableDisableReq), nil)

	d.AddOperation("disableCloudNodeControls", http.MethodPost, "/khulnasoft/controls/cloud-node/disable",
		"Disable Cloud Node Controls", "Disable controls for a cloud node",
		http.StatusOK, []string{tagControls}, bearerToken, new(CloudNodeEnableDisableReq), nil)
}

func (d *OpenAPIDocs) AddCloudNodeOperations() {
	d.AddOperation("registerCloudNodeAccount", http.MethodPost, "/khulnasoft/cloud-node/account",
		"Register Cloud Node Account", "Register Cloud Node Account and return any pending compliance scans from console",
		http.StatusOK, []string{tagCloudNodes}, bearerToken, new(CloudNodeAccountRegisterReq), new(CloudNodeAccountRegisterResp))

	d.AddOperation("deleteCloudNodeAccount", http.MethodPatch, "/khulnasoft/cloud-node/account/delete",
		"Delete Cloud Node Account", "Delete Cloud Node Account and related resources",
		http.StatusAccepted, []string{tagCloudNodes}, bearerToken, new(CloudAccountDeleteReq), nil)

	d.AddOperation("listCloudNodeAccount", http.MethodPost, "/khulnasoft/cloud-node/list/accounts",
		"List Cloud Node Accounts", "List Cloud Node Accounts registered with the console",
		http.StatusOK, []string{tagCloudNodes}, bearerToken, new(CloudNodeAccountsListReq), new(CloudNodeAccountsListResp))

	d.AddOperation("refreshCloudNodeAccount", http.MethodPost, "/khulnasoft/cloud-node/account/refresh",
		"Refresh Cloud Account", "Refresh the cloud resources in a Cloud Account",
		http.StatusNoContent, []string{tagCloudNodes}, bearerToken, new(CloudAccountRefreshReq), nil)

	d.AddOperation("listCloudProviders", http.MethodGet, "/khulnasoft/cloud-node/list/providers",
		"List Cloud Node Providers", "List Cloud Node Providers registered with the console",
		http.StatusOK, []string{tagCloudNodes}, bearerToken, new(CloudNodeProvidersListReq), new(CloudNodeProvidersListResp))
}

func (d *OpenAPIDocs) AddIngestersOperations() {
	d.AddOperation("ingestAgentReport", http.MethodPost, "/khulnasoft/ingest/report",
		"Ingest Topology Data", "Ingest data reported by one Agent",
		http.StatusOK, []string{tagTopology}, bearerToken, new(report.RawReport), new(controls.AgentBeat))

	d.AddOperation("ingestSyncAgentReport", http.MethodPost, "/khulnasoft/ingest/sync-report",
		"Ingest Topology Data", "Ingest data reported by one Agent",
		http.StatusOK, []string{tagTopology}, bearerToken, new(ingesters.ReportIngestionData), nil)

	d.AddOperation("ingestSbom", http.MethodPost, "/khulnasoft/ingest/sbom",
		"Ingest SBOM from Scan", "Ingest SBOM from Scan",
		http.StatusOK, []string{tagVulnerability}, bearerToken, new(utils.ScanSbomRequest), nil)

	d.AddOperation("ingestVulnerabilities", http.MethodPost, "/khulnasoft/ingest/vulnerabilities",
		"Ingest Vulnerabilities", "Ingest vulnerabilities found while scanning the agent host or containers",
		http.StatusOK, []string{tagVulnerability}, bearerToken, new([]ingestersUtil.Vulnerability), nil)

	d.AddOperation("ingestVulnerabilitiesScanStatus", http.MethodPost, "/khulnasoft/ingest/vulnerabilities-scan-logs",
		"Ingest Vulnerabilities Scan Status", "Ingest vulnerabilities scan status from agent",
		http.StatusOK, []string{tagVulnerability}, bearerToken, new([]ingestersUtil.VulnerabilityScanStatus), nil)

	d.AddOperation("ingestSecrets", http.MethodPost, "/khulnasoft/ingest/secrets",
		"Ingest Secrets", "Ingest secrets found while scanning the agent",
		http.StatusOK, []string{tagSecretScan}, bearerToken, new([]ingestersUtil.Secret), nil)

	d.AddOperation("ingestSecretScanStatus", http.MethodPost, "/khulnasoft/ingest/secret-scan-logs",
		"Ingest Secrets Scan Status", "Ingest secrets scan status from the agent",
		http.StatusOK, []string{tagSecretScan}, bearerToken, new([]ingestersUtil.SecretScanStatus), nil)

	d.AddOperation("ingestCompliances", http.MethodPost, "/khulnasoft/ingest/compliance",
		"Ingest Compliances", "Ingest compliance issues found while scanning the agent",
		http.StatusOK, []string{tagCompliance}, bearerToken, new([]ingestersUtil.Compliance), nil)

	d.AddOperation("ingestComplianceScanStatus", http.MethodPost, "/khulnasoft/ingest/compliance-scan-logs",
		"Ingest Compliance Scan Status", "Ingest compliance issues found while scanning the agent",
		http.StatusOK, []string{tagCompliance}, bearerToken, new([]ingestersUtil.ComplianceScanStatus), nil)

	d.AddOperation("ingestCloudCompliances", http.MethodPost, "/khulnasoft/ingest/cloud-compliance",
		"Ingest Cloud Compliances", "Ingest Cloud compliances found while scanning cloud provider",
		http.StatusOK, []string{tagCloudScanner}, bearerToken, new([]ingestersUtil.CloudCompliance), nil)

	d.AddOperation("ingestCloudComplianceScanStatus", http.MethodPost, "/khulnasoft/ingest/cloud-compliance-status",
		"Ingest Cloud Compliances scan status", "Ingest Cloud compliances found while scanning cloud provider",
		http.StatusOK, []string{tagCloudScanner}, bearerToken, new([]ingestersUtil.CloudComplianceScanStatus), nil)

	d.AddOperation("ingestMalwareScanStatus", http.MethodPost, "/khulnasoft/ingest/malware-scan-logs",
		"Ingest Malware Scan Status", "Ingest malware scan status from the agent",
		http.StatusOK, []string{tagMalwareScan}, bearerToken, new([]ingestersUtil.MalwareScanStatus), nil)

	d.AddOperation("ingestMalware", http.MethodPost, "/khulnasoft/ingest/malware",
		"Ingest Malware", "Ingest malware found while scanning the agent",
		http.StatusOK, []string{tagMalwareScan}, bearerToken, new([]ingestersUtil.Malware), nil)

	d.AddOperation("ingestCloudResources", http.MethodPost, "/khulnasoft/ingest/cloud-resources",
		"Ingest Cloud resources", "Ingest Clouds Resources found while scanning cloud provider",
		http.StatusOK, []string{tagCloudResources}, bearerToken, new([]ingestersUtil.CloudResource), nil)
}

func (d *OpenAPIDocs) AddScansOperations() {
	// List scan result fields
	d.AddOperation("getScanReportFields", http.MethodGet, "/khulnasoft/scan/results/fields",
		"Get Scan Report Fields", "Get all the fields available in all the scan reports",
		http.StatusOK, []string{tagCommon}, bearerToken, nil, new(ScanReportFieldsResponse))
	// Start scan
	d.AddOperation("startVulnerabilityScan", http.MethodPost, "/khulnasoft/scan/start/vulnerability",
		"Start Vulnerability Scan", "Start Vulnerability Scan on agent or registry",
		http.StatusAccepted, []string{tagVulnerability}, bearerToken, new(VulnerabilityScanTriggerReq), new(ScanTriggerResp))
	d.AddOperation("startSecretScan", http.MethodPost, "/khulnasoft/scan/start/secret",
		"Start Secret Scan", "Start Secret Scan on agent or registry",
		http.StatusAccepted, []string{tagSecretScan}, bearerToken, new(SecretScanTriggerReq), new(ScanTriggerResp))
	d.AddOperation("startComplianceScan", http.MethodPost, "/khulnasoft/scan/start/compliance",
		"Start Compliance Scan", "Start Compliance Scan on agent or registry",
		http.StatusAccepted, []string{tagCompliance}, bearerToken, new(ComplianceScanTriggerReq), new(ScanTriggerResp))
	d.AddOperation("startMalwareScan", http.MethodPost, "/khulnasoft/scan/start/malware",
		"Start Malware Scan", "Start Malware Scan on agent or registry",
		http.StatusAccepted, []string{tagMalwareScan}, bearerToken, new(MalwareScanTriggerReq), new(ScanTriggerResp))

	// Stop scan
	d.AddOperation("stopVulnerabilityScan", http.MethodPost, "/khulnasoft/scan/stop/vulnerability",
		"Stop Vulnerability Scan", "Stop Vulnerability Scan on agent or registry",
		http.StatusAccepted, []string{tagVulnerability}, bearerToken, new(StopScanRequest), nil)
	d.AddOperation("stopComplianceScan", http.MethodPost, "/khulnasoft/scan/stop/compliance",
		"Stop Compliance Scan", "Stop Compliance Scan on agent or registry",
		http.StatusAccepted, []string{tagCompliance}, bearerToken, new(StopScanRequest), nil)
	d.AddOperation("stopMalwareScan", http.MethodPost, "/khulnasoft/scan/stop/malware",
		"Stop Malware Scan", "Stop Malware Scan on agent or registry",
		http.StatusAccepted, []string{tagMalwareScan}, bearerToken, new(StopScanRequest), nil)
	d.AddOperation("stopSecretScan", http.MethodPost, "/khulnasoft/scan/stop/secret",
		"Stop Secret Scan", "Stop Secret Scan on agent or registry",
		http.StatusAccepted, []string{tagSecretScan}, bearerToken, new(StopScanRequest), nil)

	// Status scan
	d.AddOperation("statusVulnerabilityScan", http.MethodPost, "/khulnasoft/scan/status/vulnerability",
		"Get Vulnerability Scan Status", "Get Vulnerability Scan Status on agent or registry",
		http.StatusOK, []string{tagVulnerability}, bearerToken, new(ScanStatusReq), new(ScanStatusResp))
	d.AddOperation("statusSecretScan", http.MethodPost, "/khulnasoft/scan/status/secret",
		"Get Secret Scan Status", "Get Secret Scan Status on agent or registry",
		http.StatusOK, []string{tagSecretScan}, bearerToken, new(ScanStatusReq), new(ScanStatusResp))
	d.AddOperation("statusComplianceScan", http.MethodPost, "/khulnasoft/scan/status/compliance",
		"Get Compliance Scan Status", "Get Compliance Scan Status on agent or registry",
		http.StatusOK, []string{tagCompliance}, bearerToken, new(ScanStatusReq), new(ScanStatusResp))
	d.AddOperation("statusMalwareScan", http.MethodPost, "/khulnasoft/scan/status/malware",
		"Get Malware Scan Status", "Get Malware Scan status on agent or registry",
		http.StatusOK, []string{tagMalwareScan}, bearerToken, new(ScanStatusReq), new(ScanStatusResp))
	d.AddOperation("statusCloudComplianceScan", http.MethodPost, "/khulnasoft/scan/status/cloud-compliance",
		"Get Cloud Compliance Scan Status", "Get Cloud Compliance Scan Status on cloud node",
		http.StatusOK, []string{tagCloudScanner}, bearerToken, new(ScanStatusReq), new(ComplianceScanStatusResp))

	// List scans
	d.AddOperation("listVulnerabilityScans", http.MethodPost, "/khulnasoft/scan/list/vulnerability",
		"Get Vulnerability Scans List", "Get Vulnerability Scan list on agent or registry",
		http.StatusOK, []string{tagVulnerability}, bearerToken, new(ScanListReq), new(ScanListResp))
	d.AddOperation("listSecretScan", http.MethodPost, "/khulnasoft/scan/list/secret",
		"Get Secret Scans List", "Get Secret Scans list on agent or registry",
		http.StatusOK, []string{tagSecretScan}, bearerToken, new(ScanListReq), new(ScanListResp))
	d.AddOperation("listComplianceScan", http.MethodPost, "/khulnasoft/scan/list/compliance",
		"Get Compliance Scans List", "Get Compliance Scans list on agent or registry",
		http.StatusOK, []string{tagCompliance}, bearerToken, new(ScanListReq), new(ScanListResp))
	d.AddOperation("listMalwareScan", http.MethodPost, "/khulnasoft/scan/list/malware",
		"Get Malware Scans List", "Get Malware Scans list on agent or registry",
		http.StatusOK, []string{tagMalwareScan}, bearerToken, new(ScanListReq), new(ScanListResp))
	d.AddOperation("listCloudComplianceScan", http.MethodPost, "/khulnasoft/scan/list/cloud-compliance",
		"Get Cloud Compliance Scans List", "Get Cloud Compliance Scans list for cloud node",
		http.StatusOK, []string{tagCloudScanner}, bearerToken, new(ScanListReq), new(ScanListResp))

	// Scans' Results
	d.AddOperation("resultsVulnerabilityScans", http.MethodPost, "/khulnasoft/scan/results/vulnerability",
		"Get Vulnerability Scans Results", "Get Vulnerability Scan results on agent or registry",
		http.StatusOK, []string{tagVulnerability}, bearerToken, new(ScanResultsReq), new(VulnerabilityScanResult))
	d.AddOperation("resultsSecretScan", http.MethodPost, "/khulnasoft/scan/results/secret",
		"Get Secret Scans Results", "Get Secret Scans results on agent or registry",
		http.StatusOK, []string{tagSecretScan}, bearerToken, new(ScanResultsReq), new(SecretScanResult))
	d.AddOperation("resultsRulesSecretScan", http.MethodPost, "/khulnasoft/scan/results/secret/rules",
		"Get Secret Scans Result Rules", "Get Secret Scans detected rules names",
		http.StatusOK, []string{tagSecretScan}, bearerToken, new(ScanResultsReq), new(SecretScanResultRules))
	d.AddOperation("resultsComplianceScan", http.MethodPost, "/khulnasoft/scan/results/compliance",
		"Get Compliance Scans Results", "Get Compliance Scans results on agent or registry",
		http.StatusOK, []string{tagCompliance}, bearerToken, new(ScanResultsReq), new(ComplianceScanResult))
	d.AddOperation("resultsMalwareScan", http.MethodPost, "/khulnasoft/scan/results/malware",
		"Get Malware Scans Results", "Get Malware Scans results on agent or registry",
		http.StatusOK, []string{tagMalwareScan}, bearerToken, new(ScanResultsReq), new(MalwareScanResult))
	d.AddOperation("resultsRulesMalwareScan", http.MethodPost, "/khulnasoft/scan/results/malware/rules",
		"Get Malware Scans Result Rules", "Get Malware Scans detected rules names",
		http.StatusOK, []string{tagMalwareScan}, bearerToken, new(ScanResultsReq), new(MalwareScanResultRules))
	d.AddOperation("resultsClassMalwareScan", http.MethodPost, "/khulnasoft/scan/results/malware/class",
		"Get Malware Scans Results", "Get Malware Scans detected class names",
		http.StatusOK, []string{tagMalwareScan}, bearerToken, new(ScanResultsReq), new(MalwareScanResultClass))
	d.AddOperation("resultsCloudComplianceScan", http.MethodPost, "/khulnasoft/scan/results/cloud-compliance",
		"Get Cloud Compliance Scan Results", "Get Cloud Compliance Scan results for cloud node",
		http.StatusOK, []string{tagCloudScanner}, bearerToken, new(ScanResultsReq), new(CloudComplianceScanResult))

	// Scans results counts
	d.AddOperation("countResultsVulnerabilityScans", http.MethodPost, "/khulnasoft/scan/results/count/vulnerability",
		"Get Vulnerability Scans Results", "Get Vulnerability Scan results on agent or registry",
		http.StatusOK, []string{tagVulnerability}, bearerToken, new(ScanResultsReq), new(SearchCountResp))
	d.AddOperation("countResultsSecretScan", http.MethodPost, "/khulnasoft/scan/results/count/secret",
		"Get Secret Scans Results", "Get Secret Scans results on agent or registry",
		http.StatusOK, []string{tagSecretScan}, bearerToken, new(ScanResultsReq), new(SearchCountResp))
	d.AddOperation("countResultsComplianceScan", http.MethodPost, "/khulnasoft/scan/results/count/compliance",
		"Get Compliance Scans Results", "Get Compliance Scans results on agent or registry",
		http.StatusOK, []string{tagCompliance}, bearerToken, new(ScanResultsReq), new(SearchCountResp))
	d.AddOperation("countResultsMalwareScan", http.MethodPost, "/khulnasoft/scan/results/count/malware",
		"Get Malware Scans Results", "Get Malware Scans results on agent or registry",
		http.StatusOK, []string{tagMalwareScan}, bearerToken, new(ScanResultsReq), new(SearchCountResp))
	d.AddOperation("countResultsCloudComplianceScan", http.MethodPost, "/khulnasoft/scan/results/count/cloud-compliance",
		"Get Cloud Compliance Scan Results", "Get Cloud Compliance Scan results for cloud node",
		http.StatusOK, []string{tagCloudScanner}, bearerToken, new(ScanResultsReq), new(SearchCountResp))

	// pie chart apis
	d.AddOperation("groupResultsSecrets", http.MethodGet, "/khulnasoft/scan/results/count/group/secret",
		"Group Secret Results", "Group Secret Scans results by severity/rule",
		http.StatusOK, []string{tagSecretScan}, bearerToken, nil, new(ResultGroupResp))
	d.AddOperation("groupResultsMalwares", http.MethodGet, "/khulnasoft/scan/results/count/group/malware",
		"Group Malware Results", "Group Malware Scans results by severity/rule",
		http.StatusOK, []string{tagMalwareScan}, bearerToken, nil, new(ResultGroupResp))
	d.AddOperation("groupResultsMalwaresClass", http.MethodGet, "/khulnasoft/scan/results/count/group/malware/class",
		"Group Malware Results By Class", "Group Malware Scans results by severity/class",
		http.StatusOK, []string{tagMalwareScan}, bearerToken, nil, new(ResultGroupResp))

	d.AddOperation("getAllNodesInScanResults", http.MethodPost, "/khulnasoft/scan/nodes-in-result",
		"Get all nodes in given scan result ids", "Get all nodes in given scan result ids",
		http.StatusOK, []string{tagScanResults}, bearerToken, new(NodesInScanResultRequest), new([]ScanResultBasicNode))

	// Scan Result Actions
	d.AddOperation("maskScanResult", http.MethodPost, "/khulnasoft/scan/results/action/mask",
		"Mask Scans Results", "Mask scan results",
		http.StatusNoContent, []string{tagScanResults}, bearerToken, new(ScanResultsMaskRequest), nil)
	d.AddOperation("unmaskScanResult", http.MethodPost, "/khulnasoft/scan/results/action/unmask",
		"Unmask Scans Results", "Unmask scan results",
		http.StatusNoContent, []string{tagScanResults}, bearerToken, new(ScanResultsMaskRequest), nil)
	d.AddOperation("deleteScanResult", http.MethodPatch, "/khulnasoft/scan/results/action/delete",
		"Delete selected scan results", "Delete selected scan results",
		http.StatusNoContent, []string{tagScanResults}, bearerToken, new(ScanResultsActionRequest), nil)
	d.AddOperation("notifyScanResult", http.MethodPost, "/khulnasoft/scan/results/action/notify",
		"Notify Scans Results", "Notify scan results in connected integration channels",
		http.StatusNoContent, []string{tagScanResults}, bearerToken, new(ScanResultsActionRequest), nil)

	// Bulk Delete Scans
	d.AddOperation("bulkDeleteScans", http.MethodPost, "/khulnasoft/scans/bulk/delete",
		"Bulk Delete Scans", "Bulk delete scans along with their results for a particular scan type",
		http.StatusNoContent, []string{tagScanResults}, bearerToken, new(BulkDeleteScansRequest), nil)

	// Scan ID Actions
	d.AddOperation("downloadScanResults", http.MethodGet, "/khulnasoft/scan/{scan_type}/{scan_id}/download",
		"Download Scans Results", "Download scan results",
		http.StatusOK, []string{tagScanResults}, bearerToken, new(ScanActionRequest), new(DownloadScanResultsResponse))
	d.AddOperation("deleteScanResultsForScanID", http.MethodDelete, "/khulnasoft/scan/{scan_type}/{scan_id}",
		"Delete all scan results for a scan id", "Delete all scan results for a scan id",
		http.StatusNoContent, []string{tagScanResults}, bearerToken, new(ScanActionRequest), nil)

	// SBOM
	d.AddOperation("getSBOM", http.MethodPost, "/khulnasoft/scan/sbom",
		"Get SBOM for a node or scan id", "Get SBOM for a node or scan id",
		http.StatusOK, []string{tagVulnerability}, bearerToken, new(SbomRequest), new([]SbomResponse))
	d.AddOperation("downloadSBOM", http.MethodPost, "/khulnasoft/scan/sbom/download",
		"Download SBOM for a node or scan id", "Download SBOM for a node or scan id",
		http.StatusOK, []string{tagVulnerability}, bearerToken, new(SbomRequest), new(DownloadReportResponse))
}

func (d *OpenAPIDocs) AddDiagnosisOperations() {
	d.AddOperation("diagnosticNotification", http.MethodGet, "/khulnasoft/diagnosis/notification",
		"Get Diagnostic Notification", "Get Diagnostic Notification",
		http.StatusOK, []string{tagDiagnosis}, bearerToken, nil, new([]diagnosis.DiagnosticNotification))
	d.AddOperation("generateConsoleDiagnosticLogs", http.MethodPost, "/khulnasoft/diagnosis/console-logs",
		"Generate Console Diagnostic Logs", "Generate Console Diagnostic Logs",
		http.StatusAccepted, []string{tagDiagnosis}, bearerToken, new(diagnosis.GenerateConsoleDiagnosticLogsRequest), nil)
	d.AddOperation("generateAgentDiagnosticLogs", http.MethodPost, "/khulnasoft/diagnosis/agent-logs",
		"Generate Agent Diagnostic Logs", "Generate Agent Diagnostic Logs",
		http.StatusAccepted, []string{tagDiagnosis}, bearerToken, new(diagnosis.GenerateAgentDiagnosticLogsRequest), nil)
	d.AddOperation("updateAgentDiagnosticLogsStatus", http.MethodPut, "/khulnasoft/diagnosis/agent-logs/status/{node_id}",
		"Update Agent Diagnostic Logs Status", "Update agent diagnostic logs status",
		http.StatusNoContent, []string{tagDiagnosis}, bearerToken, new(diagnosis.DiagnosticLogsStatus), nil)
	d.AddOperation("generateCloudScannerDiagnosticLogs", http.MethodPost, "/khulnasoft/diagnosis/cloud-scanner-logs",
		"Generate Cloud Scanner Diagnostic Logs", "Generate Cloud Scanner Diagnostic Logs",
		http.StatusAccepted, []string{tagDiagnosis}, bearerToken, new(diagnosis.GenerateCloudScannerDiagnosticLogsRequest), nil)
	d.AddOperation("updateCloudScannerDiagnosticLogsStatus", http.MethodPut, "/khulnasoft/diagnosis/cloud-scanner-logs/status/{node_id}",
		"Update Cloud Scanner Diagnostic Logs Status", "Update cloud scanner diagnostic logs status",
		http.StatusNoContent, []string{tagDiagnosis}, bearerToken, new(diagnosis.DiagnosticLogsStatus), nil)
	d.AddOperation("getDiagnosticLogs", http.MethodGet, "/khulnasoft/diagnosis/diagnostic-logs",
		"Get Diagnostic Logs", "Get diagnostic logs download url links",
		http.StatusOK, []string{tagDiagnosis}, bearerToken, nil, new(diagnosis.GetDiagnosticLogsResponse))
}

func (d *OpenAPIDocs) AddRegistryOperations() {
	d.AddOperation("listRegistry", http.MethodGet, "/khulnasoft/registryaccount",
		"List Registries", "List all the added Registries",
		http.StatusOK, []string{tagRegistry}, bearerToken, new(RegistryListReq), new([]RegistryListResp))
	d.AddOperation("addRegistry", http.MethodPost, "/khulnasoft/registryaccount",
		"Add Registry", "Add a new supported registry",
		http.StatusOK, []string{tagRegistry}, bearerToken, new(RegistryAddReq), new(MessageResponse))
	d.AddOperation("addRegistryGCR", http.MethodPost, "/khulnasoft/registryaccount/gcr",
		"Add Google Container Registry", "Add a Google Container registry",
		http.StatusOK, []string{tagRegistry}, bearerToken, new(RegistryGCRAddReq), new(MessageResponse))
	d.AddOperation("updateRegistry", http.MethodPut, "/khulnasoft/registryaccount/{registry_id}",
		"Update Registry", "Update registry",
		http.StatusOK, []string{tagRegistry}, bearerToken, new(RegistryUpdateReq), new(MessageResponse))
	d.AddOperation("deleteRegistry", http.MethodDelete, "/khulnasoft/registryaccount/{registry_id}",
		"Delete Registry", "Delete registry",
		http.StatusNoContent, []string{tagRegistry}, bearerToken, new(RegistryIDPathReq), nil)
	d.AddOperation("deleteRegistryBulk", http.MethodPatch, "/khulnasoft/registryaccount/delete",
		"Batch Delete Registry", "Batch Delete registry",
		http.StatusNoContent, []string{tagRegistry}, bearerToken, new(DeleteRegistryBulkReq), nil)
	d.AddOperation("syncRegistry", http.MethodPost, "/khulnasoft/registryaccount/{registry_id}/sync",
		"Sync Registry", "synchronize registry images",
		http.StatusOK, []string{tagRegistry}, bearerToken, new(RegistryIDPathReq), new(MessageResponse))
	d.AddOperation("getSummaryAll", http.MethodGet, "/khulnasoft/registryaccount/summary",
		"Get All Registries Summary By Type", "get summary of all registries scans, images and tags by registry type",
		http.StatusOK, []string{tagRegistry}, bearerToken, nil, new(RegistrySummaryAllResp))
	d.AddOperation("getSummaryByType", http.MethodGet, "/khulnasoft/registryaccount/{registry_type}/summary-by-type",
		"Get Registry Summary By Type", "get summary of registries scans, images and tags by registry type",
		http.StatusOK, []string{tagRegistry}, bearerToken, new(RegistryTypeReq), new(Summary))
	d.AddOperation("getRegistrySummary", http.MethodGet, "/khulnasoft/registryaccount/{registry_id}/summary",
		"Get Registry Summary", "get summary of registry scans, images and tags",
		http.StatusOK, []string{tagRegistry}, bearerToken, new(RegistryIDPathReq), new(Summary))
	d.AddOperation("listImages", http.MethodPost, "/khulnasoft/registryaccount/images",
		"List Registry Images", "list images from a given registry",
		http.StatusOK, []string{tagRegistry}, bearerToken, new(RegistryImagesReq), new([]ContainerImage))
	d.AddOperation("listImageStubs", http.MethodPost, "/khulnasoft/registryaccount/stubs",
		"List Image Stubs", "list image tags for a given image and registry",
		http.StatusOK, []string{tagRegistry}, bearerToken, new(RegistryImageStubsReq), new([]ImageStub))
	d.AddOperation("CountImages", http.MethodPost, "/khulnasoft/registryaccount/count/images",
		"Count Registry Images", "count of images from a given registry",
		http.StatusOK, []string{tagRegistry}, bearerToken, new(RegistryImagesReq), new(RegistryCountResp))
	d.AddOperation("CountImageStubs", http.MethodPost, "/khulnasoft/registryaccount/count/stubs",
		"Count Image Stubs", "count of image tags for a given image and registry",
		http.StatusOK, []string{tagRegistry}, bearerToken, new(RegistryImageStubsReq), new(RegistryCountResp))
}

func (d *OpenAPIDocs) AddIntegrationOperations() {
	d.AddOperation("addIntegration", http.MethodPost, "/khulnasoft/integration",
		"Add Integration", "Add a new supported integration",
		http.StatusOK, []string{tagIntegration}, bearerToken, new(IntegrationAddReq), new(MessageResponse))
	d.AddOperation("listIntegration", http.MethodGet, "/khulnasoft/integration",
		"List Integrations", "List all the added Integrations",
		http.StatusOK, []string{tagIntegration}, bearerToken, new(IntegrationListReq), new([]IntegrationListResp))
	d.AddOperation("updateIntegration", http.MethodPut, "/khulnasoft/integration/{integration_id}",
		"Update Integration", "Update integration",
		http.StatusOK, []string{tagIntegration}, bearerToken, new(IntegrationUpdateReq), new(MessageResponse))
	d.AddOperation("deleteIntegration", http.MethodDelete, "/khulnasoft/integration/{integration_id}",
		"Delete Single Integration", "Delete single integration",
		http.StatusNoContent, []string{tagIntegration}, bearerToken, new(IntegrationIDPathReq), nil)
	d.AddOperation("deleteIntegrations", http.MethodPatch, "/khulnasoft/integration/delete",
		"Delete Integrations", "Delete integrations",
		http.StatusNoContent, []string{tagIntegration}, bearerToken, new(DeleteIntegrationReq), nil)
	d.AddOperation("addGenerativeAiIntegrationOpenAI", http.MethodPost, "/khulnasoft/generative-ai-integration/openai",
		"Add OpenAI Generative AI Integration", "Add a new OpenAI Generative AI Integration",
		http.StatusOK, []string{tagGenerativeAi}, bearerToken, new(AddGenerativeAiOpenAIIntegration), new(MessageResponse))
	d.AddOperation("addGenerativeAiIntegrationBedrock", http.MethodPost, "/khulnasoft/generative-ai-integration/bedrock",
		"Add AWS Bedrock Generative AI Integration", "Add a new AWS Bedrock Generative AI Integration",
		http.StatusOK, []string{tagGenerativeAi}, bearerToken, new(AddGenerativeAiBedrockIntegration), new(MessageResponse))
	d.AddOperation("autoAddGenerativeAiIntegration", http.MethodPost, "/khulnasoft/generative-ai-integration/auto-add",
		"Automatically add Generative AI Integration", "Automatically add Generative AI Integrations using IAM role",
		http.StatusAccepted, []string{tagGenerativeAi}, bearerToken, nil, nil)

	d.AddOperation("listGenerativeAiIntegration", http.MethodGet, "/khulnasoft/generative-ai-integration",
		"List Generative AI Integrations", "List all the added Generative AI Integrations",
		http.StatusOK, []string{tagGenerativeAi}, bearerToken, new(GenerativeAiIntegrationListRequest), new([]GenerativeAiIntegrationListResponse))
	d.AddOperation("deleteGenerativeAiIntegration", http.MethodDelete, "/khulnasoft/generative-ai-integration/{integration_id}",
		"Delete Generative AI Integration", "Delete Generative AI integration",
		http.StatusNoContent, []string{tagGenerativeAi}, bearerToken, new(IntegrationIDPathReq), nil)
	d.AddOperation("setDefaultGenerativeAiIntegration", http.MethodPut, "/khulnasoft/generative-ai-integration/{integration_id}/default",
		"Set Default Generative AI Integration", "Set Default Generative AI integration",
		http.StatusNoContent, []string{tagGenerativeAi}, bearerToken, new(IntegrationIDPathReq), nil)

	d.AddOperation("generativeAiIntegrationCloudPostureQuery", http.MethodPost, "/khulnasoft/generative-ai-integration/query/cloud-posture",
		"Send Cloud Posture query to Generative AI Integration", "Send Cloud Posture query to Generative AI Integration",
		http.StatusOK, []string{tagGenerativeAi}, bearerToken, new(GenerativeAiIntegrationCloudPostureRequest), new(string))
	d.AddOperation("generativeAiIntegrationLinuxPostureQuery", http.MethodPost, "/khulnasoft/generative-ai-integration/query/linux-posture",
		"Send Linux Posture query to Generative AI Integration", "Send Linux Posture query to Generative AI Integration",
		http.StatusOK, []string{tagGenerativeAi}, bearerToken, new(GenerativeAiIntegrationLinuxPostureRequest), new(string))
	d.AddOperation("generativeAiIntegrationKubernetesPostureQuery", http.MethodPost, "/khulnasoft/generative-ai-integration/query/kubernetes-posture",
		"Send Kubernetes Posture query to Generative AI Integration", "Send Kubernetes Posture query to Generative AI Integration",
		http.StatusOK, []string{tagGenerativeAi}, bearerToken, new(GenerativeAiIntegrationKubernetesPostureRequest), new(string))
	d.AddOperation("generativeAiIntegrationVulnerabilityQuery", http.MethodPost, "/khulnasoft/generative-ai-integration/query/vulnerability",
		"Send Vulnerability query to Generative AI Integration", "Send Vulnerability query to Generative AI Integration",
		http.StatusOK, []string{tagGenerativeAi}, bearerToken, new(GenerativeAiIntegrationVulnerabilityRequest), new(string))
	d.AddOperation("generativeAiIntegrationSecretQuery", http.MethodPost, "/khulnasoft/generative-ai-integration/query/secret",
		"Send Secret query to Generative AI Integration", "Send Secret query to Generative AI Integration",
		http.StatusOK, []string{tagGenerativeAi}, bearerToken, new(GenerativeAiIntegrationSecretRequest), new(string))
	d.AddOperation("generativeAiIntegrationMalwareQuery", http.MethodPost, "/khulnasoft/generative-ai-integration/query/malware",
		"Send Malware query to Generative AI Integration", "Send Malware query to Generative AI Integration",
		http.StatusOK, []string{tagGenerativeAi}, bearerToken, new(GenerativeAiIntegrationMalwareRequest), new(string))
}

func (d *OpenAPIDocs) AddReportsOperations() {
	d.AddOperation("generateReport", http.MethodPost, "/khulnasoft/reports",
		"Generate Report", "generate report for given type and filters",
		http.StatusOK, []string{tagReports}, bearerToken, new(GenerateReportReq), new(GenerateReportResp))
	d.AddOperation("listReports", http.MethodGet, "/khulnasoft/reports",
		"List Reports", "List all the available reports",
		http.StatusOK, []string{tagReports}, bearerToken, nil, new([]ExportReport))
	d.AddOperation("getReport", http.MethodGet, "/khulnasoft/reports/{report_id}",
		"Get Report", "get report by report_id",
		http.StatusOK, []string{tagReports}, bearerToken, new(ReportReq), new(ExportReport))
	d.AddOperation("deleteReport", http.MethodDelete, "/khulnasoft/reports/{report_id}",
		"Delete Report", "delete report for given report_id",
		http.StatusNoContent, []string{tagReports}, bearerToken, new(ReportReq), nil)
	d.AddOperation("bulkDeleteReports", http.MethodPatch, "/khulnasoft/reports/delete",
		"Bulk Delete Reports", "Bulk Delete reports",
		http.StatusNoContent, []string{tagReports}, bearerToken, new(BulkDeleteReportReq), nil)
}

func (d *OpenAPIDocs) AddSettingsOperations() {
	d.AddOperation("addEmailConfiguration", http.MethodPost, "/khulnasoft/settings/email",
		"Add Email Configuration", "This email configuration is used to send email notifications",
		http.StatusOK, []string{tagSettings}, bearerToken, new(EmailConfigurationAdd), new(MessageResponse))
	d.AddOperation("getEmailConfiguration", http.MethodGet, "/khulnasoft/settings/email",
		"Get Email Configurations", "Get Email Smtp / ses Configurations in system",
		http.StatusOK, []string{tagSettings}, bearerToken, nil, new([]EmailConfigurationResp))
	d.AddOperation("deleteEmailConfiguration", http.MethodDelete, "/khulnasoft/settings/email/{config_id}",
		"Delete Email Configurations", "Delete Email Smtp / ses Configurations in system",
		http.StatusNoContent, []string{tagSettings}, bearerToken, new(ConfigIDPathReq), nil)
	d.AddOperation("testConfiguredEmail", http.MethodPost, "/khulnasoft/settings/email/test",
		"Test Configured Email", "Test Configured Email",
		http.StatusOK, []string{tagSettings}, bearerToken, nil, new(MessageResponse))
	d.AddOperation("testUnconfiguredEmail", http.MethodPost, "/khulnasoft/settings/email/test-unconfigured",
		"Test Unconfigured Email", "Test Unconfigured Email",
		http.StatusOK, []string{tagSettings}, bearerToken, new(EmailConfigurationAdd), new(MessageResponse))
	d.AddOperation("getSettings", http.MethodGet, "/khulnasoft/settings/global-settings",
		"Get settings", "Get all settings",
		http.StatusOK, []string{tagSettings}, bearerToken, nil, new([]SettingsResponse))
	d.AddOperation("updateSetting", http.MethodPatch, "/khulnasoft/settings/global-settings/{id}",
		"Update setting", "Update setting",
		http.StatusNoContent, []string{tagSettings}, bearerToken, new(SettingUpdateRequest), nil)
	d.AddOperation("getUserAuditLogs", http.MethodPost, "/khulnasoft/settings/user-audit-log",
		"Get user audit logs", "Get audit logs for all users",
		http.StatusOK, []string{tagSettings}, bearerToken, new(GetAuditLogsRequest), new([]postgresqldb.GetAuditLogsRow))
	d.AddOperation("getUserAuditLogsCount", http.MethodGet, "/khulnasoft/settings/user-audit-log/count",
		"Get user audit logs count", "Get user audit logs count",
		http.StatusOK, []string{tagSettings}, bearerToken, nil, new(SearchCountResp))

	// Scheduled tasks
	d.AddOperation("getScheduledTasks", http.MethodGet, "/khulnasoft/scheduled-task",
		"Get scheduled tasks", "Get scheduled tasks",
		http.StatusOK, []string{tagSettings}, bearerToken, nil, new([]postgresqldb.Scheduler))
	d.AddOperation("updateScheduledTask", http.MethodPatch, "/khulnasoft/scheduled-task/{id}",
		"Update scheduled task", "Update scheduled task",
		http.StatusNoContent, []string{tagSettings}, bearerToken, new(UpdateScheduledTaskRequest), nil)
	d.AddOperation("addScheduledTask", http.MethodPost, "/khulnasoft/scheduled-task",
		"Add scheduled task", "Add scheduled task",
		http.StatusNoContent, []string{tagSettings}, bearerToken, new(AddScheduledTaskRequest), nil)
	d.AddOperation("deleteCustomScheduledTask", http.MethodDelete, "/khulnasoft/scheduled-task/{id}",
		"Delete Custom Schedule task", "Delete Custom Schedule task",
		http.StatusNoContent, []string{tagSettings}, bearerToken, new(ScheduleJobID), nil)

	d.AddOperation("uploadAgentVersion", http.MethodPut, "/khulnasoft/settings/agent/version",
		"Upload New agent version", "Upload Agent version",
		http.StatusOK, []string{tagSettings}, bearerToken, new(BinUploadRequest), nil)

	d.AddOperation("getAgentVersions", http.MethodGet, "/khulnasoft/settings/agent/versions",
		"Get available agent versions", "Get available agent versions",
		http.StatusOK, []string{tagSettings}, bearerToken, nil, new(ListAgentVersionResp))

	// Database upload
	d.AddOperation("uploadVulnerabilityDatabase", http.MethodPut, "/khulnasoft/database/vulnerability",
		"Upload Vulnerability Database", "Upload Vulnerability Database for use in vulnerability scans",
		http.StatusOK, []string{tagSettings}, bearerToken, new(threatintel.DBUploadRequest), new(MessageResponse))
	d.AddOperation("uploadSecretsRules", http.MethodPut, "/khulnasoft/database/secret",
		"Upload Secrets Rules", "Upload secrets rules for use in secrets scans",
		http.StatusOK, []string{tagSettings}, bearerToken, new(threatintel.DBUploadRequest), new(MessageResponse))
	d.AddOperation("uploadMalwareRules", http.MethodPut, "/khulnasoft/database/malware",
		"Upload Malware Rules", "Upload malware rules for use in malware scans",
		http.StatusOK, []string{tagSettings}, bearerToken, new(threatintel.DBUploadRequest), new(MessageResponse))
	d.AddOperation("uploadPostureControls", http.MethodPut, "/khulnasoft/database/posture",
		"Upload Posture Controls", "Upload posture controls for use in posture scans",
		http.StatusOK, []string{tagSettings}, bearerToken, new(threatintel.DBUploadRequest), new(MessageResponse))

	d.AddOperation("getAgentBinaryDownloadURL", http.MethodGet, "/khulnasoft/agent-deployment/binary/download-url",
		"Get agent binary download url", "Get agent binary download url",
		http.StatusOK, []string{tagSettings}, bearerToken, nil, new(GetAgentBinaryDownloadURLResponse))
}

func (d *OpenAPIDocs) AddLicenseOperations() {
	// License
	d.AddOperation("generateLicense", http.MethodPost, "/khulnasoft/license/generate",
		"Generate License Key", "Generate a new Kengine license key",
		http.StatusOK, []string{tagSettings}, bearerToken, new(GenerateLicenseRequest), new(GenerateLicenseResponse))

	d.AddOperation("registerLicense", http.MethodPost, "/khulnasoft/license",
		"Register License", "Register new license key to the console and activate",
		http.StatusOK, []string{tagSettings}, bearerToken, new(RegisterLicenseRequest), new(RegisterLicenseResponse))

	d.AddOperation("getLicense", http.MethodGet, "/khulnasoft/license",
		"Get License Details", "Get license status and expiry",
		http.StatusOK, []string{tagSettings}, bearerToken, nil, new(License))

	d.AddOperation("deleteLicense", http.MethodDelete, "/khulnasoft/license",
		"Delete License", "Delete license from the console database",
		http.StatusNoContent, []string{tagSettings}, bearerToken, nil, nil)
}

func (d *OpenAPIDocs) AddDiffAddOperations() {
	d.AddOperation("diffAddVulnerability", http.MethodPost, "/khulnasoft/diff-add/vulnerability",
		"Get Vulnerability Diff", "Get Vulnerability Diff between two scans",
		http.StatusOK, []string{tagDiffAdd}, bearerToken, new(ScanCompareReq), new(ScanCompareResVulnerability))
	d.AddOperation("diffAddSecret", http.MethodPost, "/khulnasoft/diff-add/secret",
		"Get Secret Diff", "Get Secret Diff between two scans",
		http.StatusOK, []string{tagDiffAdd}, bearerToken, new(ScanCompareReq), new(ScanCompareResSecret))
	d.AddOperation("diffAddCompliance", http.MethodPost, "/khulnasoft/diff-add/compliance",
		"Get Compliance Diff", "Get Compliance Diff between two scans",
		http.StatusOK, []string{tagDiffAdd}, bearerToken, new(ScanCompareReq), new(ScanCompareResCompliance))
	d.AddOperation("diffAddMalware", http.MethodPost, "/khulnasoft/diff-add/malware",
		"Get Malware Diff", "Get Malware Diff between two scans",
		http.StatusOK, []string{tagDiffAdd}, bearerToken, new(ScanCompareReq), new(ScanCompareResMalware))
	d.AddOperation("diffAddCloudCompliance", http.MethodPost, "/khulnasoft/diff-add/cloud-compliance",
		"Get Cloud Compliance Diff", "Get Cloud Compliance Diff between two scans",
		http.StatusOK, []string{tagDiffAdd}, bearerToken, new(ScanCompareReq), new(ScanCompareResCloudCompliance))
}

func (d *OpenAPIDocs) AddCompletionOperations() {
	d.AddOperation("completeProcessInfo", http.MethodPost, "/khulnasoft/complete/process",
		"Get Completion for process fields", "Complete process info",
		http.StatusOK, []string{tagCompletion}, bearerToken, new(CompletionNodeFieldReq), new(CompletionNodeFieldRes))
	d.AddOperation("completeVulnerabilityInfo", http.MethodPost, "/khulnasoft/complete/vulnerability",
		"Get Completion for vulnerability fields", "Complete vulnerability info",
		http.StatusOK, []string{tagCompletion}, bearerToken, new(CompletionNodeFieldReq), new(CompletionNodeFieldRes))
	d.AddOperation("completeHostInfo", http.MethodPost, "/khulnasoft/complete/host",
		"Get Completion for host fields", "Complete host info",
		http.StatusOK, []string{tagCompletion}, bearerToken, new(CompletionNodeFieldReq), new(CompletionNodeFieldRes))
	d.AddOperation("completeCloudCompliance", http.MethodPost, "/khulnasoft/complete/cloud-compliance",
		"Get Completion for cloud compliance fields", "Complete cloud compliance info",
		http.StatusOK, []string{tagCompletion}, bearerToken, new(CompletionNodeFieldReq), new(CompletionNodeFieldRes))
	d.AddOperation("completeComplianceInfo", http.MethodPost, "/khulnasoft/complete/compliance",
		"Get Completion for compliance fields", "Complete compliance info",
		http.StatusOK, []string{tagCompletion}, bearerToken, new(CompletionNodeFieldReq), new(CompletionNodeFieldRes))
	d.AddOperation("completePodInfo", http.MethodPost, "/khulnasoft/complete/pod",
		"Get Completion for Pod fields", "Complete Pod info",
		http.StatusOK, []string{tagCompletion}, bearerToken, new(CompletionNodeFieldReq), new(CompletionNodeFieldRes))
	d.AddOperation("completeContainerInfo", http.MethodPost, "/khulnasoft/complete/container",
		"Get Completion for Container fields", "Complete Container info",
		http.StatusOK, []string{tagCompletion}, bearerToken, new(CompletionNodeFieldReq), new(CompletionNodeFieldRes))
}
