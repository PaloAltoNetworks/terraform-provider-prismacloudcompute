resource "prismacloudcompute_group" "mygroup" {
  name  = "my group"
  users = ["george"]
}
