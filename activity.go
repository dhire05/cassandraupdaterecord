package cassandraupdaterecord

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/gocql/gocql"
)

// THIS IS ADDED
// log is the default package logger which we'll use to log
var log = logger.GetLogger("activity-CassandraUpdateRecord")

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
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {
	// Get the activity data from the context
	clusterIP := context.GetInput("ClusterIP").(string)
	keySpace := context.GetInput("Keyspace").(string)
	tableName := context.GetInput("TableName").(string)	
	setElements := context.GetInput("SET").(string)
	whereClause := context.GetInput("Where").(string)

	// Use the log object to log the greeting
	log.Debugf("The Flogo engine says connect to [%s] with  [%s] with table [%s]", clusterIP, keySpace, tableName)
	log.Debugf("Flogo is about to update [%s] from table [%s].[%s] where [%s] on cluster [%s]", setElements, keySpace, tableName, whereClause, clusterIP)
	
	fmt.Println("The Flogo engine says connect to "+clusterIP+" with  "+keySpace+" with table "+tableName)
	fmt.Println("Flogo is about to update "+setElements+" from table "+keySpace+"."+tableName+" where "+whereClause+" on cluster "+clusterIP)

	// Provide the cassandra cluster instance here.
	cluster := gocql.NewCluster(clusterIP)

	// gocql requires the keyspace to be provided before the session is created.
	// In future there might be provisions to do this later.
	cluster.Keyspace = keySpace

	session, err := cluster.CreateSession()
	log.Debugf("Session Created Sucessfully")

	if err != nil {
		log.Debugf("Could not connect to cassandra cluster : ", err)
		fmt.Println("Could not connect to cassandra cluster : ", err)
	}
	log.Debugf("Session : ", session)
	log.Debugf("Cluster : ", clusterIP)
	log.Debugf("Keyspace : ", keySpace)
	log.Debugf("Session Timeout : ", cluster.Timeout)
	log.Debugf("TableName : ", tableName)
	log.Debugf("SET Element : ", setElements)
	log.Debugf("Where Clause : ", whereClause)

	log.Debugf("Next Step is Update Query Execution")

	//fmt.Println("Session : ", session)
	//fmt.Println("Cluster : ", clusterIP)
	//fmt.Println("Keyspace : ", keySpace)
	//fmt.Println("Session Timeout : ", cluster.Timeout)
	//fmt.Println("TableName : ", tableName)
	//fmt.Println("SET Element : ", setElements)
	//fmt.Println("Where Clause : ", whereClause)

	fmt.Println("Next Step is Update Query Execution")

	updateString := "UPDATE " + tableName + " SET " + setElements
	if whereClause != "" {
		updateString += " WHERE " + whereClause
	}
	log.Debugf("Update string: [%s]", updateString)

	fmt.Println("Update string: " + updateString)

	if err := session.Query(updateString).Exec(); err != nil {
		log.Debugf("Error is update query : ", err)
		fmt.Println("Error is update query : ", err)
	}
	
	// Set the result as part of the context
	context.SetOutput("result", "Record Updated SuccessFully")
	
	fmt.Println("Record Updated SuccessFully")

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}
