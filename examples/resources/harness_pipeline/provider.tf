terraform {
  required_providers {
    harness = {
      source = "suranc/harness/harness"
    }
  }
}

provider "harness" {
  endpoint   = "http://localhost:7143"
  account_id = "7UEOqlXfQaCH5AxFeDT0_A"
  ng_api_key    = "sat.621b00b1dbd0795304f6458c.hQp1AEk94RDiJOC4F9rv"
  api_key    = "sat.621b00b1dbd0795304f6458c.hQp1AEk94RDiJOC4F9rv"
}

