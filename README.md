## Usage
### Copy JS files
Copy the JS files from the `/js/` directory of this repo to a directory of your web app

### Build wasm binary
`GOOS=js GOARCH=wasm go build -o path/to/output/wasm/binary .`

### Embed script into the page
```html
<!--This only needs to be in the top of the page once-->
<script src="wasm_exec.js"></script>

<!--For each script embed this-->
<script src="wasm.js" bin="<path/to/wasm/binary>"></script>
```