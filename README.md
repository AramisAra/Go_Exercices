<head>
  <link rel="stylesheet" href="READMEstyle.css">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Madimi+One&display=swap" rel="stylesheet">
  <link rel="preconnect" href="https://fonts.googleapis.com">
  <link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
  <link href="https://fonts.googleapis.com/css2?family=Madimi+One&family=Noto+Sans+KR:wght@100..900&display=swap" rel="stylesheet">

</head>
<body>
  <h1> This repo is the exericices and project I did to learn to use Golang </h1>

<section class="Startsection">
  <h2> Start </h2>
  <p class="p"> I started with simple task to level and fully understand the go syntx with the use of to project of task that create functions to allow to use the language</p>
  <p class="p"> Files that fall under this level are:</p>
  <ul>
      <li> functions_nested_loops </li>
      <li> variables_if_else_while </li>
  </ul>
  <h3>This is an example of the first section's code level</h3>


  <code class="code">
    
    package main

    func print_comb5() {
    var num_a int
    var num_b int

    for num_a = 0; num_a < 99; num_a++ {
      for num_b = num_a + 1; num_b < 100; num_b++ {
        _putchar(false, '1', (num_a / 10))
        _putchar(false, '1', (num_a % 10))
        _putchar(true, ' ', 1)
        _putchar(false, '1', (num_b / 10))
        _putchar(false, '1', (num_b % 10))
          
        if (num_a != 98 || num_b != 99) {
          _putchar(true, ',', 1)
          _putchar(true, ' ', 1)
        }
      }
    }
    _putchar(true, '\n', 1)
    }
  <p> This code can be found in the variables_if_else_while dir it the problem print_comb5.</p></code>
<p class="p"> These are the result of the of the code above. The test case is in the main.go file <p>
<pre class="result">
    00 01, 00 02, 00 03, 00 04, 00 05, 00 06, 00 07, 00 08, 00 09, 00 10, 00 11, 00 12, 00 13, 00 14, 00 15, 00 16, 00 17, 00 18, 00 19, 00 20, 00 21, 00 22, 00 23, 00 24, 00 25, 00 26, 00 27, 00 28, 00 29, 00 30, 00 31, 00 32, 00 33, 00 34, 00 35, 00 36, 00 37, 00 38, 00 39, 00 40, 00 41, 00 42, 00 43, 00 44, 00 45, 00 46, 00 47, 00 48, 00 49, 00 50, 00 51, 00 52, 00 53, 00 54, 00 55, 00 56, 00 57, 00 58, 00 59, 00 60, 00 61, 00 62, 00 63, 00 64, 00 65, 00 66, 00 67, 00 68, 00 69, 00 70, 00 71, 00 72, 00 73, 00 74, 00 75, 00 76, 00 77, 00 78, 00 79, 00 80, 00 81, 00 82, 00 83, 00 84, 00 85, 00 86, 00 87, 00 88, 00 89, 00 90, 00 91, 00 92, ... 94 98, 94 99, 95 96, 95 97, 95 98, 95 99, 96 97, 96 98, 96 99, 97 98, 97 99, 98 99
  </pre>
  <p> For more overall examples of the two project go to the folders with have their own README's explain what the code does </p>
</section>
<section class="API">
<h2> API </h2>
<p> In this section I am taking the time to full build and understand API </p>
<h3> First Step </h3>
<p> Understanding have go and gin work together to connect end points </p>
</section>
</body>