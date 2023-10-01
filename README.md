# Guide Golang
    package main
    import ("fmt")

    func myFunction(fullName string) {
        fmt.Printf("%v Fullname: ", fullName)
    }

    func main() {
        ## This is a comment
        fmt.Println("Hello, World!");

        // Variables
        var myNum = 50;
        var myWord = "Hello!";
        var x = 5;
        var y = 10;
        var z = x + y;
        var a, b, c = 5, 10, 15;

        // Data Types
        var dataNumber int = 50;
        var dataString string= "Hello!";
        var dataBool bool= true;

        // Arrays
        var cars = [4] string {"Volvo", "BMW", "Ford", "Mazda"};
        cars[0] = "Opel";

        // Conditions
        if x > y { 
            fmt.Print("False") 
        }

        if x == y { 
            fmt.Print("Draw") 
        } else if x > y { 
            fmt.Print("Yes") 
        } else { 
            fmt.Print("No") 
        }

        //Switch case
        switch day { 
            case (1): 
                fmt.Print("Saturday");
            case (2): 
                fmt.Print("Sunday"); 
            default : 
                fmt.Print("Weekday");
        }

        // Loop
        for i:=0; i < 6; i++ { 
            if i==4 { //  Jump directly to the next value if i = 4
                continue;
            }
            if i==5 {  // Stop the loop if i = 5
                break;
            }
        }

        myFunction("Nguyen Xuan Thao");

    }
