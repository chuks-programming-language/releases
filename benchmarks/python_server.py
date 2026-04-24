import http.server
import socketserver
import json
import os

class Handler(http.server.BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/":
            self.send_response(200)
            self.send_header("Content-Type", "text/plain")
            self.end_headers()
            self.wfile.write(b"ok")
        elif self.path == "/json":
            self.send_response(200)
            self.send_header("Content-Type", "application/json")
            self.end_headers()
            self.wfile.write(b'{"status":"ok"}')
        else:
            self.send_response(404)
            self.end_headers()
            self.wfile.write(b"not found")

    def log_message(self, format, *args):
        pass  # Suppress request logging for benchmark

if __name__ == "__main__":
    with socketserver.TCPServer(("", 9005), Handler) as httpd:
        httpd.serve_forever()
