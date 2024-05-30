package utils

import (
	"errors"
	"fmt"
	"time"
)

// PregnancyPeriod 包含怀孕期信息的结构体
type PregnancyPeriod struct {
	TotalDays      int    // 总天数
	NumberOfWeeks  int    // 总周数
	NumberOfYears  int    // 总年数
	NumberOfMonths int    // 总月数
	NumberOfDays   int    // 剩余天数
	Type           string // 类型，可以是 "age"（年龄）或 "pregnancy"（怀孕期）
}

// WeekInfo 存储每个年龄类型的周刊信息
type WeekInfo struct {
	AgeType     string
	TitleFormat string
}

// Age types
const (
	AgeTypeYunqi   = "yunqi"
	AgeTypeYinger  = "yinger"
	AgeTypeYouer   = "youer"
	AgeTypeXueqian = "xueqian"
)

// WeekInfos 存储不同年龄类型的周刊信息
var WeekInfos = map[int]WeekInfo{
	1: {AgeTypeYunqi, "预产期孕%d周"},
	2: {AgeTypeYinger, "婴儿期%s"},
	3: {AgeTypeYouer, "幼儿期1岁%d个月"},
	4: {AgeTypeXueqian, "学前%d岁%s个月"},
}

// GetWeekURLByBirthday 根据宝宝生日获取相应孕育周刊的 URL 和相关信息
func GetWeekURLByBirthday(babyBirthday time.Time) (map[string]string, error) {
	webServer := ""
	stageID, err := GetTimeLineStageIDByBirthday(babyBirthday)
	if err != nil {
		return nil, err
	}

	weekInfo, ok := WeekInfos[stageID]
	if !ok {
		return nil, errors.New("invalid stageID")
	}

	tmp := stageID - stageIDOffset[weekInfo.AgeType]

	url := fmt.Sprintf("%s/type=%s_%d_1", webServer, weekInfo.AgeType, tmp)
	weeklyTitle := fmt.Sprintf(weekInfo.TitleFormat, tmp)

	ret := map[string]string{
		"url":      url,
		"age_type": weekInfo.AgeType,
		"type":     fmt.Sprintf("%s_%d_1", weekInfo.AgeType, tmp),
		"title":    weeklyTitle,
	}

	return ret, nil
}

// GetTimeLineStageIDByBirthday 根据宝宝生日计算时间线阶段 ID
func GetTimeLineStageIDByBirthday(babyBirthday time.Time) (int, error) {
	now := time.Now()
	if babyBirthday.After(now) {
		return 0, errors.New("baby's birthday cannot be in the future")
	}

	period, err := DateDiffBetweenDates(babyBirthday, now)
	if err != nil {
		return 0, err
	}

	return calculateStageID(period), nil
}

// DateDiff 包含日期差异信息的结构体
type DateDiff struct {
	NumberOfYears  int
	NumberOfMonths int
	TotalDays      int
	NumberOfWeeks  int
	NumberOfDays   int
	LeaveOfDays    int
	KnowUsedWeek   int // 已知周数
}

// DateDiffBetweenDates 计算两个日期之间的差异
func DateDiffBetweenDates(startDate, endDate time.Time) (*DateDiff, error) {
	if endDate.Before(startDate) {
		return nil, errors.New("start date cannot be after end date")
	}

	var dateDiff DateDiff

	duration := endDate.Sub(startDate)

	// 计算总天数
	dateDiff.NumberOfDays = int(duration.Hours() / 24)

	// 计算总周数
	dateDiff.NumberOfWeeks = dateDiff.NumberOfDays / 7

	// 计算总年数和总月数
	years, months, _ := endDate.Date()
	startYears, startMonths, _ := startDate.Date()

	dateDiff.NumberOfYears = years - startYears
	dateDiff.NumberOfMonths = int(months - startMonths)

	// 计算已知周数
	dateDiff.KnowUsedWeek = calculateKnownWeeks(&dateDiff)

	// 计算剩余天数
	dateDiff.LeaveOfDays = dateDiff.NumberOfDays % 7

	return &dateDiff, nil
}

// calculateStageID 根据 DateDiff 计算时间线阶段 ID
func calculateStageID(period *DateDiff) int {
	if period.NumberOfYears < 1 {
		return 1000 + period.KnowUsedWeek
	} else if period.NumberOfMonths < 25 {
		return 2000 + (period.NumberOfMonths - 12)
	} else {
		return 3000 + ((period.NumberOfMonths - 24) / 3)
	}
}

// calculateKnownWeeks 根据 DateDiff 计算已知周数
func calculateKnownWeeks(period *DateDiff) int {
	if period.NumberOfYears < 1 {
		if period.KnowUsedWeek < 49 {
			return period.KnowUsedWeek + 1
		}
	} else if period.NumberOfMonths < 25 {
		return period.NumberOfWeeks + 1
	}
	return period.NumberOfWeeks + 1
}

// 初始化不同年龄类型的偏移值
var stageIDOffset = map[string]int{
	AgeTypeYunqi:   100,
	AgeTypeYinger:  1000,
	AgeTypeYouer:   2000,
	AgeTypeXueqian: 3000,
}
