package reports

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/khulnasoft/kengine/khulnasoft_utils/log"
	"github.com/khulnasoft/kengine/khulnasoft_utils/utils"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

func secretPDF(ctx context.Context, params utils.ReportParams) (core.Document, error) {

	log := log.WithCtx(ctx)

	data, err := getSecretData(ctx, params)
	if err != nil {
		log.Error().Err(err).Msg("failed to get secrets info")
		return nil, err
	}

	log.Info().Msgf("report id %s has %d records",
		params.ReportID, data.NodeWiseData.RecordCount)

	// get new instance of marato
	m := getMarato()

	// applied filter page
	filtersPage := getFiltersPage(
		data.Title,
		data.ScanType,
		data.AppliedFilters.NodeType,
		fmt.Sprintf("%s - %s", data.StartTime, data.EndTime),
		strings.Join(data.AppliedFilters.SeverityOrCheckType, ","),
		data.AppliedFilters.AdvancedReportFilters.String(),
	)

	// summary table
	summaryPage := getSummaryPage(&data.NodeWiseData.SeverityCount)

	cellStyle := &props.Cell{
		BackgroundColor: &props.Color{Red: 255, Green: 255, Blue: 255},
		BorderType:      border.Full,
		BorderColor:     &props.Color{Red: 0, Green: 0, Blue: 0},
		BorderThickness: 0.1,
	}

	resultHeaderProps := props.Text{
		Size:  10,
		Left:  1,
		Top:   1,
		Align: align.Center,
		Style: fontstyle.Bold,
		Color: &props.Color{Red: 0, Green: 0, Blue: 200},
	}

	// page per scan
	resultPages := []core.Page{}
	for i, d := range data.NodeWiseData.ScanData {

		// skip if there are no results
		if len(d.ScanResults) == 0 {
			continue
		}

		p := page.New()
		p.Add(text.NewRow(10, fmt.Sprintf("%s - Scan Details", i), resultHeaderProps))
		p.Add(row.New(6).Add(
			text.NewCol(1, "No.", resultHeaderProps).WithStyle(cellStyle),
			text.NewCol(2, "Rule Name", resultHeaderProps).WithStyle(cellStyle),
			text.NewCol(3, "File Name", resultHeaderProps).WithStyle(cellStyle),
			text.NewCol(1, "Severity", resultHeaderProps).WithStyle(cellStyle),
			text.NewCol(4, "Matched Content", resultHeaderProps).WithStyle(cellStyle),
			text.NewCol(1, "Part", resultHeaderProps).WithStyle(cellStyle),
		))

		resultRows := []core.Row{}
		for k, v := range d.ScanResults {
			resultRows = append(
				resultRows,
				row.New(15).Add(
					text.NewCol(1, strconv.Itoa(k),
						props.Text{Size: 10, Top: 1, Align: align.Center}).
						WithStyle(cellStyle),
					text.NewCol(2, v.Name,
						props.Text{Size: 10, Top: 1, Align: align.Center, BreakLineStrategy: breakline.DashStrategy}).WithStyle(cellStyle),
					text.NewCol(3, v.FullFilename,
						props.Text{Size: 10, Left: 1, Top: 1, BreakLineStrategy: breakline.DashStrategy}).
						WithStyle(cellStyle),
					text.NewCol(1, v.Level,
						props.Text{Size: 10, Top: 1, Align: align.Center, Style: fontstyle.Bold, Color: colors[v.Level]}).
						WithStyle(cellStyle),
					text.NewCol(4, truncateText(v.MatchedContent, 80),
						props.Text{Size: 10, Left: 1, Top: 1, BreakLineStrategy: breakline.DashStrategy}).
						WithStyle(cellStyle),
					text.NewCol(1, v.Part,
						props.Text{Size: 10, Top: 1, Align: align.Center}).
						WithStyle(cellStyle),
				),
			)
		}
		p.Add(resultRows...)
		resultPages = append(resultPages, p)
	}

	// add all pages
	m.AddPages(filtersPage)
	m.AddPages(summaryPage)
	m.AddPages(resultPages...)

	return m.Generate()
}
