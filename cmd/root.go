package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "analisis-algoritmos",
	Short: "Cositas de análisis de algoritmos",
	Run:   root,
}

func root(cmd *cobra.Command, args []string) {
	fmt.Println("Hola, por favor dame un comando")
}

func Execute() {
	if rootCmd.Execute() != nil {
		os.Exit(1)
	}
}
