package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/auth"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/collection"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/policy"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/settings"
)

func main() {
	credsFile, err := os.Open("creds.json")
	if err != nil {
		fmt.Printf("error opening creds file: %v", err)
	}
	defer credsFile.Close()

	fileContent, err := io.ReadAll(credsFile)
	if err != nil {
		fmt.Printf("error reading creds file: %v", err)
		return
	}
	var config api.APIClientConfig
	if err := json.Unmarshal(fileContent, &config); err != nil {
		fmt.Printf("error unmarshalling creds file: %v", err)
		return
	}

	client, err := api.APIClient(config)
	if err != nil {
		fmt.Printf("failed creating API client: %v", err)
		return
	}

	/*
		COLLECTIONS
	*/
	fmt.Printf("create collection\n")
	newColl := collection.Collection{Name: "My Collection"}
	err = collection.CreateCollection(*client, newColl)
	if err != nil {
		fmt.Printf("Failed to create collection: %s\n", err)
		return
	}

	fmt.Printf("\nlist collections:\n")
	colls, err := collection.ListCollections(*client)
	if err != nil {
		fmt.Printf("Failed to list collections: %s\n", err)
	}

	for _, v := range colls {
		fmt.Printf("* %s %s\n", v.Name, v.Color)
	}

	fmt.Printf("\nget collection:\n")
	coll, err := collection.GetCollection(*client, "My Collection")
	if err != nil {
		fmt.Printf("Failed to get collection: %s\n", err)
	}
	fmt.Printf("* %s %s\n", coll.Name, coll.Color)

	fmt.Printf("\nupdate collection\n")
	existingColl := collection.Collection{Name: "My Collection", Color: "#FFFFFF"}
	err = collection.UpdateCollection(*client, existingColl)
	if err != nil {
		fmt.Printf("Failed to update collection: %s\n", err)
	}

	fmt.Printf("\nlist collections:\n")
	colls, err = collection.ListCollections(*client)
	if err != nil {
		fmt.Printf("Failed to list collections: %s\n", err)
	}

	for _, v := range colls {
		fmt.Printf("* %s %s\n", v.Name, v.Color)
	}

	fmt.Printf("\ndelete collection\n")
	err = collection.DeleteCollection(*client, "My Collection")
	if err != nil {
		fmt.Printf("Failed to delete collection: %s\n", err)
	}

	fmt.Printf("\nlist collections:\n")
	colls, err = collection.ListCollections(*client)
	if err != nil {
		fmt.Printf("Failed to list collections: %s\n", err)
	}

	for _, v := range colls {
		fmt.Printf("* %s %s\n", v.Name, v.Color)
	}

	/*
		CI IMAGE COMPLIANCE
	*/
	complianceCiImageColl := collection.Collection{Name: "All"}
	complianceCiImageCheck1 := policy.ComplianceCheck{Id: 41, Block: false}
	complianceCiImageCheck2 := policy.ComplianceCheck{Id: 422, Block: true}
	complianceCiImageConditions := policy.ComplianceConditions{Checks: []policy.ComplianceCheck{complianceCiImageCheck1, complianceCiImageCheck2}}
	complianceCiImageRule := policy.ComplianceRule{Name: "example ci image compliance rule", Effect: "alert, block", Collections: []collection.Collection{complianceCiImageColl}, Conditions: complianceCiImageConditions}
	complianceCiImageRules := []policy.ComplianceRule{complianceCiImageRule}
	complianceCiImagePolicy := policy.CompliancePolicy{Type: "ciImagesCompliance", Rules: complianceCiImageRules}

	fmt.Printf("\nupdate CI image compliance policy\n")
	complianceCiImageErr := policy.UpdateComplianceCiImage(*client, complianceCiImagePolicy)
	if complianceCiImageErr != nil {
		fmt.Printf("\nFailed to update CI image compliance policy: %v\n", complianceCiImageErr)
	}

	fmt.Printf("\nget CI image compliance policy:\n")
	retrievedCompliancePolicy, complianceCiImageErr := policy.GetComplianceCiImage(*client)
	if complianceCiImageErr != nil {
		fmt.Printf("failed to get CI image compliance policy: %s\n", complianceCiImageErr)
	}
	fmt.Printf("* ci image compliance: %v\n", retrievedCompliancePolicy.Rules)

	fmt.Printf("\nupdate CI image compliance policy\n")
	complianceCiImageRule.Name = "name change"
	complianceCiImagePolicy = policy.CompliancePolicy{Type: "ciImagesCompliance", Rules: []policy.ComplianceRule{complianceCiImageRule}}
	complianceCiImageErr = policy.UpdateComplianceCiImage(*client, complianceCiImagePolicy)
	if complianceCiImageErr != nil {
		fmt.Printf("failed to update CI image compliance policy: %s\n", complianceCiImageErr)
	}

	fmt.Printf("\nget CI image compliance policy:\n")
	retrievedCompliancePolicy, complianceCiImageErr = policy.GetComplianceCiImage(*client)
	if complianceCiImageErr != nil {
		fmt.Printf("failed to get CI image compliance policy: %s\n", complianceCiImageErr)
	}
	fmt.Printf("* ci image compliance: %v\n", retrievedCompliancePolicy.Rules)

	/*
		CONTAINER COMPLIANCE
	*/
	complianceContainerColl := collection.Collection{Name: "All"}
	complianceContainerCheck1 := policy.ComplianceCheck{Id: 41, Block: false}
	complianceContainerCheck2 := policy.ComplianceCheck{Id: 422, Block: true}
	complianceContainerConditions := policy.ComplianceConditions{Checks: []policy.ComplianceCheck{complianceContainerCheck1, complianceContainerCheck2}}
	complianceContainerRule := policy.ComplianceRule{Name: "example container compliance rule", Effect: "alert, block", Collections: []collection.Collection{complianceContainerColl}, Conditions: complianceContainerConditions}
	complianceContainerRules := []policy.ComplianceRule{complianceContainerRule}
	complianceContainerPolicy := policy.CompliancePolicy{Type: "containerCompliance", Rules: complianceContainerRules}

	fmt.Printf("\nupdate container compliance policy\n")
	complianceContainerErr := policy.UpdateComplianceContainer(*client, complianceContainerPolicy)
	if complianceContainerErr != nil {
		fmt.Printf("\nFailed to update container compliance policy: %v\n", complianceContainerErr)
	}

	fmt.Printf("\nget container compliance policy:\n")
	retrievedCompliancePolicy, complianceContainerErr = policy.GetComplianceContainer(*client)
	if complianceContainerErr != nil {
		fmt.Printf("failed to get container compliance policy: %s\n", complianceContainerErr)
	}
	fmt.Printf("* container compliance: %v\n", retrievedCompliancePolicy.Rules)

	fmt.Printf("\nupdate container compliance policy\n")
	complianceContainerRule.Name = "name change"
	complianceContainerPolicy = policy.CompliancePolicy{Type: "containerCompliance", Rules: []policy.ComplianceRule{complianceContainerRule}}
	complianceContainerErr = policy.UpdateComplianceContainer(*client, complianceContainerPolicy)
	if complianceContainerErr != nil {
		fmt.Printf("failed to update container compliance policy: %s\n", complianceContainerErr)
	}

	fmt.Printf("\nget container compliance policy:\n")
	retrievedCompliancePolicy, complianceContainerErr = policy.GetComplianceContainer(*client)
	if complianceContainerErr != nil {
		fmt.Printf("failed to get container compliance policy: %s\n", complianceContainerErr)
	}
	fmt.Printf("* container compliance: %v\n", retrievedCompliancePolicy.Rules)

	/*
		HOST COMPLIANCE
	*/
	complianceHostColl := collection.Collection{Name: "All"}
	complianceHostCheck1 := policy.ComplianceCheck{Id: 11, Block: false}
	complianceHostCheck2 := policy.ComplianceCheck{Id: 112, Block: true}
	complianceHostConditions := policy.ComplianceConditions{Checks: []policy.ComplianceCheck{complianceHostCheck1, complianceHostCheck2}}
	complianceHostRule := policy.ComplianceRule{Name: "example host compliance rule", Effect: "alert, block", Collections: []collection.Collection{complianceHostColl}, Conditions: complianceHostConditions}
	complianceHostRules := []policy.ComplianceRule{complianceHostRule}
	complianceHostPolicy := policy.CompliancePolicy{Type: "hostCompliance", Rules: complianceHostRules}

	fmt.Printf("\nupdate host compliance policy\n")
	complianceHostErr := policy.UpdateComplianceHost(*client, complianceHostPolicy)
	if complianceHostErr != nil {
		fmt.Printf("\nFailed to update host compliance policy: %v\n", complianceHostErr)
	}

	fmt.Printf("\nget host compliance policy:\n")
	retrievedCompliancePolicy, complianceHostErr = policy.GetComplianceHost(*client)
	if complianceHostErr != nil {
		fmt.Printf("failed to get host compliance policy: %s\n", complianceHostErr)
	}
	fmt.Printf("* host compliance: %v\n", retrievedCompliancePolicy.Rules)

	fmt.Printf("\nupdate host compliance policy\n")
	complianceHostRule.Name = "name change"
	complianceHostPolicy = policy.CompliancePolicy{Type: "hostCompliance", Rules: []policy.ComplianceRule{complianceHostRule}}
	complianceHostErr = policy.UpdateComplianceHost(*client, complianceHostPolicy)
	if complianceHostErr != nil {
		fmt.Printf("failed to update host compliance policy: %s\n", complianceHostErr)
	}

	fmt.Printf("\nget host compliance policy:\n")
	retrievedCompliancePolicy, complianceHostErr = policy.GetComplianceHost(*client)
	if complianceHostErr != nil {
		fmt.Printf("failed to get host compliance policy: %s\n", complianceHostErr)
	}
	fmt.Printf("* host compliance: %v\n", retrievedCompliancePolicy.Rules)

	/*
		CONTAINER RUNTIME
	*/
	runtimeContainerColl := collection.Collection{Name: "All"}
	runtimeContainerRule := policy.RuntimeContainerRule{
		Name:               "example container runtime rule",
		Collections:        []collection.Collection{runtimeContainerColl},
		AdvancedProtection: true,
		Dns:                policy.RuntimeContainerDns{DenyEffect: "disable"},
		Filesystem:         policy.RuntimeContainerFilesystem{DenyEffect: "alert"},
		Network:            policy.RuntimeContainerNetwork{DenyEffect: "alert"},
		Processes:          policy.RuntimeContainerProcesses{DenyEffect: "alert"},
		WildFireAnalysis:   "alert",
	}
	runtimeContainerRules := []policy.RuntimeContainerRule{runtimeContainerRule}
	runtimeContainerPolicy := policy.RuntimeContainerPolicy{LearningDisabled: false, Rules: runtimeContainerRules}

	fmt.Printf("\nupdate container runtime policy\n")
	runtimeContainerErr := policy.UpdateRuntimeContainer(*client, runtimeContainerPolicy)
	if runtimeContainerErr != nil {
		fmt.Printf("\nfailed to update container runtime policy: %v\n", runtimeContainerErr)
	}

	fmt.Printf("\nget container runtime policy:\n")
	retrievedRuntimeContainerPolicy, runtimeContainerErr := policy.GetRuntimeContainer(*client)
	if runtimeContainerErr != nil {
		fmt.Printf("failed to get container runtime policy: %s\n", runtimeContainerErr)
	}
	fmt.Printf("* container runtime: %v\n", retrievedRuntimeContainerPolicy.Rules)

	fmt.Printf("\nupdate container runtime policy\n")
	runtimeContainerRule.Name = "name change"
	runtimeContainerPolicy = policy.RuntimeContainerPolicy{LearningDisabled: false, Rules: []policy.RuntimeContainerRule{runtimeContainerRule}}
	runtimeContainerErr = policy.UpdateRuntimeContainer(*client, runtimeContainerPolicy)
	if runtimeContainerErr != nil {
		fmt.Printf("failed to update container runtime policy: %s\n", runtimeContainerErr)
	}

	fmt.Printf("\nget container runtime policy:\n")
	retrievedRuntimeContainerPolicy, runtimeContainerErr = policy.GetRuntimeContainer(*client)
	if runtimeContainerErr != nil {
		fmt.Printf("failed to get container runtime policy: %s\n", runtimeContainerErr)
	}
	fmt.Printf("* container runtime: %v\n", retrievedRuntimeContainerPolicy.Rules)

	/*
		HOST RUNTIME
	*/
	runtimeHostColl := collection.Collection{Name: "All"}
	runtimeHostForensic := policy.RuntimeHostForensic{
		ActivitiesDisabled:       false,
		SshdEnabled:              false,
		SudoEnabled:              false,
		ServiceActivitiesEnabled: false,
		DockerEnabled:            false,
		ReadonlyDockerEnabled:    false,
	}
	runtimeHostNetwork := policy.RuntimeHostNetwork{DenyEffect: "alert", CustomFeed: "alert", IntelligenceFeed: "alert"}
	runtimeHostDNS := policy.RuntimeHostDns{DenyEffect: "disable", IntelligenceFeed: "disable"}

	runtimeHostDeniedProcesses := policy.RuntimeHostDeniedProcesses{Effect: "alert"}
	runtimeHostAntiMalware := policy.RuntimeHostAntiMalware{
		DeniedProcesses:            runtimeHostDeniedProcesses,
		CryptoMiner:                "alert",
		ServiceUnknownOriginBinary: "alert",
		UserUnknownOriginBinary:    "alert",
		EncryptedBinaries:          "alert",
		SuspiciousElfHeaders:       "alert",
		TempFsProcesses:            "alert",
		ReverseShell:               "alert",
		WebShell:                   "alert",
		ExecutionFlowHijack:        "alert",
		CustomFeed:                 "alert",
		IntelligenceFeed:           "alert",
		WildFireAnalysis:           "alert",
	}

	runtimeHostRule := policy.RuntimeHostRule{Name: "example host runtime rule", Collections: []collection.Collection{runtimeHostColl}, Forensic: runtimeHostForensic, Network: runtimeHostNetwork, Dns: runtimeHostDNS, AntiMalware: runtimeHostAntiMalware}
	runtimeHostRules := []policy.RuntimeHostRule{runtimeHostRule}
	runtimeHostPolicy := policy.RuntimeHostPolicy{Rules: runtimeHostRules}

	fmt.Printf("\nupdate host runtime policy\n")
	runtimeHostErr := policy.UpdateRuntimeHost(*client, runtimeHostPolicy)
	if runtimeHostErr != nil {
		fmt.Printf("\nfailed to update host runtime policy: %v\n", runtimeHostErr)
	}

	fmt.Printf("\nget host runtime policy:\n")
	retrievedRuntimeHostPolicy, runtimeHostErr := policy.GetRuntimeHost(*client)
	if runtimeHostErr != nil {
		fmt.Printf("failed to get host runtime policy: %s\n", runtimeHostErr)
	}
	fmt.Printf("* host runtime: %v\n", retrievedRuntimeHostPolicy.Rules)

	fmt.Printf("\nupdate host runtime policy\n")
	runtimeHostRule.Name = "name change"
	runtimeHostPolicy = policy.RuntimeHostPolicy{Rules: []policy.RuntimeHostRule{runtimeHostRule}}
	runtimeHostErr = policy.UpdateRuntimeHost(*client, runtimeHostPolicy)
	if runtimeHostErr != nil {
		fmt.Printf("failed to update host runtime policy: %s\n", runtimeHostErr)
	}

	fmt.Printf("\nget host runtime policy:\n")
	retrievedRuntimeHostPolicy, runtimeHostErr = policy.GetRuntimeHost(*client)
	if runtimeHostErr != nil {
		fmt.Printf("failed to get host runtime policy: %s\n", runtimeHostErr)
	}
	fmt.Printf("* host runtime: %v\n", retrievedRuntimeHostPolicy.Rules)

	/*
		CI IMAGE VULNERABILITY
	*/
	vulnerabilityCiImageColl := collection.Collection{Name: "All"}
	vulnerabilityCiImageRule := policy.VulnerabilityImageRule{Name: "example CI image vulnerability rule", Collections: []collection.Collection{vulnerabilityCiImageColl}, Effect: "alert"}
	vulnerabilityCiImagePolicy := policy.VulnerabilityImagePolicy{Type: "ciImagesVulnerability", Rules: []policy.VulnerabilityImageRule{vulnerabilityCiImageRule}}

	fmt.Printf("\nupdate CI image vulnerability policy\n")
	vulnerabilityCiImageErr := policy.UpdateVulnerabilityCiImage(*client, vulnerabilityCiImagePolicy)
	if vulnerabilityCiImageErr != nil {
		fmt.Printf("failed to update CI image vulnerability policy: %s\n", vulnerabilityCiImageErr)
	}

	fmt.Printf("\nget CI image vulnerability policy:\n")
	retrievedVulnerabilityImagePolicy, vulnerabilityCiImageErr := policy.GetVulnerabilityCiImage(*client)
	if vulnerabilityCiImageErr != nil {
		fmt.Printf("failed to get CI image vulnerability policy: %s\n", vulnerabilityCiImageErr)
	}
	fmt.Printf("* ci image vulnerability: %v\n", retrievedVulnerabilityImagePolicy.Rules)

	fmt.Printf("\nupdate CI image vulnerability policy\n")
	vulnerabilityCiImageRule.Name = "name change"
	vulnerabilityCiImagePolicy = policy.VulnerabilityImagePolicy{Type: "ciImagesVulnerability", Rules: []policy.VulnerabilityImageRule{vulnerabilityCiImageRule}}
	vulnerabilityCiImageErr = policy.UpdateVulnerabilityCiImage(*client, vulnerabilityCiImagePolicy)
	if vulnerabilityCiImageErr != nil {
		fmt.Printf("failed to update CI image vulnerability policy: %s\n", vulnerabilityCiImageErr)
	}

	fmt.Printf("\nget CI image vulnerability policy:\n")
	retrievedVulnerabilityImagePolicy, vulnerabilityCiImageErr = policy.GetVulnerabilityCiImage(*client)
	if vulnerabilityCiImageErr != nil {
		fmt.Printf("failed to get CI image vulnerability policy: %s\n", vulnerabilityCiImageErr)
	}
	fmt.Printf("* ci image vulnerability: %v\n", retrievedVulnerabilityImagePolicy.Rules)

	/*
		HOST VULNERABILITY
	*/
	vulnerabilityHostColl := collection.Collection{Name: "All"}
	vulnerabilityHostRule := policy.VulnerabilityHostRule{Name: "example host vulnerability rule", Collections: []collection.Collection{vulnerabilityHostColl}, Effect: "alert"}
	vulnerabilityHostPolicy := policy.VulnerabilityHostPolicy{Type: "hostVulnerability", Rules: []policy.VulnerabilityHostRule{vulnerabilityHostRule}}

	fmt.Printf("\nupdate host vulnerability policy\n")
	vulnerabilityHostErr := policy.UpdateVulnerabilityHost(*client, vulnerabilityHostPolicy)
	if vulnerabilityHostErr != nil {
		fmt.Printf("failed to update host vulnerability policy: %s\n", vulnerabilityHostErr)
	}

	fmt.Printf("\nget host vulnerability policy:\n")
	retrievedVulnerabilityHostPolicy, vulnerabilityHostErr := policy.GetVulnerabilityHost(*client)
	if vulnerabilityHostErr != nil {
		fmt.Printf("failed to get host vulnerability policy: %s\n", vulnerabilityHostErr)
	}
	fmt.Printf("* host vulnerability: %v\n", retrievedVulnerabilityHostPolicy.Rules)

	fmt.Printf("\nupdate host vulnerability policy\n")
	vulnerabilityHostRule.Name = "name change"
	vulnerabilityHostPolicy = policy.VulnerabilityHostPolicy{Type: "hostVulnerability", Rules: []policy.VulnerabilityHostRule{vulnerabilityHostRule}}
	vulnerabilityHostErr = policy.UpdateVulnerabilityHost(*client, vulnerabilityHostPolicy)
	if vulnerabilityHostErr != nil {
		fmt.Printf("failed to update host vulnerability policy: %s\n", vulnerabilityHostErr)
	}

	fmt.Printf("\nget host vulnerability policy:\n")
	retrievedVulnerabilityHostPolicy, vulnerabilityHostErr = policy.GetVulnerabilityHost(*client)
	if vulnerabilityHostErr != nil {
		fmt.Printf("failed to get host vulnerability policy: %s\n", vulnerabilityHostErr)
	}
	fmt.Printf("* host vulnerability: %v\n", retrievedVulnerabilityHostPolicy.Rules)

	/*
		IMAGE VULNERABILITY
	*/
	vulnerabilityimageColl := collection.Collection{Name: "All"}
	vulnerabilityimageRule := policy.VulnerabilityImageRule{Name: "example image vulnerability rule", Collections: []collection.Collection{vulnerabilityimageColl}, Effect: "alert"}
	vulnerabilityimagePolicy := policy.VulnerabilityImagePolicy{Type: "containerVulnerability", Rules: []policy.VulnerabilityImageRule{vulnerabilityimageRule}}
	fmt.Printf("\nupdate image vulnerability policy\n")
	vulnerabilityimageErr := policy.UpdateVulnerabilityImage(*client, vulnerabilityimagePolicy)
	if vulnerabilityimageErr != nil {
		fmt.Printf("failed to update image vulnerability policy: %s\n", vulnerabilityimageErr)
	}

	fmt.Printf("\nget image vulnerability policy:\n")
	retrievedVulnerabilityImagePolicy, vulnerabilityimageErr = policy.GetVulnerabilityImage(*client)
	if vulnerabilityimageErr != nil {
		fmt.Printf("failed to get image vulnerability policy: %s\n", vulnerabilityimageErr)
	}
	fmt.Printf("* image vulnerability: %v\n", retrievedVulnerabilityImagePolicy.Rules)

	fmt.Printf("\nupdate image vulnerability policy\n")
	vulnerabilityimageRule.Name = "name change"
	vulnerabilityimagePolicy = policy.VulnerabilityImagePolicy{Type: "containerVulnerability", Rules: []policy.VulnerabilityImageRule{vulnerabilityimageRule}}
	vulnerabilityimageErr = policy.UpdateVulnerabilityImage(*client, vulnerabilityimagePolicy)
	if vulnerabilityimageErr != nil {
		fmt.Printf("failed to update image vulnerability policy: %s\n", vulnerabilityimageErr)
	}

	fmt.Printf("\nget image vulnerability policy:\n")
	retrievedVulnerabilityImagePolicy, vulnerabilityimageErr = policy.GetVulnerabilityImage(*client)
	if vulnerabilityimageErr != nil {
		fmt.Printf("failed to get image vulnerability policy: %s\n", vulnerabilityimageErr)
	}
	fmt.Printf("* image vulnerability: %v\n", retrievedVulnerabilityImagePolicy.Rules)

	/*
		REGISTRY SETTINGS
	*/
	registrySpec := settings.RegistrySpecification{
		Version:     "2",
		Registry:    "",
		Os:          "linux",
		Cap:         5,
		Scanners:    2,
		Repository:  "library/ubuntu",
		Tag:         "20.04",
		Collections: []string{"All"},
	}
	reg := settings.RegistrySettings{Specifications: []settings.RegistrySpecification{registrySpec}}

	fmt.Printf("\ncreate registry settings\n")
	registryErr := settings.UpdateRegistrySettings(*client, reg)
	if registryErr != nil {
		fmt.Printf("failed to create registry settings: %s\n", registryErr)
	}

	fmt.Printf("\nget registry settings:\n")
	retrievedRegistry, registryErr := settings.GetRegistrySettings(*client)
	if registryErr != nil {
		fmt.Printf("failed to get registry settings: %s\n", registryErr)
	}
	fmt.Printf("* %v\n", retrievedRegistry)

	fmt.Printf("\nupdate registry settings\n")
	registrySpec.Tag = "21.04"
	reg = settings.RegistrySettings{Specifications: []settings.RegistrySpecification{registrySpec}}
	registryErr = settings.UpdateRegistrySettings(*client, reg)
	if registryErr != nil {
		fmt.Printf("failed to update registry settings: %s\n", registryErr)
	}

	fmt.Printf("\nget registry settings:\n")
	retrievedRegistry, registryErr = settings.GetRegistrySettings(*client)
	if registryErr != nil {
		fmt.Printf("failed to get registry settings: %s\n", registryErr)
	}
	fmt.Printf("* %v\n", retrievedRegistry)

	/*
		USERS
	*/
	/*	user := auth.User{
			Username: "test user",
			Password: "test password",
			AuthType: "basic",
			Role:     "user",
		}

		fmt.Printf("\ncreate user\n")
		userErr := auth.CreateUser(*client, user)
		if userErr != nil {
			fmt.Printf("failed to create user: %s\n", userErr)
		}

		fmt.Printf("\nlist users:\n")
		retrievedUsers, userErr := auth.ListUsers(*client)
		if userErr != nil {
			fmt.Printf("failed to get users: %s\n", userErr)
		}
		fmt.Printf("* %+v\n", retrievedUsers)

		fmt.Printf("\nupdate user\n")
		user.Role = "vulnerabilityManager"
		userErr = auth.UpdateUser(*client, user)
		if userErr != nil {
			fmt.Printf("failed to update user: %s\n", userErr)
		}

		fmt.Printf("\nlist users:\n")
		retrievedUsers, userErr = auth.ListUsers(*client)
		if userErr != nil {
			fmt.Printf("failed to get users: %s\n", userErr)
		}
		fmt.Printf("* %+v\n", retrievedUsers)

		fmt.Printf("\ndelete user\n")
		userErr = auth.DeleteUser(*client, user.Username)
		if userErr != nil {
			fmt.Printf("failed to delete user: %s\n", userErr)
		}

		fmt.Printf("\nlist users:\n")
		retrievedUsers, userErr = auth.ListUsers(*client)
		if userErr != nil {
			fmt.Printf("failed to get users: %s\n", userErr)
		}
		fmt.Printf("* %+v\n", retrievedUsers)
	*/
	/*
		GROUPS
	*/
	/*	group := auth.Group{
			Name: "test group",
			Users: []auth.GroupUser{
				{
					Username: "admin",
				},
			},
		}

		fmt.Printf("\ncreate group\n")
		groupErr := auth.CreateGroup(*client, group)
		if groupErr != nil {
			fmt.Printf("failed to create group: %s\n", groupErr)
		}

		fmt.Printf("\nlist groups:\n")
		retrievedGroups, groupErr := auth.ListGroups(*client)
		if groupErr != nil {
			fmt.Printf("failed to get groups: %s\n", groupErr)
		}
		fmt.Printf("* %+v\n", retrievedGroups)

		fmt.Printf("\nupdate group\n")
		group.Users = make([]auth.GroupUser, 0)
		groupErr = auth.UpdateGroup(*client, group)
		if groupErr != nil {
			fmt.Printf("failed to update group: %s\n", groupErr)
		}

		fmt.Printf("\nlist groups:\n")
		retrievedGroups, groupErr = auth.ListGroups(*client)
		if groupErr != nil {
			fmt.Printf("failed to get groups: %s\n", groupErr)
		}
		fmt.Printf("* %+v\n", retrievedGroups)

		fmt.Printf("\ndelete group\n")
		groupErr = auth.DeleteGroup(*client, group.Name)
		if groupErr != nil {
			fmt.Printf("failed to delete group: %s\n", groupErr)
		}

		fmt.Printf("\nlist groups:\n")
		retrievedGroups, groupErr = auth.ListGroups(*client)
		if groupErr != nil {
			fmt.Printf("failed to get groups: %s\n", groupErr)
		}
		fmt.Printf("* %+v\n", retrievedGroups)
	*/
	/*
		ROLES
	*/
	/*	role := auth.Role{
			Name: "test role",
			Permissions: []auth.RolePermission{
				{
					Name:      "radarsCloud",
					ReadWrite: false,
				},
				{
					Name:      "user",
					ReadWrite: true,
				},
			},
		}

		fmt.Printf("\ncreate role\n")
		roleErr := auth.CreateRole(*client, role)
		if roleErr != nil {
			fmt.Printf("failed to create role: %s\n", roleErr)
		}

		fmt.Printf("\nget roles:\n")
		retrievedRoles, roleErr := auth.GetRole(*client, role.Name)
		if roleErr != nil {
			fmt.Printf("failed to get role: %s\n", roleErr)
		}
		fmt.Printf("* %+v\n", retrievedRoles)

		fmt.Printf("\nupdate role\n")
		role.Permissions = append(role.Permissions, auth.RolePermission{
			Name:      "accessUI",
			ReadWrite: false,
		})
		roleErr = auth.UpdateRole(*client, role)
		if roleErr != nil {
			fmt.Printf("failed to update role: %s\n", roleErr)
		}

		fmt.Printf("\nget roles:\n")
		retrievedRoles, roleErr = auth.GetRole(*client, role.Name)
		if roleErr != nil {
			fmt.Printf("failed to get role: %s\n", roleErr)
		}
		fmt.Printf("* %+v\n", retrievedRoles)

		fmt.Printf("\ndelete role\n")
		roleErr = auth.DeleteRole(*client, role.Name)
		if roleErr != nil {
			fmt.Printf("failed to delete role: %s\n", roleErr)
		}

		fmt.Printf("\nget roles:\n")
		retrievedRoles, roleErr = auth.GetRole(*client, role.Name)
		if roleErr != nil {
			fmt.Printf("failed to get role: %s\n", roleErr)
		}
		fmt.Printf("* %+v\n", retrievedRoles)
	*/

	/*
		CREDENTIALS
	*/
	credential := auth.Credential{
		Secret: auth.Secret{
			Encrypted: "",
			Plain:     "test",
		},
		//	    ServiceAccount: {},
		Type:        "basic",
		Description: "",
		SkipVerify:  false,
		Id:          "test",
		AccountID:   "test",
	}

	fmt.Printf("\ncreate credential\n")
	credentialErr := auth.UpdateCredential(*client, credential)
	if credentialErr != nil {
		fmt.Printf("failed to create credential: %s\n", credentialErr)
	}

	fmt.Printf("\nlist credentials:\n")
	credentials, err := auth.ListCredentials(*client)
	if err != nil {
		fmt.Printf("Failed to list credentials: %s\n", err)
	}
	for _, v := range credentials {
		fmt.Printf("* %s\n", v.Id)
	}

	fmt.Printf("\nupdate credential\n")
	credential.AccountID = "test update"
	credentialErr = auth.UpdateCredential(*client, credential)
	if credentialErr != nil {
		fmt.Printf("failed to update credential: %s\n", credentialErr)
	}

	fmt.Printf("\ndelete credential\n")
	credentialErr = auth.DeleteCredential(*client, credential.Id)
	if credentialErr != nil {
		fmt.Printf("failed to delete credential: %s\n", credentialErr)
	}

	fmt.Printf("\nlist credentials:\n")
	credentials, err = auth.ListCredentials(*client)
	if err != nil {
		fmt.Printf("Failed to list credentials: %s\n", err)
	}
	for _, v := range credentials {
		fmt.Printf("* %s\n", v.Id)
	}
}
