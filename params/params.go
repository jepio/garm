// Copyright 2022 Cloudbase Solutions SRL
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package params

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"time"

	commonParams "github.com/cloudbase/garm-provider-common/params"

	"github.com/cloudbase/garm/util/appdefaults"

	"github.com/google/go-github/v57/github"
	"github.com/google/uuid"
)

type (
	PoolType            string
	EventType           string
	EventLevel          string
	ProviderType        string
	JobStatus           string
	RunnerStatus        string
	WebhookEndpointType string
)

const (
	// LXDProvider represents the LXD provider.
	LXDProvider ProviderType = "lxd"
	// ExternalProvider represents an external provider.
	ExternalProvider ProviderType = "external"
)

const (
	// WebhookEndpointDirect instructs garm that it should attempt to create a webhook
	// in the target entity, using the callback URL defined in the config as a target.
	WebhookEndpointDirect WebhookEndpointType = "direct"
	// WebhookEndpointTunnel instructs garm that it should attempt to create a webhook
	// in the target entity, using the tunnel URL as a base for the webhook URL.
	// This is defined for future use.
	WebhookEndpointTunnel WebhookEndpointType = "tunnel"
)

const (
	JobStatusQueued     JobStatus = "queued"
	JobStatusInProgress JobStatus = "in_progress"
	JobStatusCompleted  JobStatus = "completed"
)

const (
	RepositoryPool   PoolType = "repository"
	OrganizationPool PoolType = "organization"
	EnterprisePool   PoolType = "enterprise"
)

const (
	StatusEvent     EventType = "status"
	FetchTokenEvent EventType = "fetchToken"
)

const (
	EventInfo    EventLevel = "info"
	EventWarning EventLevel = "warning"
	EventError   EventLevel = "error"
)

const (
	RunnerIdle       RunnerStatus = "idle"
	RunnerPending    RunnerStatus = "pending"
	RunnerTerminated RunnerStatus = "terminated"
	RunnerInstalling RunnerStatus = "installing"
	RunnerFailed     RunnerStatus = "failed"
	RunnerActive     RunnerStatus = "active"
)

type StatusMessage struct {
	CreatedAt  time.Time  `json:"created_at"`
	Message    string     `json:"message"`
	EventType  EventType  `json:"event_type"`
	EventLevel EventLevel `json:"event_level"`
}

type Instance struct {
	// ID is the database ID of this instance.
	ID string `json:"id,omitempty"`

	// PeoviderID is the unique ID the provider associated
	// with the compute instance. We use this to identify the
	// instance in the provider.
	ProviderID string `json:"provider_id,omitempty"`

	// AgentID is the github runner agent ID.
	AgentID int64 `json:"agent_id"`

	// Name is the name associated with an instance. Depending on
	// the provider, this may or may not be useful in the context of
	// the provider, but we can use it internally to identify the
	// instance.
	Name string `json:"name,omitempty"`

	// OSType is the operating system type. For now, only Linux and
	// Windows are supported.
	OSType commonParams.OSType `json:"os_type,omitempty"`

	// OSName is the name of the OS. Eg: ubuntu, centos, etc.
	OSName string `json:"os_name,omitempty"`

	// OSVersion is the version of the operating system.
	OSVersion string `json:"os_version,omitempty"`

	// OSArch is the operating system architecture.
	OSArch commonParams.OSArch `json:"os_arch,omitempty"`

	// Addresses is a list of IP addresses the provider reports
	// for this instance.
	Addresses []commonParams.Address `json:"addresses,omitempty"`

	// Status is the status of the instance inside the provider (eg: running, stopped, etc)
	Status commonParams.InstanceStatus `json:"status,omitempty"`

	// RunnerStatus is the github runner status as it appears on GitHub.
	RunnerStatus RunnerStatus `json:"runner_status,omitempty"`

	// PoolID is the ID of the garm pool to which a runner belongs.
	PoolID string `json:"pool_id,omitempty"`

	// ProviderFault holds any error messages captured from the IaaS provider that is
	// responsible for managing the lifecycle of the runner.
	ProviderFault []byte `json:"provider_fault,omitempty"`

	// StatusMessages is a list of status messages sent back by the runner as it sets itself
	// up.
	StatusMessages []StatusMessage `json:"status_messages,omitempty"`

	// UpdatedAt is the timestamp of the last update to this runner.
	UpdatedAt time.Time `json:"updated_at"`

	// GithubRunnerGroup is the github runner group to which the runner belongs.
	// The runner group must be created by someone with access to the enterprise.
	GitHubRunnerGroup string `json:"github-runner-group"`

	// Do not serialize sensitive info.
	CallbackURL      string            `json:"-"`
	MetadataURL      string            `json:"-"`
	CreateAttempt    int               `json:"-"`
	TokenFetched     bool              `json:"-"`
	AditionalLabels  []string          `json:"-"`
	JitConfiguration map[string]string `json:"-"`
}

