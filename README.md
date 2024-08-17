It seems that there will never be a single programming language. Unified APIs are important to create sdks that can be used in multiple languages.

What we will cover:
- compile a TypeScipt program into WebAssembly, which exports two functions, one to add two numbers and and another to concatenate a string
- call both functions in a separate JavaScript program
- call both functions in a separate Golang program

Gotchas?
Syntatically, strings look like modern objects, such as strings in python or javascript; under the hood, they are treated as byte arrays, such as in C and C++. (See String interpretation)


``` bash
npm init -y
npm install --save-dev typescript
npx tsc --init
```

Set up a basic `tsconfig.json`
``` json
{
  "compilerOptions": {
    "target": "ES6",
    "module": "CommonJS",
    "strict": true,
    "esModuleInterop": true,
    "skipLibCheck": true,
    "forceConsistentCasingInFileNames": true,
    "outDir": "./dist"
  },
  "include": ["src"]
}
```

``` bash
mkdir src
touch src/index.ts
```

Write the following simple script
``` typescript
// src/index.ts
function greet(name: string): string {
    return `Hello, ${name}! Welcome to TypeScript.`;
}

const userName: string = "Jamal";
console.log(greet(userName));
```

Compile it
``` bash
npx tsc
```

Run it
``` bash
node dist/index.js
```

## Set up WebAssembly

Install tools
``` bash
npm install --save-dev assemblyscript
```

Initialize AssemblyScript
``` bash
npx asinit .
```

AssemblyScript is a subset of TypeScript; hence, some code adjustments might be necessary. 

Compile to WebAssembly
``` bash
npm run asbuild
```

Install AssemblyScript loader
``` bash
npm install @assemblyscript/loader
```

In WebAssembly, strings are not handled as high-level objects like in JavaScript. Instead, they are sequences of bytes stored in linear memory. When you return a string from a WebAssembly function, what you get in JavaScript is typically a pointer (an integer) to the location of the string in WebAssembly memory, not the string itself.


## Call From Golang
``` bash
mkdir go
cd go/
go mod init jiaming/webassembly-test
go mod tidy
```


### String Interpretation

The string data in the result array shows a pattern where each character is followed by a 0x00 byte (e.g., H = 0x48, e = 0x65, l = 0x6c). This is typical for UTF-16LE (little-endian) encoding, where each character is represented by two bytes. The second byte in this encoding is usually 0x00 for standard ASCII characters.

Example:

The first few characters of the string in the result array:
H = 0x48 0x00
e = 0x65 0x00
l = 0x6c 0x00
l = 0x6c 0x00
o = 0x6f 0x00
...
This pattern suggests that the string "Hello, Jamal! Welcome..." is stored in memory in a format where each character is two bytes long due to UTF-16 encoding.
