import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpExchange;
import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;

public class JavaServer {
    public static void main(String[] args) throws IOException {
        HttpServer server = HttpServer.create(new InetSocketAddress(9002), 0);

        server.createContext("/", exchange -> {
            String path = exchange.getRequestURI().getPath();
            byte[] response;
            int code;
            String contentType;

            if (path.equals("/")) {
                code = 200;
                contentType = "text/plain";
                response = "ok".getBytes();
            } else if (path.equals("/json")) {
                code = 200;
                contentType = "application/json";
                response = "{\"status\":\"ok\"}".getBytes();
            } else {
                code = 404;
                contentType = "text/plain";
                response = "not found".getBytes();
            }

            exchange.getResponseHeaders().set("Content-Type", contentType);
            exchange.sendResponseHeaders(code, response.length);
            try (OutputStream os = exchange.getResponseBody()) {
                os.write(response);
            }
        });

        server.setExecutor(java.util.concurrent.Executors.newFixedThreadPool(
            Runtime.getRuntime().availableProcessors()
        ));
        server.start();
    }
}
