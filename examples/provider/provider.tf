terraform {
  required_providers {
    harness = {
      source = "harness/harness-platform"
    }
  }
}

provider "harness" {
  endpoint   = "https://app.harness.io/gateway"
  account_id = "...."
  api_key    = "......"
}

