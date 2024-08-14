
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