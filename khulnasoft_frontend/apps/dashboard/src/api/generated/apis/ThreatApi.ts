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


import * as runtime from '../runtime';
import type {
  ApiDocsBadRequestResponse,
  ApiDocsFailureResponse,
  GraphIndividualThreatGraph,
  GraphIndividualThreatGraphRequest,
  GraphProviderThreatGraph,
  GraphThreatFilters,
} from '../models';
import {
    ApiDocsBadRequestResponseFromJSON,
    ApiDocsBadRequestResponseToJSON,
    ApiDocsFailureResponseFromJSON,
    ApiDocsFailureResponseToJSON,
    GraphIndividualThreatGraphFromJSON,
    GraphIndividualThreatGraphToJSON,
    GraphIndividualThreatGraphRequestFromJSON,
    GraphIndividualThreatGraphRequestToJSON,
    GraphProviderThreatGraphFromJSON,
    GraphProviderThreatGraphToJSON,
    GraphThreatFiltersFromJSON,
    GraphThreatFiltersToJSON,
} from '../models';

export interface GetIndividualThreatGraphRequest {
    graphIndividualThreatGraphRequest?: GraphIndividualThreatGraphRequest;
}

export interface GetThreatGraphRequest {
    graphThreatFilters?: GraphThreatFilters;
}

/**
 * ThreatApi - interface
 * 
 * @export
 * @interface ThreatApiInterface
 */
export interface ThreatApiInterface {
    /**
     * Retrieve threat graph associated with vulnerabilities
     * @summary Get Vulnerability Threat Graph
     * @param {GraphIndividualThreatGraphRequest} [graphIndividualThreatGraphRequest] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ThreatApiInterface
     */
    getIndividualThreatGraphRaw(requestParameters: GetIndividualThreatGraphRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<GraphIndividualThreatGraph>>>;

    /**
     * Retrieve threat graph associated with vulnerabilities
     * Get Vulnerability Threat Graph
     */
    getIndividualThreatGraph(requestParameters: GetIndividualThreatGraphRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<GraphIndividualThreatGraph>>;

    /**
     * Retrieve the full threat graph associated with the account
     * @summary Get Threat Graph
     * @param {GraphThreatFilters} [graphThreatFilters] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof ThreatApiInterface
     */
    getThreatGraphRaw(requestParameters: GetThreatGraphRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{ [key: string]: GraphProviderThreatGraph; }>>;

    /**
     * Retrieve the full threat graph associated with the account
     * Get Threat Graph
     */
    getThreatGraph(requestParameters: GetThreatGraphRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{ [key: string]: GraphProviderThreatGraph; }>;

}

/**
 * 
 */
export class ThreatApi extends runtime.BaseAPI implements ThreatApiInterface {

    /**
     * Retrieve threat graph associated with vulnerabilities
     * Get Vulnerability Threat Graph
     */
    async getIndividualThreatGraphRaw(requestParameters: GetIndividualThreatGraphRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<GraphIndividualThreatGraph>>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearer_token", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/khulnasoft/graph/threat/individual`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: GraphIndividualThreatGraphRequestToJSON(requestParameters.graphIndividualThreatGraphRequest),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(GraphIndividualThreatGraphFromJSON));
    }

    /**
     * Retrieve threat graph associated with vulnerabilities
     * Get Vulnerability Threat Graph
     */
    async getIndividualThreatGraph(requestParameters: GetIndividualThreatGraphRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<GraphIndividualThreatGraph>> {
        const response = await this.getIndividualThreatGraphRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Retrieve the full threat graph associated with the account
     * Get Threat Graph
     */
    async getThreatGraphRaw(requestParameters: GetThreatGraphRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<{ [key: string]: GraphProviderThreatGraph; }>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        if (this.configuration && this.configuration.accessToken) {
            const token = this.configuration.accessToken;
            const tokenString = await token("bearer_token", []);

            if (tokenString) {
                headerParameters["Authorization"] = `Bearer ${tokenString}`;
            }
        }
        const response = await this.request({
            path: `/khulnasoft/graph/threat`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: GraphThreatFiltersToJSON(requestParameters.graphThreatFilters),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => runtime.mapValues(jsonValue, GraphProviderThreatGraphFromJSON));
    }

    /**
     * Retrieve the full threat graph associated with the account
     * Get Threat Graph
     */
    async getThreatGraph(requestParameters: GetThreatGraphRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<{ [key: string]: GraphProviderThreatGraph; }> {
        const response = await this.getThreatGraphRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
