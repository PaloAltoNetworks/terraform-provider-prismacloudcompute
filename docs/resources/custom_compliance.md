# prismacloudcompute_custom_compliance (Resource)


## Example Usage

```terraform
resource "prismacloudcompute_custom_compliance" "defender" {
  name = "Check defender exist"
  title = "Check if the defender file exists in the filesystem"
  severity = "critical"
  script = <<EOT
if [ ! -f /usr/local/bin/defender ]; then
    echo "Defender not found!"
    exit 1
fi
EOT
}
```

## Schema

### Required

- `name` (String) Free-form text description of the custom Compliance.
- `script` (String) Script of this custom compliance
- `severity` (String) Severity of this custom compliance
- `title` (String) Description of the custom compliance

### Read-Only

- `id` (String) ID of the custom Compliance.
- `prisma_id` (Number) Prisma Cloud Compute ID of the custom rule.


