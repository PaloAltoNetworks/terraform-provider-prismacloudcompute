---
page_title: "Prisma Cloud: prismacloudcompute_collection"
---

# prismacloudcompute_collection

Manage a cloud compute collection.

## Example Usage

```hcl
data "prismacloudcompute_collection" "example1" {
    name = "New Collection"
    color = "#FF0000"
}
```

## Argument Reference

* `name` - (Required) Unique collection name.
* `accountids` - List of account IDs.
* `appids` - List of application IDs.
* `clusters` - List of Kubernetes cluster names.
* `coderepos` - List of code repositories.
* `color` - A hex color code for a collection.
* `containers` - List of containers.
* `description` - A free-form text description of the collection.
* `functions` - List of functions.
* `hosts` - List of hosts.
* `images` - List of images.
* `labels` - List of labels.
* `modified`- Date/time when the collection was last modified.
* `namespaces` - List of Kubernetes namespaces.
* `owner` - User who created or last modified the collection.
* `prisma` - If set to `true`, this collection originates from Prisma Cloud.
* `system` - If set to `true`, this collection was created by the system (i.e., a non-user). Otherwise it was created by a real user.