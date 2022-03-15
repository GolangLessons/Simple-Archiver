package cmd

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"archiver/lib/comression"
	"archiver/lib/comression/vlc"
	"archiver/lib/comression/vlc/table/shannon_fano"
)

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

// TODO: take extension from the file
const unpackedExtension = "txt"

func unpack(cmd *cobra.Command, args []string) {
	var decoder comression.Decoder

	method := cmd.Flag("method").Value.String()

	switch method {
	case "vlc":
		tblGen := shannon_fano.NewGenerator()
		decoder = vlc.New(tblGen)
	default:
		cmd.PrintErrln("unknown method")
		return
	}

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer func() { _ = r.Close() }()

	data, err := io.ReadAll(r)
	if err != nil {
		handleErr(err)
	}

	packed, err := decoder.Decode(data)
	if err != nil {
		panic("can't decode data: " + err.Error())
	}

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
