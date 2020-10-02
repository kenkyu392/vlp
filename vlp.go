package vlp

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/zalando/go-keyring"
)

// Run the VLP
func Run(ctx context.Context, argv []string, outStream, errStream io.Writer) error {
	log.SetFlags(log.LstdFlags)
	log.SetOutput(errStream)
	log.SetPrefix("[vlp] ")

	fs := flag.NewFlagSet("vlp", flag.ContinueOnError)
	fs.SetOutput(ioutil.Discard)
	fs.Usage = func() {
		name := fs.Name()
		out := fs.Output()
		fmt.Fprint(out, "VLP is a CLI tool for extracting passwords stored in credentials.\n")
		fmt.Fprintf(out, "Usage: %s <options>\n", name)
		fs.PrintDefaults()
	}

	hf := fs.Bool("help", false, "Prints the help.")
	vf := fs.Bool("version", false, "Prints the version.")
	sf := fs.String("service", "", "Service name.")
	uf := fs.String("username", "", "User name.")

	if err := fs.Parse(argv); err != nil {
		return err
	}
	fs.SetOutput(errStream)

	if *hf {
		fs.Usage()
		return flag.ErrHelp
	}

	if *vf {
		_, err := fmt.Fprintf(outStream, "v%s\n", version)
		return err
	}

	if *sf == "" || *uf == "" {
		fs.Usage()
		return flag.ErrHelp
	}

	password, err := keyring.Get(*sf, *uf)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(map[string]string{
		"service":  *sf,
		"username": *uf,
		"password": password,
	})
}
