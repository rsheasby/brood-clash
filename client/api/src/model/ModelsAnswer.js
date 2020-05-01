/*
 * Brood-Clash API
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.13
 *
 * Do not edit the class manually.
 *
 */

import {ApiClient} from '../ApiClient';

/**
 * The ModelsAnswer model module.
 * @module model/ModelsAnswer
 * @version 1.0
 */
export class ModelsAnswer {
  /**
   * Constructs a new <code>ModelsAnswer</code>.
   * @alias module:model/ModelsAnswer
   * @class
   */
  constructor() {
  }

  /**
   * Constructs a <code>ModelsAnswer</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/ModelsAnswer} obj Optional instance to populate.
   * @return {module:model/ModelsAnswer} The populated <code>ModelsAnswer</code> instance.
   */
  static constructFromObject(data, obj) {
    if (data) {
      obj = obj || new ModelsAnswer();
      if (data.hasOwnProperty('id'))
        obj.id = ApiClient.convertToType(data['id'], 'String');
      if (data.hasOwnProperty('points'))
        obj.points = ApiClient.convertToType(data['points'], 'Number');
      if (data.hasOwnProperty('revealed'))
        obj.revealed = ApiClient.convertToType(data['revealed'], 'Boolean');
      if (data.hasOwnProperty('text'))
        obj.text = ApiClient.convertToType(data['text'], 'String');
    }
    return obj;
  }
}

/**
 * @member {String} id
 */
ModelsAnswer.prototype.id = undefined;

/**
 * @member {Number} points
 */
ModelsAnswer.prototype.points = undefined;

/**
 * @member {Boolean} revealed
 */
ModelsAnswer.prototype.revealed = undefined;

/**
 * @member {String} text
 */
ModelsAnswer.prototype.text = undefined;