func (i Instance) GetName() string {
	return i.Name
}

func (i Instance) GetID() string {
	return i.ID
}

// used by swagger client generated code
type Instances []Instance

type BootstrapInstance struct {
	Name  string                              `json:"name"`
	Tools []*github.RunnerApplicationDownload `json:"tools"`
	// RepoURL is the URL the github runner agent needs to configure itself.
	RepoURL string `json:"repo_url"`
	// CallbackUrl is the URL where the instance can send a post, signaling
	// progress or status.
	CallbackURL string `json:"callback-url"`
	// MetadataURL is the URL where instances can fetch information needed to set themselves up.
	MetadataURL string `json:"metadata-url"`
	// InstanceToken is the token that needs to be set by the instance in the headers
	// in order to send updated back to the garm via CallbackURL.
	InstanceToken string `json:"instance-token"`
	// SSHKeys are the ssh public keys we may want to inject inside the runners, if the
	// provider supports it.
	SSHKeys []string `json:"ssh-keys"`
	// ExtraSpecs is an opaque raw json that gets sent to the provider
	// as part of the bootstrap params for instances. It can contain
	// any kind of data needed by providers. The contents of this field means
	// nothing to garm itself. We don't act on the information in this field at
	// all. We only validate that it's a proper json.
	ExtraSpecs json.RawMessage `json:"extra_specs,omitempty"`

	// GitHubRunnerGroup is the github runner group in which the newly installed runner
	// should be added to. The runner group must be created by someone with access to the
	// enterprise.
	GitHubRunnerGroup string `json:"github-runner-group"`

	// CACertBundle is a CA certificate bundle which will be sent to instances and which
	// will tipically be installed as a system wide trusted root CA. by either cloud-init
	// or whatever mechanism the provider will use to set up the runner.
	CACertBundle []byte `json:"ca-cert-bundle"`

	// OSArch is the target OS CPU architecture of the runner.
	OSArch commonParams.OSArch `json:"arch"`

	// OSType is the target OS platform of the runner (windows, linux).
	OSType commonParams.OSType `json:"os_type"`

	// Flavor is the platform specific abstraction that defines what resources will be allocated
	// to the runner (CPU, RAM, disk space, etc). This field is meaningful to the provider which
	// handles the actual creation.
	Flavor string `json:"flavor"`

	// Image is the platform specific identifier of the operating system template that will be used
	// to spin up a new machine.
	Image string `json:"image"`

	// Labels are a list of github runner labels that will be added to the runner.
	Labels []string `json:"labels"`

	// PoolID is the ID of the garm pool to which this runner belongs.
	PoolID string `json:"pool_id"`

	// UserDataOptions are the options for the user data generation.
	UserDataOptions UserDataOptions `json:"user_data_options"`
}

