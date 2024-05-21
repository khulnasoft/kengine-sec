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
 * @interface ModelInviteUserResponse
 */
export interface ModelInviteUserResponse {
    /**
     * 
     * @type {number}
     * @memberof ModelInviteUserResponse
     */
    invite_expiry_hours?: number;
    /**
     * 
     * @type {string}
     * @memberof ModelInviteUserResponse
     */
    invite_url?: string;
    /**
     * 
     * @type {string}
     * @memberof ModelInviteUserResponse
     */
    message?: string;
}

/**
 * Check if a given object implements the ModelInviteUserResponse interface.
 */
export function instanceOfModelInviteUserResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function ModelInviteUserResponseFromJSON(json: any): ModelInviteUserResponse {
    return ModelInviteUserResponseFromJSONTyped(json, false);
}

export function ModelInviteUserResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelInviteUserResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'invite_expiry_hours': !exists(json, 'invite_expiry_hours') ? undefined : json['invite_expiry_hours'],
        'invite_url': !exists(json, 'invite_url') ? undefined : json['invite_url'],
        'message': !exists(json, 'message') ? undefined : json['message'],
    };
}

export function ModelInviteUserResponseToJSON(value?: ModelInviteUserResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'invite_expiry_hours': value.invite_expiry_hours,
        'invite_url': value.invite_url,
        'message': value.message,
    };
}

