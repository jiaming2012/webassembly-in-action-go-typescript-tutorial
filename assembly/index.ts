// assembly/index.ts

// The entry file of your WebAssembly module.

export function add(a: i32, b: i32): i32 {
  return a + b;
}

export function greet(name: string): string {
  return `Welcome to AssemblyScript, ${name}`;
}

// AssemblyScript: Exporting an allocate function
export function allocate(size: i32): i32 {
  return heap.alloc(size) as i32;
}

// AssemblyScript: Exporting a deallocate function
export function deallocate(ptr: i32): void {
  heap.free(ptr);
}