package mysql

// nolint: lll
var dbmsARMTemplateBytes = []byte(`
	{
		"$schema": "http://schema.management.azure.com/schemas/2014-04-01-preview/deploymentTemplate.json#",
		"contentVersion": "1.0.0.0",
		"parameters": {
			"tags": {
				"type": "object"
			}
		},
		"variables": {
			"DBforMySQLapiVersion": "2017-12-01"
		},
		"resources": [
			{
				"apiVersion": "[variables('DBforMySQLapiVersion')]",
				"kind": "",
				"location": "{{.location}}",
				"name": "{{ .serverName }}",
				"properties": {
					"version": "{{.version}}",
					"administratorLogin": "azureuser",
					"administratorLoginPassword": "{{ .administratorLoginPassword }}",
					"storageProfile": {
						"storageMB": {{.storage}},
						{{ if .geoRedundantBackup }}
						"geoRedundantBackup": "Enabled",
						{{ end }}
						"backupRetentionDays": {{.backupRetention}}
					},
					"sslEnforcement": "{{ .sslEnforcement }}"
				},
				"sku": {
					"name": "{{.sku}}",
					"tier": "{{.tier}}",
					"capacity": "{{.cores}}",
					"size": "{{.storage}}",
					"family": "Gen5"
				},
				"type": "Microsoft.DBforMySQL/servers",
				"tags": "[parameters('tags')]",
				"resources": [
					{{ $root := . }}
					{{$count:= sub (len .firewallRules)  1}}
					{{range $i, $rule := .firewallRules}}
					{
						"type": "firewallrules",
						"apiVersion": "[variables('DBforMySQLapiVersion')]",
						"dependsOn": [
							"Microsoft.DBforMySQL/servers/{{ $.serverName }}"
						],
						"location": "{{$root.location}}",
						"name": "{{$rule.name}}",
						"properties": {
							"startIpAddress": "{{$rule.startIPAddress}}",
							"endIpAddress": "{{$rule.endIPAddress}}"
						}
					}{{if lt $i $count}},{{end}}
					{{end}}
				]
			}
		],
		"outputs": {
			"fullyQualifiedDomainName": {
				"type": "string",
				"value": "[reference('{{ .serverName }}').fullyQualifiedDomainName]"
			}
		}
	}
	`)
