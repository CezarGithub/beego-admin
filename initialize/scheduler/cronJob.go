package scheduler

type Interval struct {
	//0-6 - repeat each week
	DayOfWeek int `orm:"column(day_of_week);default(-1)" description:"0-6 - repeat each week" json:"day_of_week" i18n:"admin.cron.dayOfWeek"`
	//1-12
	MonthOfTheYear int `orm:"column(month_of_year);default(-1)" description:"1-12" json:"month_of_year" i18n:"admin.cron.monthOfTheYear"`
	//1-31 - in combination with MonthOfTheYear
	DayOfTheMonth int `orm:"column(day_of_month);default(-1)" description:"1-31 - in combination with MonthOfTheYear" json:"day_of_month" i18n:"admin.cron.dayOfMonth"`
	// 0-23 - start hour
	Hour int `orm:"column(hour);default(-1)" description:" 0-23 - start hour" json:"hour" i18n:"admin.cron.hour"`
	//0-59 - start hour + minutes
	Minute int `orm:"column(minute);default(-1)" description:"0-59 - start hour + minutes" json:"minute" i18n:"admin.cron.minute"`
	
	//0-23
	//RepeatHours int `orm:"column(repeat_hours);default(-1)" description:"0-23" json:"repeat_hours" i18n:"admin.cron.repeatH"`
	//0-59 if both equal zero RunOnce will be set to true
	//RepeatMinutes int `orm:"column(repeat_minutes);default(-1)" description:"0-59 if both equal zero RunOnce will be set to true" json:"repeat_minutes" i18n:"admin.cron.repeatM"`
}

type cronJob struct {
	Name        string //unique
	Description string
	Module      string
	Func        func() error
}
