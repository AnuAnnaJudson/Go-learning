package stores

type Storer interface {
    Create(usr User) error
    Update(usr User) error
    Delete(usr User) error
}
 
// var StorerInterface Storer

type Service struct {
    Storer //embedding interface inside a struct (automatically implements the interface)
}

func NewService(storer Storer) Service {
    s := Service{Storer: storer} //without having to use a global variable which can be changed further
                                 //create an instance of the Storer interface when the function is called and 
                                 //instance of a type is passed to it from main
    return s
}