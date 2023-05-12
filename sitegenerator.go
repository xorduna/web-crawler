package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	rootFolder := "website"

	// Create root folder
	err := os.Mkdir(rootFolder, 0755)
	if err != nil {
		fmt.Println("Failed to create root folder:", err)
		return
	}

	// Create subfolders
	subfolders := []string{"folder1", "folder2", "folder3"}
	for _, folder := range subfolders {
		err = os.Mkdir(filepath.Join(rootFolder, folder), 0755)
		if err != nil {
			fmt.Printf("Failed to create %s folder: %v\n", folder, err)
			return
		}

		// Create main.html file in each subfolder
		mainHTML := generateMainHTML(subfolders, folder)
		err = ioutil.WriteFile(filepath.Join(rootFolder, folder, "main.html"), []byte(mainHTML), 0644)
		if err != nil {
			fmt.Printf("Failed to write main.html file in %s folder: %v\n", folder, err)
			return
		}

		// Create pages in each subfolder
		for i := 1; i <= 3; i++ {
			pageHTML := generatePageHTML(subfolders, folder, i)
			err = ioutil.WriteFile(filepath.Join(rootFolder, folder, fmt.Sprintf("page%d.html", i)), []byte(pageHTML), 0644)
			if err != nil {
				fmt.Printf("Failed to write page%d.html file in %s folder: %v\n", i, folder, err)
				return
			}
		}
	}

	// Create index.html file in the root folder
	indexHTML := generateIndexHTML(subfolders)
	err = ioutil.WriteFile(filepath.Join(rootFolder, "index.html"), []byte(indexHTML), 0644)
	if err != nil {
		fmt.Println("Failed to write index.html file:", err)
		return
	}

	fmt.Println("HTML files generated successfully!")
}

func generateMainHTML(subfolders []string, currentFolder string) string {
	var links string
	for _, folder := range subfolders {
		link := fmt.Sprintf(`<a href="../%s/main.html">Link to %s</a><br>`, folder, folder)
		links += link
	}

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Main Page - %s</title>
</head>
<body>
    <h1>Main Page - %s</h1>
    %s
</body>
</html>
`, currentFolder, currentFolder, links)
}

var FOOTER = `
	<a href="https://www.facebook.com">Facebook</a><br>
	<a href="https://www.twitter.com">Facebook</a><br>
`

func generatePageHTML(subfolders []string, currentFolder string, pageNum int) string {
	var links string
	for _, folder := range subfolders {
		link := fmt.Sprintf(`<a href="../%s/main.html">Link to %s</a><br>`, folder, folder)
		links += link
	}

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Page %d - %s</title>
</head>
<body>
    <h1>Page %d - %s</h1>
    %s
	%s
</body>
</html>
`, pageNum, currentFolder, pageNum, currentFolder, links, FOOTER)
}

func generateIndexHTML(subfolders []string) string {
	var mainLinks string
	for _, folder := range subfolders {
		mainLink := fmt.Sprintf(`<a href="%s/main.html">Link to %s</a><br>`, folder, folder)
		mainLinks += mainLink
	}

	return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Index Page</title>
</head>
<body>
    <h1>Index Page</h1>
    <h2>Main Pages:</h2>
    %s
	%s
</body>
</html>
`, mainLinks, FOOTER)
}
