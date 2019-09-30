package models

type URLMonitor struct {
	URL    		string
	StatusOK 	bool
}

type URLMonitors struct{
    URLMonitors []URLMonitor
}