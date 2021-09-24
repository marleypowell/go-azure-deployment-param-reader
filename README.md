# go-azure-deployment-param-reader

## Build

To create the `go-azure-deployment-param-reader.exe` you can run:

```shell
go build
```

## Usage

```shell
./go-azure-deployment-param-reader.exe <RESOURCE_GROUP>
```

### Example

Command:

```shell
./go-azure-deployment-param-reader.exe ratsmail7-environment
```

Response:

```json
{
  "name": "AIODeploy-20210624-090705-ff75",
  "resourceGroup": "ratsmail7-environment",
  "properties": {
    "outputs": {
      "adssaConnectionString": {
        "type": "String",
        "value": ""
      },
      "auditEventHubNameSpace": {
        "type": "String",
        "value": ""
      },
      "auditHubConnectionString": {
        "type": "String",
        "value": ""
      },
      "auditHubReceiveKey": {
        "type": "String",
        "value": ""
      },
      "auditHubSendKey": {
        "type": "String",
        "value": ""
      },
      "blwlsaConnectionString": {
        "type": "String",
        "value": ""
      },
      "csisaConnectionString": {
        "type": "String",
        "value": ""
      },
      "gSuiteSAConnectionString": {
        "type": "String",
        "value": ""
      },
      "gffsaConnectionString": {
        "type": "String",
        "value": ""
      },
      "imprintSAConnectionString": {
        "type": "String",
        "value": ""
      },
      "logsSAConnectionString": {
        "type": "String",
        "value": ""
      },
      "photosSAConnectionString": {
        "type": "String",
        "value": ""
      },
      "storageSAConnectionString": {
        "type": "String",
        "value": ""
      },
      "vmDnsName": {
        "type": "String",
        "value": ""
      },
      "vmPublicIp": {
        "type": "String",
        "value": ""
      }
    },
    "parameters": {
      "auditEventHubNameSpace": {
        "type": "String",
        "value": ""
      },
      "auditHubConnectionString": {
        "type": "String",
        "value": ""
      },
      "auditHubReceiveKey": {
        "type": "String",
        "value": ""
      },
      "auditHubSendKey": {
        "type": "String",
        "value": ""
      },
      "dataCenter": {
        "type": "String",
        "value": ""
      },
      "keyVaultClientId": {
        "type": "String",
        "value": ""
      },
      "location": {
        "type": "String",
        "value": ""
      },
      "publicIPName": {
        "type": "String",
        "value": ""
      },
      "tenantId": {
        "type": "String",
        "value": ""
      },
      "userPassword": {
        "type": "String",
        "value": ""
      },
      "windowsOSVersion": {
        "type": "String",
        "value": ""
      }
    }
  }
}
```
