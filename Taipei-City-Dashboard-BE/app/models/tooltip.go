package models

import "fmt"

func FetchTooltipDataByDistrict(district string) (map[string]string, error) {
	type TooltipRow struct {
        Category string
        Total    int
    }

    var results []TooltipRow

    err := DBDashboard.
        Raw(`
            SELECT
                category,
                SUM(value::int) AS total
            FROM friendly_store,
            LATERAL (
                VALUES
                    ('跨文化友善', f_lang),
                    ('跨文化友善', f_wifi),
                    ('跨文化友善', f_muslim),
                    ('多元共榮友善', f_sex),
                    ('多元共榮友善', f_lactation),
                    ('多元共榮友善', f_mc),
                    ('日常照護友善', f_pay),
                    ('日常照護友善', f_veg),
                    ('日常照護友善', f_toilet),
                    ('空間可及友善', f_moblie),
                    ('空間可及友善', f_acc),
                    ('空間可及友善', f_bike)
            ) AS categories(category, value)
            WHERE zone = ?
            GROUP BY category
            ORDER BY category
        `, district).
        Scan(&results).Error

		if err != nil {
			return nil, err
		}
	
		result := make(map[string]string)
		for _, row := range results {
			result[row.Category] = fmt.Sprintf("%d 個", row.Total)
		}
	
		return result, nil
}