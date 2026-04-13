module.exports = {
  proxy: process.env.BROWSER_SYNC_PROXY || "http://127.0.0.1:8080",
  port: 3000,
  ui: false,
  notify: false,
  open: false,
  reloadDebounce: 300,
  files: [
    "internal/forumsite/templates/*.html",
    "internal/forumsite/templates/*.css",
    "tmp/gobbs-serve",
  ],
  watchOptions: {
    ignoreInitial: true,
  },
};
