// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package app

import "log"

func PrintAdminRoutes(baseURL, adminKeyShutdown string) {
	log.Printf("shutdown server: %s/admin/shutdown-server/%s\n", baseURL, adminKeyShutdown)
}
