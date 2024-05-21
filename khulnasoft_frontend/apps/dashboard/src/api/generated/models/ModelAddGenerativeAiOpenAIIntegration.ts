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
 * @interface ModelAddGenerativeAiOpenAIIntegration
 */
export interface ModelAddGenerativeAiOpenAIIntegration {
    /**
     * 
     * @type {string}
     * @memberof ModelAddGenerativeAiOpenAIIntegration
     */
    api_key: string;
    /**
     * 
     * @type {string}
     * @memberof ModelAddGenerativeAiOpenAIIntegration
     */
    model_id: ModelAddGenerativeAiOpenAIIntegrationModelIdEnum;
}


/**
 * @export
 */
export const ModelAddGenerativeAiOpenAIIntegrationModelIdEnum = {
    Gpt4: 'gpt-4'
} as const;
export type ModelAddGenerativeAiOpenAIIntegrationModelIdEnum = typeof ModelAddGenerativeAiOpenAIIntegrationModelIdEnum[keyof typeof ModelAddGenerativeAiOpenAIIntegrationModelIdEnum];


/**
 * Check if a given object implements the ModelAddGenerativeAiOpenAIIntegration interface.
 */
export function instanceOfModelAddGenerativeAiOpenAIIntegration(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "api_key" in value;
    isInstance = isInstance && "model_id" in value;

    return isInstance;
}

export function ModelAddGenerativeAiOpenAIIntegrationFromJSON(json: any): ModelAddGenerativeAiOpenAIIntegration {
    return ModelAddGenerativeAiOpenAIIntegrationFromJSONTyped(json, false);
}

export function ModelAddGenerativeAiOpenAIIntegrationFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelAddGenerativeAiOpenAIIntegration {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'api_key': json['api_key'],
        'model_id': json['model_id'],
    };
}

export function ModelAddGenerativeAiOpenAIIntegrationToJSON(value?: ModelAddGenerativeAiOpenAIIntegration | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'api_key': value.api_key,
        'model_id': value.model_id,
    };
}

