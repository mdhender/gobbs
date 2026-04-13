module.exports = {
  host: process.env.BROWSER_SYNC_HOST || "gobbs.test",
  proxy: process.env.BROWSER_SYNC_PROXY || "http://gobbs.test:8080",
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
