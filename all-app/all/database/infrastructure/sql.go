package infrastructure

const SQL = "create view spreadsheet as " +
	"select " +
	"t.id,\n    " +
	"t.create_time,\n    " +
	"t.update_time,\n    " +
	"o.organization_name,\n    " +
	"t.subject,\n    " +
	"t.status,\n    " +
	"t.priority,\n    " +
	"r.request_name,\n    " +
	"a.assignee_name,\n    " +
	"t.tag\n" +
	"from tickets as t\n" +
	"join (\n        " +
	"select\n            " +
	"t.id,\n            " +
	"u.name as assignee_name\n        " +
	"from tickets as t\n        " +
	"join users as u\n        " +
	"on t.assignee_id = u.id\n    " +
	") as a\n" +
	"on t.id = a.id\n" +
	"join (\n        " +
	"select \n            " +
	"t.id,\n            " +
	"u.name as request_name\n        " +
	"from tickets as t\n        " +
	"join users as u\n        " +
	"on t.requester_id = u.id\n    " +
	") as r\n" +
	"on t.id = r.id\n" +
	"join (\n        " +
	"select\n            " +
	"t.id,\n            " +
	"o.name as organization_name\n        " +
	"from tickets as t\n        " +
	"join organizations as o\n        " +
	"on t.organization_id = o.id\n    " +
	") as o\n" +
	"on t.id = o.id;"
