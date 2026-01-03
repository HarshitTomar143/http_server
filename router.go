package main

func handleRoute(req Request) string {
	if req.Path == "/"{
		return "<h1>Home Page</h1>"
	}

	if req.Path == "/about" {
		return "<h1>About Page</h1>"
	}

	if req.Path == "/contact"{
		return "<h1>Contact Page</h1>"
	} 

	return "<h1>404 - Not Found</h1>"
}