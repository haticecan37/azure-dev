// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"context"
	"github.com/azure/azure-dev/cli/azd/cmd/actions"
	"github.com/azure/azure-dev/cli/azd/internal"
	"github.com/azure/azure-dev/cli/azd/pkg/account"
	"github.com/azure/azure-dev/cli/azd/pkg/auth"
	"github.com/azure/azure-dev/cli/azd/pkg/config"
	"github.com/azure/azure-dev/cli/azd/pkg/input"
	"github.com/azure/azure-dev/cli/azd/pkg/output"
	"github.com/azure/azure-dev/cli/azd/pkg/templates"
	"github.com/azure/azure-dev/cli/azd/pkg/tools/git"
	"github.com/spf13/cobra"
)

import (
	_ "github.com/azure/azure-dev/cli/azd/pkg/infra/provisioning/bicep"
	_ "github.com/azure/azure-dev/cli/azd/pkg/infra/provisioning/terraform"
)

// Injectors from wire.go:

func initConsole(cmd *cobra.Command, o *internal.GlobalCommandOptions) (input.Console, error) {
	formatter, err := output.GetCommandFormatter(cmd)
	if err != nil {
		return nil, err
	}
	console := newConsoleFromOptions(o, formatter, cmd)
	return console, nil
}

func initDeployAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags deployFlags, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	cmdDeployAction, err := newDeployAction(flags, azdContext, console, formatter, writer)
	if err != nil {
		return nil, err
	}
	return cmdDeployAction, nil
}

func initInitAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags initFlags, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	manager := config.NewManager()
	commandRunner := newCommandRunnerFromConsole(console)
	userConfigManager := config.NewUserConfigManager()
	authManager, err := auth.NewManager(userConfigManager)
	if err != nil {
		return nil, err
	}
	tokenCredential, err := newCredential(ctx, authManager)
	if err != nil {
		return nil, err
	}
	azCli := newAzCliFromOptions(o, commandRunner, tokenCredential)
	accountManager, err := account.NewManager(manager, azCli)
	if err != nil {
		return nil, err
	}
	gitCli := git.NewGitCliFromRunner(commandRunner)
	cmdInitAction, err := newInitAction(azdContext, accountManager, commandRunner, console, gitCli, flags)
	if err != nil {
		return nil, err
	}
	return cmdInitAction, nil
}

func initLoginAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags loginFlags, args []string) (actions.Action, error) {
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	userConfigManager := config.NewUserConfigManager()
	manager, err := auth.NewManager(userConfigManager)
	if err != nil {
		return nil, err
	}
	cmdLoginAction := newLoginAction(formatter, writer, manager, flags, console)
	return cmdLoginAction, nil
}

func initLogoutAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags struct{}, args []string) (actions.Action, error) {
	userConfigManager := config.NewUserConfigManager()
	manager, err := auth.NewManager(userConfigManager)
	if err != nil {
		return nil, err
	}
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	cmdLogoutAction := newLogoutAction(manager, formatter, writer)
	return cmdLogoutAction, nil
}

func initUpAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags upFlags, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	manager := config.NewManager()
	commandRunner := newCommandRunnerFromConsole(console)
	userConfigManager := config.NewUserConfigManager()
	authManager, err := auth.NewManager(userConfigManager)
	if err != nil {
		return nil, err
	}
	tokenCredential, err := newCredential(ctx, authManager)
	if err != nil {
		return nil, err
	}
	azCli := newAzCliFromOptions(o, commandRunner, tokenCredential)
	accountManager, err := account.NewManager(manager, azCli)
	if err != nil {
		return nil, err
	}
	gitCli := git.NewGitCliFromRunner(commandRunner)
	cmdInitFlags := flags.initFlags
	cmdInitAction, err := newInitAction(azdContext, accountManager, commandRunner, console, gitCli, cmdInitFlags)
	if err != nil {
		return nil, err
	}
	cmdInfraCreateFlags := flags.infraCreateFlags
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	cmdInfraCreateAction := newInfraCreateAction(cmdInfraCreateFlags, azdContext, console, formatter, writer)
	cmdDeployFlags := flags.deployFlags
	cmdDeployAction, err := newDeployAction(cmdDeployFlags, azdContext, console, formatter, writer)
	if err != nil {
		return nil, err
	}
	cmdUpAction := newUpAction(cmdInitAction, cmdInfraCreateAction, cmdDeployAction, console)
	return cmdUpAction, nil
}

func initMonitorAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags monitorFlags, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	commandRunner := newCommandRunnerFromConsole(console)
	userConfigManager := config.NewUserConfigManager()
	manager, err := auth.NewManager(userConfigManager)
	if err != nil {
		return nil, err
	}
	tokenCredential, err := newCredential(ctx, manager)
	if err != nil {
		return nil, err
	}
	azCli := newAzCliFromOptions(o, commandRunner, tokenCredential)
	cmdMonitorAction := newMonitorAction(azdContext, azCli, console, flags)
	return cmdMonitorAction, nil
}

func initRestoreAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags restoreFlags, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	cmdRestoreAction := newRestoreAction(flags, console, azdContext)
	return cmdRestoreAction, nil
}

func initShowAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags showFlags, args []string) (actions.Action, error) {
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	cmdShowAction := newShowAction(console, formatter, writer, azdContext, flags)
	return cmdShowAction, nil
}

func initVersionAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags versionFlags, args []string) (actions.Action, error) {
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	cmdVersionAction := newVersionAction(flags, formatter, writer, console)
	return cmdVersionAction, nil
}

func initInfraCreateAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags infraCreateFlags, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	cmdInfraCreateAction := newInfraCreateAction(flags, azdContext, console, formatter, writer)
	return cmdInfraCreateAction, nil
}

func initInfraDeleteAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags infraDeleteFlags, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	cmdInfraDeleteAction := newInfraDeleteAction(flags, azdContext, console)
	return cmdInfraDeleteAction, nil
}

func initEnvSetAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags struct{}, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	commandRunner := newCommandRunnerFromConsole(console)
	userConfigManager := config.NewUserConfigManager()
	manager, err := auth.NewManager(userConfigManager)
	if err != nil {
		return nil, err
	}
	tokenCredential, err := newCredential(ctx, manager)
	if err != nil {
		return nil, err
	}
	azCli := newAzCliFromOptions(o, commandRunner, tokenCredential)
	cmdEnvSetAction := newEnvSetAction(azdContext, azCli, console, o, args)
	return cmdEnvSetAction, nil
}

func initEnvSelectAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags struct{}, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	cmdEnvSelectAction := newEnvSelectAction(azdContext, args)
	return cmdEnvSelectAction, nil
}

func initEnvListAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags struct{}, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	cmdEnvListAction := newEnvListAction(azdContext, formatter, writer)
	return cmdEnvListAction, nil
}

func initEnvNewAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags envNewFlags, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	commandRunner := newCommandRunnerFromConsole(console)
	userConfigManager := config.NewUserConfigManager()
	manager, err := auth.NewManager(userConfigManager)
	if err != nil {
		return nil, err
	}
	tokenCredential, err := newCredential(ctx, manager)
	if err != nil {
		return nil, err
	}
	azCli := newAzCliFromOptions(o, commandRunner, tokenCredential)
	cmdEnvNewAction := newEnvNewAction(azdContext, azCli, flags, args, console)
	return cmdEnvNewAction, nil
}

func initEnvRefreshAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags struct{}, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	commandRunner := newCommandRunnerFromConsole(console)
	userConfigManager := config.NewUserConfigManager()
	manager, err := auth.NewManager(userConfigManager)
	if err != nil {
		return nil, err
	}
	tokenCredential, err := newCredential(ctx, manager)
	if err != nil {
		return nil, err
	}
	azCli := newAzCliFromOptions(o, commandRunner, tokenCredential)
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	cmdEnvRefreshAction := newEnvRefreshAction(azdContext, azCli, o, console, formatter, writer)
	return cmdEnvRefreshAction, nil
}

func initEnvGetValuesAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags struct{}, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	commandRunner := newCommandRunnerFromConsole(console)
	userConfigManager := config.NewUserConfigManager()
	manager, err := auth.NewManager(userConfigManager)
	if err != nil {
		return nil, err
	}
	tokenCredential, err := newCredential(ctx, manager)
	if err != nil {
		return nil, err
	}
	azCli := newAzCliFromOptions(o, commandRunner, tokenCredential)
	cmdEnvGetValuesAction := newEnvGetValuesAction(azdContext, console, formatter, writer, azCli, o)
	return cmdEnvGetValuesAction, nil
}

func initPipelineConfigAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags pipelineConfigFlags, args []string) (actions.Action, error) {
	azdContext, err := newAzdContext()
	if err != nil {
		return nil, err
	}
	commandRunner := newCommandRunnerFromConsole(console)
	userConfigManager := config.NewUserConfigManager()
	manager, err := auth.NewManager(userConfigManager)
	if err != nil {
		return nil, err
	}
	tokenCredential, err := newCredential(ctx, manager)
	if err != nil {
		return nil, err
	}
	azCli := newAzCliFromOptions(o, commandRunner, tokenCredential)
	cmdPipelineConfigAction := newPipelineConfigAction(azdContext, console, flags, azCli)
	return cmdPipelineConfigAction, nil
}

func initTemplatesListAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags templatesListFlags, args []string) (actions.Action, error) {
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	templateManager := templates.NewTemplateManager()
	cmdTemplatesListAction := newTemplatesListAction(flags, formatter, writer, templateManager)
	return cmdTemplatesListAction, nil
}

func initTemplatesShowAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags struct{}, args []string) (actions.Action, error) {
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	templateManager := templates.NewTemplateManager()
	cmdTemplatesShowAction := newTemplatesShowAction(formatter, writer, templateManager, args)
	return cmdTemplatesShowAction, nil
}

func initConfigListAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags struct{}, args []string) (actions.Action, error) {
	userConfigManager := config.NewUserConfigManager()
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	cmdConfigListAction := newConfigListAction(userConfigManager, formatter, writer)
	return cmdConfigListAction, nil
}

func initConfigGetAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags struct{}, args []string) (actions.Action, error) {
	userConfigManager := config.NewUserConfigManager()
	formatter := newFormatterFromConsole(console)
	writer := newOutputWriter(console)
	cmdConfigGetAction := newConfigGetAction(userConfigManager, formatter, writer, args)
	return cmdConfigGetAction, nil
}

func initConfigSetAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags struct{}, args []string) (actions.Action, error) {
	userConfigManager := config.NewUserConfigManager()
	cmdConfigSetAction := newConfigSetAction(userConfigManager, args)
	return cmdConfigSetAction, nil
}

func initConfigUnsetAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags struct{}, args []string) (actions.Action, error) {
	userConfigManager := config.NewUserConfigManager()
	cmdConfigUnsetAction := newConfigUnsetAction(userConfigManager, args)
	return cmdConfigUnsetAction, nil
}

func initConfigResetAction(console input.Console, ctx context.Context, o *internal.GlobalCommandOptions, flags struct{}, args []string) (actions.Action, error) {
	userConfigManager := config.NewUserConfigManager()
	cmdConfigResetAction := newConfigResetAction(userConfigManager, args)
	return cmdConfigResetAction, nil
}
