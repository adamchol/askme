/*
Copyright © 2024 ADAM CHOLEWIŃSKI
*/
package cmd

import (
	"os"
	"strings"

	"github.com/adamchol/askme/internal"
	"github.com/adamchol/askme/internal/services"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var ModelFlag string

var rootCmd = &cobra.Command{
	Use:     "ask [prompt]",
	Version: "v0.0.1",
	Short:   "AskMe CLI - Terminal interface for popular LLMs.",
	Long:    `Use AI models from OpenAI, Google, Anthropic and more with your own API keys.`,

	Run: runRoot,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate("AskMe CLI {{ .Version }}\n")
	rootCmd.Flags().StringVarP(&ModelFlag, "model", "m", "gpt", "Model to use for complation. Works with aliases like \"claude\" or \"gpt\".")
}

func runRoot(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		return
	}

	config, err := internal.NewConfigService()
	if err != nil {
		log.Error("Failed to initialize config", "error", err)
	}

	userMsg := strings.Join(args, " ")

	p := tea.NewProgram(&services.UIModel{
		Input: services.CompletionInput{
			Prompt: userMsg,
			Model:  "gpt-4o-mini",
		},
		Config: *config,
	})

	if _, err := p.Run(); err != nil {
		log.Error("Failed to run a program", "error", err)
	}

}
