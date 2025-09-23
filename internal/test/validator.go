package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// toCamelCase converts a string from CamelCase to a JSON-compatible camelCase format.
// It correctly handles Go's convention for acronyms in variable names.
func toCamelCase(s string) string {

	if s == "" {
		return ""
	}

	runes := []rune(s)

	// Check for a full-word acronym (e.g., "ID", "URL", "GPS") and convert to lowercase.
	// This handles the special case where a field name is only an acronym.
	isAcronym := true
	for _, r := range runes {
		if !unicode.IsUpper(r) {
			isAcronym = false
			break
		}
	}
	if isAcronym {
		return strings.ToLower(s)
	}

	// For names that start with an uppercase letter, convert the first letter to lowercase.
	// This covers "FirstName" -> "firstName" and names with acronyms in the middle,
	// such as "UserID" -> "userId" and "AvatarURL" -> "avatarURL".
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run validator.go <path/to/models>")
		os.Exit(1)
	}

	// مسیر دایرکتوری را از آرگومان خط فرمان دریافت می‌کنیم.
	dirPath := os.Args[1]
	fmt.Printf("Starting validation for directory: %s\n", dirPath)

	// شروع به بررسی فایل‌ها می‌کنیم.
	err := filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(d.Name(), ".go") {
			if err := validateFile(path); err != nil {
				fmt.Printf("Validation failed for %s:\n%v\n", path, err)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("An error occurred during directory traversal: %v\n", err)
	}

	fmt.Println("Validation finished.")
}

func validateFile(filePath string) error {

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	var validationErrors []string

	// از ast.Inspect برای پیمایش درخت نحو انتزاعی (AST) استفاده می‌کنیم.
	ast.Inspect(node, func(n ast.Node) bool {

		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true // به پیمایش ادامه می‌دهیم.
		}

		// بررسی می‌کنیم که آیا نوع داده یک struct است.
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true // به پیمایش ادامه می‌دهیم.
		}

		structName := typeSpec.Name.Name

		// فیلدهای struct را بررسی می‌کنیم.
		for _, field := range structType.Fields.List {
			if field.Names == nil || len(field.Names) == 0 {
				continue // فیلدهایی که نام ندارند (مانند embedding) را نادیده می‌گیریم.
			}

			fieldName := field.Names[0].Name

			// قانون ۱: نام فیلد باید با حرف بزرگ شروع شود.
			if !unicode.IsUpper(rune(fieldName[0])) {
				validationErrors = append(validationErrors,
					fmt.Sprintf("ERROR: Struct '%s', Field '%s': Field name must start with an uppercase letter.", structName, fieldName))
			}

			// قانون ۲: تگ JSON باید با نام فیلد مطابقت داشته باشد.
			if field.Tag != nil {
				tag := strings.Trim(field.Tag.Value, "`")
				tagParts := strings.Split(tag, " ")

				if len(tagParts) > 0 {
					jsonTagPart := tagParts[0]

					if strings.HasPrefix(jsonTagPart, "json:\"") {
						jsonTag := strings.TrimPrefix(jsonTagPart, "json:\"")
						jsonTag = strings.TrimSuffix(jsonTag, "\"")

						// Split the tag by comma to handle options like `,omitempty`
						jsonTag = strings.Split(jsonTag, ",")[0]

						expectedTag := toCamelCase(fieldName)
						if jsonTag != expectedTag {
							validationErrors = append(validationErrors,
								fmt.Sprintf("ERROR: Struct '%s', Field '%s': JSON tag '%s' does not match the expected format '%s'.", structName, fieldName, jsonTag, expectedTag))
						}
					}
				}
			}
		}

		return true
	})

	if len(validationErrors) > 0 {
		return fmt.Errorf("\n%s", strings.Join(validationErrors, "\n"))
	}
	return nil
}
