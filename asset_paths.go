package config

var assetPaths = map[string]string{
	"/css/admin.css": "/css/admin_e66ccbafd592363984bc3414eb03dd71.css",
	"/css/app.css": "/css/app_dda5ca159f88792dd8e31faeb8b8f155.css",
	"/css/lib.css": "/css/lib_3428b0ac4cfefcf1dd9b15164d5e19b4.css",
	
	"/js/admin.js": "/js/admin_57c1f5cfe8e71f7c8e4fc9f624fb53dc.js",
	"/js/app.js": "/js/app_560ae07a24b8302f3d4700ade0707561.js",
	"/js/init.js": "/js/init_7999d01c88456e63af071e1e085d5a5b.js",
	
}

func AssetPath(name string) string {
	if path, ok := assetPaths[name]; ok {
		return path
	}

	return name
}
