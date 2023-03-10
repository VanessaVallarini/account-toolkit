package avros

const (
	AccountCreateOrUpdateSubject = "com.account.create.or.update"
	AccountCreateOrUpdateAvro    = `{
		"type":"record",
		"name":"Account_Create_Or_Update",
		"namespace":"com.account.create.or.update",
		"fields":[
			 {
				"name":"email",
				"type":"string"
			 },
			 {
				"name":"full_number",
				"type":"string"
			 },
			 {
				"name":"alias",
				"type":"string"
			 },
			 {
				"name":"city",
				"type":"string"
			 },
			 {
				"name":"district",
				"type":"string"
			 },
			 {
				"name":"name",
				"type":"string"
			 },
			 {
				"name":"public_place",
				"type":"string"
			 },
			 {
				"name":"status",
				"type":"string"
			 },
			 {
				"name":"zip_code",
				"type":"string"
			 }		   
		]
	 }`
	AccountDeleteSubject = "com.account.delete"
	AccountDeleteAvro    = `{
		"type":"record",
		"name":"Account_Delete",
		"namespace":"com.account.delete",
		"fields":[
			{
				"name":"email",
				"type":"string"
			 }	   
		]
	 }`
)
