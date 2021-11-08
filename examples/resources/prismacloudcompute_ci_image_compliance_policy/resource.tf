resource "prismacloudcompute_ci_image_compliance_policy" "ruleset" {
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
      id    = 422
    }
  }
}
