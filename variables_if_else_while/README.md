<h1> This project is my process of learning go </h1>
<h2> The Variables if else while </h2>
<p> This folder teach the use of if, if else, and while statements </p>
<pre>

        package main
        import (
            "fmt"
        )

        func positive_or_negative() {
        R_number := RandInt(-30000, 30000)
        if R_number > 0 {
            fmt.Printf("%d is positive\n", R_number)
        } else if R_number < 0 {
            fmt.Printf("%d is negative\n", R_number)
        } else {
            fmt.Printf("%d is zero\n", R_number)
        }
        }
</pre>
<p> This is the first task where make a function that can check if it's a positive, negative, or zero </p>
# Go_Exercices
