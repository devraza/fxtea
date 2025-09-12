package main

import (
	"fmt"
	"fxtea/fx"
	"github.com/charmbracelet/bubbles/table"
	"strconv"
	"strings"
)

func quadraticView(m model) string {
	headerContent := fmt.Sprintf("Enter the values for a quadratic in the form %s", codeStyle.Render("ax² + bx + c"))
	helpText := header(headerContent, []string{help("q", "return")})

	arguments := strings.Split(m.TextInput.Value(), " ")
	m.TextInput.Placeholder = "a b c"

	var result string

	if len(arguments) == 3 {
		var coefficients []float64

		for i := range arguments {
			parsed, _ := strconv.ParseFloat(arguments[i], 64)
			coefficients = append(coefficients, parsed)
		}

		roots := fx.Quadratic(coefficients[0], coefficients[1], coefficients[2])

		if roots[0] != roots[1] {
			result = fmt.Sprintf(
				"The roots are %v and %v",
				keywordStyle.Render(roots[0]),
				keywordStyle.Render(roots[1]),
			)
		} else {
			result = fmt.Sprintf(
				"The root is %v",
				keywordStyle.Render(roots[0]),
			)
		}
	}

	content := fmt.Sprintf(
		"%s\n\n%v",
		m.TextInput.View(),
		result,
	)

	return fmt.Sprintf(helpText, content)
}

func poissonView(m model) string {
	helpText := header("Enter the rate and the value of x", []string{help("q", "return")})

	arguments := strings.Split(m.TextInput.Value(), " ")
	m.TextInput.Placeholder = "λ x"

	var result string

	if len(arguments) == 2 {
		lambda, _ := strconv.ParseFloat(arguments[0], 64)
		x, _ := strconv.ParseUint(arguments[1], 10, 64)

		result = fmt.Sprintf(
			"The cumulative probability is: %s",
			keywordStyle.Render(fmt.Sprintf("%.4f", fx.PoissonCD(lambda, x))),
		)
	}

	content := fmt.Sprintf(
		"%s\n\n%v",
		m.TextInput.View(),
		result,
	)

	return fmt.Sprintf(helpText, content)
}

func chiView(m model) string {
	helpText := header("Enter the degrees of freedom and the significance level", []string{help("q", "return")})

	arguments := strings.Split(strings.TrimSpace(m.TextInput.Value()), " ")
	m.TextInput.Placeholder = "ν α"

	var result string

	if len(arguments) == 2 {
		var floatArgs []float64

		for i := range arguments {
			parsed, _ := strconv.ParseFloat(arguments[i], 64)
			floatArgs = append(floatArgs, parsed)
		}

		critical := fx.ChiCritical(floatArgs[0], floatArgs[1])
		result = fmt.Sprintf(
			"The critical value is %v",
			keywordStyle.Render(fx.FormatFloat(critical)),
		)
	}

	content := fmt.Sprintf("%s\n\n%v", m.TextInput.View(), result)

	return fmt.Sprintf(helpText, content)
}

func fibonacciView(m model) string {
	helpText := header("Enter the length of the generated sequence", []string{help("q", "return")})

	m.TextInput.Placeholder = "n"

	parsed, _ := strconv.ParseUint(m.TextInput.Value(), 10, 64)
	sequence := fx.Fibonacci(parsed)

	formattedSequence := ""
	for _, v := range sequence {
		formattedSequence += listStyle.Render(strconv.FormatUint(v, 10)) + " "
	}

	content := fmt.Sprintf("%s\n\n%s", m.TextInput.View(), formattedSequence)

	return fmt.Sprintf(helpText, content)
}

func binarySearchView(m model) string {
	headerText := fmt.Sprintf(
		"%s %s",
		"Enter the query in the format",
		codeStyle.Render("<query> <list>"),
	)
	helpText := header(headerText, []string{help("q", "return")})

	arguments := strings.Split(strings.TrimSpace(m.TextInput.Value()), " ")
	m.TextInput.Placeholder = "3 1 2 3 4 5"

	query, _ := strconv.ParseInt(arguments[0], 10, 64)

	var integerList []int64

	for _, v := range arguments[1:] {
		parsed, _ := strconv.ParseInt(v, 10, 64)
		integerList = append(integerList, parsed)
	}

	integerList = fx.QuickSort(integerList)

	result := fx.BinarySearch(integerList, query)

	sortedText := ""
	if len(integerList) > 0 {
		sortedText = "The sorted list is "
		for _, v := range integerList {
			sortedText += listStyle.Render(strconv.FormatInt(v, 10)) + " "
		}
	}

	resultText := errorStyle.Render("The element is not in the list")
	if result >= 0 {
		resultText = fmt.Sprintf(
			"Element %v found at index %v",
			keywordStyle.Render(strconv.FormatInt(query, 10)),
			resultStyle.Render(strconv.FormatInt(result, 10)),
		)
	}

	content := fmt.Sprintf("%s\n\n%s\n\n%s", m.TextInput.View(), sortedText, resultText)

	return fmt.Sprintf(helpText, content)
}

func pmccView(m model) string {
	headerText := "Add rows to the table"
	helpText := header(headerText, []string{help("enter", "add row to table"), help("q", "return")})

	pmccColumns := []table.Column{
		{Title: "x", Width: 10},
		{Title: "y", Width: 10},
	}

	m.TextInput.Placeholder = "x y"

	pmccTable := table.New(
		table.WithColumns(pmccColumns),
		table.WithFocused(false),
		table.WithHeight(10),
	)

	style := table.DefaultStyles()
	style.Header = tableHeaderStyle
	style.Selected = tableSelectedStyle

	pmccTable.SetStyles(style)

	pmccTable.FromValues(m.PMCCData, " ")

	var resultText string
	rows_str := strings.Split(m.PMCCData, "\n")
	if len(rows_str) < 3 {
		resultText = "Enter the data first"
	} else {
		var x, y []float64
		var err error

		for _, row := range rows_str {
			values := strings.Split(row, " ")
			if len(values) < 2 {
				break
			}
			var valX, valY float64
			valX, err = strconv.ParseFloat(values[0], 64)
			valY, err = strconv.ParseFloat(values[1], 64)
			if err == nil {
				x = append(x, valX)
				y = append(y, valY)
			}
		}

		r, err := fx.PMCC(x, y)

		if err != nil {
			resultText = fmt.Sprintf("ERR: %s", err)
		} else {
			resultText = fmt.Sprintf("r = %.4f", r)
		}
	}

	content := fmt.Sprintf("%s\n\n%s\n\n%s", m.TextInput.View(), resultText, pmccTable.View())

	return fmt.Sprintf(helpText, content)
}
