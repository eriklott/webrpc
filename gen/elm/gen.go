//go:generate statik -src=./templates -dest=. -f -Z -p=embed
package elm

import (
	"bytes"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/goware/statik/fs"
	"github.com/webrpc/webrpc/gen"
	"github.com/webrpc/webrpc/gen/elm/embed"
	"github.com/webrpc/webrpc/schema"
)

func init() {
	gen.Register("elm", &generator{})
}

type generator struct{}

func (g *generator) Gen(proto *schema.WebRPCSchema, opts gen.TargetOptions) (string, error) {
	// Get templates from `embed` asset package
	// NOTE: make sure to `go generate` whenever you change the files in `templates/` folder
	templates, err := getTemplates()
	if err != nil {
		return "", err
	}

	// TODO: we can move a bunch of this code to the core gen package at githb.com/webrpc/webrpc/gen
	// .. then typescript gen, and others can use it too..

	// Load templates
	tmpl := template.
		New("webrpc-gen-elm").
		Delims("<%", "%>").
		Funcs(templateFuncMap)

	for _, tmplData := range templates {
		_, err = tmpl.Parse(tmplData)
		if err != nil {
			return "", err
		}
	}

	vars := struct {
		*schema.WebRPCSchema
		TargetOpts gen.TargetOptions
	}{
		proto, opts,
	}

	// Generate the template
	genBuf := bytes.NewBuffer(nil)
	err = tmpl.ExecuteTemplate(genBuf, "proto", vars)
	if err != nil {
		return "", err
	}

	return string(genBuf.Bytes()), nil
}

func getTemplates() (map[string]string, error) {
	data := map[string]string{}

	statikFS, err := fs.New(embed.Asset)
	if err != nil {
		return nil, err
	}

	fs.Walk(statikFS, "/", func(path string, info os.FileInfo, err error) error {
		if path == "/" {
			return nil
		}
		f, err := statikFS.Open(path)
		if err != nil {
			return err
		}
		buf, err := ioutil.ReadAll(f)
		if err != nil {
			return err
		}
		data[path] = string(buf)
		return nil
	})

	return data, nil
}
