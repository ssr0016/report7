package config

import "strconv"

type ConfigPagination struct {
	Pagination PaginationConfig
}

type PaginationConfig struct {
	Page      int
	PageLimit int
}

func ParsePage(pageStr string) int {
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		return 1 // Default to page 1 if invalid or not provided
	}
	return page
}

// Helper function to parse per_page from query parameter
func ParsePerPage(perPageStr string) int {
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		return 10 // Default to 10 per page if invalid or not provided
	}
	return perPage
}
