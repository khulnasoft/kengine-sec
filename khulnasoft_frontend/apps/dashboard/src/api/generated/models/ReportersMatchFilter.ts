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
 * @interface ReportersMatchFilter
 */
export interface ReportersMatchFilter {
    /**
     * 
     * @type {{ [key: string]: Array<any>; }}
     * @memberof ReportersMatchFilter
     */
    filter_in: { [key: string]: Array<any>; } | null;
}

/**
 * Check if a given object implements the ReportersMatchFilter interface.
 */
export function instanceOfReportersMatchFilter(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "filter_in" in value;

    return isInstance;
}

export function ReportersMatchFilterFromJSON(json: any): ReportersMatchFilter {
    return ReportersMatchFilterFromJSONTyped(json, false);
}

export function ReportersMatchFilterFromJSONTyped(json: any, ignoreDiscriminator: boolean): ReportersMatchFilter {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'filter_in': json['filter_in'],
    };
}

export function ReportersMatchFilterToJSON(value?: ReportersMatchFilter | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'filter_in': value.filter_in,
    };
}

