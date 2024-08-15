// The entry file of your WebAssembly module.

export function add(a: i32, b: i32): i32 {
  return a + b;
}

export function greet(name: string): string {
  return `Welcome to AssemblyScript, ${name}`;
}

// AssemblyScript: Exporting an allocate function
export function allocate(size: i32): i32 {
  // Implementation of memory allocation
  // For simplicity, this could just return a pointer to a block of memory
  return heap.alloc(size) as i32;
}