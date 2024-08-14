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

console.log(greetString);
