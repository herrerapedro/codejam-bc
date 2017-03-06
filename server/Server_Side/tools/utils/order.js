'use strict';

const Util = require('./util.js');
const hfc = require('hfc');

class Order {

    constructor() {
        //this.usersToSecurityContext = usersToSecurityContext;
        this.chain = hfc.getChain('myChain'); //TODO: Make this a config param?
    }

    create(ackcus, ackxom, ackcarrier, cost) {
        //let securityContext = this.usersToSecurityContext[userId];
        let sernum = Order.newSerNum();

        return this.doesSerNumExist(userId, sernum)
        .then(function() {
            return Util.invokeChaincode(securityContext, 'CreateOrder', [ sernum, ackcus, ackxom, ackcarrier, cost])
            .then(function() {
                return sernum;
            });
        });
    }

    /*transfer(userId, buyer, functionName, v5cID) {
        return this.updateAttribute(userId, functionName , buyer, v5cID);
    }*/

    updateAttribute(userId, functionName, value, sernum) {
        let securityContext = this.usersToSecurityContext[userId];
        return Util.invokeChaincode(securityContext, functionName, [ value, sernum ]);
    }

    doeSerNumExist(userId, sernum) {
        let securityContext = this.usersToSecurityContext[userId];
        return Util.queryChaincode(securityContext, 'ReadOrder', [ sernum ]);
    }

    static newSerNum() {
        let numbers = '1234567890';
        let characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
        let sernum = '';
        for(let i = 0; i < 7; i++)
            {
            sernum += numbers.charAt(Math.floor(Math.random() * numbers.length));
        }
        sernum = characters.charAt(Math.floor(Math.random() * characters.length)) + sernum;
        sernum = characters.charAt(Math.floor(Math.random() * characters.length)) + sernum;
        return sernum;
    }
}

module.exports = Order;
