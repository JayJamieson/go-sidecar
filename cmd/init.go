package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var ProjectExists = errors.New("project already exists")

var configuration = []byte(`# Make sure to check the documentation at https://github.com/JayJamieson/go-sidecar/wiki/Configuration-reference
name: gosidecar # Name of application, used to name functions when deploying. Defaults to gosidecar
timeout: 300 # Default Lambda timeout
memory: 128 # Default Lambda memory
storage: 512 # Default Lambda ephemeral storage for /tmp directory 

# List all functions you want deployed here. This is auto update for functions added using add command.
# try and ensure function file name matches function name, less trouble this way
functions:
    - sample 
`)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Setup GoSidecar directory and files",
	Long:  "Initialize (gosidecar init) will create a new directory (./gosidecar), package files and configuration files.",
	Run: func(cmd *cobra.Command, args []string) {
		projectPath, err := initializeProject()
		cobra.CheckErr(err)

		fmt.Printf("Your GoSidecar project is ready at:\n%s\n", projectPath)
	},
}

func initializeProject() (string, error) {
	wd, err := os.Getwd()

	if err != nil {
		return "", err
	}
	projectPath := filepath.Join(wd, "gosidecar")

	if _, err = os.Stat(projectPath); !os.IsNotExist(err) && err == nil {
		return "", ProjectExists
	}

	err = os.MkdirAll(projectPath, 0754)

	if err != nil {
		return "", err
	}

	err = os.WriteFile(filepath.Join(projectPath, "gosidecar.yaml"), configuration, 0754)

	if err != nil {
		return "", err
	}

	return projectPath, nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}
