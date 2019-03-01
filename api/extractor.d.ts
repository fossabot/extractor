import { DependencyManagementFile } from "./deps"
import { Client, ServerUnaryCall, ClientUnaryCall, ServiceDefinition, ChannelCredentials } from "grpc";

export interface ExtractRequest {
    repositoryUrl: string
}

export interface ExtractResponse {
    repositoryUrl: string
    managementFiles: DependencyManagementFile[]
}

export interface IDependencyExtractor {
    extract(call: ServerUnaryCall<ExtractRequest>, callback: (error: Error, response: ExtractResponse) => void): void;
}

export class DependencyExtractor extends Client {
    static service: ServiceDefinition<IDependencyExtractor>;

    constructor(address: string, credentials: ChannelCredentials, options?: object);

    extract(request: ExtractRequest, callback: (error: Error, response: ExtractResponse) => void): void;
}
