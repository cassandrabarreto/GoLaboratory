package prices

import (
	"fmt"
	"taxation/conversion"
	"taxation/filemanager"
)

type TaxIncludedPrice struct {
	IOManager filemanager.FileManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_rate"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPrice) Process(doneChan chan bool, errorChan chan error){
	err := job.LoadData()

	if err != nil {
		errorChan <- err
		return
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		TaxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", TaxIncludedPrice)
	}
	job.TaxIncludedPrices = result
	job.IOManager.WriteResult(job)
	doneChan <- true
}

func (job *TaxIncludedPrice) LoadData() error {
	lines, err := job.IOManager.ReadLines()
 
	prices, err := conversion.StringsToFloats(lines)

	if err != nil{
		fmt.Println(err)
		return err
	}
	job.InputPrices = prices
}

func NewTaxIncludedPrice(fm filemanager.FileManager, taxRate float64) *TaxIncludedPrice{
	return &TaxIncludedPrice{
		IOManager: fm,
		InputPrices : []float64{10,20,30},
		TaxRate : taxRate,
	}
}