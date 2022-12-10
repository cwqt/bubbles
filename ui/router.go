package UI

type Routes map[string]Route

type Route struct {
	Component FC
	Children  Routes
}

// Private to Create() closure -- its internal model
type routerModel struct {
	Path   []string // active router path
	Routes Routes   // routing structure with factory functions
}

// Public methods
type Router struct {
	GetPath  func() *[]string
	Navigate func(paths ...string)
	Outlet   *Outlet // the root outlet
}

type RouteChange struct {
	Path []string
}

// Router instance -- one per Program
func CreateRouter(routes Routes, initialPath []string) Router {
	router := routerModel{
		Path:   make([]string, 0),
		Routes: routes,
	}

	// Set starting route
	if len(initialPath) > 0 {
		router.Path = append(router.Path, initialPath...)
	}

	return Router{
		Outlet: CreateOutlet(router.Routes, &router.Path, 0),
		GetPath: func() *[]string {
			return &router.Path
		},
		Navigate: func(path ...string) {
			// Copy path into router.Path slice, need to make a
			// new slice to expand original length to accomodate
			router.Path = make([]string, len(path))
			copy(router.Path, path)
		},
	}
}
