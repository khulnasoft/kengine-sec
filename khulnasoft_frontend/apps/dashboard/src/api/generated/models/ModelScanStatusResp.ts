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
import type { ModelScanInfo } from './ModelScanInfo';
import {
    ModelScanInfoFromJSON,
    ModelScanInfoFromJSONTyped,
    ModelScanInfoToJSON,
} from './ModelScanInfo';

/**
 * 
 * @export
 * @interface ModelScanStatusResp
 */
export interface ModelScanStatusResp {
    /**
     * 
     * @type {{ [key: string]: ModelScanInfo; }}
     * @memberof ModelScanStatusResp
     */
    statuses: { [key: string]: ModelScanInfo; } | null;
}

/**
 * Check if a given object implements the ModelScanStatusResp interface.
 */
export function instanceOfModelScanStatusResp(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "statuses" in value;

    return isInstance;
}

export function ModelScanStatusRespFromJSON(json: any): ModelScanStatusResp {
    return ModelScanStatusRespFromJSONTyped(json, false);
}

export function ModelScanStatusRespFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelScanStatusResp {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'statuses': (json['statuses'] === null ? null : mapValues(json['statuses'], ModelScanInfoFromJSON)),
    };
}

export function ModelScanStatusRespToJSON(value?: ModelScanStatusResp | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'statuses': (value.statuses === null ? null : mapValues(value.statuses, ModelScanInfoToJSON)),
    };
}

