package main

import (
	"flag"
	"google.golang.org/protobuf/compiler/protogen"
)

const Version = "1.0.0"

func main() {
	var (
		flags       flag.FlagSet
		showVersion bool
	)
	flags.BoolVar(&showVersion, "version", false, "print the version and exit")
	pbOpts := protogen.Options{
		ParamFunc: flags.Set,
	}
	pbOpts.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			GenerateFile(gen, f)
		}
		return nil
	})
}
