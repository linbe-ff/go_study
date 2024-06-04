package excel

import (
	"bytes"
	"github.com/xuri/excelize/v2"
	"os"
	"strconv"
	"testing"
)

func TestOutput(t *testing.T) {

	var (
		err    error
		index  int = 2
		list       = make(GetPurchaseGreyFabricDataForOutPutList, 0)
		f          = excelize.NewFile()
		sheet      = "坯布采购"
		buffer     = &bytes.Buffer{}
	)

	titleNames := []interface{}{
		"订单编号", "供应商", "收货单位", "营销体系", "收货日期", "采购日期", "单据金额", "创建人", "创建时间", "修改人",
		"修改时间", "审核人", "审核时间", "状态",
		"坯布编号", "坯布名称", "客户", "幅宽", "克重", "针寸数",
		"总针数", "织坯颜色", "纱批", "机台号", "坯布等级", "织厂名称", "匹数", "均重", "数量总计", "单位",
		"单价", "金额", "备注",
	}
	lineMap := ArrangeA2Z(len(titleNames))

	// 样式封装函数
	f, err = setPriceExcelStyle(f, sheet, titleNames, lineMap)
	if err != nil {
		return
	}
	for _, order := range list {
		var (
			indexStr           = strconv.Itoa(index)
			itemLen            = len(order.ItemData)
			lineIndex          = 0
			indexAddItemLenStr = strconv.Itoa(index + itemLen - 1)
		)
		// 合并列
		f.MergeCell(sheet, lineMap[lineIndex]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.OrderNo})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.SupplierName})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.ReceiveUnitName})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.SaleSystemName})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.ReceiveDate})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.PurchaseDate})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.TotalPrice})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.CreatorName})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.CreateTime})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.UpdaterName})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.UpdateTime})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.AuditorName})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.AuditTime})
		f.MergeCell(sheet, lineMap[*AddOne(&lineIndex)]+indexStr, lineMap[lineIndex]+indexAddItemLenStr)
		f.SetSheetRow(sheet, lineMap[lineIndex]+indexStr, &[]interface{}{order.StatusName})
		for _, item := range order.ItemData {
			var (
				copyLineIndex = lineIndex
			)
			indexStr = strconv.Itoa(index)
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.GreyFabricCode})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.GreyFabricName})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.CustomerName})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.GreyFabricWidthAndUnitName})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.GreyFabricGramWeightAndUnitName})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.NeedleSize})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.TotalNeedleSize})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.ColorName})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.YarnBatch})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.MachineCombinationNumber})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.GreyFabricLevelName})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.WeaveFactoryName})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.Roll})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.AvgWeight})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.TotalWeight})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.MeasurementUnitName})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.UnitPrice})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.TotalPrice})
			f.SetSheetRow(sheet, lineMap[*AddOne(&copyLineIndex)]+indexStr, &[]interface{}{item.Remark})
			index++
		}
	}
	// 删除Sheet1表
	f.DeleteSheet("Sheet1")
	buffer, err = f.WriteToBuffer()
	if err != nil {
		return
	}

	file, err := os.Create("D:\\" + sheet + ".xlsx")
	if err != nil {
		return
	}
	defer file.Close()

	_, err = buffer.WriteTo(file)
	if err != nil {
		return
	}
}

func setPriceExcelStyle(f *excelize.File, sheet string, titleNames []interface{}, lineIndex []string) (*excelize.File, error) {
	var (
		err           error
		lenTitleNames = len(titleNames)
		strOne        = strconv.Itoa(1)
	)
	columnWidths := []float64{
		20, 15, 15, 30, 13, 13, 13, 13, 20, 13,
		20, 13, 20, 10, 13, 20, 20, 13, 20, 13,
		13, 20, 20, 20, 13, 20, 13, 13, 13, 13,
		20, 10, 26,
	} // 示例宽度数组，您可以自定义每个值

	f.NewSheet(sheet)
	f.SetRowHeight("坯布采购", 1, 20)

	// 垂直居中样式 字体加粗
	// 定义一个居中对齐的单元格样式
	style, _ := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	for i := 0; i < lenTitleNames; i++ {
		f.SetSheetRow(sheet, lineIndex[i]+strOne, &[]interface{}{titleNames[i]})
		f.SetColStyle(sheet, lineIndex[i], style)
		f.SetColWidth(sheet, lineIndex[i], lineIndex[i], columnWidths[i])
	}

	return f, err
}
