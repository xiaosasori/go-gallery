package views

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
)

var ViewsDir string = "views"
var LayoutDir string = ViewsDir + "/layouts"
var ViewsExt string = ".gohtml"

func NewView(layout string, files ...string) *View {
	files = addViewsDirPrefix(files)
	files = addViewsExtSuffix(files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	var vd Data
	switch d := data.(type) {
	case Data:
		vd = d
	default:
		vd = Data{
			Yield: data,
		}
	}
	var buf bytes.Buffer
	if err := v.Template.ExecuteTemplate(&buf, v.Layout, vd); err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong. If the problem persists, please email support@lenslocked.com", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, r, nil)
}

// This function takes a slice of strings and prepends the
// ViewsDir directory to the file paths.
//
// Eg the input ["home"] would result in the output ["views/home"]
func addViewsDirPrefix(files []string) []string {
	ret := make([]string, len(files))
	for i, f := range files {
		ret[i] = ViewsDir + "/" + f
	}
	return ret
}

// This function takes a slice of strings and appends the
// ViewsExt suffix to the file paths.
//
// Eg the input ["home"] would result in the output ["home.gohtml"]
func addViewsExtSuffix(files []string) []string {
	ret := make([]string, len(files))
	for i, f := range files {
		ret[i] = f + ViewsExt
	}
	return ret
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*" + ViewsExt)
	if err != nil {
		panic(err)
	}
	return files
}
