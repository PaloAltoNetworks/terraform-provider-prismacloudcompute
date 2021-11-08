resource "prismacloudcompute_registry_settings" "registry" {
  specification {
    type        = "2"
    registry    = ""
    os          = "linux"
    cap         = 5
    scanners    = 2
    repository  = "library/ubuntu"
    tag         = "20.04"
    collections = ["All"]
  }
  specification {
    type        = "2"
    registry    = ""
    os          = "linux"
    cap         = 5
    scanners    = 2
    repository  = "library/ubuntu"
    tag         = "21.04"
    collections = ["All"]
  }
}
