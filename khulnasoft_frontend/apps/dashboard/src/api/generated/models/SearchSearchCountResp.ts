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
 * @interface SearchSearchCountResp
 */
export interface SearchSearchCountResp {
    /**
     * 
     * @type {{ [key: string]: number; }}
     * @memberof SearchSearchCountResp
     */
    categories: { [key: string]: number; } | null;
    /**
     * 
     * @type {number}
     * @memberof SearchSearchCountResp
     */
    count: number;
}

/**
 * Check if a given object implements the SearchSearchCountResp interface.
 */
export function instanceOfSearchSearchCountResp(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "categories" in value;
    isInstance = isInstance && "count" in value;

    return isInstance;
}

export function SearchSearchCountRespFromJSON(json: any): SearchSearchCountResp {
    return SearchSearchCountRespFromJSONTyped(json, false);
}

export function SearchSearchCountRespFromJSONTyped(json: any, ignoreDiscriminator: boolean): SearchSearchCountResp {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'categories': json['categories'],
        'count': json['count'],
    };
}

export function SearchSearchCountRespToJSON(value?: SearchSearchCountResp | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'categories': value.categories,
        'count': value.count,
    };
}

