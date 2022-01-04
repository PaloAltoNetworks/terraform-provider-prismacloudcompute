# Changelog

## Version 0.4.0 - 2022-01-03
#### Added
- Support for grace days policy in vulnerability policies ([#42](https://github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/pull/42), @hi-artem)

#### Fixed
- Updated SDK reference for fix in github.com/paloaltonetworks/prisma-cloud-compute-go@v0.4.2.

## Version 0.3.0 - 2021-12-16
#### Added
- Projects support ([#40](https://github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/pull/40), @wfg).
- Custom runtime rule data source ([#39](https://github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/pull/39), @hi-artem).

#### Fixed
- Typo in host runtime policy parser ([#36](https://github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/pull/36), @hi-artem).

## Version 0.2.0 - 2021-12-13
#### Added
- Admission policy resource type ([#33](https://github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/pull/33), @hi-artem).
- Custom runtime rule resource type ([#35](https://github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/pull/35), @hi-artem).

#### Fixed
- You can now specify config items directly instead of using config file ([#31](https://github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/pull/31), @hi-artem).

## Version 0.1.0 - 2021-11-08
#### Added
- User, group, role, and credential resource types.

#### Changed
- Updated many of the resource arguments to have reasonable names.

#### Removed
- Cut many of the data sources that didn't make much sense.
Also removed the ones that do make sense from the provider until they can be properly developed and tested.
