```mermaid
flowchart
subgraph cubelife
    Itsybitsy--SPI-->Panel1[HUB75 #1]
    subgraph panels
    Panel1-->Panel2[HUB75 #2]
    Panel2-->Panel3[HUB75 #3]
    Panel3-->Panel4[HUB75 #4]
    Panel4-->Panel5[HUB75 #5]
    Panel5-->Panel6[HUB75 #6]
    end
    Panel6--SPI-->Itsybitsy
end
```
