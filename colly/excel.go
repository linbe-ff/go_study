package colly

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	. "go_study/mysql"
	"strconv"
	"strings"
)

func Read(fileName string, sheetName string) (PageInfoList, error) {
	var (
		list = make(PageInfoList, 0)
	)
	// 读取Excel
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
	}

	// _是索引号，从0开始，row是整行数据
	for i, row := range rows {
		if i == 0 {
			continue // 跳过表头
		}
		page, _ := strconv.Atoi(row[0])
		list = append(list, &PageInfo{
			Page:  page,
			Title: row[1],
			Href:  row[2],
			Src:   row[3],
		})
	}
	return list, nil
}

func CreateExcel(list PageInfoList, fileName string, sheetName string) error {
	var (
		err error
	)
	f := excelize.NewFile()
	// sheet改名
	f.SetSheetName("Sheet1", sheetName)

	// 设置单元格的值
	err = f.SetCellValue(sheetName, "A1", "page")
	err = f.SetCellValue(sheetName, "B1", "title")
	err = f.SetCellValue(sheetName, "C1", "href")
	err = f.SetCellValue(sheetName, "D1", "src")
	if err != nil {
		return err
	}
	for i, info := range list {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), info.Page)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), info.Title)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), info.Href)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), info.Src)
	}

	// 保存Excel
	if err := f.SaveAs(fileName); err != nil {
		return err
	}
	return nil
}

func ExcelAppend(list PageInfoList, fileName string, sheetName string) error {
	var err error
	f, err := excelize.OpenFile(fileName)
	if err != nil && !strings.Contains(err.Error(), "cannot find") {
		fmt.Println(err)
		return err
	}
	if f == nil {
		err = CreateExcel(list, fileName, sheetName)
		return err
	}
	cols, err := f.GetCols(sheetName)
	if err != nil {
		return err
	}
	var index int
	if len(cols) == 0 {
		index = 1
	} else {
		index = len(cols[0]) + 1
	}
	for i, info := range list {
		f.SetCellValue(sheetName, "A"+strconv.Itoa(index+i), info.Page)
		f.SetCellValue(sheetName, "B"+strconv.Itoa(index+i), info.Title)
		f.SetCellValue(sheetName, "C"+strconv.Itoa(index+i), info.Href)
		f.SetCellValue(sheetName, "D"+strconv.Itoa(index+i), info.Src)
	}
	err = f.SaveAs(fileName)
	if err != nil {
		return err
	}
	return nil
}
