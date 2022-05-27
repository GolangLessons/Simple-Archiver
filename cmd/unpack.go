package cmd

import (
	"archiver/lib/compression/vlc/table/shannon_fano"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"archiver/lib/compression"
	"archiver/lib/compression/vlc"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

// TODO: take extension from the file
const unpackedExtension = "txt"

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		decoder = vlc.New(shannon_fano.Generator{})
	default:
		cmd.PrintErr("unknown method")
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed := decoder.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

// TODO: refactor this
func unpackedFileName(path string) string {
	fileName := filepath.Base(path)

	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + unpackedExtension
}

func init() {
	rootCmd.AddCommand(unpackCmd)

	unpackCmd.Flags().StringP("method", "m", "", "decompression method: vlc")

	if err := unpackCmd.MarkFlagRequired("method"); err != nil {
		panic(err)
	}
}
