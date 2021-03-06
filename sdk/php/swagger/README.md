# swagger
A cloud native access control server providing best-practice patterns (RBAC, ABAC, ACL, AWS IAM Policies, Kubernetes Roles, ...) via REST APIs.

This PHP package is automatically generated by the [Swagger Codegen](https://github.com/swagger-api/swagger-codegen) project:

- API version: Latest
- Build package: io.swagger.codegen.languages.PhpClientCodegen
For more information, please visit [https://www.ory.sh](https://www.ory.sh)

## Requirements

PHP 5.4.0 and later

## Installation & Usage
### Composer

To install the bindings via [Composer](http://getcomposer.org/), add the following to `composer.json`:

```
{
  "repositories": [
    {
      "type": "git",
      "url": "https://github.com/ory/swagger.git"
    }
  ],
  "require": {
    "ory/swagger": "*@dev"
  }
}
```

Then run `composer install`

### Manual Installation

Download the files and include `autoload.php`:

```php
    require_once('/path/to/swagger/autoload.php');
```

## Tests

To run the unit tests:

```
composer install
./vendor/bin/phpunit
```

## Getting Started

Please follow the [installation procedure](#installation--usage) and then run the following:

```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$api_instance = new ketoSDK\Api\EnginesApi();
$flavor = "flavor_example"; // string | The ORY Access Control Policy flavor. Can be \"regex\", \"glob\", and \"exact\".
$id = "id_example"; // string | The ID of the ORY Access Control Policy Role.
$body = new \ketoSDK\Model\AddOryAccessControlPolicyRoleMembersBody(); // \ketoSDK\Model\AddOryAccessControlPolicyRoleMembersBody | 

try {
    $result = $api_instance->addOryAccessControlPolicyRoleMembers($flavor, $id, $body);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling EnginesApi->addOryAccessControlPolicyRoleMembers: ', $e->getMessage(), PHP_EOL;
}

?>
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*EnginesApi* | [**addOryAccessControlPolicyRoleMembers**](docs/Api/EnginesApi.md#addoryaccesscontrolpolicyrolemembers) | **PUT** /engines/acp/ory/{flavor}/roles/{id}/members | Add a member to an ORY Access Control Policy Role
*EnginesApi* | [**deleteOryAccessControlPolicy**](docs/Api/EnginesApi.md#deleteoryaccesscontrolpolicy) | **DELETE** /engines/acp/ory/{flavor}/policies/{id} | 
*EnginesApi* | [**deleteOryAccessControlPolicyRole**](docs/Api/EnginesApi.md#deleteoryaccesscontrolpolicyrole) | **DELETE** /engines/acp/ory/{flavor}/roles/{id} | Delete an ORY Access Control Policy Role
*EnginesApi* | [**doOryAccessControlPoliciesAllow**](docs/Api/EnginesApi.md#dooryaccesscontrolpoliciesallow) | **POST** /engines/acp/ory/{flavor}/allowed | Check if a request is allowed
*EnginesApi* | [**getOryAccessControlPolicy**](docs/Api/EnginesApi.md#getoryaccesscontrolpolicy) | **GET** /engines/acp/ory/{flavor}/policies/{id} | 
*EnginesApi* | [**getOryAccessControlPolicyRole**](docs/Api/EnginesApi.md#getoryaccesscontrolpolicyrole) | **GET** /engines/acp/ory/{flavor}/roles/{id} | Get an ORY Access Control Policy Role
*EnginesApi* | [**listOryAccessControlPolicies**](docs/Api/EnginesApi.md#listoryaccesscontrolpolicies) | **GET** /engines/acp/ory/{flavor}/policies | 
*EnginesApi* | [**listOryAccessControlPolicyRoles**](docs/Api/EnginesApi.md#listoryaccesscontrolpolicyroles) | **GET** /engines/acp/ory/{flavor}/roles | List ORY Access Control Policy Roles
*EnginesApi* | [**removeOryAccessControlPolicyRoleMembers**](docs/Api/EnginesApi.md#removeoryaccesscontrolpolicyrolemembers) | **DELETE** /engines/acp/ory/{flavor}/roles/{id}/members/{member} | Remove a member from an ORY Access Control Policy Role
*EnginesApi* | [**upsertOryAccessControlPolicy**](docs/Api/EnginesApi.md#upsertoryaccesscontrolpolicy) | **PUT** /engines/acp/ory/{flavor}/policies | 
*EnginesApi* | [**upsertOryAccessControlPolicyRole**](docs/Api/EnginesApi.md#upsertoryaccesscontrolpolicyrole) | **PUT** /engines/acp/ory/{flavor}/roles | Upsert an ORY Access Control Policy Role
*HealthApi* | [**isInstanceAlive**](docs/Api/HealthApi.md#isinstancealive) | **GET** /health/alive | Check alive status
*HealthApi* | [**isInstanceReady**](docs/Api/HealthApi.md#isinstanceready) | **GET** /health/ready | Check readiness status
*VersionApi* | [**getVersion**](docs/Api/VersionApi.md#getversion) | **GET** /version | Get service version


## Documentation For Models

 - [AddOryAccessControlPolicyRoleMembers](docs/Model/AddOryAccessControlPolicyRoleMembers.md)
 - [AddOryAccessControlPolicyRoleMembersBody](docs/Model/AddOryAccessControlPolicyRoleMembersBody.md)
 - [AddOryAccessControlPolicyRoleMembersInternalServerError](docs/Model/AddOryAccessControlPolicyRoleMembersInternalServerError.md)
 - [AddOryAccessControlPolicyRoleMembersInternalServerErrorBody](docs/Model/AddOryAccessControlPolicyRoleMembersInternalServerErrorBody.md)
 - [AddOryAccessControlPolicyRoleMembersOK](docs/Model/AddOryAccessControlPolicyRoleMembersOK.md)
 - [AddOryAccessControlPolicyRoleMembersReader](docs/Model/AddOryAccessControlPolicyRoleMembersReader.md)
 - [AuthorizationResult](docs/Model/AuthorizationResult.md)
 - [Context](docs/Model/Context.md)
 - [DeleteOryAccessControlPolicy](docs/Model/DeleteOryAccessControlPolicy.md)
 - [DeleteOryAccessControlPolicyCreated](docs/Model/DeleteOryAccessControlPolicyCreated.md)
 - [DeleteOryAccessControlPolicyInternalServerError](docs/Model/DeleteOryAccessControlPolicyInternalServerError.md)
 - [DeleteOryAccessControlPolicyInternalServerErrorBody](docs/Model/DeleteOryAccessControlPolicyInternalServerErrorBody.md)
 - [DeleteOryAccessControlPolicyNoContent](docs/Model/DeleteOryAccessControlPolicyNoContent.md)
 - [DeleteOryAccessControlPolicyReader](docs/Model/DeleteOryAccessControlPolicyReader.md)
 - [DeleteOryAccessControlPolicyRole](docs/Model/DeleteOryAccessControlPolicyRole.md)
 - [DeleteOryAccessControlPolicyRoleCreated](docs/Model/DeleteOryAccessControlPolicyRoleCreated.md)
 - [DeleteOryAccessControlPolicyRoleInternalServerError](docs/Model/DeleteOryAccessControlPolicyRoleInternalServerError.md)
 - [DeleteOryAccessControlPolicyRoleInternalServerErrorBody](docs/Model/DeleteOryAccessControlPolicyRoleInternalServerErrorBody.md)
 - [DeleteOryAccessControlPolicyRoleNoContent](docs/Model/DeleteOryAccessControlPolicyRoleNoContent.md)
 - [DeleteOryAccessControlPolicyRoleReader](docs/Model/DeleteOryAccessControlPolicyRoleReader.md)
 - [DoOryAccessControlPoliciesAllow](docs/Model/DoOryAccessControlPoliciesAllow.md)
 - [DoOryAccessControlPoliciesAllowForbidden](docs/Model/DoOryAccessControlPoliciesAllowForbidden.md)
 - [DoOryAccessControlPoliciesAllowInternalServerError](docs/Model/DoOryAccessControlPoliciesAllowInternalServerError.md)
 - [DoOryAccessControlPoliciesAllowInternalServerErrorBody](docs/Model/DoOryAccessControlPoliciesAllowInternalServerErrorBody.md)
 - [DoOryAccessControlPoliciesAllowOK](docs/Model/DoOryAccessControlPoliciesAllowOK.md)
 - [DoOryAccessControlPoliciesAllowReader](docs/Model/DoOryAccessControlPoliciesAllowReader.md)
 - [GetOryAccessControlPolicy](docs/Model/GetOryAccessControlPolicy.md)
 - [GetOryAccessControlPolicyInternalServerError](docs/Model/GetOryAccessControlPolicyInternalServerError.md)
 - [GetOryAccessControlPolicyInternalServerErrorBody](docs/Model/GetOryAccessControlPolicyInternalServerErrorBody.md)
 - [GetOryAccessControlPolicyNotFound](docs/Model/GetOryAccessControlPolicyNotFound.md)
 - [GetOryAccessControlPolicyNotFoundBody](docs/Model/GetOryAccessControlPolicyNotFoundBody.md)
 - [GetOryAccessControlPolicyOK](docs/Model/GetOryAccessControlPolicyOK.md)
 - [GetOryAccessControlPolicyReader](docs/Model/GetOryAccessControlPolicyReader.md)
 - [GetOryAccessControlPolicyRole](docs/Model/GetOryAccessControlPolicyRole.md)
 - [GetOryAccessControlPolicyRoleInternalServerError](docs/Model/GetOryAccessControlPolicyRoleInternalServerError.md)
 - [GetOryAccessControlPolicyRoleInternalServerErrorBody](docs/Model/GetOryAccessControlPolicyRoleInternalServerErrorBody.md)
 - [GetOryAccessControlPolicyRoleNotFound](docs/Model/GetOryAccessControlPolicyRoleNotFound.md)
 - [GetOryAccessControlPolicyRoleNotFoundBody](docs/Model/GetOryAccessControlPolicyRoleNotFoundBody.md)
 - [GetOryAccessControlPolicyRoleOK](docs/Model/GetOryAccessControlPolicyRoleOK.md)
 - [GetOryAccessControlPolicyRoleReader](docs/Model/GetOryAccessControlPolicyRoleReader.md)
 - [HealthNotReadyStatus](docs/Model/HealthNotReadyStatus.md)
 - [HealthStatus](docs/Model/HealthStatus.md)
 - [InlineResponse500](docs/Model/InlineResponse500.md)
 - [Input](docs/Model/Input.md)
 - [IsInstanceAliveInternalServerError](docs/Model/IsInstanceAliveInternalServerError.md)
 - [IsInstanceAliveInternalServerErrorBody](docs/Model/IsInstanceAliveInternalServerErrorBody.md)
 - [IsInstanceAliveOK](docs/Model/IsInstanceAliveOK.md)
 - [IsInstanceAliveReader](docs/Model/IsInstanceAliveReader.md)
 - [ListOryAccessControlPolicies](docs/Model/ListOryAccessControlPolicies.md)
 - [ListOryAccessControlPoliciesInternalServerError](docs/Model/ListOryAccessControlPoliciesInternalServerError.md)
 - [ListOryAccessControlPoliciesInternalServerErrorBody](docs/Model/ListOryAccessControlPoliciesInternalServerErrorBody.md)
 - [ListOryAccessControlPoliciesOK](docs/Model/ListOryAccessControlPoliciesOK.md)
 - [ListOryAccessControlPoliciesReader](docs/Model/ListOryAccessControlPoliciesReader.md)
 - [ListOryAccessControlPolicyRoles](docs/Model/ListOryAccessControlPolicyRoles.md)
 - [ListOryAccessControlPolicyRolesInternalServerError](docs/Model/ListOryAccessControlPolicyRolesInternalServerError.md)
 - [ListOryAccessControlPolicyRolesInternalServerErrorBody](docs/Model/ListOryAccessControlPolicyRolesInternalServerErrorBody.md)
 - [ListOryAccessControlPolicyRolesOK](docs/Model/ListOryAccessControlPolicyRolesOK.md)
 - [ListOryAccessControlPolicyRolesReader](docs/Model/ListOryAccessControlPolicyRolesReader.md)
 - [OryAccessControlPolicies](docs/Model/OryAccessControlPolicies.md)
 - [OryAccessControlPolicy](docs/Model/OryAccessControlPolicy.md)
 - [OryAccessControlPolicyAllowedInput](docs/Model/OryAccessControlPolicyAllowedInput.md)
 - [OryAccessControlPolicyRole](docs/Model/OryAccessControlPolicyRole.md)
 - [OryAccessControlPolicyRoles](docs/Model/OryAccessControlPolicyRoles.md)
 - [Policies](docs/Model/Policies.md)
 - [Policy](docs/Model/Policy.md)
 - [RemoveOryAccessControlPolicyRoleMembers](docs/Model/RemoveOryAccessControlPolicyRoleMembers.md)
 - [RemoveOryAccessControlPolicyRoleMembersCreated](docs/Model/RemoveOryAccessControlPolicyRoleMembersCreated.md)
 - [RemoveOryAccessControlPolicyRoleMembersInternalServerError](docs/Model/RemoveOryAccessControlPolicyRoleMembersInternalServerError.md)
 - [RemoveOryAccessControlPolicyRoleMembersInternalServerErrorBody](docs/Model/RemoveOryAccessControlPolicyRoleMembersInternalServerErrorBody.md)
 - [RemoveOryAccessControlPolicyRoleMembersReader](docs/Model/RemoveOryAccessControlPolicyRoleMembersReader.md)
 - [Role](docs/Model/Role.md)
 - [Roles](docs/Model/Roles.md)
 - [SwaggerHealthStatus](docs/Model/SwaggerHealthStatus.md)
 - [SwaggerNotReadyStatus](docs/Model/SwaggerNotReadyStatus.md)
 - [SwaggerVersion](docs/Model/SwaggerVersion.md)
 - [UpsertOryAccessControlPolicy](docs/Model/UpsertOryAccessControlPolicy.md)
 - [UpsertOryAccessControlPolicyInternalServerError](docs/Model/UpsertOryAccessControlPolicyInternalServerError.md)
 - [UpsertOryAccessControlPolicyInternalServerErrorBody](docs/Model/UpsertOryAccessControlPolicyInternalServerErrorBody.md)
 - [UpsertOryAccessControlPolicyOK](docs/Model/UpsertOryAccessControlPolicyOK.md)
 - [UpsertOryAccessControlPolicyReader](docs/Model/UpsertOryAccessControlPolicyReader.md)
 - [UpsertOryAccessControlPolicyRole](docs/Model/UpsertOryAccessControlPolicyRole.md)
 - [UpsertOryAccessControlPolicyRoleInternalServerError](docs/Model/UpsertOryAccessControlPolicyRoleInternalServerError.md)
 - [UpsertOryAccessControlPolicyRoleInternalServerErrorBody](docs/Model/UpsertOryAccessControlPolicyRoleInternalServerErrorBody.md)
 - [UpsertOryAccessControlPolicyRoleOK](docs/Model/UpsertOryAccessControlPolicyRoleOK.md)
 - [UpsertOryAccessControlPolicyRoleReader](docs/Model/UpsertOryAccessControlPolicyRoleReader.md)
 - [Version](docs/Model/Version.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author

hi@ory.sh


