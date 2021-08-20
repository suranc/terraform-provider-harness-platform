---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "harness_user Resource - harness"
subcategory: ""
description: |-
  Resource for creating a Harness user
---

# harness_user (Resource)

Resource for creating a Harness user



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **email** (String) The email of the user.
- **name** (String) The name of the user.

### Read-Only

- **id** (String) Unique identifier of the user.
- **is_email_verified** (Boolean) Flag indicating whether or not the users email has been verified.
- **is_imported_from_identity_provider** (Boolean) Flag indicating whether or not the user was imported from an identity provider.
- **is_password_expired** (Boolean) Flag indicating whether or not the users password has expired.
- **is_two_factor_auth_enabled** (Boolean) Flag indicating whether or not two-factor authentication is enabled for the user.
- **is_user_locked** (Boolean) Flag indicating whether or not the user is locked out.

