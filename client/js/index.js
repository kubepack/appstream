// Code generated by ./hack/browserify.py
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for AppsCode Seed gRPC JSON APIs.
*/

var _ = require('lodash');

var apis = _.merge({},
    require('./apis/app/v1beta1/metdata.gw.js'),
    require('./apis/app/v1beta1/version.gw.js')
);
module.exports = apis.appscode;
