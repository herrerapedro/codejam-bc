'use strict';

// let request = require('request');
// let configFile = require(__dirname+'/../../../../configurations/configuration.js');
let tracing = require(__dirname+'/../../../tools/traces/trace');
let map_ID = require(__dirname+'/../../../tools/map_ID/map_ID');
let Util = require(__dirname+'/../../../tools/utils/util');

let user_id;
let securityContext;

function get_all_orders(req, res, next, usersToSecurityContext, sernum)
{

    tracing.create('ENTER', 'GET blockchain/orders', {});

    /*if(typeof req.cookies.user !== 'undefined')
    {
        req.session.user = req.cookies.user;
        req.session.identity = map_ID.user_to_id(req.cookies.user);
    }*/
    //user_id = req.session.identity;
    securityContext = usersToSecurityContext[user_id];

    return Util.queryChaincode(securityContext, 'ReadOrder', [sernum])
    .then(function(data) {
        let order = JSON.parse(data.toString());
        console.log(car);
        res.write(JSON.stringify(order));
        
        /*orders.forEach(function(order) {
            tracing.create('INFO', 'GET blockchain/orders', JSON.stringify(order));
            res.write(JSON.stringify(order)+'&&');
        });*/
        
        tracing.create('EXIT', 'GET blockchain/orders', {});
        res.end('');
    })
    .catch(function(err) {
        res.status(400);
        let error = {};
        error.error = true;
        error.message = err;
        tracing.create('ERROR', 'GET blockchain/orders', err);
        res.end(JSON.stringify(error));
    });
}

exports.read = get_all_orders;
