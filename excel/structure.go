package excel

type (
	GetPurchaseGreyFabricDataForOutPut struct {
		OrderNo         string                                      `json:"order_no"`          // 采购单号
		SupplierName    string                                      `json:"supplier_name"`     // 供应商名称
		ReceiveUnitName string                                      `json:"receive_unit_name"` // 收货单位
		SaleSystemName  string                                      `json:"sale_system_name"`  // 营销体系名称
		ReceiveDate     string                                      `json:"receive_date"`      // 收货日期
		PurchaseDate    string                                      `json:"purchase_date"`     // 采购日期
		TotalPrice      int                                         `json:"total_price"`       // 总价
		CreatorName     string                                      `json:"creator_name"`      // 创建人
		CreateTime      string                                      `json:"create_time"`       // 创建时间
		UpdaterName     string                                      `json:"updater_name"`      // 更新人
		UpdateTime      string                                      `json:"update_time"`       // 更新时间
		AuditorName     string                                      `json:"auditor_name"`      // 审核人
		AuditTime       string                                      `json:"audit_time"`        // 审核时间
		StatusName      string                                      `json:"status_name"`       // 状态
		ItemData        []*GetPurchaseGreyFabricDetailDataForOutPut `json:"item_data"`         // 坯布信息
	}
	GetPurchaseGreyFabricDetailDataForOutPut struct {
		//structure_base.GreyFabricWidthAndWightUnit
		GreyFabricCode                  string `json:"code"`                                  // 坯布编号，必
		GreyFabricName                  string `json:"name"`                                  // 坯布名，必
		GreyFabricWidthAndUnitName      string `json:"grey_fabric_width_and_unit_name"`       // 坯布幅宽及单位名称
		GreyFabricGramWeightAndUnitName string `json:"grey_fabric_gram_weight_and_unit_name"` // 坯布克重及单位名称
		CustomerName                    string `json:"customer_name"`                         // 客人名称
		NeedleSize                      string `json:"needle_size"`                           // 针寸数
		TotalNeedleSize                 string `json:"total_needle_size"`                     // 总针寸数
		ColorName                       string `json:"color_name"`                            // 颜色
		YarnBatch                       string `json:"yarn_batch"`                            // 纱批
		MachineCombinationNumber        string `json:"machine_combination_number"`            // 机台号
		GreyFabricLevelName             string `json:"grey_fabric_level_name"`                // 坯布等级
		WeaveFactoryName                string `json:"weave_factory_name"`                    // 织厂
		Roll                            int    `json:"roll"`                                  // 卷数
		AvgWeight                       int    `json:"avg_weight"`                            // 平均重
		TotalWeight                     int    `json:"total_weight"`                          // 总重
		MeasurementUnitName             string `json:"measurement_unit_name"`                 // 计量单位名称
		UnitPrice                       int    `json:"unit_price"`                            // 单价
		TotalPrice                      int    `json:"total_price"`                           // 总价
		Remark                          string `json:"remark"`                                // 备注
	}
	GetPurchaseGreyFabricDataForOutPutList []*GetPurchaseGreyFabricDataForOutPut
)
