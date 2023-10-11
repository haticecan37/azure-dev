package devcenter

import (
	"context"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
	"github.com/azure/azure-dev/cli/azd/pkg/azsdk"
	"github.com/azure/azure-dev/cli/azd/pkg/devcentersdk"
	"github.com/azure/azure-dev/cli/azd/pkg/environment"
	"github.com/azure/azure-dev/cli/azd/pkg/input"
	"github.com/azure/azure-dev/cli/azd/test/mocks"
	"github.com/azure/azure-dev/cli/azd/test/mocks/mockdevcentersdk"
	"github.com/stretchr/testify/require"
)

func Test_Prompt_DevCenter(t *testing.T) {
	mockContext := mocks.NewMockContext(context.Background())
	selectedDevCenterIndex := 0

	manager := &mockDevCenterManager{}
	manager.
		On("WritableProjects", *mockContext.Context).
		Return(mockProjects, nil)

	prompter := newPrompterForTest(t, mockContext, &Config{}, manager)

	mockContext.Console.WhenSelect(func(options input.ConsoleOptions) bool {
		return strings.Contains(options.Message, "Select a Dev Center")
	}).RespondFn(func(options input.ConsoleOptions) (any, error) {
		return selectedDevCenterIndex, nil
	})

	selectedDevCenter, err := prompter.PromptDevCenter(*mockContext.Context)
	require.NoError(t, err)
	require.NotNil(t, selectedDevCenter)
	require.Equal(t, mockDevCenterList[selectedDevCenterIndex], selectedDevCenter)
}

func Test_Prompt_Project(t *testing.T) {
	mockContext := mocks.NewMockContext(context.Background())
	selectedDevCenter := mockDevCenterList[0]
	selectedProjectIndex := 1

	manager := &mockDevCenterManager{}
	manager.
		On("WritableProjects", *mockContext.Context).
		Return(mockProjects, nil)

	prompter := newPrompterForTest(t, mockContext, &Config{}, manager)

	mockContext.Console.WhenSelect(func(options input.ConsoleOptions) bool {
		return strings.Contains(options.Message, "Select a project")
	}).RespondFn(func(options input.ConsoleOptions) (any, error) {
		return selectedProjectIndex, nil
	})

	selectedProject, err := prompter.PromptProject(*mockContext.Context, selectedDevCenter.Name)
	require.NoError(t, err)
	require.NotNil(t, selectedProject)
	require.Equal(t, mockProjects[selectedProjectIndex], selectedProject)
}

func Test_Prompt_Catalog(t *testing.T) {
	mockContext := mocks.NewMockContext(context.Background())
	selectedDevCenter := mockDevCenterList[0]
	selectedProject := mockProjects[1]
	selectedCatalogIndex := 1

	mockdevcentersdk.MockDevCenterGraphQuery(mockContext, mockDevCenterList)
	mockdevcentersdk.MockListCatalogs(mockContext, selectedProject.Name, mockCatalogs)

	manager := &mockDevCenterManager{}
	manager.
		On("WritableProjects", *mockContext.Context).
		Return(mockProjects, nil)

	prompter := newPrompterForTest(t, mockContext, &Config{}, manager)

	mockContext.Console.WhenSelect(func(options input.ConsoleOptions) bool {
		return strings.Contains(options.Message, "Select a catalog")
	}).RespondFn(func(options input.ConsoleOptions) (any, error) {
		return selectedCatalogIndex, nil
	})

	selectedCatalog, err := prompter.PromptCatalog(
		*mockContext.Context,
		selectedDevCenter.Name,
		selectedProject.Name,
	)
	require.NoError(t, err)
	require.NotNil(t, selectedCatalog)
	require.Equal(t, mockCatalogs[selectedCatalogIndex], selectedCatalog)
}

