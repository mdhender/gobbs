// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package app

import (
	"net/http"
	"os"
)

func Routes(baseURL, assets, components string, adminShutdownKey string, stop chan os.Signal) http.Handler {
	r := NewRouter()

	r.Get("/about-us", servePage(assets, "/about-us.html"))
	r.Get("/admin/shutdown-server/{key}", serveAdminShutdownServer(adminShutdownKey, stop))

	r.Get("/index.php", getIndex(components))
	r.Get("/showthread.php", getShowthread(components))
	r.Get("/task.php", getTasks())

	// serve assets, redirecting "/" to the index page.
	r.Get("/", serveAssets(assets, true, "/index.php"))

	return r
}
