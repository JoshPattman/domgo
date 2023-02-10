const go = new Go();

WebAssembly.instantiateStreaming(fetch(document.currentScript.getAttribute('bin')+".wasm"), go.importObject)
    .then((result) => {
        go.run(result.instance);
    });