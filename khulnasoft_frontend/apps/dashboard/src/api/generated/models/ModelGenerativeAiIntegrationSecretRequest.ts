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
 * @interface ModelGenerativeAiIntegrationSecretRequest
 */
export interface ModelGenerativeAiIntegrationSecretRequest {
    /**
     * 
     * @type {number}
     * @memberof ModelGenerativeAiIntegrationSecretRequest
     */
    integration_id?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelGenerativeAiIntegrationSecretRequest
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof ModelGenerativeAiIntegrationSecretRequest
     */
    query_type: ModelGenerativeAiIntegrationSecretRequestQueryTypeEnum;
}


/**
 * @export
 */
export const ModelGenerativeAiIntegrationSecretRequestQueryTypeEnum = {
    Remediation: 'remediation'
} as const;
export type ModelGenerativeAiIntegrationSecretRequestQueryTypeEnum = typeof ModelGenerativeAiIntegrationSecretRequestQueryTypeEnum[keyof typeof ModelGenerativeAiIntegrationSecretRequestQueryTypeEnum];


/**
 * Check if a given object implements the ModelGenerativeAiIntegrationSecretRequest interface.
 */
export function instanceOfModelGenerativeAiIntegrationSecretRequest(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "query_type" in value;

    return isInstance;
}

export function ModelGenerativeAiIntegrationSecretRequestFromJSON(json: any): ModelGenerativeAiIntegrationSecretRequest {
    return ModelGenerativeAiIntegrationSecretRequestFromJSONTyped(json, false);
}

export function ModelGenerativeAiIntegrationSecretRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelGenerativeAiIntegrationSecretRequest {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'integration_id': !exists(json, 'integration_id') ? undefined : json['integration_id'],
        'name': json['name'],
        'query_type': json['query_type'],
    };
}

export function ModelGenerativeAiIntegrationSecretRequestToJSON(value?: ModelGenerativeAiIntegrationSecretRequest | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'integration_id': value.integration_id,
        'name': value.name,
        'query_type': value.query_type,
    };
}

