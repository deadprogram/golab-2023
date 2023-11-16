```mermaid
flowchart
subgraph cloudlife
subgraph Multiverse
    HTTPM[HTTP endpoints]-->MC[multiverse.wasm];
end
subgraph Universe
    MC-->HTTPU;
    HTTPU[HTTP endpoints]-->UC[universe.wasm];
    UC-->D[DistributedUniverse];
end
subgraph Spin
KV[Key/Value Store]
end
MC-->KV;
UC-->KV;
end
```
