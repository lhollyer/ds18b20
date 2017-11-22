package ds18b20

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/yryz/ds18b20"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	sensors, err := ds18b20.Sensors()
    	if err != nil {
       		panic(err)
    	}

    	for _, sensor := range sensors {
        	t, err := ds18b20.Temperature(sensor)
        	if err == nil {
            		log.Debugf("sensor: %s temperature: %.2f°C\n", sensor, t)
        	}
    	}


	return true, nil
}
