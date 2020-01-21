/*
 * Brood Clash
 * Backend API for Brood Clash
 *
 * OpenAPI spec version: 0.1.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.4.10
 *
 * Do not edit the class manually.
 *
 */

import {ApiClient} from './ApiClient';
import {Answer} from './model/Answer';
import {Question} from './model/Question';
import {DefaultApi} from './api/DefaultApi';


/**
* Backend_API_for_Brood_Clash.<br>
* The <code>index</code> module provides access to constructors for all the classes which comprise the public API.
* <p>
* An AMD (recommended!) or CommonJS application will generally do something equivalent to the following:
* <pre>
* var BroodClash = require('index'); // See note below*.
* var xxxSvc = new BroodClash.XxxApi(); // Allocate the API class we're going to use.
* var yyyModel = new BroodClash.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* <em>*NOTE: For a top-level AMD script, use require(['index'], function(){...})
* and put the application logic within the callback function.</em>
* </p>
* <p>
* A non-AMD browser application (discouraged) might do something like this:
* <pre>
* var xxxSvc = new BroodClash.XxxApi(); // Allocate the API class we're going to use.
* var yyy = new BroodClash.Yyy(); // Construct a model instance.
* yyyModel.someProperty = 'someValue';
* ...
* var zzz = xxxSvc.doSomething(yyyModel); // Invoke the service.
* ...
* </pre>
* </p>
* @module index
* @version 0.1.0
*/
export {
    /**
     * The ApiClient constructor.
     * @property {module:ApiClient}
     */
    ApiClient,

    /**
     * The Answer model constructor.
     * @property {module:model/Answer}
     */
    Answer,

    /**
     * The Question model constructor.
     * @property {module:model/Question}
     */
    Question,

    /**
    * The DefaultApi service constructor.
    * @property {module:api/DefaultApi}
    */
    DefaultApi
};