func Test_Prompt_EnvironmentType(t *testing.T) {
	mockContext := mocks.NewMockContext(context.Background())
	selectedDevCenter := mockDevCenterList[0]
	selectedProject := mockProjects[1]

	selectedIndex := 3

	mockdevcentersdk.MockDevCenterGraphQuery(mockContext, mockDevCenterList)
	mockdevcentersdk.MockListEnvironmentTypes(mockContext, selectedProject.Name, mockEnvironmentTypes)

	manager := &mockDevCenterManager{}
	manager.
		On("WritableProjects", *mockContext.Context).
		Return(mockProjects, nil)

	prompter := newPrompterForTest(t, mockContext, &Config{}, manager)

	mockContext.Console.WhenSelect(func(options input.ConsoleOptions) bool {
		return strings.Contains(options.Message, "Select an environment type")
	}).RespondFn(func(options input.ConsoleOptions) (any, error) {
		return selectedIndex, nil
	})

	selectedEnvironmentType, err := prompter.PromptEnvironmentType(
		*mockContext.Context,
		selectedDevCenter.Name,
		selectedProject.Name,
	)
	require.NoError(t, err)
	require.NotNil(t, selectedEnvironmentType)
	require.Equal(t, mockEnvironmentTypes[selectedIndex], selectedEnvironmentType)
}

func Test_Prompt_EnvironmentDefinitions(t *testing.T) {
	mockContext := mocks.NewMockContext(context.Background())
	selectedDevCenter := mockDevCenterList[0]
	selectedProject := mockProjects[1]

	selectedIndex := 2

	mockdevcentersdk.MockDevCenterGraphQuery(mockContext, mockDevCenterList)
	mockdevcentersdk.MockListEnvironmentDefinitions(mockContext, selectedProject.Name, mockEnvDefinitions)

	manager := &mockDevCenterManager{}
	manager.
		On("WritableProjects", *mockContext.Context).
		Return(mockProjects, nil)

	prompter := newPrompterForTest(t, mockContext, &Config{}, manager)

	mockContext.Console.WhenSelect(func(options input.ConsoleOptions) bool {
		return strings.Contains(options.Message, "Select an environment definition")
	}).RespondFn(func(options input.ConsoleOptions) (any, error) {
		return selectedIndex, nil
	})

	selectedEnvironmentType, err := prompter.PromptEnvironmentDefinition(
		*mockContext.Context,
		selectedDevCenter.Name,
		selectedProject.Name,
	)
	require.NoError(t, err)
	require.NotNil(t, selectedEnvironmentType)
	require.Equal(t, mockEnvDefinitions[selectedIndex], selectedEnvironmentType)
}

