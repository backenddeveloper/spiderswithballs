FROM nginx:1.17.0

COPY ["index.html", "wasm_exec.js", "spiderswithballs.wasm"] /usr/share/nginx/html

EXPOSE 80
