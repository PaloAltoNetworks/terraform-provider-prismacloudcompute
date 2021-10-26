resource "prismacloudcompute_collection" "example1" {
  name              = "my collection"
  description       = "collection created with Terraform"
  color             = "#FF0000"
  application_ids   = ["app1"]
  code_repositories = ["coderepo1", "prefix1*"]
  images            = ["prefix2*", "prefix3*"]
  labels            = ["env:development", "env:staging"]
  namespaces        = ["hamilton"]
}
