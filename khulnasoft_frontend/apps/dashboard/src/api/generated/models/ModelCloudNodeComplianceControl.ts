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
/**
 * 
 * @export
 * @interface ModelCloudNodeComplianceControl
 */
export interface ModelCloudNodeComplianceControl {
    /**
     * 
     * @type {Array<string>}
     * @memberof ModelCloudNodeComplianceControl
     */
    category_hierarchy?: Array<string> | null;
    /**
     * 
     * @type {string}
     * @memberof ModelCloudNodeComplianceControl
     */
    control_id?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelCloudNodeComplianceControl
     */
    description?: string;
    /**
     * 
     * @type {boolean}
     * @memberof ModelCloudNodeComplianceControl
     */
    enabled?: boolean;
    /**
     * 
     * @type {string}
     * @memberof ModelCloudNodeComplianceControl
     */
    service?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelCloudNodeComplianceControl
     */
    title?: string;
}

/**
 * Check if a given object implements the ModelCloudNodeComplianceControl interface.
 */
export function instanceOfModelCloudNodeComplianceControl(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ModelCloudNodeComplianceControlFromJSON(json: any): ModelCloudNodeComplianceControl {
    return ModelCloudNodeComplianceControlFromJSONTyped(json, false);
}

export function ModelCloudNodeComplianceControlFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelCloudNodeComplianceControl {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'category_hierarchy': !exists(json, 'category_hierarchy') ? undefined : json['category_hierarchy'],
        'control_id': !exists(json, 'control_id') ? undefined : json['control_id'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'enabled': !exists(json, 'enabled') ? undefined : json['enabled'],
        'service': !exists(json, 'service') ? undefined : json['service'],
        'title': !exists(json, 'title') ? undefined : json['title'],
    };
}

export function ModelCloudNodeComplianceControlToJSON(value?: ModelCloudNodeComplianceControl | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'category_hierarchy': value.category_hierarchy,
        'control_id': value.control_id,
        'description': value.description,
        'enabled': value.enabled,
        'service': value.service,
        'title': value.title,
    };
}

