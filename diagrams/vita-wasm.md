```mermaid
flowchart LR
    subgraph Application
    direction LR
    subgraph Page
    direction LR
        Form[HTML Form elements]-->Javascript;
        Canvas[HTML Canvas element]-->Javascript;
    end
    subgraph vita.wasm
        Javascript-->main[main.go];
        main-->universe[game.Universe]
    end
    end
```
