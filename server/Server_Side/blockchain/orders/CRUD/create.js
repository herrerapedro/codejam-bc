'use strict';
let tracing = require(__dirname+'/../../../tools/traces/trace.js');
let map_ID = require(__dirname+'/../../../tools/map_ID/map_ID.js');
let Util = require(__dirname+'/../../../tools/utils/util');
let Order = require(__dirname+'/../../../tools/utils/order');

function create (req, res, next) {
    let user_id;

    /*if(typeof req.cookies.user !== 'undefined')
    {
        req.session.user = req.cookies.user;
        req.session.identity = map_ID.user_to_id(req.cookies.user);
    }*/
    
    //user_id = req.session.identity;
    ackcus = req.body.order.ackcus;
    ackxom = req.body.order.ackxom; 
    ackcarrier = req.body.order.ackcarrier
    cost = req.body.order.cost;

    let orderData = new Order();

    return orderData.create(ackcus, ackxom, ackcarrier, cost)
    .then(function(sernum) {
        tracing.create('INFO', 'POST blockchain/orders', 'Created order');
        let result = {};
        result.message = 'Order creation Confirmed';
        result.sernum = sernum;
        res.end(JSON.stringify(result));
    })
    .catch(function(err) {
        tracing.create('ERROR', 'POST blockchain/orders', err.stack);
        res.send(JSON.stringify({'message':err.stack}));
    });
}

exports.create = create;