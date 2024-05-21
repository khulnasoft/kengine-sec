/* tslint:disable */
/* eslint-disable */
/**
 * Khulnasoft Kengine
 * Khulnasoft Runtime API provides programmatic control over Khulnasoft microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.
 *
 * The version of the OpenAPI document: v2.2.1
 * Contact: community@khulnasoft.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { ModelNodeIdentifier } from './ModelNodeIdentifier';
import {
    ModelNodeIdentifierFromJSON,
    ModelNodeIdentifierFromJSONTyped,
    ModelNodeIdentifierToJSON,
} from './ModelNodeIdentifier';

/**
 * 
 * @export
 * @interface ModelTopologyDeltaResponse
 */
export interface ModelTopologyDeltaResponse {
    /**
     * 
     * @type {number}
     * @memberof ModelTopologyDeltaResponse
     */
    addition_timestamp?: number;
    /**
     * 
     * @type {Array<ModelNodeIdentifier>}
     * @memberof ModelTopologyDeltaResponse
     */
    additons?: Array<ModelNodeIdentifier> | null;
    /**
     * 
     * @type {number}
     * @memberof ModelTopologyDeltaResponse
     */
    deletion_timestamp?: number;
    /**
     * 
     * @type {Array<ModelNodeIdentifier>}
     * @memberof ModelTopologyDeltaResponse
     */
    deletions?: Array<ModelNodeIdentifier> | null;
}

/**
 * Check if a given object implements the ModelTopologyDeltaResponse interface.
 */
export function instanceOfModelTopologyDeltaResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ModelTopologyDeltaResponseFromJSON(json: any): ModelTopologyDeltaResponse {
    return ModelTopologyDeltaResponseFromJSONTyped(json, false);
}

export function ModelTopologyDeltaResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelTopologyDeltaResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'addition_timestamp': !exists(json, 'addition_timestamp') ? undefined : json['addition_timestamp'],
        'additons': !exists(json, 'additons') ? undefined : (json['additons'] === null ? null : (json['additons'] as Array<any>).map(ModelNodeIdentifierFromJSON)),
        'deletion_timestamp': !exists(json, 'deletion_timestamp') ? undefined : json['deletion_timestamp'],
        'deletions': !exists(json, 'deletions') ? undefined : (json['deletions'] === null ? null : (json['deletions'] as Array<any>).map(ModelNodeIdentifierFromJSON)),
    };
}

export function ModelTopologyDeltaResponseToJSON(value?: ModelTopologyDeltaResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'addition_timestamp': value.addition_timestamp,
        'additons': value.additons === undefined ? undefined : (value.additons === null ? null : (value.additons as Array<any>).map(ModelNodeIdentifierToJSON)),
        'deletion_timestamp': value.deletion_timestamp,
        'deletions': value.deletions === undefined ? undefined : (value.deletions === null ? null : (value.deletions as Array<any>).map(ModelNodeIdentifierToJSON)),
    };
}

