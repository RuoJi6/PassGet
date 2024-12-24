package browser

import (
	"PassGet/modules/browser/pick"
	br "PassGet/modules/utils/browser"
	"PassGet/modules/utils/browser/fileutil"
)

func Get() error {
	browsers, err := browser.PickBrowsers(br.BrowserName, br.ProfilePath)
	if err != nil {
		//log.Errorf("pick browsers %v", err)
		return err
	}
	for _, b := range browsers {
		data, err := b.BrowsingData(true)
		if err != nil {
			//log.Errorf("get browsing data error %v", err)
			continue
		}
		data.Output(br.OutputDir, b.Name(), br.OutputFormat)
	}
	if err = fileutil.CompressDir(br.OutputDir); err != nil {
		//log.Errorf("compress error %v", err)
	}
	//log.Debug("compress success")
	return nil
}
