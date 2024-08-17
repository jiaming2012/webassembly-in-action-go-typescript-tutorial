// src/index.ts
function greet(name: string): string {
    return `Hello, ${name}! Welcome to TypeScript.`;
}

function add(a: number, b: number): number {
    return a + b;
}

const userName: string = "Jamal";
console.log(greet(userName));
console.log('The sum of 3 + 4 is ', add(3, 4));