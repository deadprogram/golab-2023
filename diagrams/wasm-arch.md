```mermaid
flowchart
subgraph Browser
    subgraph "V8"
    HTML-->Javascript;
    Javascript-->JSShim[Javascript WASM Loader];
    Javascript-->JSRuntime[Browser JS Runtime];
    end
    subgraph "WebAssembly engine"
    WASM[Compiled .WASM file]-->WASMRuntime[Browser WASM Runtime];
    end
    JSShim-->WASM
end
```
