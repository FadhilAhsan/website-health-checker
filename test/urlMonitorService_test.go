package test

import(
	"testing"
	"github.com/FadhilAhsan/website-health-checker/internal/app/website-health-checker/services"
)


func TestIsURL(t *testing.T){
	result := services.IsURL("@detik!..com")
	if !result {
		t.Log("IsURL(string) bad scnario PASSED")
	}else{
		t.Errorf("IsURL(string) bad scnario Failed, expected %s but actualy %s ","false","true")
	}

	result = services.IsURL("http://detik.com")
	if result {
		t.Log("IsURL(string) good scnario PASSED")
	}else{
		t.Errorf("IsURL(string) good scnario Failed, expected %s but actualy %s ","true","false")
	}
}


func TestCheckHealthURL(t *testing.T){
	result,err := services.CheckHealthURL("http:/taraktak123.com")
	if err != nil {
		t.Errorf("CheckHealthURL(string) bad scnario Failed, expected expected %s but actualy %s : %v ","not getting error","getting error",err)
	}else{
		if !result.StatusOK {
			t.Log("CheckHealthURL(string) bad scnario PASSED")
		}else{
			t.Errorf("CheckHealthURL(string) bad scnario Failed, expected expected StatusOK is %s but actualy %s ","false","true")
		}
	}
	

	result,err = services.CheckHealthURL("http://google.com")
	if err != nil {
		t.Errorf("CheckHealthURL(string) good scnario Failed, expected expected %s but actualy %s : %v","not getting error","getting error",err)
	}else{
		if result.StatusOK {
			t.Log("CheckHealthURL(string) good scnario PASSED")
		}else{
			t.Errorf("CheckHealthURL(string) good scnario Failed, expected StatusOK is %s but actualy %s ","true","false")
		}
	}
	
}