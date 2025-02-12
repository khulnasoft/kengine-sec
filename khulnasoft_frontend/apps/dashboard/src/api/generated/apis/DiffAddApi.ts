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
  ModelScanCompareReq,
  ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance,
  ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCompliance,
  ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelMalware,
  ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelSecret,
  ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelVulnerability,
} from '../models';
import {
    ApiDocsBadRequestResponseFromJSON,
    ApiDocsBadRequestResponseToJSON,
    ApiDocsFailureResponseFromJSON,
    ApiDocsFailureResponseToJSON,
    ModelScanCompareReqFromJSON,
    ModelScanCompareReqToJSON,
    ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudComplianceFromJSON,
    ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudComplianceToJSON,
    ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelComplianceFromJSON,
    ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelComplianceToJSON,
    ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelMalwareFromJSON,
    ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelMalwareToJSON,
    ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelSecretFromJSON,
    ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelSecretToJSON,
    ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelVulnerabilityFromJSON,
    ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelVulnerabilityToJSON,
} from '../models';

export interface DiffAddCloudComplianceRequest {
    modelScanCompareReq?: ModelScanCompareReq;
}

export interface DiffAddComplianceRequest {
    modelScanCompareReq?: ModelScanCompareReq;
}

export interface DiffAddMalwareRequest {
    modelScanCompareReq?: ModelScanCompareReq;
}

export interface DiffAddSecretRequest {
    modelScanCompareReq?: ModelScanCompareReq;
}

export interface DiffAddVulnerabilityRequest {
    modelScanCompareReq?: ModelScanCompareReq;
}

/**
 * DiffAddApi - interface
 * 
 * @export
 * @interface DiffAddApiInterface
 */
export interface DiffAddApiInterface {
    /**
     * Get Cloud Compliance Diff between two scans
     * @summary Get Cloud Compliance Diff
     * @param {ModelScanCompareReq} [modelScanCompareReq] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DiffAddApiInterface
     */
    diffAddCloudComplianceRaw(requestParameters: DiffAddCloudComplianceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance>>;

    /**
     * Get Cloud Compliance Diff between two scans
     * Get Cloud Compliance Diff
     */
    diffAddCloudCompliance(requestParameters: DiffAddCloudComplianceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance>;

    /**
     * Get Compliance Diff between two scans
     * @summary Get Compliance Diff
     * @param {ModelScanCompareReq} [modelScanCompareReq] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DiffAddApiInterface
     */
    diffAddComplianceRaw(requestParameters: DiffAddComplianceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCompliance>>;

    /**
     * Get Compliance Diff between two scans
     * Get Compliance Diff
     */
    diffAddCompliance(requestParameters: DiffAddComplianceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCompliance>;

    /**
     * Get Malware Diff between two scans
     * @summary Get Malware Diff
     * @param {ModelScanCompareReq} [modelScanCompareReq] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DiffAddApiInterface
     */
    diffAddMalwareRaw(requestParameters: DiffAddMalwareRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelMalware>>;

    /**
     * Get Malware Diff between two scans
     * Get Malware Diff
     */
    diffAddMalware(requestParameters: DiffAddMalwareRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelMalware>;

    /**
     * Get Secret Diff between two scans
     * @summary Get Secret Diff
     * @param {ModelScanCompareReq} [modelScanCompareReq] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DiffAddApiInterface
     */
    diffAddSecretRaw(requestParameters: DiffAddSecretRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelSecret>>;

    /**
     * Get Secret Diff between two scans
     * Get Secret Diff
     */
    diffAddSecret(requestParameters: DiffAddSecretRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelSecret>;

    /**
     * Get Vulnerability Diff between two scans
     * @summary Get Vulnerability Diff
     * @param {ModelScanCompareReq} [modelScanCompareReq] 
     * @param {*} [options] Override http request option.
     * @throws {RequiredError}
     * @memberof DiffAddApiInterface
     */
    diffAddVulnerabilityRaw(requestParameters: DiffAddVulnerabilityRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelVulnerability>>;

    /**
     * Get Vulnerability Diff between two scans
     * Get Vulnerability Diff
     */
    diffAddVulnerability(requestParameters: DiffAddVulnerabilityRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelVulnerability>;

}

/**
 * 
 */
export class DiffAddApi extends runtime.BaseAPI implements DiffAddApiInterface {

    /**
     * Get Cloud Compliance Diff between two scans
     * Get Cloud Compliance Diff
     */
    async diffAddCloudComplianceRaw(requestParameters: DiffAddCloudComplianceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance>> {
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
            path: `/khulnasoft/diff-add/cloud-compliance`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ModelScanCompareReqToJSON(requestParameters.modelScanCompareReq),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudComplianceFromJSON(jsonValue));
    }

    /**
     * Get Cloud Compliance Diff between two scans
     * Get Cloud Compliance Diff
     */
    async diffAddCloudCompliance(requestParameters: DiffAddCloudComplianceRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance> {
        const response = await this.diffAddCloudComplianceRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get Compliance Diff between two scans
     * Get Compliance Diff
     */
    async diffAddComplianceRaw(requestParameters: DiffAddComplianceRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCompliance>> {
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
            path: `/khulnasoft/diff-add/compliance`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ModelScanCompareReqToJSON(requestParameters.modelScanCompareReq),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelComplianceFromJSON(jsonValue));
    }

    /**
     * Get Compliance Diff between two scans
     * Get Compliance Diff
     */
    async diffAddCompliance(requestParameters: DiffAddComplianceRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCompliance> {
        const response = await this.diffAddComplianceRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get Malware Diff between two scans
     * Get Malware Diff
     */
    async diffAddMalwareRaw(requestParameters: DiffAddMalwareRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelMalware>> {
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
            path: `/khulnasoft/diff-add/malware`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ModelScanCompareReqToJSON(requestParameters.modelScanCompareReq),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelMalwareFromJSON(jsonValue));
    }

    /**
     * Get Malware Diff between two scans
     * Get Malware Diff
     */
    async diffAddMalware(requestParameters: DiffAddMalwareRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelMalware> {
        const response = await this.diffAddMalwareRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get Secret Diff between two scans
     * Get Secret Diff
     */
    async diffAddSecretRaw(requestParameters: DiffAddSecretRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelSecret>> {
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
            path: `/khulnasoft/diff-add/secret`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ModelScanCompareReqToJSON(requestParameters.modelScanCompareReq),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelSecretFromJSON(jsonValue));
    }

    /**
     * Get Secret Diff between two scans
     * Get Secret Diff
     */
    async diffAddSecret(requestParameters: DiffAddSecretRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelSecret> {
        const response = await this.diffAddSecretRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get Vulnerability Diff between two scans
     * Get Vulnerability Diff
     */
    async diffAddVulnerabilityRaw(requestParameters: DiffAddVulnerabilityRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelVulnerability>> {
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
            path: `/khulnasoft/diff-add/vulnerability`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ModelScanCompareReqToJSON(requestParameters.modelScanCompareReq),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelVulnerabilityFromJSON(jsonValue));
    }

    /**
     * Get Vulnerability Diff between two scans
     * Get Vulnerability Diff
     */
    async diffAddVulnerability(requestParameters: DiffAddVulnerabilityRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelVulnerability> {
        const response = await this.diffAddVulnerabilityRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
