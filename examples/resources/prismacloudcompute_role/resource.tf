resource "prismacloudcompute_role" "myrole" {
  name = "my role"
  permission {
    name       = "monitorVuln"
    read_write = false
  }
  permission {
    name       = "monitorCompliance"
    read_write = false
  }
  permission {
    name       = "user"
    read_write = true
  }
}
