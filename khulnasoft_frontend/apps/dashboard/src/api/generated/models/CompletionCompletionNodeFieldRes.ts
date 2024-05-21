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
 * @interface CompletionCompletionNodeFieldRes
 */
export interface CompletionCompletionNodeFieldRes {
    /**
     * 
     * @type {Array<string>}
     * @memberof CompletionCompletionNodeFieldRes
     */
    possible_values: Array<string> | null;
}

/**
 * Check if a given object implements the CompletionCompletionNodeFieldRes interface.
 */
export function instanceOfCompletionCompletionNodeFieldRes(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "possible_values" in value;

    return isInstance;
}

export function CompletionCompletionNodeFieldResFromJSON(json: any): CompletionCompletionNodeFieldRes {
    return CompletionCompletionNodeFieldResFromJSONTyped(json, false);
}

export function CompletionCompletionNodeFieldResFromJSONTyped(json: any, ignoreDiscriminator: boolean): CompletionCompletionNodeFieldRes {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'possible_values': json['possible_values'],
    };
}

export function CompletionCompletionNodeFieldResToJSON(value?: CompletionCompletionNodeFieldRes | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'possible_values': value.possible_values,
    };
}

