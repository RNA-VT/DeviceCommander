package utilities

import "encoding/json"

//LabelString -
func LabelString(key string, value string) string {
	return "\n\t[" + key + "]:     " + value
}

//JSON Marshals an object into a byte array
func JSON(obj interface{}) (out []byte, err error) {
	out, err = json.Marshal(obj)
	if err != nil {
		return
	}
	return
}

//StringJSON returns a stringified of a json string representing obj
func StringJSON(obj interface{}) (out string, err error) {
	bytes, err := JSON(obj)
	if err != nil {
		out = "ObjectFailedToParse"
		return
	}
	return string(bytes), err
}

/*
TODO: Something still needs to do the below work to support deregistration (on graceful exit) and stop all GoFire (safety event)

//ClusterError - Log the errors, warn the others and then panic.
func (e *ErrorHandler) ClusterError(panicAfterWarning bool, panicCluster bool, notGoodThings ...BadThing) {
	//Errors that render this microcontroller unusable, but do not effect the rest of the cluster
	e.UhOh(notGoodThings...)
	if panicCluster {
		e.EverybodyPanic(*e...)
	} else {
		e.WarnTheOthers(*e...)
	}
	if panicAfterWarning {
		panic(e)
	}
}

//WarnTheOthers - POST Error(s) to cluster.
func (e *ErrorHandler) WarnTheOthers(notGoodThings ...BadThing) {
	//This path should be used for errors that make this instance of GoFire unavailable
	e.tellTheOthers("/errors/warn", notGoodThings...)
}

//EverybodyPanic - Meant for Errors that should stop the entire cluster
func (e *ErrorHandler) EverybodyPanic(notGoodThings ...BadThing) {
	e.tellTheOthers("errors/panic", notGoodThings...)
}

func (e *ErrorHandler) tellTheOthers(path string, notGoodThings ...BadThing) {
	/* 	c := *ClusterRef
	   	c.UpdatePeers(
	   		path,
	   		PeerErrorMessage{
	   			Source: c.Me,
	   			Errors: notGoodThings,
	   		},
	   		[]microcontroller.Microcontroller{c.Me}) */ /*
	}

*/
