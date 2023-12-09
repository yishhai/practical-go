// run `go vet game.go` for static analysis
package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type mover interface {
	Move(x, y int)
	// Move(int, int)
}

type Item struct {
	X int
	Y int
}

type Player struct {
	Name string
	Item // Embedded struct
	// Item Item // we can use the syntax above or this one
	Keys []Key
}

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("invalid key: %s", k)
	}

	l := len(p.Keys)

	if l > 0 {
		for _, ks := range p.Keys {
			if k == ks {
				return fmt.Errorf("key exist: %s", k)
			}
		}
	}
	p.Keys = append(p.Keys, k)
	return nil
}

func (p *Player) FoundKey2(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("invalid key: %s", k)
	}

	// if containsKey(p.Keys, k) {
	// 	return fmt.Errorf("key exist: %s", k)
	// }
	// the above or this
	if slices.Contains(p.Keys, k) {
		return fmt.Errorf("key exist: %s", k)
	}

	p.Keys = append(p.Keys, k)
	return nil
}

func containsKey(Keys []Key, k Key) bool {
	for _, ks := range Keys {
		if k == ks {
			return true
		}
	}
	return false
}

const (
	maxX = 1000
	maxY = 600
	min  = 0
)

type Key byte

// iota is Go's version of enums,
// they are auto incremented by one and starts from 0.
const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // not exported - small caps
)

func main() {
	var i1 Item
	i1.X = 10
	i1.Y = 20
	fmt.Printf("%#v\n", i1)

	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, -20))

	// i1 in this case is not a pointer,
	// but the Go compiler is smart enough to pass a pointer to the `Move` receiver function since that is what it accepts
	i1.Move(30, 40)
	fmt.Println("i1:", i1)

	i2, _ := NewItem(100, 200)
	// i1 in this case is not a pointer, it was created with `NewItem()` which returns an Item struct pointer
	i2.Move(250, 500)
	fmt.Println("i2:", i2)

	p1 := Player{
		Name: "Jesse",
		Item: Item{44, 902},
	}
	// The properties of the embedded Item struct is accessible in the top layer of the Player struct,
	// so we can do `p1.X` instead of `p1.Item.X`,
	// however if the struct has a field with the same name or property there will be a conflict, thus we have to use `p1.Item.X`
	// e.g
	// type Player struct {
	// 	Name string
	//	X // this field we cause a conflict if we try to use `p1.X` and not `p1.Item.X`
	// 	Item
	// }
	// This is also the case if there are other embedded structs with same field name
	fmt.Printf("p1.X: %v\n", p1.X)
	// We can also do
	p1.Move(50, 100) // this is not inheritance but embedding
	fmt.Printf("p1 after move: %v\n", p1)

	ms := []mover{&i1, &p1, i2}
	moveAll(ms, 1, 2)
	for _, m := range ms {
		fmt.Printf("%#v\n", m)
	}

	k := Jade
	fmt.Println(k)       // prints "Jade" because of the String() function and the Stringer interface
	fmt.Println(Key(19)) // prints <Key 19> because we didn't handle it in the String function

	// json.NewEncoder(os.Stdout).Encode(time.Now())

	err := p1.FoundKey(Jade)
	err = p1.FoundKey(Copper)
	err = p1.FoundKey2(Crystal)
	// err = p1.FoundKey2(Crystal) // key exist: Crystal
	// err = p1.FoundKey(Copper) // key exist: Copper
	// err = p1.FoundKey(Crystal) // key exist: Crystal
	// err = p1.FoundKey(Jade) // key exist: Jade
	// err = p1.FoundKey(100) // invalid key: <Key 100>
	fmt.Println(err)
	fmt.Printf("add key for p1: %v\n", p1)

}

// does not work
// func FoundKey(k Key, p *Player) error {

// 	switch k {
// 	case Jade:
// 		if len(p.Keys) > 0 {
// 			for _, ks := range p.Keys {
// 				if k == ks {
// 					break
// 				}
//
// 			}
// 		} else if len(p.Keys) > 0 {
// 			p.Keys = append(p.Keys, k)
// 		}
//
// 	case Copper:
// 		if len(p.Keys) > 0 {
// 			for _, ks := range p.Keys {
// 				if k == ks {
// 					break
// 				}

// 			}
// 		} else if len(p.Keys) > 0{
// 			p.Keys = append(p.Keys, k)
// 		}
// 	case Crystal:
// 		if len(p.Keys) > 0 {
// 			for _, ks := range p.Keys {
// 				if k == ks {
// 					break
// 				}
//
// 			}
// 		} else {
// 			p.Keys = append(p.Keys, k)
// 		}
// 	default:
// 		return fmt.Errorf("Key not found %d", k)
// 	}
// 	return nil
// }

// String implements the fmt.Stringer interface for the Key type.
// When Key is used with fmt.Printf, fmt.Println, or any other print-related method in the fmt package,
// Go will automatically call this String method to get a string representation of Key.
// This allows for a customized textual representation of Key values when printing.
func (k Key) String() string {
	switch k {
	case Jade:
		return "Jade"
	case Copper:
		return "Copper" // Fixed the return value to match the case
	case Crystal:
		return "Crystal"
	}

	// Return a default string representation for Key values not listed in the switch.
	return fmt.Sprintf("<Key %d>", k)
}

// Creates a new Item struct and returns a pointer to it
/*
Why would you want to return a pointer to a struct?

Efficiency with Large Structures: If Item were a large struct, returning a pointer would be more efficient because it avoids copying the entire struct when the function returns.
Instead, only the address (a pointer) is copied, which is much smaller in size. This can significantly reduce memory usage and increase performance,
especially in functions that are called frequently or return large data structures.

Mutability: Returning a pointer allows the caller to modify the struct directly. When you return a struct, you're returning a copy,
so any changes made to it won't affect the original. But when you return a pointer, modifications to the pointed-to struct are reflected in the original, allowing for mutable behavior.

Indicating Failure: In Go, it's a common pattern to return a pointer to a struct along with an error. When an error occurs, the pointer can be nil,
clearly indicating a failure to create the struct. This pattern is widely used for functions that might fail to produce a valid result.

Consistency: Sometimes, functions return pointers for consistency with other parts of the API or codebase,
even if it's not strictly necessary for performance or mutability reasons. Consistent APIs can be easier to understand and use.

Heap Allocation: As noted in your comment, when you return a pointer to a local variable,
Go's escape analysis determines that the variable must be allocated on the heap rather than the stack. This ensures the variable persists after the function returns.
While heap allocations can be more costly than stack allocations, they are necessary for data that must outlive the function call.
*/
func NewItem(x, y int) (*Item, error) {
	if x < min || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds (max: %d/%d - min: %d/%d)", x, y, maxX, maxY, min, min)
	}

	i := Item{x, y}

	// `i` is originally created on the stack,
	// but because we're making a reference to `i` below
	// the compiler will move `i` from the stack to the heap
	// this can affect program performance
	// run: `go run -gcflags=-m game.go` for escape analysis
	// println(&i) // uncomment this to print the actual pointer to `i`
	return &i, nil
}

// We use a pointer so we can mutate the struct item
func (ir *Item) Move(x, y int) {
	ir.X = x
	ir.Y = y
}

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

// Generic function example
// Go >=1.18 supports generics
// func NewNumber[T int | float64](kind string) T {
// 	if kind == "int" {
// 		return 0
// 	}
// 	return 0.0
// }
