// Bind values to names in Monkey
let age = 1;
let name = "Monkey";
let result = 10 * (20 / 2);

// Monkey interpreter also supports arrays and hashes
let myArray = [1, 2, 3, 4, 5];
let thorsten = {"name": "Thorsten", "age": 20};

// Accessing the elements in arrays and hashes is done with index expressions

let value = myArray[0]      // 1
let name = thorsten["name"] // "Thorsten

// The 'let' statements can also be used to bind functions to names

let add = fn(a, b) { return a + b; };

// Implicit returns are also possible, which means we can leave out the return if we want to

let add = fn(a, b) { a + b; };
let result = add(1, 2);

// Complex function with recursive calls

let fibonacci = fn(x) {
    if (x == 0) {
        0
    } else {
        if (x == 1) {
            1
        } else {
            fibonacci(x - 1) + fibonacci(x - 2);
        }
    }
};

// Monkey also supports a special type of function, called higher order functions
// These are functions that take other functions as arguments
// Functions in Monkey are just values, like integers or strings
// This feature is called ‘first class functions’

let twice = fn(f, x) {
    return f(f(x));
};

let addTwo = fn(x) {
    return x + 2;
};

let result = twice (addTwo, 2); // 6