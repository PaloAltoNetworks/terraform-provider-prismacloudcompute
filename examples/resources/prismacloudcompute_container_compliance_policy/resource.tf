resource "prismacloudcompute_container_compliance_policy" "ruleset" {
  rule {
    name        = "Default - ignore Twistlock components"
    effect      = "alert"
    collections = ["All"]
    compliance_check {
      block = false
      id    = 56
    }
    compliance_check {
      block = false
      id    = 57
    }
  }
  rule {
    name        = "Default - alert on critical and high"
    effect      = "alert"
    collections = ["All"]
    compliance_check {
      block = false
      id    = 41
    }
    compliance_check {
      block = false
      id    = 51
    }
  }
}
