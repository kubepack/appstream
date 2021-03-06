// Code generated by protoc-gen-grpc-js-client
// source: metdata.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function metadataGet(p, conf) {
    path = '/api/apps/v1beta1/metadata/' + p['type']
    delete p['type']
    return xhr(path, 'GET', conf, p);
}

var services = {
    metadata: {
        get: metadataGet
    }
};

module.exports = {appscode: {apps: {v1beta1: services}}};
