# go-gallery

An awesome photo gallery application written in Go!

```go
  // ...string means we can pass one or many files NewView("f1", "f2")
  func NewView(files ...string) *View {
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

  func layoutFiles() []string {
    files, err := filepath.Glob(LayoutDir + "/*.gohtml")
    if err != nil {
      panic(err)
    }
    return files
  }

```sql
create table users(
id SERIAL PRIMARY KEY,
name TEXT,
email TEXT NOT NULL
)
```