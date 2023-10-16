// Simplified code for illustration purposes
// Import necessary packages
// ...

func main() {
    if err := buildStaticSite(); err != nil {
        log.Fatalf("Error: %v", err)
    }
}

func buildStaticSite() error {
    // Parse the Markdown file once
    markdownContent, err := os.ReadFile(readmePath)
    if err != nil {
        return fmt.Errorf("read markdown file: %w", err)
    }

    markdownHTML, err := markdown.ToHTML(markdownContent)
    if err != nil {
        return fmt.Errorf("convert markdown to HTML: %w", err)
    }

    // Process categories and links concurrently
    categories, err := extractCategories(markdownHTML)
    if err != nil {
        return fmt.Errorf("extract categories: %w", err)
    }

    // Render categories concurrently
    var categoryErrors []error
    var wg sync.WaitGroup

    for _, category := range categories {
        wg.Add(1)
        go func(category Category) {
            defer wg.Done()

            categoryDir := filepath.Join(outDir, category.Slug)
            if err := mkdirAll(categoryDir); err != nil {
                categoryErrors = append(categoryErrors, err)
                return
            }

            categoryIndexFilename := filepath.Join(categoryDir, "index.html")
            if err := renderCategory(category, categoryIndexFilename); err != nil {
                categoryErrors = append(categoryErrors, err)
                return
            }
        }(category)
    }

    wg.Wait()

    // Handle errors from concurrent category rendering
    if len(categoryErrors) > 0 {
        return categoryErrors[0] // Return the first error
    }

    // Render the main index file
    if err := renderIndex(markdownHTML, outIndexFile); err != nil {
        return fmt.Errorf("render index: %w", err)
    }

    // ... Continue with other tasks
    return nil
}

// Separate functions for rendering, error handling, and concurrency
// ...

