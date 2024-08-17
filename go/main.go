package main

import (
	"encoding/binary"
	"fmt"
	"os"
	"unicode/utf16"

	"github.com/wasmerio/wasmer-go/wasmer"
)

func readMemoryString(data []byte, offset int32) string {
	var runes []rune
	size := int32(int(offset) + len(data))
	for i := offset; i < size; i += 2 {
		// Assuming data is in UTF-16LE and we're dealing with ASCII characters,
		// the second byte (i+1) should be 0x00, and the first byte (i) is the character.
		if data[i+1] != 0x00 {
			// Handle the case where the second byte is not 0x00
			continue
		}
		if data[i] == 0x00 && data[i+1] == 0x00 {
			// Indicate the end of the string using double null-termination.
			break
		}
		runes = append(runes, rune(data[i]))
	}
	return string(runes)
}

func writeNameToMemory(name string, instance *wasmer.Instance, memory *wasmer.Memory, allocateFn func(...interface{}) (interface{}, error)) (int32, error) {
	utf16CodeUnits := utf16.Encode([]rune(name))

	// `allocate` function needs to be implemented in your AssemblyScript code
	// It should allocate enough space for the input string and return a pointer to the start of the block
	inputPointer, err := allocateFn(len(utf16CodeUnits) * 2)
	if err != nil {
		return 0, fmt.Errorf("failed to allocate memory for the input string: %w", err)
	}

	inputOffset := inputPointer.(int32)
	memoryData := memory.Data()[inputOffset:]

	for i, codeUnit := range utf16CodeUnits {
		binary.LittleEndian.PutUint16(memoryData[i*2:], codeUnit)
	}

	return inputOffset, nil
}

func main() {
	// Read the WebAssembly module file
	wasmBytes, err := os.ReadFile("../build/release.wasm")
	if err != nil {
		panic(fmt.Errorf("failed to read the WebAssembly module file: %v", err))
	}

	// Create an engine
	engine := wasmer.NewEngine()

	// Create a store
	store := wasmer.NewStore(engine)

	// Compile the module
	module, err := wasmer.NewModule(store, wasmBytes)
	if err != nil {
		panic(fmt.Errorf("failed to compile module: %v", err))
	}

	// Instantiate the module
	var instance *wasmer.Instance
	importObject := wasmer.NewImportObject()
	envNamespace := map[string]wasmer.IntoExtern{
		"abort": wasmer.NewFunction(
			store,
			wasmer.NewFunctionType(wasmer.NewValueTypes(wasmer.I32, wasmer.I32, wasmer.I32, wasmer.I32), wasmer.NewValueTypes()),
			func(args []wasmer.Value) ([]wasmer.Value, error) {
				messagePtr := args[0].I32()
				fileNamePtr := args[1].I32()
				lineNumber := args[2].I32()
				columnNumber := args[3].I32()

				// Ensure `instance` and its memory are accessible here. You might need to adjust this part.
				// For demonstration, assuming `instance` is accessible and has a method `Memory()` to get its memory.
				// Convert the pointer to an offset we can use in Go
				memory, err := instance.Exports.GetMemory("memory")
				if err != nil {
					return nil, fmt.Errorf("failed to get the memory: %w", err)
				}

				// Read strings from memory
				message := readMemoryString(memory.Data(), messagePtr)
				fileName := readMemoryString(memory.Data(), fileNamePtr)

				fmt.Printf("Wasm called abort! Message: %s, File: %s, Line: %d, Column: %d\n", message, fileName, lineNumber, columnNumber)

				return []wasmer.Value{}, nil
			},
		),
	}

	importObject.Register("env", envNamespace)

	// Instantiate the module
	instance, err = wasmer.NewInstance(module, importObject)
	if err != nil {
		panic(fmt.Errorf("failed to instantiate the module: %v", err))
	}

	// Convert the pointer to an offset we can use in Go
	memory, err := instance.Exports.GetMemory("memory")
	if err != nil {
		panic(fmt.Errorf("failed to get the memory: %v", err))
	}

	// Allocate memory for the input string in the WebAssembly module
	allocate, err := instance.Exports.GetFunction("allocate")
	if err != nil {
		panic(fmt.Errorf("failed to get the `allocate` function: %w", err))
	}

	inputOffset, err := writeNameToMemory("Jamal", instance, memory, allocate)
	if err != nil {
		panic(fmt.Errorf("failed to write the input string to memory: %v", err))
	}

	// Call the `greet` function
	greet, err := instance.Exports.GetFunction("greet")
	if err != nil {
		panic(fmt.Errorf("failed to get the `greet` function: %v", err))
	}

	outputPointer, err := greet(inputOffset)
	if err != nil {
		panic(fmt.Errorf("failed to call the `greet` function: %v", err))
	}

	outputString := readMemoryString(memory.Data(), outputPointer.(int32))

	// Deallocate the memory
	deallocate, err := instance.Exports.GetFunction("deallocate")
	if err != nil {
		panic(fmt.Errorf("failed to get the `deallocate` function: %v", err))
	}

	_, err = deallocate(inputOffset)
	if err != nil {
		panic(fmt.Errorf("failed to deallocate the input string: %v", err))
	}

	fmt.Println("Golang - WebAssembly Example")

	fmt.Println("[greet] Output:", outputString)

	// Call the `add` function
	add, err := instance.Exports.GetFunction("add")
	if err != nil {
		panic(fmt.Errorf("failed to get the `add` function: %v", err))
	}

	result, err := add(3, 4)
	if err != nil {
		panic(fmt.Errorf("failed to call the `add` function: %v", err))
	}

	fmt.Println("[add] Result:", result)
}
