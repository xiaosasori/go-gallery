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
type UserDB interface {
	// Methods for querying for single users
	ByID(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRemember(token string) (*User, error)

	// Methods for altering users
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error

	// Used to close a DB connection
	Close() error

	// Migration helpers
	AutoMigrate() error
	DestructiveReset() error
}

type UserService interface {
	Authenticate(email, password string) (*User, error)
	UserDB
}

func NewUserService(connectionInfo string) (UserService, error) {
	ug, err := newUserGorm(connectionInfo)
	if err != nil {
		return nil, err
	}
	uv := newUserValidator(ug)
	return &userService{
		UserDB: uv,
	}, nil
}

var _ UserService = &userService{}

type userService struct {
	UserDB
}

// validator layer
type userValFunc func(*User) error

func runUserValFuncs(user *User, fns ...userValFunc) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}

var _ UserDB = &userValidator{}

func newUserValidator(udb UserDB) *userValidator {
	return &userValidator{
		UserDB:     udb,
	}
}

type userValidator struct {
	UserDB
}

func (uv *userValidator) ByEmail(email string) (*User, error) {
	user := User{
		Email: email,
	}
	if err := runUserValFuncs(&user, uv.normalizeEmail); err != nil {
		return nil, err
	}
	return uv.UserDB.ByEmail(user.Email)
}

func (uv *userValidator) normalizeEmail(user *User) error {
	user.Email = strings.ToLower(user.Email)
	user.Email = strings.TrimSpace(user.Email)
	return nil
}

// Delete will delete the user with the provided ID
func (uv *userValidator) Delete(id uint) error {
	// user := User{
	//   Model: gorm.Model{ID: id,},
	// }
	var user User
	user.ID = id
	err := runUserValFuncs(&user, uv.idGreaterThan(0))
	if err != nil {
		return err
	}
	return uv.UserDB.Delete(id)
}

func (uv *userValidator) idGreaterThan(n uint) userValFunc {
	return userValFunc(func(user *User) error {
		if user.ID <= n {
			return ErrIDInvalid
		}
		return nil
	})
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
