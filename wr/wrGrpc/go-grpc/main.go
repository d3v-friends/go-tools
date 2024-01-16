package main

import (
	"fmt"
	"github.com/d3v-friends/go-tools/fn/fnFile"
	"github.com/d3v-friends/go-tools/fn/fnPanic"
	"github.com/d3v-friends/go-tools/fn/fnYaml"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

func main() {
	var err error
	cmd := &cobra.Command{
		Use: "go-grpc",
	}

	var yamlPath = "grpc.yml"
	var yamlPathFlag = "config"

	cmd.Flags().StringVar(&yamlPath, yamlPathFlag, yamlPath, "--config grpc.yml")

	cmd.Run = func(c *cobra.Command, args []string) {
		yamlPath = fnPanic.Get(c.Flags().GetString(yamlPathFlag))

		var fp = fnPanic.OnValue(fnFile.NewPathBuilderWithWD())
		fp.Join(yamlPath)
		yamlPath = fp.String()

		var yaml = &protocYaml{}
		if err = fnYaml.Open(yamlPath, yaml); err != nil {
			panic(err)
		}

		var outPath string
		if outPath, err = yaml.Out.Path(); err != nil {
			panic(err)
		}

		if err = os.MkdirAll(outPath, os.ModePerm); err != nil {
			panic(err)
		}

		yaml.generate("go")
		yaml.generate("go-grpc")

		fmt.Printf("generated!\n")
	}

	fnPanic.On(cmd.Execute())
	return
}

type protocYaml struct {
	Proto   fnFile.Path   `yaml:"proto"`
	Out     fnFile.Path   `yaml:"out"`
	Package string        `yaml:"package"`
	Include []fnFile.Path `yaml:"include"`
	Import  []fnFile.Path `yaml:"import"`
}

func (x *protocYaml) generate(opt string) {
	var protoPath = fnPanic.OnValue(x.Proto.Path())
	var outPath = fnPanic.OnValue(x.Out.Path())

	for _, protoNm := range x.Include {
		var strCmd = make([]string, 0)
		strCmd = append(strCmd,
			fmt.Sprintf("--proto_path=%s", protoPath),
			fmt.Sprintf("--%s_out=%s", opt, outPath),
			fmt.Sprintf("--%s_opt=M%s=/%s", opt, protoNm, x.Package),
		)

		var importFileLs = make([]string, 0)
		for _, importProtoNm := range x.Import {
			strCmd = append(strCmd,
				fmt.Sprintf("--%s_opt=M%s=/%s", opt, importProtoNm.String(), x.Package),
			)

			importFileLs = append(importFileLs, importProtoNm.String())
		}

		strCmd = append(strCmd, protoNm.String())
		strCmd = append(strCmd, importFileLs...)

		var cmd = exec.Command("protoc", strCmd...)
		log.Println(fmt.Sprintf("cmd: %s", cmd.String()))
		fnPanic.On(cmd.Run())
	}
}
