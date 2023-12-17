package project_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mongodb/terraform-provider-mongodbatlas/internal/common/conversion"
	"github.com/mongodb/terraform-provider-mongodbatlas/internal/service/project"
	"go.mongodb.org/atlas-sdk/v20231115002/admin"
)

var (
	roles        = []string{"GROUP_DATA_ACCESS_READ_ONLY", "GROUP_CLUSTER_MANAGER"}
	teamRolesSDK = []admin.TeamRole{
		{
			TeamId:    conversion.StringPtr("teamId"),
			RoleNames: roles,
		},
	}
)

func TestTeamRoleListTFtoSDK(t *testing.T) {
	var rolesSet, _ = types.SetValueFrom(context.Background(), types.StringType, roles)
	teamsTF := []project.TfTeamModel{
		{
			TeamID:    types.StringValue("teamId"),
			RoleNames: rolesSet,
		},
	}
	testCases := []struct {
		name           string
		expectedResult *[]admin.TeamRole
		teamRolesTF    []project.TfTeamModel
	}{
		{
			name:           "Team roles",
			teamRolesTF:    teamsTF,
			expectedResult: &teamRolesSDK,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultModel := project.NewTeamRoleList(context.Background(), tc.teamRolesTF)
			if !reflect.DeepEqual(resultModel, tc.expectedResult) {
				t.Errorf("created terraform model did not match expected output")
			}
		})
	}
}

func TestTeamModelMapTF(t *testing.T) {
	teams := []project.TfTeamModel{
		{
			TeamID: types.StringValue("id1"),
		},
		{
			TeamID: types.StringValue("id2"),
		},
	}
	testCases := []struct {
		name           string
		expectedResult map[types.String]project.TfTeamModel
		teamRolesTF    []project.TfTeamModel
	}{
		{
			name:        "Team roles",
			teamRolesTF: teams,
			expectedResult: map[types.String]project.TfTeamModel{
				types.StringValue("id1"): teams[0],
				types.StringValue("id2"): teams[1],
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultModel := project.NewTfTeamModelMap(tc.teamRolesTF)
			if !reflect.DeepEqual(resultModel, tc.expectedResult) {
				t.Errorf("created terraform model did not match expected output")
			}
		})
	}
}

func TestLimitModelMapTF(t *testing.T) {
	limits := []project.TfLimitModel{
		{
			Name: types.StringValue("limit1"),
		},
		{
			Name: types.StringValue("limit2"),
		},
	}
	testCases := []struct {
		name           string
		expectedResult map[types.String]project.TfLimitModel
		limitsTF       []project.TfLimitModel
	}{
		{
			name:     "Limits",
			limitsTF: limits,
			expectedResult: map[types.String]project.TfLimitModel{
				types.StringValue("limit1"): limits[0],
				types.StringValue("limit2"): limits[1],
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			resultModel := project.NewTfLimitModelMap(tc.limitsTF)
			if !reflect.DeepEqual(resultModel, tc.expectedResult) {
				t.Errorf("created terraform model did not match expected output")
			}
		})
	}
}
