package OOP

import "fmt"

// The parent struct
type Animal struct {
    Name   string
    Origin string
}

// The child struct which embeds Animal
type Bird struct {
    Animal
    CanFly bool
}

func inheritance() {
    // Creating an instance of Bird
    b := Bird{
        Animal: Animal{Name: "Ostrich", Origin: "Africa"},
        CanFly: false,
    }

    // Accessing the fields of the embedded struct
    fmt.Printf("Name: %s\n", b.Name)       // Output: Name: Ostrich
    fmt.Printf("Origin: %s\n", b.Origin)   // Output: Origin: Africa
    fmt.Printf("CanFly: %v\n", b.CanFly)   // Output: CanFly: false
}
