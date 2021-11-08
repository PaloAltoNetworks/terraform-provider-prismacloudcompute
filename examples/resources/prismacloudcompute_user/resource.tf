resource "prismacloudcompute_user" "myuser" {
  authentication_type = "basic"
  username            = "george"
  password            = "thepossum"
  role                = "user"
}
