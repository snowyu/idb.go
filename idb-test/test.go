// Hello is a trivial example of a main package.
package main

import (
        "fmt"
        //"xattr"
        "iDB"
        "flag"
)

func main() {
    print := flag.Bool("p", false, "print")
    long := flag.Bool("l", false, "long")
    write := flag.Bool("w", false, "write")
    delete := flag.Bool("d", false, "delete")
    flag.Parse()
    switch {
        case *print:
            name := flag.Arg(0)
            for _, file := range flag.Args()[1:] {
                    value, err := iDB.GetXattr(file, name)
                    if *long {
                            fmt.Printf("%s: %s\n", name, value)
                    } else {
                            fmt.Println(value)
                    }
                    if err != nil {
                        fmt.Println(err)
                    }
            }
        case *write:
                name, value := flag.Arg(0), flag.Arg(1)
                for _, file := range flag.Args()[2:] {
                        iDB.SetXattr(file, name, value)
        }
        case *delete:
                name := flag.Arg(0)
                for _, file := range flag.Args()[1:] {
                        iDB.DeleteXattr(file, name)
        }
        default:
            for _, file := range flag.Args() {
                    names, _ := iDB.ListXattr(file)
                    fmt.Println("file:",file)
                    for _, name := range names {
                            if *long {
                                    value, _ := iDB.GetXattr(file, name)
                                    fmt.Printf("%s: %s\n", name, value)
                            } else {
                                    fmt.Println(name)
                            }
                    }
            }
    }
}
