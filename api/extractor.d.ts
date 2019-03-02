import {ChannelCredentials, Client, ServerUnaryCall, ServiceDefinition} from "grpc";
import {DependencyManagementFile} from "./deps";

export interface ExtractRequest {
    repositoryUrl: string;
}

export interface ExtractResponse {
    repositoryUrl: string;
    managementFiles: DependencyManagementFile[];
}

export interface IDependencyExtractor {
    extract(call: ServerUnaryCall<ExtractRequest>, callback: (error: Error, response: ExtractResponse) => void): void;
}

export class DependencyExtractor extends Client {
    public static service: ServiceDefinition<IDependencyExtractor>;

    constructor(address: string, credentials: ChannelCredentials, options?: object);

    public extract(request: ExtractRequest, callback: (error: Error, response: ExtractResponse) => void): void;
}
