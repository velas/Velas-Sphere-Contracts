pragma solidity >=0.4.22 <0.6.0;

contract VelasTest {
    struct Pricing {
        uint keepPerByte;
        uint writePerByte;
        uint GPUTPerCycle;
        uint CPUTtPerCycle;
        bool isChanged; //was it changed by the customer;
    }

    Pricing defaultPricing; //default market pricing

    constructor() public {
        defaultPricing.keepPerByte = 1;
        defaultPricing.writePerByte = 1;
        defaultPricing.GPUTPerCycle = 1;
        defaultPricing.CPUTtPerCycle = 1;
        defaultPricing.isChanged = false;
    }

    struct Node {
        address addr;
        uint balance;
        bool active;
    }

    mapping(address => Node) nodes;

    struct Customer {
        Pricing pricing;
        mapping(address => Node) nodes;
        uint balance;
    }

    mapping(address => Customer) customers;

    function() external payable {
        Customer storage current = customers[msg.sender];
        current.balance += msg.value;
        current.pricing = defaultPricing;
    }

    //Customer may want to increase the price to be first in list
    function proposePricing(uint _keepPerByte, uint _writePerByte, uint _GPUTPerCycle, uint _CPUTtPerCycle) public {
        Customer storage current = customers[msg.sender];
        current.pricing.keepPerByte = _keepPerByte;
        current.pricing.writePerByte = _writePerByte;
        current.pricing.GPUTPerCycle = _GPUTPerCycle;
        current.pricing.CPUTtPerCycle = _CPUTtPerCycle;
        defaultPricing.isChanged = true;

        //TODO need to implement the sorting of customers to provide the priority list for resources;
    }

    function depositWithNodes(address[] memory _nodes) public payable {
        Customer storage current = customers[msg.sender];
        current.balance += msg.value;

        if (defaultPricing.isChanged == false)
            current.pricing = defaultPricing;

        //current.nodes = nodes;
    }

    struct Invoice {
        Pricing height_start;
        Pricing height_end;
        address user;
        Pricing pricing;
    }

    //node sends the invoice to decrease the balance of the customer
    function createInvoice(uint height_start, uint height_end, address user, uint keepPerByte, uint writePerByte, uint GPUTPerCycle, uint CPUTtPerCycle) public  {
            //TODO here need to implement the logic of node voting and result verification by 2/3 of voices
    }

    function registerNode(address addr) public {
        //TODO need to check if it exists
        nodes[addr].active = true;
    }

    function withdraw(address addr) public {
        Node storage node = nodes[addr];
        require(node.balance > 0);
        msg.sender.transfer(node.balance);
        node.balance = 0;
    }
}