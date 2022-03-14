---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_connector Resource - harness"
subcategory: ""
description: |-
  [NG] - Resource for creating a connector. This resource is part of the Harness nextgen platform.
---

# harness_connector (Resource)

[NG] - Resource for creating a connector. This resource is part of the Harness nextgen platform.

## Example Usage

```terraform
# Create a kubernetes connector

resource "harness_connector" "test" {
  identifier = "k8s"
  name       = "Test Kubernetes Cluster"

  k8s_cluster {
    service_account {
      master_url                = "https://kubernetes.example.com"
      service_account_token_ref = "account.k8s_service_account_token"
    }
  }
}

# Create a docker registry connector

resource "harness_connector" "test" {
  identifier = "dockerhub"
  name       = "Docker Hub"

  docker_registry {
    type = "DockerHub"
    url  = "https://hub.docker.io"

    credentials {
      username     = "admin"
      password_ref = "account.docker_registry_password"
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **identifier** (String) The unique identifier for the connector.
- **name** (String) The name of the connector.

### Optional

- **app_dynamics** (Block List, Max: 1) App Dynamics connector (see [below for nested schema](#nestedblock--app_dynamics))
- **artifactory** (Block List, Max: 1) Artifactory connector. (see [below for nested schema](#nestedblock--artifactory))
- **aws** (Block List, Max: 1) Aws account configuration. (see [below for nested schema](#nestedblock--aws))
- **aws_cloudcost** (Block List, Max: 1) Aws cloud cost account configuration. (see [below for nested schema](#nestedblock--aws_cloudcost))
- **aws_kms** (Block List, Max: 1) The AWS KMS configuration. (see [below for nested schema](#nestedblock--aws_kms))
- **aws_secret_manager** (Block List, Max: 1) The AWS Secret Manager configuration. (see [below for nested schema](#nestedblock--aws_secret_manager))
- **bitbucket** (Block List, Max: 1) BitBucket connector (see [below for nested schema](#nestedblock--bitbucket))
- **branch** (String) The branch to use for the connector.
- **datadog** (Block List, Max: 1) Datadog connector (see [below for nested schema](#nestedblock--datadog))
- **description** (String) The description of the connector.
- **docker_registry** (Block List, Max: 1) The docker registry to use for the connector. (see [below for nested schema](#nestedblock--docker_registry))
- **dynatrace** (Block List, Max: 1) Dynatrace connector (see [below for nested schema](#nestedblock--dynatrace))
- **gcp** (Block List, Max: 1) Gcp connector configuration. (see [below for nested schema](#nestedblock--gcp))
- **git** (Block List, Max: 1) Git connector (see [below for nested schema](#nestedblock--git))
- **github** (Block List, Max: 1) Github connector (see [below for nested schema](#nestedblock--github))
- **gitlab** (Block List, Max: 1) Gitlab connector (see [below for nested schema](#nestedblock--gitlab))
- **http_helm** (Block List, Max: 1) Helm connector. (see [below for nested schema](#nestedblock--http_helm))
- **id** (String) The ID of this resource.
- **jira** (Block List, Max: 1) Jira connector (see [below for nested schema](#nestedblock--jira))
- **k8s_cluster** (Block List, Max: 1) The k8s cluster to use for the connector. (see [below for nested schema](#nestedblock--k8s_cluster))
- **newrelic** (Block List, Max: 1) NewRelic connector (see [below for nested schema](#nestedblock--newrelic))
- **nexus** (Block List, Max: 1) Nexus connector. (see [below for nested schema](#nestedblock--nexus))
- **org_id** (String) The unique identifier for the organization.
- **pagerduty** (Block List, Max: 1) PagerDuty connector (see [below for nested schema](#nestedblock--pagerduty))
- **project_id** (String) The unique identifier for the project.
- **prometheus** (Block List, Max: 1) Prometheus connector (see [below for nested schema](#nestedblock--prometheus))
- **repo_id** (String) The unique identifier for the repository.
- **splunk** (Block List, Max: 1) Splunk connector (see [below for nested schema](#nestedblock--splunk))
- **sumologic** (Block List, Max: 1) SumoLogic connector (see [below for nested schema](#nestedblock--sumologic))
- **tags** (Set of String) Tags associated with the connector.

<a id="nestedblock--app_dynamics"></a>
### Nested Schema for `app_dynamics`

Required:

- **account_name** (String) The App Dynamics account name.
- **url** (String) Url of the App Dynamics controller.

Optional:

- **api_token** (Block List, Max: 1) Authenticate to App Dynamics using api token. (see [below for nested schema](#nestedblock--app_dynamics--api_token))
- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.
- **username_password** (Block List, Max: 1) Authenticate to App Dynamics using username and password. (see [below for nested schema](#nestedblock--app_dynamics--username_password))

<a id="nestedblock--app_dynamics--api_token"></a>
### Nested Schema for `app_dynamics.api_token`

Required:

- **client_id** (String) The client id used for connecting to App Dynamics.
- **client_secret_ref** (String) Reference to the Harness secret containing the App Dynamics client secret.


<a id="nestedblock--app_dynamics--username_password"></a>
### Nested Schema for `app_dynamics.username_password`

Required:

- **password_ref** (String) Reference to a secret containing the password to use for authentication.
- **username** (String) Username to use for authentication.



<a id="nestedblock--artifactory"></a>
### Nested Schema for `artifactory`

Required:

- **url** (String) URL of the Artifactory server.

Optional:

- **credentials** (Block List, Max: 1) Credentials to use for authentication. (see [below for nested schema](#nestedblock--artifactory--credentials))
- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.

<a id="nestedblock--artifactory--credentials"></a>
### Nested Schema for `artifactory.credentials`

Required:

- **password_ref** (String) Reference to a secret containing the password to use for authentication.

Optional:

- **username** (String) Username to use for authentication.
- **username_ref** (String) Reference to a secret containing the username to use for authentication.



<a id="nestedblock--aws"></a>
### Nested Schema for `aws`

Optional:

- **cross_account_access** (Block List, Max: 1) Select this option if you want to use one AWS account for the connection, but you want to deploy or build in a different AWS account. In this scenario, the AWS account used for AWS access in Credentials will assume the IAM role you specify in Cross-account role ARN setting. This option uses the AWS Security Token Service (STS) feature. (see [below for nested schema](#nestedblock--aws--cross_account_access))
- **inherit_from_delegate** (Block List, Max: 1) Inherit credentials from the delegate. (see [below for nested schema](#nestedblock--aws--inherit_from_delegate))
- **irsa** (Block List, Max: 1) Use IAM role for service accounts. (see [below for nested schema](#nestedblock--aws--irsa))
- **manual** (Block List, Max: 1) Use IAM role for service accounts. (see [below for nested schema](#nestedblock--aws--manual))

<a id="nestedblock--aws--cross_account_access"></a>
### Nested Schema for `aws.cross_account_access`

Required:

- **role_arn** (String) The Amazon Resource Name (ARN) of the role that you want to assume. This is an IAM role in the target AWS account.

Optional:

- **external_id** (String) If the administrator of the account to which the role belongs provided you with an external ID, then enter that value.


<a id="nestedblock--aws--inherit_from_delegate"></a>
### Nested Schema for `aws.inherit_from_delegate`

Required:

- **delegate_selectors** (Set of String) The delegates to inherit the credentials from.


<a id="nestedblock--aws--irsa"></a>
### Nested Schema for `aws.irsa`

Required:

- **delegate_selectors** (Set of String) The delegates to inherit the credentials from.


<a id="nestedblock--aws--manual"></a>
### Nested Schema for `aws.manual`

Required:

- **secret_key_ref** (String) Reference to the Harness secret containing the aws secret key.

Optional:

- **access_key** (String) AWS access key.
- **access_key_ref** (String) Reference to the Harness secret containing the aws access key.
- **delegate_selectors** (Set of String) Connect only use delegates with these tags.



<a id="nestedblock--aws_cloudcost"></a>
### Nested Schema for `aws_cloudcost`

Required:

- **account_id** (String) The AWS account id.
- **cross_account_access** (Block List, Min: 1, Max: 1) Harness uses the secure cross-account role to access your AWS account. The role includes a restricted policy to access the cost and usage reports and resources for the sole purpose of cost analysis and cost optimization. (see [below for nested schema](#nestedblock--aws_cloudcost--cross_account_access))
- **features_enabled** (Set of String) The features enabled for the connector. Valid values are BILLING, OPTIMIZATION, VISIBILITY.
- **report_name** (String) The cost and usage report name. Provided in the delivery options when the template is opened in the AWS console.
- **s3_bucket** (String) The name of s3 bucket.

<a id="nestedblock--aws_cloudcost--cross_account_access"></a>
### Nested Schema for `aws_cloudcost.cross_account_access`

Required:

- **external_id** (String) The external id of the role to use for cross-account access. This is a random unique value to provide additional secure authentication.
- **role_arn** (String) The ARN of the role to use for cross-account access.



<a id="nestedblock--aws_kms"></a>
### Nested Schema for `aws_kms`

Required:

- **arn_ref** (String) A reference to the Harness secret containing the ARN of the AWS KMS.
- **credentials** (Block List, Min: 1, Max: 1) The credentials to use for connecting to aws. (see [below for nested schema](#nestedblock--aws_kms--credentials))
- **region** (String) The AWS region where the AWS Secret Manager is.

Optional:

- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.

<a id="nestedblock--aws_kms--credentials"></a>
### Nested Schema for `aws_kms.credentials`

Optional:

- **assume_role** (Block List, Max: 1) Connect using STS assume role. (see [below for nested schema](#nestedblock--aws_kms--credentials--assume_role))
- **inherit_from_delegate** (Boolean) Inherit the credentials from from the delegate.
- **manual** (Block List, Max: 1) Specify the AWS key and secret used for authenticating. (see [below for nested schema](#nestedblock--aws_kms--credentials--manual))

<a id="nestedblock--aws_kms--credentials--assume_role"></a>
### Nested Schema for `aws_kms.credentials.assume_role`

Required:

- **duration** (Number) The duration, in seconds, of the role session. The value can range from 900 seconds (15 minutes) to 3600 seconds (1 hour). By default, the value is set to 3600 seconds. An expiration can also be specified in the client request body. The minimum value is 1 hour.
- **role_arn** (String) The ARN of the role to assume.

Optional:

- **external_id** (String) If the administrator of the account to which the role belongs provided you with an external ID, then enter that value.


<a id="nestedblock--aws_kms--credentials--manual"></a>
### Nested Schema for `aws_kms.credentials.manual`

Required:

- **access_key_ref** (String) The reference to the Harness secret containing the AWS access key.
- **secret_key_ref** (String) The reference to the Harness secret containing the AWS secret key.




<a id="nestedblock--aws_secret_manager"></a>
### Nested Schema for `aws_secret_manager`

Required:

- **credentials** (Block List, Min: 1, Max: 1) The credentials to use for connecting to aws. (see [below for nested schema](#nestedblock--aws_secret_manager--credentials))
- **region** (String) The AWS region where the AWS Secret Manager is.

Optional:

- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.
- **secret_name_prefix** (String) A prefix to be added to all secrets.

<a id="nestedblock--aws_secret_manager--credentials"></a>
### Nested Schema for `aws_secret_manager.credentials`

Optional:

- **assume_role** (Block List, Max: 1) Connect using STS assume role. (see [below for nested schema](#nestedblock--aws_secret_manager--credentials--assume_role))
- **inherit_from_delegate** (Boolean) Inherit the credentials from from the delegate.
- **manual** (Block List, Max: 1) Specify the AWS key and secret used for authenticating. (see [below for nested schema](#nestedblock--aws_secret_manager--credentials--manual))

<a id="nestedblock--aws_secret_manager--credentials--assume_role"></a>
### Nested Schema for `aws_secret_manager.credentials.assume_role`

Required:

- **duration** (Number) The duration, in seconds, of the role session. The value can range from 900 seconds (15 minutes) to 3600 seconds (1 hour). By default, the value is set to 3600 seconds. An expiration can also be specified in the client request body. The minimum value is 1 hour.
- **role_arn** (String) The ARN of the role to assume.

Optional:

- **external_id** (String) If the administrator of the account to which the role belongs provided you with an external ID, then enter that value.


<a id="nestedblock--aws_secret_manager--credentials--manual"></a>
### Nested Schema for `aws_secret_manager.credentials.manual`

Required:

- **access_key_ref** (String) The reference to the Harness secret containing the AWS access key.
- **secret_key_ref** (String) The reference to the Harness secret containing the AWS secret key.




<a id="nestedblock--bitbucket"></a>
### Nested Schema for `bitbucket`

Required:

- **connection_type** (String) Whether the connection we're making is to a BitBucket repository or a BitBucket account. Valid values are Account, Repo.
- **credentials** (Block List, Min: 1, Max: 1) Credentials to use for the connection. (see [below for nested schema](#nestedblock--bitbucket--credentials))
- **url** (String) Url of the BitBucket repository or account.

Optional:

- **api_authentication** (Block List, Max: 1) Configuration for using the BitBucket api. API Access is required for using “Git Experience”, for creation of Git based triggers, Webhooks management and updating Git statuses. (see [below for nested schema](#nestedblock--bitbucket--api_authentication))
- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.
- **validation_repo** (String) Repository to test the connection with. This is only used when `connection_type` is `Account`.

<a id="nestedblock--bitbucket--credentials"></a>
### Nested Schema for `bitbucket.credentials`

Optional:

- **http** (Block List, Max: 1) Authenticate using Username and password over http(s) for the connection. (see [below for nested schema](#nestedblock--bitbucket--credentials--http))
- **ssh** (Block List, Max: 1) Authenticate using SSH for the connection. (see [below for nested schema](#nestedblock--bitbucket--credentials--ssh))

<a id="nestedblock--bitbucket--credentials--http"></a>
### Nested Schema for `bitbucket.credentials.http`

Optional:

- **password_ref** (String) Reference to a secret containing the password to use for authentication.
- **username** (String) Username to use for authentication.
- **username_ref** (String) Reference to a secret containing the username to use for authentication.


<a id="nestedblock--bitbucket--credentials--ssh"></a>
### Nested Schema for `bitbucket.credentials.ssh`

Required:

- **ssh_key_ref** (String) Reference to the Harness secret containing the ssh key.



<a id="nestedblock--bitbucket--api_authentication"></a>
### Nested Schema for `bitbucket.api_authentication`

Required:

- **token_ref** (String) Personal access token for interacting with the BitBucket api.

Optional:

- **username** (String) The username used for connecting to the api.
- **username_ref** (String) The name of the Harness secret containing the username.



<a id="nestedblock--datadog"></a>
### Nested Schema for `datadog`

Required:

- **api_key_ref** (String) Reference to the Harness secret containing the api key.
- **application_key_ref** (String) Reference to the Harness secret containing the application key.
- **url** (String) Url of the Datadog server.

Optional:

- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.


<a id="nestedblock--docker_registry"></a>
### Nested Schema for `docker_registry`

Required:

- **type** (String) The type of the docker registry. Valid options are DockerHub, Harbor, Other, Quay
- **url** (String) The url of the docker registry.

Optional:

- **credentials** (Block List, Max: 1) The credentials to use for the docker registry. If not specified then the connection is made to the registry anonymously. (see [below for nested schema](#nestedblock--docker_registry--credentials))
- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.

<a id="nestedblock--docker_registry--credentials"></a>
### Nested Schema for `docker_registry.credentials`

Required:

- **password_ref** (String) The reference to the password to use for the docker registry.

Optional:

- **username** (String) The username to use for the docker registry.
- **username_ref** (String) The reference to the username to use for the docker registry.



<a id="nestedblock--dynatrace"></a>
### Nested Schema for `dynatrace`

Required:

- **api_token_ref** (String) The reference to the Harness secret containing the api token.
- **url** (String) Url of the Dynatrace server.

Optional:

- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.


<a id="nestedblock--gcp"></a>
### Nested Schema for `gcp`

Optional:

- **inherit_from_delegate** (Block List) Inherit configuration from delegate. (see [below for nested schema](#nestedblock--gcp--inherit_from_delegate))
- **manual** (Block List, Max: 1) Manual credential configuration. (see [below for nested schema](#nestedblock--gcp--manual))

<a id="nestedblock--gcp--inherit_from_delegate"></a>
### Nested Schema for `gcp.inherit_from_delegate`

Required:

- **delegate_selectors** (Set of String) The delegates to inherit the credentials from.


<a id="nestedblock--gcp--manual"></a>
### Nested Schema for `gcp.manual`

Required:

- **delegate_selectors** (Set of String) The delegates to connect with.
- **secret_key_ref** (String) Reference to the Harness secret containing the secret key.



<a id="nestedblock--git"></a>
### Nested Schema for `git`

Required:

- **connection_type** (String) Whether the connection we're making is to a git repository or a git account. Valid values are Account, Repo.
- **credentials** (Block List, Min: 1, Max: 1) Credentials to use for the connection. (see [below for nested schema](#nestedblock--git--credentials))
- **url** (String) Url of the git repository or account.

Optional:

- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.
- **validation_repo** (String) Repository to test the connection with. This is only used when `connection_type` is `Account`.

<a id="nestedblock--git--credentials"></a>
### Nested Schema for `git.credentials`

Optional:

- **http** (Block List, Max: 1) Authenticate using Username and password over http(s) for the connection. (see [below for nested schema](#nestedblock--git--credentials--http))
- **ssh** (Block List, Max: 1) Authenticate using SSH for the connection. (see [below for nested schema](#nestedblock--git--credentials--ssh))

<a id="nestedblock--git--credentials--http"></a>
### Nested Schema for `git.credentials.http`

Required:

- **password_ref** (String) Reference to a secret containing the password to use for authentication.

Optional:

- **username** (String) Username to use for authentication.
- **username_ref** (String) Reference to a secret containing the username to use for authentication.


<a id="nestedblock--git--credentials--ssh"></a>
### Nested Schema for `git.credentials.ssh`

Required:

- **ssh_key_ref** (String) Reference to the Harness secret containing the ssh key.




<a id="nestedblock--github"></a>
### Nested Schema for `github`

Required:

- **connection_type** (String) Whether the connection we're making is to a github repository or a github account. Valid values are Account, Repo.
- **credentials** (Block List, Min: 1, Max: 1) Credentials to use for the connection. (see [below for nested schema](#nestedblock--github--credentials))
- **url** (String) Url of the github repository or account.

Optional:

- **api_authentication** (Block List, Max: 1) Configuration for using the github api. API Access is required for using “Git Experience”, for creation of Git based triggers, Webhooks management and updating Git statuses. (see [below for nested schema](#nestedblock--github--api_authentication))
- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.
- **validation_repo** (String) Repository to test the connection with. This is only used when `connection_type` is `Account`.

<a id="nestedblock--github--credentials"></a>
### Nested Schema for `github.credentials`

Optional:

- **http** (Block List, Max: 1) Authenticate using Username and password over http(s) for the connection. (see [below for nested schema](#nestedblock--github--credentials--http))
- **ssh** (Block List, Max: 1) Authenticate using SSH for the connection. (see [below for nested schema](#nestedblock--github--credentials--ssh))

<a id="nestedblock--github--credentials--http"></a>
### Nested Schema for `github.credentials.http`

Required:

- **token_ref** (String) Reference to a secret containing the personal access to use for authentication.

Optional:

- **username** (String) Username to use for authentication.
- **username_ref** (String) Reference to a secret containing the username to use for authentication.


<a id="nestedblock--github--credentials--ssh"></a>
### Nested Schema for `github.credentials.ssh`

Required:

- **ssh_key_ref** (String) Reference to the Harness secret containing the ssh key.



<a id="nestedblock--github--api_authentication"></a>
### Nested Schema for `github.api_authentication`

Optional:

- **github_app** (Block List, Max: 1) Configuration for using the github app for interacting with the github api. (see [below for nested schema](#nestedblock--github--api_authentication--github_app))
- **token_ref** (String) Personal access token for interacting with the github api.

<a id="nestedblock--github--api_authentication--github_app"></a>
### Nested Schema for `github.api_authentication.github_app`

Required:

- **application_id** (String) Enter the GitHub App ID from the GitHub App General tab.
- **installation_id** (String) Enter the Installation ID located in the URL of the installed GitHub App.
- **private_key_ref** (String) Reference to the secret containing the private key.




<a id="nestedblock--gitlab"></a>
### Nested Schema for `gitlab`

Required:

- **connection_type** (String) Whether the connection we're making is to a gitlab repository or a gitlab account. Valid values are Account, Repo.
- **credentials** (Block List, Min: 1, Max: 1) Credentials to use for the connection. (see [below for nested schema](#nestedblock--gitlab--credentials))
- **url** (String) Url of the gitlab repository or account.

Optional:

- **api_authentication** (Block List, Max: 1) Configuration for using the gitlab api. API Access is required for using “Git Experience”, for creation of Git based triggers, Webhooks management and updating Git statuses. (see [below for nested schema](#nestedblock--gitlab--api_authentication))
- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.
- **validation_repo** (String) Repository to test the connection with. This is only used when `connection_type` is `Account`.

<a id="nestedblock--gitlab--credentials"></a>
### Nested Schema for `gitlab.credentials`

Optional:

- **http** (Block List, Max: 1) Authenticate using Username and password over http(s) for the connection. (see [below for nested schema](#nestedblock--gitlab--credentials--http))
- **ssh** (Block List, Max: 1) Authenticate using SSH for the connection. (see [below for nested schema](#nestedblock--gitlab--credentials--ssh))

<a id="nestedblock--gitlab--credentials--http"></a>
### Nested Schema for `gitlab.credentials.http`

Optional:

- **password_ref** (String) Reference to a secret containing the password to use for authentication.
- **token_ref** (String) Reference to a secret containing the personal access to use for authentication.
- **username** (String) Username to use for authentication.
- **username_ref** (String) Reference to a secret containing the username to use for authentication.


<a id="nestedblock--gitlab--credentials--ssh"></a>
### Nested Schema for `gitlab.credentials.ssh`

Required:

- **ssh_key_ref** (String) Reference to the Harness secret containing the ssh key.



<a id="nestedblock--gitlab--api_authentication"></a>
### Nested Schema for `gitlab.api_authentication`

Required:

- **token_ref** (String) Personal access token for interacting with the gitlab api.



<a id="nestedblock--http_helm"></a>
### Nested Schema for `http_helm`

Required:

- **url** (String) URL of the helm server.

Optional:

- **credentials** (Block List, Max: 1) Credentials to use for authentication. (see [below for nested schema](#nestedblock--http_helm--credentials))
- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.

<a id="nestedblock--http_helm--credentials"></a>
### Nested Schema for `http_helm.credentials`

Required:

- **password_ref** (String) Reference to a secret containing the password to use for authentication.

Optional:

- **username** (String) Username to use for authentication.
- **username_ref** (String) Reference to a secret containing the username to use for authentication.



<a id="nestedblock--jira"></a>
### Nested Schema for `jira`

Required:

- **password_ref** (String) Reference to a secret containing the password to use for authentication.
- **url** (String) Url of the Jira server.

Optional:

- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.
- **username** (String) Username to use for authentication.
- **username_ref** (String) Reference to a secret containing the username to use for authentication.


<a id="nestedblock--k8s_cluster"></a>
### Nested Schema for `k8s_cluster`

Optional:

- **client_key_cert** (Block List, Max: 1) Client key and certificate config for the connector. (see [below for nested schema](#nestedblock--k8s_cluster--client_key_cert))
- **inherit_from_delegate** (Block List, Max: 1) Credentials are inherited from the delegate. (see [below for nested schema](#nestedblock--k8s_cluster--inherit_from_delegate))
- **openid_connect** (Block List, Max: 1) OpenID configuration for the connector. (see [below for nested schema](#nestedblock--k8s_cluster--openid_connect))
- **service_account** (Block List, Max: 1) Service account for the connector. (see [below for nested schema](#nestedblock--k8s_cluster--service_account))
- **username_password** (Block List, Max: 1) Username and password for the connector. (see [below for nested schema](#nestedblock--k8s_cluster--username_password))

<a id="nestedblock--k8s_cluster--client_key_cert"></a>
### Nested Schema for `k8s_cluster.client_key_cert`

Required:

- **client_cert_ref** (String) Reference to the secret containing the client certificate for the connector.
- **client_key_algorithm** (String) The algorithm used to generate the client key for the connector. Valid values are RSA, EC
- **client_key_ref** (String) Reference to the secret containing the client key for the connector.
- **master_url** (String) The URL of the Kubernetes cluster.

Optional:

- **ca_cert_ref** (String) Reference to the secret containing the CA certificate for the connector.
- **client_key_passphrase_ref** (String) Reference to the secret containing the client key passphrase for the connector.


<a id="nestedblock--k8s_cluster--inherit_from_delegate"></a>
### Nested Schema for `k8s_cluster.inherit_from_delegate`

Required:

- **delegate_selectors** (Set of String) Selectors to use for the delegate.


<a id="nestedblock--k8s_cluster--openid_connect"></a>
### Nested Schema for `k8s_cluster.openid_connect`

Required:

- **client_id_ref** (String) Reference to the secret containing the client ID for the connector.
- **issuer_url** (String) The URL of the OpenID Connect issuer.
- **master_url** (String) The URL of the Kubernetes cluster.
- **password_ref** (String) Reference to the secret containing the password for the connector.

Optional:

- **scopes** (List of String) Scopes to request for the connector.
- **secret_ref** (String) Reference to the secret containing the client secret for the connector.
- **username** (String) Username for the connector.
- **username_ref** (String) Reference to the secret containing the username for the connector.


<a id="nestedblock--k8s_cluster--service_account"></a>
### Nested Schema for `k8s_cluster.service_account`

Required:

- **master_url** (String) The URL of the Kubernetes cluster.
- **service_account_token_ref** (String) Reference to the secret containing the service account token for the connector.


<a id="nestedblock--k8s_cluster--username_password"></a>
### Nested Schema for `k8s_cluster.username_password`

Required:

- **master_url** (String) The URL of the Kubernetes cluster.
- **password_ref** (String) Reference to the secret containing the password for the connector.

Optional:

- **username** (String) Username for the connector.
- **username_ref** (String) Reference to the secret containing the username for the connector.



<a id="nestedblock--newrelic"></a>
### Nested Schema for `newrelic`

Required:

- **account_id** (String) Account ID of the NewRelic account.
- **api_key_ref** (String) Reference to the Harness secret containing the api key.
- **url** (String) Url of the NewRelic server.

Optional:

- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.


<a id="nestedblock--nexus"></a>
### Nested Schema for `nexus`

Required:

- **url** (String) URL of the Nexus server.
- **version** (String) Version of the Nexus server. Valid values are 2.x, 3.x

Optional:

- **credentials** (Block List, Max: 1) Credentials to use for authentication. (see [below for nested schema](#nestedblock--nexus--credentials))
- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.

<a id="nestedblock--nexus--credentials"></a>
### Nested Schema for `nexus.credentials`

Required:

- **password_ref** (String) Reference to a secret containing the password to use for authentication.

Optional:

- **username** (String) Username to use for authentication.
- **username_ref** (String) Reference to a secret containing the username to use for authentication.



<a id="nestedblock--pagerduty"></a>
### Nested Schema for `pagerduty`

Required:

- **api_token_ref** (String) Reference to the Harness secret containing the api token.

Optional:

- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.


<a id="nestedblock--prometheus"></a>
### Nested Schema for `prometheus`

Required:

- **url** (String) Url of the Prometheus server.

Optional:

- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.


<a id="nestedblock--splunk"></a>
### Nested Schema for `splunk`

Required:

- **account_id** (String) Splunk account id.
- **password_ref** (String) The reference to the Harness secret containing the Splunk password.
- **url** (String) Url of the Splunk server.
- **username** (String) The username used for connecting to Splunk.

Optional:

- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.


<a id="nestedblock--sumologic"></a>
### Nested Schema for `sumologic`

Required:

- **access_id_ref** (String) Reference to the Harness secret containing the access id.
- **access_key_ref** (String) Reference to the Harness secret containing the access key.
- **url** (String) Url of the SumoLogic server.

Optional:

- **delegate_selectors** (Set of String) Connect using only the delegates which have these tags.

## Import

Import is supported using the following syntax:

```shell
# Import an account level connector
terraform import harness_connector.example <connector_id>

# Import an organization level connector
terraform import harness_connector.example <organization_id>/<connector_id>

# Import an project level connector
terraform import harness_connector.example <organization_id>/<project_id>/<connector_id>
```