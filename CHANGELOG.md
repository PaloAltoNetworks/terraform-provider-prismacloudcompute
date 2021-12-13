# Changelog

## Unreleased
## Version 0.2.0 - 2021-12-13
### Added
- Admission policy resource type ([#33](https://github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/pull/33), @hi-artem).
- Custom runtime rule resource type ([#35](https://github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/pull/35), @hi-artem).

### Fixed
- You can now specify config items directly instead of using config file ([#31](https://github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/pull/31), @hi-artem).

## Version 0.1.0 - 2021-11-08
### Added
- User, group, role, and credential resource types.

### Changed
- Updated many of the resource arguments to have reasonable names.

### Removed
- Cut many of the data sources that didn't make much sense.
Also removed the ones that do make sense from the provider until they can be properly developed and tested.
