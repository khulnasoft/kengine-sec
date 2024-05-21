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
import type { ModelScanResultsCommon } from './ModelScanResultsCommon';
import {
    ModelScanResultsCommonFromJSON,
    ModelScanResultsCommonFromJSONTyped,
    ModelScanResultsCommonToJSON,
} from './ModelScanResultsCommon';

/**
 * 
 * @export
 * @interface ModelDownloadScanResultsResponse
 */
export interface ModelDownloadScanResultsResponse {
    /**
     * 
     * @type {ModelScanResultsCommon}
     * @memberof ModelDownloadScanResultsResponse
     */
    scan_info?: ModelScanResultsCommon;
    /**
     * 
     * @type {Array<any>}
     * @memberof ModelDownloadScanResultsResponse
     */
    scan_results?: Array<any> | null;
}

/**
 * Check if a given object implements the ModelDownloadScanResultsResponse interface.
 */
export function instanceOfModelDownloadScanResultsResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ModelDownloadScanResultsResponseFromJSON(json: any): ModelDownloadScanResultsResponse {
    return ModelDownloadScanResultsResponseFromJSONTyped(json, false);
}

export function ModelDownloadScanResultsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelDownloadScanResultsResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'scan_info': !exists(json, 'scan_info') ? undefined : ModelScanResultsCommonFromJSON(json['scan_info']),
        'scan_results': !exists(json, 'scan_results') ? undefined : json['scan_results'],
    };
}

export function ModelDownloadScanResultsResponseToJSON(value?: ModelDownloadScanResultsResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'scan_info': ModelScanResultsCommonToJSON(value.scan_info),
        'scan_results': value.scan_results,
    };
}

