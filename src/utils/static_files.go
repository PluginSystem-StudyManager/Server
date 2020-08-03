package utils

const STATIC_PATH = "../dist/"

func StaticFile(path string) string {
	return STATIC_PATH + path
}
