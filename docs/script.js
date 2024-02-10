"use strict";

window.addEventListener('load', (event) => {
  const go = new Go();

  WebAssembly
    .instantiateStreaming(fetch('/scram-sha-256/encrypt.wasm'), go.importObject)
    .then((result) => go.run(result.instance));
});
