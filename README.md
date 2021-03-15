# go-gallery

An awesome photo gallery application written in Go!

## Embedding & Chaining
```go
  type Cat struct{}
  func (c Cat) Speak() {
    fmt.Println("Meow")
  }

  type Dog struct{
    Tail bool
  }

  type Husky struct{
    Dog
    Speaker
  }

  type Speaker interface{
    Speak()
  }

  type SpeakerPrefixer struct {
    Speaker
  }

  func (sp SpeakerPrefixer) Speak() {
    fmt.Print("Prefix: ")
    sp.Speaker.Speak()
  }

  func main() {
    // h := Husky{Dog{Tail: true}, Cat{}}
    // h.Speak() // Meow
    //. fmt.Println(h.Tail) // true
    h := Husky{Dog{Tail: true}, SpeakerPrefixer{Cat{}}}
    h.Speak() // Prefix: Meow
  }
```
```go
  var ViewsDir string = "views"
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
```
