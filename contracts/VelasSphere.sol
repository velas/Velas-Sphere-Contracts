//TODO
// Make the upgradable contract
// Impelemnt the voting for default Pricing
// Implement invoice execution based of 2/3 of majority execution
// Check how ofter the node inside the majority. in case when it is offline for long time exclude it. it losses membershipFee

pragma solidity >=0.4.22 <0.6.0;

contract VelasSphere {
    uint membershipFee = 100000000000;
    uint nodeCount;
    uint customerCount;
    //TODO add grace period logic?
    uint gracePeriod;

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
 
    struct Location {
        uint pool; //Pool of 94 nodes joined together.
        uint place; //One of 94 position in that generation. Once all of 94 positon is busy in current generation, so need to move to next generation
    }


    struct Node {
        address addr;
        uint balance;
        bool active;
        uint generation; 
        Location location; 
    }

    mapping(address => Node) nodes;

    struct Customer {
        Pricing pricing;
        uint places; // bit positions
        uint balance;
        bool registered;
    }

    mapping(address => Customer) customers;

    function() external payable {
        deposit();
    }

    //Customer may want to increase the price to be first in list
    function proposePricing(uint _keepPerByte, uint _writePerByte, uint _GPUTPerCycle, uint _CPUTtPerCycle) public {
        //Customer current = customers[msg.sender];
        customers[msg.sender].pricing.keepPerByte = _keepPerByte;
        customers[msg.sender].pricing.writePerByte = _writePerByte;
        customers[msg.sender].pricing.GPUTPerCycle = _GPUTPerCycle;
        customers[msg.sender].pricing.CPUTtPerCycle = _CPUTtPerCycle;
        customers[msg.sender].pricing.isChanged = true;
    }

    function deposit() internal {
        //Customer current = customers[msg.sender];
        require(msg.value > 0);
        customers[msg.sender].balance += msg.value;
        if (customers[msg.sender].registered == false) {
        customers[msg.sender].pricing = defaultPricing;
            customerCount += 1;
        }
        customers[msg.sender].registered = true;
    }

    //bit places
    function depositWithNodes(uint _places) public payable {
        deposit();
        //Customer current = customers[msg.sender];
        customers[msg.sender].places = _places;
    }

    struct Invoice {
        uint height_start;
        uint height_end;
        address user;
        Pricing pricing;
        uint voices;
    }

    mapping(address => Invoice) invoices;

    //node sends the invoice to decrease the balance of the customer
    function createInvoice(uint height_start, uint height_end, address user, uint keepPerByte, uint writePerByte, uint GPUTPerCycle, uint CPUTtPerCycle) public  {
            //Node node = nodes[msg.sender];
            require(nodes[msg.sender].active);
            //TODO here need to implement the logic of node voting and result verification by 2/3 of voices
            if (invoices[user].voices == 0) {
               invoices[user].height_start = height_start;
               invoices[user].height_end = height_end;
               invoices[user].user = user;
            }
            //TODO check if invoice details are the same as previous
            invoices[user].voices += 1;

            //TODO count minNodesVoices
            uint minNodesVoices;
            if (invoices[user].voices >= minNodesVoices) {
                uint price;
                //TODO calculate price
                closeInvoice(user, price);
            }
    }

    function closeInvoice(address user, uint price) internal {
        //TODO what if balance < invoice?
        customers[user].balance -= price;
        delete invoices[user];

    }

    function getNextBitPosition() internal returns (uint) {
        //TODO need to calculate nodeCount based on current pool
        if (nodeCount > 94)
            return 0;
        return (1 << nodeCount);
    }

    function registerNode(address addr) public payable {
        //TODO need to check if it exists
        require(nodes[addr].active == false);
        nodes[addr].active = true;
        require(msg.value == membershipFee);
        //TODO nodes[node].position - implement next bit position
        nodes[addr].location.place = getNextBitPosition();
        nodeCount += 1;
    }

    function withdraw(address payable addr) public {
        //Node node = nodes[addr];
        require(nodes[addr].balance > 0);
        addr.transfer(nodes[addr].balance);
        nodes[addr].balance = 0;
    }
}
