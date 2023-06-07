// set server host
routemamba.registerServerHost("http://localhost:3000/");
// set meta content
routemamba.registerMetaUrl("components/meta");

routemamba.register_http_routes([
  {
    method: "GET",
    meta_loader: true,
    content_url: "components/home",
    data: {},
    preloader: "loading...",
    error_content: "error",
    http_url_change: false,
    http_url: "/",
  },
  {
    method: "GET",
    meta_loader: true,
    content_url: "components/about",
    data: {},
    preloader: "loading...",
    error_content: "error",
    http_url_change: false,
    http_url: "about",
  },

]);

// Set pages headers
routemamba.register_routes_headers([
  {
    content_url: "components/header",

    preloader: "loading...",
    error_content: "error",
    http_url: ["/", "about"],
  },
]);

routemamba.register_routes_footers([
  {
    content_url: "components/footer",

    preloader: "loading...",
    error_content: "error",
    http_url: ["/", "about"],
  },
]);

routemamba.render();
