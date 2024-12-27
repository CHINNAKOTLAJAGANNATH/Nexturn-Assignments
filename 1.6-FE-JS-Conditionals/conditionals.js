//Conditional statements

//if statements
//if statements executes a block of code if a specified condition is true

var age=20;
if(age >= 18)
{
    console.log("You are an adult");
}
if (age<18)
{
    console.log("You are a minor");

}

//if-else statements
//if-else statements executes a block of code if a specified
//condition is true and another block of code if that condition if false

var user_age=12;
if(user_age>=18)
{
    console.log("You are an adult");
}
else
{
    console.log('You are a minor');
}

// Login Validation using if..else block
// if user is logged in, display welcome message
// if user is not logged in, display login message
var isUserLoggedIn=false;
if(isUserLoggedIn==true)
{
    console.log("Welcome, User!");
}
else{
    console.log("Please login to continue");
}

//if-else-if statements
//if-else-if statements executes a block of code if a specified
//condition is true, another block of code if that condition is false
// and aother block of code if another condition is true

var user_age=22;
if(user_age > 18)
{
    console.log("You are an adult");
}
else if(user_age < 18)
{
    console.log("You are a minor");
}
else
{
    console.log("You are 18 years old");
}

// Exercise: Grading System using if-else-if statement
// if marks are greater than 90, display A grade
// if marks are greater than 80, display B grade
// if marks are greater than 70, diaplay C grade
// otherwise, display D grade

var marks=90;
if (marks > 90)
{
    console.log("A Grade");  
}
else if(marks > 80)
{
    console.log("B Grade");
}
else if(marks > 70)
{
    console.log("C Grade");
}
else
{
    console.log("D Grade");
}

//Check Order status Example using if-else-if statements
//if order status is pending, display order is pending
// if order status is shipped, display order is shipped
// if order status is delivered, diaplay order is delivered
// otherwise, diaplay order status is unknown

var order_status="shipped";
if(order_status=="pending")
{
    console.log("Order is pending");
}
else if(order_status=="shipped")
{
    console.log("Order is shipped");
}
else if(order_status=="delivered")
{
    console.log("Order is delivered");
}
else
{
    console.log("Order status is unknown");
}

//Switch Case

var order_status="shipped";
switch(order_status)
{
    case "pending":
        console.log("Order is pending");
        break
    case "shipped":
        console.log("Order is shipped");
        break
    case "delivered":
        console.log("Order status is unknown");
        break
}

//Exercise: Grading System using switch case
//if marks are greater than 90, display A grade
//if marks are greater than 80, display B grade
//if marks are greater than 70, display C grade
//otherwise, diaplay D grade

var marks=98;
switch(true)
{
    case marks > 90:
        console.log("A Grade");
        break
    case marks > 80:
        console.log("B Grade");
        break
    case marks > 70:
        console.log("C Grade");
        break
    default:
        console.log("D Grade");
        break

}

// Ternary Operator (Short Hand if-else)
 var age = 20;
 if (age>=18)
 {
    console.log("You are an adult");
 }
 else
 {
    console.log("You are a minor");
 }

 // Ternary operator
var age=17;
var result=(age >= 18)? "You are an adult" : "You are a minor";
console.log(result);

// Syntax of Ternary Operator
// Condition ? true-expression : false-expression

// Exercise: Divisibility by 5 using Ternary Operator
// if number is divisible by 5, display divisible by 5
// otherwise, display not divisible by 5
var number=59;
var result = number%5==0 ? "Divisible by 5" : "Not Divisible by 5";
console.log(result);


