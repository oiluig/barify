# Barify 🍸
Reduce html,css,js space and generate base64 or qrcode when possible. 

```
go run . <dir>
```

## Features
- Pack every file in one base64 string
- Minify html, js, css
- Decompress with DecompressionStream

## Todo
- rename css classes
- rename js functions/classes
- Compress imgs


## Sources

- https://github.com/PuerkitoBio/goquery
- https://github.com/tdewolff/minify
- https://pkg.go.dev/github.com/skip2/go-qrcode
- https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/Data_URLs
- https://developer.mozilla.org/en-US/docs/Web/API/DecompressionStream

- https://stackoverflow.com/questions/8067546/before-deployment-is-there-tool-to-compress-html-class-attribute-and-css-select
- https://github.com/JPeer264/node-rcs-core#plugins

- https://www.freecodecamp.org/news/reducing-css-bundle-size-70-by-cutting-the-class-names-and-using-scope-isolation-625440de600b

- https://stackoverflow.com/questions/38024631/finding-all-class-names-used-in-html-dom