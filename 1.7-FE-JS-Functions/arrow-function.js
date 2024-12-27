
// Arrow Function
// Explaination: Arrow function is known as inline function.
// It is a shorter way to write a function.

// Arrow Function without parameters
var greet = () => console.log('Hello, World!');

greet();

// Arrow Function with return statement
var sayHello = () => 'Say Hello';

console.log(sayHello());

// Arrow Function with one parameters
var greet = name => 'Hello, ' + name;

console.log(greet('John Doe'));

// Arrow Function with two parameters
var getFullName = (firstName, lastName) => 'Welcome ' + firstName + ' ' + lastName;

console.log(getFullName('John', 'Doe')); // Welcome John Doe

// Arrow Function with Multiple Statements
var calcArea = (width, height) => {
var area = width * height;
return area;
}

console.log(calcArea(10, 20)); // 200

// Exercise
// Create an arrow function to calculate area of a circle and return the result
// Arrow Function with Arrays
var circlearea = (radius)=>{
    var area=3.14*radius*radius;
    return area;
    }
var numbers = [1, 2, 3, 4, 5];
// Using map() method
var doubleNumbers = numbers.map(number => number * 2);
// print the result
console.log(doubleNumbers); // [2, 4, 6, 8, 10]

/* ************************************ */

// Arrow Function with Arrays
var numbers = [1, 2, 3, 4, 5];
// Using filter() method
var evenNumbers = numbers.filter(number => number % 2 === 0);
// print the result
console.log(evenNumbers); // [2, 4]

/* ************************************ */

// Arrow Function with Arrays
var numbers = [1, 2, 3, 4, 5];
// Using reduce() method
var sum = numbers.reduce((total, number) => total + number, 0);
// print the result
console.log(sum); // 15

//Exercise-- You have an array of strings representing people's names. 
// You want to create a new array that contains only names longer than 4 characters, 
// and convert them all to uppercase.
var names = ['John', 'Jane', 'Bob', 'Alice', 'Charlie'];
var longNames = names.filter(name => name.length > 4)
                    .map(name => name.toUpperCase());
console.log(longNames);
