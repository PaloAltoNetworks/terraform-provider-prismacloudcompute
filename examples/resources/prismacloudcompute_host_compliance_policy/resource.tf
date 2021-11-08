resource "prismacloudcompute_host_compliance_policy" "ruleset" {
  rule {
    name        = "Default - alert on critical and high"
    effect      = "alert, block"
    collections = ["All"]
    compliance_check {
      block = false
      id    = 16
    }
    compliance_check {
      block = false
      id    = 21
    }
  }
}
