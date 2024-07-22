package dsk

import (
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/soumya92/barista/format"
	"github.com/soumya92/barista/bar"
	"github.com/soumya92/barista/modules/diskio"
	"github.com/soumya92/barista/modules/diskspace"
	"github.com/soumya92/barista/outputs"
	"github.com/soumya92/barista/pango"
	"github.com/glebtv/custom_barista/utils"
)

func AddTo(modules []bar.Module) []bar.Module {
	modules = append(modules, diskspace.New("/"))

	path, err := exec.LookPath("findmnt")
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(path, "/")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		parts := strings.Split(string(out), "\n")
		info := strings.Fields(parts[1])
		dn := strings.Split(info[1], "/")
		name := dn[len(dn)-1]

		diskio.RefreshInterval(2 * time.Second)

		sda := diskio.New(name).
			Output(func(io diskio.IO) bar.Output {
				//spew.Dump(io)
				return outputs.Pango(
					pango.Textf("io "),
					pango.Textf("%9s", format.IByterate(io.Input)),
					utils.Spacer,
					pango.Textf("%9s", format.IByterate(io.Output)),
				)
			})

		modules = append(modules, sda)
	}

	return modules
}

