'use strict';
// let request = require('request');
// let configFile = require(__dirname+'/../../../../../configurations/configuration.js');
let tracing = require(__dirname+'/../../../../tools/traces/trace.js');
let map_ID = require(__dirname+'/../../../../tools/map_ID/map_ID.js');
let Util = require(__dirname+'/../../../../tools/utils/util');

let user_id;
let securityContext;

let read = function (req,res,next,usersToSecurityContext)
{
    let sernum = req.params.sernum;

    tracing.create('ENTER', 'GET blockchain/orders/order/'+sernum, {});
    if(typeof req.cookies.user != 'undefined')
    {
        req.session.user = req.cookies.user;
        req.session.identity = map_ID.user_to_id(req.cookies.user);
    }

    user_id = req.session.identity;
    securityContext = usersToSecurityContext[user_id];

    return Util.queryChaincode(securityContext, 'GetOrder', [ sernum ])
    .then(function(data) {
        let order = JSON.parse(data.toString());
        let result = {};
        result.order = order;
        tracing.create('EXIT', 'GET blockchain/orders/order/'+sernum, result);
        res.send(result.order);
    })
    .catch(function(err) {
        res.status(400);
        tracing.create('ERROR', 'GET blockchain/orders/order/'+sernum, 'Unable to get vehicle. v5cID: '+ sernum);
        let error = {};
        error.message = err;
        error.sernum = sernum;
        error.error = true;
        tracing.create('ERROR', 'GET blockchain/orders/order/'+sernum, error);
        res.send(error);
    });
};

exports.read = read;
