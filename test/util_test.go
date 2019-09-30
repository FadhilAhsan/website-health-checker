package test

import(
	"testing"
	"github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/utils"
	"github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/models"
)


func TestLoadFromFile(t *testing.T){
	var URLList models.URLMonitors
	err := utils.LoadFromFile("util_test-21.json",&URLList)
	if err != nil {
		t.Log("LoadFromFile(string) load file bad scnario PASSED")
	}else{
		t.Errorf("LoadFromFile(string) load file bad scnario Failed, expected %s but actualy %s ","Failed to load file and get error","Success to read file")
	}

	err = utils.LoadFromFile("util_test-2.json",&URLList)
	if err == nil {
		t.Log("LoadFromFile(string) load file good scnario PASSED")
	}else{
		t.Errorf("LoadFromFile(string) load file good scnario Failed, expected %s but actualy %s ","Success to load file","Failed to load file")
	}

	if len(URLList.URLMonitors) == 0 {
		t.Log("LoadFromFile(string) get json from file bad scnario PASSED")
	}else{
		t.Errorf("LoadFromFile(string) get json from file bad scnario Failed, expected %s but actualy %s ","length must 0 ","but length not 0")
	}

	err = utils.LoadFromFile("util_test-1.json",&URLList)
	if len(URLList.URLMonitors) > 0 {
		t.Log("LoadFromFile(string) get json from file good scnario PASSED")
	}else{
		t.Errorf("LoadFromFile(string) get json from filee good scnario Failed, expected %s but actualy %s ","length must greater than 0"," length less than 1")
	}
}