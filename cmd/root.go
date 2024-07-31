/*
Copyright © 2024 ADAM CHOLEWIŃSKI
*/
package cmd

import (
	"os"
	"strings"

	"github.com/adamchol/askme/internal/services"
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

	config, err := services.NewConfigService()
	if err != nil {
		log.Error("Failed to initialize config", "error", err)
	}

	userMsg := strings.Join(args, " ")

	LLMService := services.NewLLMAIService(config)
	switch ModelFlag {
	case "gpt":
		err = LLMService.ShowGPTMessage(userMsg)
	case "claude":
		err = LLMService.ShowClaudeMessage(userMsg)
	}
	if err != nil {
		log.Error("Failed to show the completion", "error", err)
	}

}