func Test_Prompt_Parameters(t *testing.T) {
	type paramWithValue struct {
		devcentersdk.Parameter
		userValue     any
		expectedValue any
	}

	t.Run("MultipleParameters", func(t *testing.T) {
		mockContext := mocks.NewMockContext(context.Background())
		prompter := newPrompterForTest(t, mockContext, &Config{}, nil)
		promptedParams := map[string]bool{}

		expectedValues := map[string]paramWithValue{
			"param1": {
				Parameter:     devcentersdk.Parameter{Id: "param1", Name: "Param 1", Type: devcentersdk.ParameterTypeString},
				userValue:     "value1",
				expectedValue: "value1",
			},
			"param2": {
				Parameter:     devcentersdk.Parameter{Id: "param2", Name: "Param 2", Type: devcentersdk.ParameterTypeString},
				userValue:     "value2",
				expectedValue: "value2",
			},
			"param3": {
				Parameter:     devcentersdk.Parameter{Id: "param3", Name: "Param 3", Type: devcentersdk.ParameterTypeBool},
				userValue:     true,
				expectedValue: true,
			},
			"param4": {
				Parameter:     devcentersdk.Parameter{Id: "param4", Name: "Param 4", Type: devcentersdk.ParameterTypeInt},
				userValue:     "123",
				expectedValue: 123,
			},
		}

		var mockPrompt = func(key string, param paramWithValue) {
			if param.Type == devcentersdk.ParameterTypeBool {
				mockContext.Console.WhenConfirm(func(options input.ConsoleOptions) bool {
					return strings.Contains(options.Message, param.Name)
				}).RespondFn(func(options input.ConsoleOptions) (any, error) {
					promptedParams[key] = true
					return param.userValue, nil
				})
			} else {
				mockContext.Console.WhenPrompt(func(options input.ConsoleOptions) bool {
					return strings.Contains(options.Message, param.Name)
				}).RespondFn(func(options input.ConsoleOptions) (any, error) {
					promptedParams[key] = true
					return param.userValue, nil
				})
			}
		}

		for key, param := range expectedValues {
			mockPrompt(key, param)
		}

		env := environment.New("Test")
		envDefinition := &devcentersdk.EnvironmentDefinition{
			Parameters: []devcentersdk.Parameter{
				{
					Id:   "param1",
					Name: "Param 1",
					Type: devcentersdk.ParameterTypeString,
				},
				{
					Id:   "param2",
					Name: "Param 2",
					Type: devcentersdk.ParameterTypeString,
				},
				{
					Id:   "param3",
					Name: "Param 3",
					Type: devcentersdk.ParameterTypeBool,
				},
				{
					Id:   "param4",
					Name: "Param 4",
					Type: devcentersdk.ParameterTypeInt,
				},
			},
		}

		values, err := prompter.PromptParameters(*mockContext.Context, env, envDefinition)
		require.NoError(t, err)

		for key, value := range values {
			require.Equal(t, expectedValues[key].expectedValue, value)
		}
	})

	t.Run("WithSomeSetValues", func(t *testing.T) {
		mockContext := mocks.NewMockContext(context.Background())
		prompter := newPrompterForTest(t, mockContext, &Config{}, nil)
		promptCalled := false

		// Only mock response for param 3
		mockContext.Console.WhenPrompt(func(options input.ConsoleOptions) bool {
			return strings.Contains(options.Message, "Param 3")
		}).RespondFn(func(options input.ConsoleOptions) (any, error) {
			promptCalled = true
			return "value3", nil
		})

		env := environment.New("Test")
		envDefinition := &devcentersdk.EnvironmentDefinition{
			Parameters: []devcentersdk.Parameter{
				{
					Id:   "param1",
					Name: "Param 1",
					Type: devcentersdk.ParameterTypeString,
				},
				{
					Id:   "param2",
					Name: "Param 2",
					Type: devcentersdk.ParameterTypeString,
				},
				{
					Id:   "param3",
					Name: "Param 3",
					Type: devcentersdk.ParameterTypeString,
				},
			},
		}

		_ = env.Config.Set("provision.parameters.param1", "value1")
		_ = env.Config.Set("provision.parameters.param2", "value2")

		values, err := prompter.PromptParameters(*mockContext.Context, env, envDefinition)
		require.NoError(t, err)
		require.True(t, promptCalled)
		require.Equal(t, "value1", values["param1"])
		require.Equal(t, "value2", values["param2"])
		require.Equal(t, "value3", values["param3"])
	})

	t.Run("WithAllSetValues", func(t *testing.T) {
		mockContext := mocks.NewMockContext(context.Background())
		prompter := newPrompterForTest(t, mockContext, &Config{}, nil)

		env := environment.New("Test")
		envDefinition := &devcentersdk.EnvironmentDefinition{
			Parameters: []devcentersdk.Parameter{
				{
					Id:   "param1",
					Name: "Param 1",
					Type: devcentersdk.ParameterTypeString,
				},
				{
					Id:   "param2",
					Name: "Param 2",
					Type: devcentersdk.ParameterTypeString,
				},
			},
		}

		_ = env.Config.Set("provision.parameters.param1", "value1")
		_ = env.Config.Set("provision.parameters.param2", "value2")

		values, err := prompter.PromptParameters(*mockContext.Context, env, envDefinition)
		require.NoError(t, err)
		require.Equal(t, "value1", values["param1"])
		require.Equal(t, "value2", values["param2"])
	})
}

func newPrompterForTest(t *testing.T, mockContext *mocks.MockContext, config *Config, manager Manager) *Prompter {
	coreOptions := azsdk.
		DefaultClientOptionsBuilder(*mockContext.Context, mockContext.HttpClient, "azd").
		BuildCoreClientOptions()

	armOptions := azsdk.
		DefaultClientOptionsBuilder(*mockContext.Context, mockContext.HttpClient, "azd").
		BuildArmClientOptions()

	resourceGraphClient, err := armresourcegraph.NewClient(mockContext.Credentials, armOptions)
	require.NoError(t, err)

	devCenterClient, err := devcentersdk.NewDevCenterClient(
		mockContext.Credentials,
		coreOptions,
		resourceGraphClient,
	)

	require.NoError(t, err)

	return NewPrompter(config, mockContext.Console, manager, devCenterClient)
}