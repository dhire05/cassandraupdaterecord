---
title: CassandraDB Update Activity
---

# CassandraDB Update Activity
This activity allows you to update a record to a particular table from the CassandraDB server

## Installation
### Flogo CLI
```bash
flogo install github.com/dhire05/cassandraupdaterecord
```

## Schema
Inputs and Outputs:

```json
{   
  "inputs":[
    {
      "name": "ClusterIP",
      "type": "string",
	  "required": true      
    },
	{
      "name": "Keyspace",
      "type": "string",
      "required": true
    },
	{
      "name": "TableName",
      "type": "string",
      "required": true
    },
	 {
      "name": "SET",
      "type": "string",
      "required": true
    },
    {
      "name": "Where",
      "type": "string",
      "required": true
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
 }
```
## Settings
| Setting        | Required | Description |
|:---------------|:---------|:------------|
| ClusterIP      | True     | The CassandraDB cluster instance |         
| Keyspace       | True     | The name of the Keyspace
| TableName      | True     | The name of table to update record
| SET 	         | True     | To set the particular values
| Where          | True     | The where clause or condition |


## Example
The below example is to insert a record into CassandraDB

```json
{
  "id": "CassandraDB_1",
  "name": "CassandraDB connector",
  "description": "Insert record into CassandraDB",
  "activity": {
    "ref": "github.com/dhire05/cassandraupdaterecord",
    "input": {
      "ClusterIP": "127.0.0.1",
      "Keyspace": "sample",
      "TableName": "employee",
	  "SET": "salary = 3500",
	  "Where": "empid = 105",
    }
  }
}
```