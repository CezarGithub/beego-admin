package i18n

import "quince/utils/i18n"

//Load languages files
func init() {
	i18n.AddLocaleFiles("modules/crm/static/i18n/")
}
