package cmd

import (
	"embed"
	"log"
	"strings"

	"github.com/Diniboy1123/usque/config"
	"github.com/Diniboy1123/usque/internal"
	"github.com/spf13/cobra"
)

//go:embed assets/*
var assets embed.FS

var rootCmd = &cobra.Command{
	Use:   "usque",
	Short: "Usque Warp CLI",
	Long:  "An unofficial Cloudflare Warp CLI that uses the MASQUE protocol and exposes the tunnel as various different services.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			log.Fatalf("Failed to get config path: %v", err)
		}

		if configPath != "" {
			if asset, found := strings.CutPrefix(configPath, "my:"); found {
				file, err := assets.Open(asset)
				if err != nil {
					log.Fatalf("failed to open config file: %v", err)
				} else {
					defer func() { _ = file.Close() }()
					if err := config.LoadConfigFile(file); err != nil {
						log.Printf("Config file not found: %v", err)
						log.Printf("You may only use the register command to generate one.")
					}
				}
			} else {
				if err := config.LoadConfig(configPath); err != nil {
					log.Printf("Config file not found: %v", err)
					log.Printf("You may only use the register command to generate one.")
				}
			}
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	internal.InstallDefaultLogTZStamp()
	rootCmd.PersistentFlags().StringP("config", "c", "config.json", "config file (default is config.json)")
}
