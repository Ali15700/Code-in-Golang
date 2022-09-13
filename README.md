
# Getting started with Golang
Today was my first day to learn Go Language.I shared all the details what I learned today. 

Golang is the most common tool among DevOps professionals.The popular DevOps tools are written in Go, such as Docker, and even the open-source container orchestration system Kubernetes is built on go.

**Golang** has comparitively fast compile time and simple while **C/C++**(My Previous Programming Language) has complex system type and slow compile time.












## Installation

👉 Download Go [Click here](https://go.dev/doc/install)

👉 Open MSI file you downloaded from the above link ☝️

👉 By default, The installer will install go to Program files or Program file(x86)
 
👉 Open VS Code then go to Extensions and install *Go extension* 

👉 create file with **.go** extension and you will see popups to install go dependencies in folder.Click on all install

## Key Features in Golang

👉 Simple and minimalist

👉 Fast Compile time

👉 Garbage Collected

👉 Built-in concurrency

## Data Types in Go

Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.

This is how you can use data types
#### Different datatypes example

  var a bool = true     // Boolean

  var b int = 5         // Integer

  var c float32 = 3.14  // Floating point number

  var d string = "Hi!"  // String


## Switch in Go

The difference in **switch statement in GO** and other PL's is that it only runs the matched case so it does not need a break statement.

#### Working 

👉 *Switch Expression*'s value  is compared with the values of each case

👉 If it matches with corresponding value then associated code block will be executed otherwise it will run **default** code block.

👉 **All the case values should have the same type as the switch expression** otherwise it will through an **error**
## IF ELSE Condition

A condition can be either **true** or **false**

👉 **If** code block will be always executed if condition is **true** otherwise **else**  code block will be executed.

👉 **else if** will be used when you have to check multiple conditions and executed **true** condition.

## Loops

👉 When you have to run same piece code over and over then you use loops.

👉 Each execution of a loop is called an iteration.

👉 **For** loop is the only loop in *GO*
## Function in Go

Function is used to repeatedly call the same block of code.

Function will be executed when we call them.

### Types of Function

👉 Create/Call Function [Click here](https://www.w3schools.com/go/go_functions.php)


👉 Parameters/Arguments [Click here](https://www.w3schools.com/go/go_function_parameters.php)


👉 Return Value from Function [Click here](https://www.w3schools.com/go/go_function_returns.php)


👉 Recursive Functions [Click here](https://www.w3schools.com/go/go_function_recursion.php)

## Struct in Go

👉 Struct method is used to store multiple values of different data types into a single variable.

👉 Dot operator (.) b/w the variable name and member is used to access any member of structure

👉 You can also pass stuctname as an argument in function
## Maps in Go

👉 Maps store data values with keys pairs

👉 A map is an unordered and changeable collection that does not allow duplicates.

👉 Maps hold references to an underlying hash table.
# Web Dev in Go

I've heard Go has excellent built-in support for web dev.

#### lets explore together 😎




## Basic Https 

👉 Registers the handler function for the given pattern in the DefaultServeMux.

👉 ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections

👉 ServeMux is an HTTP request multiplexer. It matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

👉 Handler function that responds to client http requests
## Template Parsing

👉 template ParseFiles(Filename) will used to add Path of the file.


### Inserting Data 

👉 {{.}}	Renders the root element

👉 {{.Attribute}}	Renders the Attribute of struct that you send in ExecuteTemplate function.


## Range
👉 {{range .List}} {{.}} {{end}}	Loops over all “List” field and renders each using {{.}}

#### Range with Indexing
👉 {{range $index, $element := .}}

👉 First Argument always be index and next Argument will be Value/element.

## Template Nesting

👉 In Nested Template you will define a template into another template.

👉 For further clarification you can see code on nested template.

## Gin

👉 Gin is a web framework written in Go (Golang

👉 Gin allows you to build web applications and microservices in Go. It contains a set of commonly used functionalities (e.g., routing, middleware support, rendering, etc.)
## Deployment

To run these projects
#### Basic codes
```bash
  go run findAge.go
```

```bash
  go run switchDays.go
```

```bash
  go run largestNumber.go
```
```bash
  go run empolyeeData.go
```
```bash
  go run carInfo.go
```
#### Web Dev codes
```bash
  go run basicHttps.go
```
```bash
  go run main.go    // To run the template parsing code
```
```bash
  go run insertData.go    
```
```bash
  go run range.go    
```
```bash
  go run indexrange.go    
```
```bash
  go run nestedtemp.go    
```
### Output

```bash
PS C:\Users\Beast\Desktop\MyProjects\go> go run findAge.go
Enter the year of birth:2000
Your age is: 22
```
```bash
PS C:\Users\Beast\Desktop\MyProjects\go> go run switchDays.go
Enter the Name of Week Day:Thursday
Test the code
```
```bash
PS C:\Users\Beast\Desktop\MyProjects\go> go run largestNumber.go
Largest Value:4325
```
```bash
PS C:\Users\Beast\Desktop\MyProjects\go> go run empolyeeData.go
Empolyee Name:AliRaza
Empolyee Age:22
Empolyee job:SE
Empolyee salary:7000


-----------Empolyee Details---------
Name:  AliRaza
Age:  22
Job:  SE
Salary:  7000
```
```bash
PS C:\Users\Beast\Desktop\MyProjects\go> go run carInfo.go      
 Details map[Car Owner:Ali Raza Car brand:Bugatti Car model:Chiron Make year:2021]
 ```
```bash
 go run basicHttps.go
```
<a href="https://ibb.co/tQGQjb1"><img src="https://i.ibb.co/tQGQjb1/screenshot.png" alt="screenshot" border="0"  /></a>

```bash
 go run main.go
```
<a href="https://ibb.co/bH6Bz6t"><img src="https://i.ibb.co/TKw8mwy/parseee.png" alt="parseee" border="0"></a>

```bash
 go run insertData.go
```
<a href="https://ibb.co/r6S2F1G"><img src="https://i.ibb.co/Sn8w59X/mobile.png" alt="mobile" border="0"></a>
## 🔗 Connect with me
[![portfolio](https://img.shields.io/badge/my_portfolio-000?style=for-the-badge&logo=ko-fi&logoColor=white)](https://github.com/Ali15700)

[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/ali-reza-2308b3194/)

[![twitter](https://img.shields.io/badge/twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://twitter.com/devilOps66)