type UserDataOptions struct {
	DisableUpdatesOnBoot bool     `json:"disable_updates_on_boot"`
	ExtraPackages        []string `json:"extra_packages"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Pool struct {
	RunnerPrefix

	ID                     string              `json:"id"`
	ProviderName           string              `json:"provider_name"`
	MaxRunners             uint                `json:"max_runners"`
	MinIdleRunners         uint                `json:"min_idle_runners"`
	Image                  string              `json:"image"`
	Flavor                 string              `json:"flavor"`
	OSType                 commonParams.OSType `json:"os_type"`
	OSArch                 commonParams.OSArch `json:"os_arch"`
	Tags                   []Tag               `json:"tags"`
	Enabled                bool                `json:"enabled"`
	Instances              []Instance          `json:"instances"`
	RepoID                 string              `json:"repo_id,omitempty"`
	RepoName               string              `json:"repo_name,omitempty"`
	OrgID                  string              `json:"org_id,omitempty"`
	OrgName                string              `json:"org_name,omitempty"`
	EnterpriseID           string              `json:"enterprise_id,omitempty"`
	EnterpriseName         string              `json:"enterprise_name,omitempty"`
	RunnerBootstrapTimeout uint                `json:"runner_bootstrap_timeout"`
	// ExtraSpecs is an opaque raw json that gets sent to the provider
	// as part of the bootstrap params for instances. It can contain
	// any kind of data needed by providers. The contents of this field means
	// nothing to garm itself. We don't act on the information in this field at
	// all. We only validate that it's a proper json.
	ExtraSpecs json.RawMessage `json:"extra_specs,omitempty"`
	// GithubRunnerGroup is the github runner group in which the runners will be added.
	// The runner group must be created by someone with access to the enterprise.
	GitHubRunnerGroup string `json:"github-runner-group"`
}

func (p Pool) GetID() string {
	return p.ID
}

func (p *Pool) RunnerTimeout() uint {
	if p.RunnerBootstrapTimeout == 0 {
		return appdefaults.DefaultRunnerBootstrapTimeout
	}
	return p.RunnerBootstrapTimeout
}

func (p *Pool) PoolType() PoolType {
	if p.RepoID != "" {
		return RepositoryPool
	} else if p.OrgID != "" {
		return OrganizationPool
	} else if p.EnterpriseID != "" {
		return EnterprisePool
	}
	return ""
}

func (p *Pool) HasRequiredLabels(set []string) bool {
	asMap := make(map[string]struct{}, len(p.Tags))
	for _, t := range p.Tags {
		asMap[t.Name] = struct{}{}
	}

	for _, l := range set {
		if _, ok := asMap[l]; !ok {
			return false
		}
	}
	return true
}

// used by swagger client generated code
type Pools []Pool

type Internal struct {
	OAuth2Token          string `json:"oauth2"`
	ControllerID         string `json:"controller_id"`
	InstanceCallbackURL  string `json:"instance_callback_url"`
	InstanceMetadataURL  string `json:"instance_metadata_url"`
	BaseWebhookURL       string `json:"base_webhook_url"`
	ControllerWebhookURL string `json:"controller_webhook_url"`

	JWTSecret string `json:"jwt_secret"`
	// GithubCredentialsDetails contains all info about the credentials, except the
	// token, which is added above.
	GithubCredentialsDetails GithubCredentials `json:"gh_creds_details"`
}

type Repository struct {
	ID                string            `json:"id"`
	Owner             string            `json:"owner"`
	Name              string            `json:"name"`
	Pools             []Pool            `json:"pool,omitempty"`
	CredentialsName   string            `json:"credentials_name"`
	PoolManagerStatus PoolManagerStatus `json:"pool_manager_status,omitempty"`
	// Do not serialize sensitive info.
	WebhookSecret string `json:"-"`
}

func (r Repository) GetName() string {
	return r.Name
}

func (r Repository) GetID() string {
	return r.ID
}

// used by swagger client generated code
type Repositories []Repository

type Organization struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Pools             []Pool            `json:"pool,omitempty"`
	CredentialsName   string            `json:"credentials_name"`
	PoolManagerStatus PoolManagerStatus `json:"pool_manager_status,omitempty"`
	// Do not serialize sensitive info.
	WebhookSecret string `json:"-"`
}

func (o Organization) GetName() string {
	return o.Name
}

func (o Organization) GetID() string {
	return o.ID
}

// used by swagger client generated code
type Organizations []Organization

type Enterprise struct {
	ID                string            `json:"id"`
	Name              string            `json:"name"`
	Pools             []Pool            `json:"pool,omitempty"`
	CredentialsName   string            `json:"credentials_name"`
	PoolManagerStatus PoolManagerStatus `json:"pool_manager_status,omitempty"`
	// Do not serialize sensitive info.
	WebhookSecret string `json:"-"`
}

func (e Enterprise) GetName() string {
	return e.Name
}

func (e Enterprise) GetID() string {
	return e.ID
}

// used by swagger client generated code
type Enterprises []Enterprise

// Users holds information about a particular user
type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Password  string    `json:"-"`
	Enabled   bool      `json:"enabled"`
	IsAdmin   bool      `json:"is_admin"`
}

// JWTResponse holds the JWT token returned as a result of a
// successful auth
type JWTResponse struct {
	Token string `json:"token"`
}

type ControllerInfo struct {
	ControllerID         uuid.UUID `json:"controller_id"`
	Hostname             string    `json:"hostname"`
	MetadataURL          string    `json:"metadata_url"`
	CallbackURL          string    `json:"callback_url"`
	WebhookURL           string    `json:"webhook_url"`
	ControllerWebhookURL string    `json:"controller_webhook_url"`
}

type GithubCredentials struct {
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	BaseURL       string `json:"base_url"`
	APIBaseURL    string `json:"api_base_url"`
	UploadBaseURL string `json:"upload_base_url"`
	CABundle      []byte `json:"ca_bundle,omitempty"`
}

func (g GithubCredentials) RootCertificateBundle() (CertificateBundle, error) {
	if len(g.CABundle) == 0 {
		return CertificateBundle{}, nil
	}

	ret := map[string][]byte{}

	var block *pem.Block
	var rest []byte = g.CABundle
	for {
		block, rest = pem.Decode(rest)
		if block == nil {
			break
		}
		pub, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return CertificateBundle{}, err
		}
		out := &bytes.Buffer{}
		if err := pem.Encode(out, &pem.Block{Type: "CERTIFICATE", Bytes: block.Bytes}); err != nil {
			return CertificateBundle{}, err
		}
		ret[fmt.Sprintf("%d", pub.SerialNumber)] = out.Bytes()
	}

	return CertificateBundle{
		RootCertificates: ret,
	}, nil
}

// used by swagger client generated code
type Credentials []GithubCredentials

type Provider struct {
	Name         string       `json:"name"`
	ProviderType ProviderType `json:"type"`
	Description  string       `json:"description"`
}

// used by swagger client generated code
type Providers []Provider

type UpdatePoolStateParams struct {
	WebhookSecret  string
	InternalConfig *Internal
}

type PoolManagerStatus struct {
	IsRunning     bool   `json:"running"`
	FailureReason string `json:"failure_reason,omitempty"`
}

type RunnerInfo struct {
	Name   string
	Labels []string
}

type RunnerPrefix struct {
	Prefix string `json:"runner_prefix"`
}

func (p RunnerPrefix) GetRunnerPrefix() string {
	if p.Prefix == "" {
		return DefaultRunnerPrefix
	}
	return p.Prefix
}

type Job struct {
	// ID is the ID of the job.
	ID int64 `json:"id"`
	// RunID is the ID of the workflow run. A run may have multiple jobs.
	RunID int64 `json:"run_id"`
	// Action is the specific activity that triggered the event.
	Action string `json:"action"`
	// Conclusion is the outcome of the job.
	// Possible values: "success", "failure", "neutral", "cancelled", "skipped",
	// "timed_out", "action_required"
	Conclusion string `json:"conclusion"`
	// Status is the phase of the lifecycle that the job is currently in.
	// "queued", "in_progress" and "completed".
	Status string `json:"status"`
	// Name is the name if the job that was triggered.
	Name string `json:"name"`

	StartedAt   time.Time
	CompletedAt time.Time

	GithubRunnerID  int64  `json:"runner_id"`
	RunnerName      string `json:"runner_name"`
	RunnerGroupID   int64  `json:"runner_group_id"`
	RunnerGroupName string `json:"runner_group_name"`

	// repository in which the job was triggered.
	RepositoryName  string
	RepositoryOwner string

	Labels []string

	// The entity that received the hook.
	//
	// Webhooks may be configured on the repo, the org and/or the enterprise.
	// If we only configure a repo to use garm, we'll only ever receive a
	// webhook from the repo. But if we configure the parent org of the repo and
	// the parent enterprise of the org to use garm, a webhook will be sent for each
	// entity type, in response to one workflow event. Thus, we will get 3 webhooks
	// with the same run_id and job id. Record all involved entities in the same job
	// if we have them configured in garm.
	RepoID       *uuid.UUID `json:"repo_id,omitempty"`
	OrgID        *uuid.UUID `json:"org_id,omitempty"`
	EnterpriseID *uuid.UUID `json:"enterprise_id,omitempty"`

	LockedBy uuid.UUID

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// used by swagger client generated code
type Jobs []Job

type InstallWebhookParams struct {
	WebhookEndpointType WebhookEndpointType `json:"webhook_endpoint_type"`
	InsecureSSL         bool                `json:"insecure_ssl"`
}

type HookInfo struct {
	ID          int64    `json:"id"`
	URL         string   `json:"url"`
	Events      []string `json:"events"`
	Active      bool     `json:"active"`
	InsecureSSL bool     `json:"insecure_ssl"`
}

type CertificateBundle struct {
	RootCertificates map[string][]byte `json:"root_certificates"`
}

type UpdateSystemInfoParams struct {
	OSName    string `json:"os_name,omitempty"`
	OSVersion string `json:"os_version,omitempty"`
	AgentID   *int64 `json:"agent_id,omitempty"`
}
