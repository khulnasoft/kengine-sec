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
 * @interface ModelListAgentVersionResp
 */
export interface ModelListAgentVersionResp {
    /**
     * 
     * @type {Array<string>}
     * @memberof ModelListAgentVersionResp
     */
    versions: Array<string> | null;
}

/**
 * Check if a given object implements the ModelListAgentVersionResp interface.
 */
export function instanceOfModelListAgentVersionResp(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "versions" in value;

    return isInstance;
}

export function ModelListAgentVersionRespFromJSON(json: any): ModelListAgentVersionResp {
    return ModelListAgentVersionRespFromJSONTyped(json, false);
}

export function ModelListAgentVersionRespFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelListAgentVersionResp {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'versions': json['versions'],
    };
}

export function ModelListAgentVersionRespToJSON(value?: ModelListAgentVersionResp | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'versions': value.versions,
    };
}

