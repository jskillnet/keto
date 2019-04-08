/**
 * ORY Keto
 * A cloud native access control server providing best-practice patterns (RBAC, ABAC, ACL, AWS IAM Policies, Kubernetes Roles, ...) via REST APIs.
 *
 * OpenAPI spec version: Latest
 * Contact: hi@ory.sh
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.2.3
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.OryKeto) {
      root.OryKeto = {};
    }
    root.OryKeto.DeleteOryAccessControlPolicyCreated = factory(root.OryKeto.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The DeleteOryAccessControlPolicyCreated model module.
   * @module model/DeleteOryAccessControlPolicyCreated
   * @version Latest
   */

  /**
   * Constructs a new <code>DeleteOryAccessControlPolicyCreated</code>.
   * An empty response
   * @alias module:model/DeleteOryAccessControlPolicyCreated
   * @class
   */
  var exports = function() {
    var _this = this;

  };

  /**
   * Constructs a <code>DeleteOryAccessControlPolicyCreated</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/DeleteOryAccessControlPolicyCreated} obj Optional instance to populate.
   * @return {module:model/DeleteOryAccessControlPolicyCreated} The populated <code>DeleteOryAccessControlPolicyCreated</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

    }
    return obj;
  }




  return exports;
}));


