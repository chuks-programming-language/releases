const http = require("http");

const server = http.createServer((req, res) => {
  if (req.url === "/") {
    res.writeHead(200, { "Content-Type": "text/plain" });
    res.end("ok");
  } else if (req.url === "/json") {
    res.writeHead(200, { "Content-Type": "application/json" });
    res.end('{"status":"ok"}');
  } else {
    res.writeHead(404);
    res.end("not found");
  }
});

server.listen(9004);
