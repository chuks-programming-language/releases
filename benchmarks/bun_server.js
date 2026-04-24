const server = Bun.serve({
  port: 9003,
  fetch(req) {
    const url = new URL(req.url);
    if (url.pathname === "/") {
      return new Response("ok");
    }
    if (url.pathname === "/json") {
      return Response.json({ status: "ok" });
    }
    return new Response("not found", { status: 404 });
  },
});
