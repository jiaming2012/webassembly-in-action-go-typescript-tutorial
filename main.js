// main.js
import fs from "fs";
import loader from "@assemblyscript/loader";

// Load the WebAssembly module
const wasmModule = loader.instantiateSync(fs.readFileSync("./build/release.wasm"), { /* imports */ });

// Convert the JavaScript string to a WebAssembly string
const namePointer = wasmModule.exports.__newString("Jamal");

const greetPointer = wasmModule.exports.greet(namePointer);

// Convert the WebAssembly string back to a JavaScript string
const greetString = wasmModule.exports.__getString(greetPointer);

// Add 3 and 4
const result = wasmModule.exports.add(3, 4);

console.log('NodeJS - WebAssembly Example');
console.log('[greet] Output:', greetString);
console.log('[add] Result:', result);





