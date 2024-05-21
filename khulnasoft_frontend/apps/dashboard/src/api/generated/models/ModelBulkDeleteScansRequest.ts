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
import type { ReportersFieldsFilters } from './ReportersFieldsFilters';
import {
    ReportersFieldsFiltersFromJSON,
    ReportersFieldsFiltersFromJSONTyped,
    ReportersFieldsFiltersToJSON,
} from './ReportersFieldsFilters';

/**
 * 
 * @export
 * @interface ModelBulkDeleteScansRequest
 */
export interface ModelBulkDeleteScansRequest {
    /**
     * 
     * @type {ReportersFieldsFilters}
     * @memberof ModelBulkDeleteScansRequest
     */
    filters: ReportersFieldsFilters;
    /**
     * 
     * @type {string}
     * @memberof ModelBulkDeleteScansRequest
     */
    scan_type: ModelBulkDeleteScansRequestScanTypeEnum;
}


/**
 * @export
 */
export const ModelBulkDeleteScansRequestScanTypeEnum = {
    Secret: 'Secret',
    Vulnerability: 'Vulnerability',
    Malware: 'Malware',
    Compliance: 'Compliance',
    CloudCompliance: 'CloudCompliance'
} as const;
export type ModelBulkDeleteScansRequestScanTypeEnum = typeof ModelBulkDeleteScansRequestScanTypeEnum[keyof typeof ModelBulkDeleteScansRequestScanTypeEnum];


/**
 * Check if a given object implements the ModelBulkDeleteScansRequest interface.
 */
export function instanceOfModelBulkDeleteScansRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "filters" in value;
    isInstance = isInstance && "scan_type" in value;

    return isInstance;
}

export function ModelBulkDeleteScansRequestFromJSON(json: any): ModelBulkDeleteScansRequest {
    return ModelBulkDeleteScansRequestFromJSONTyped(json, false);
}

export function ModelBulkDeleteScansRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelBulkDeleteScansRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'filters': ReportersFieldsFiltersFromJSON(json['filters']),
        'scan_type': json['scan_type'],
    };
}

export function ModelBulkDeleteScansRequestToJSON(value?: ModelBulkDeleteScansRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'filters': ReportersFieldsFiltersToJSON(value.filters),
        'scan_type': value.scan_type,
    };
}

