
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

