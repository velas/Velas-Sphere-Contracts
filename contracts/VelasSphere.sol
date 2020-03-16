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
        Location location; 
    }

    mapping(address => Node) nodes;

    struct Customer {
        Pricing pricing;
        Location location;
        uint balance;
        bool registered;
    }

    mapping(address => Customer) customers;

    function() external payable {
        deposit();
    }

    //Customer may want to increase the price to be first in list
    function proposePricing(uint _keepPerByte, uint _writePerByte, uint _GPUTPerCycle, uint _CPUTtPerCycle) public {
        //TODO maybe storage?
        Customer current = customers[msg.sender];
        current.pricing.keepPerByte = _keepPerByte;
        current.pricing.writePerByte = _writePerByte;
        current.pricing.GPUTPerCycle = _GPUTPerCycle;
        current.pricing.CPUTtPerCycle = _CPUTtPerCycle;
        current.pricing.isChanged = true;
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

    //pull - user can define a specific pool. if he defines 0 then all pools
    //_places - user can define a specific places in a pool. if 0 all places
    function depositWithNodes(uint _pull, uint _places) public payable {
        deposit();
        Customer current = customers[msg.sender];
        current.location.place = _places;
        current.location.pull = _pull;
        
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
        Node node = nodes[addr];
        //TODO need to check if it exists
        require(node.active == false);
        node.active = true;
        require(msg.value == membershipFee);
        
        node.location.pull = 0; //TODO. Increment pool when place is 95
        node.location.place = getNextBitPosition();
        nodeCount += 1;
    }

    function withdraw(address payable addr) public {
        Node node = nodes[addr];
        require(node.balance > 0);
        addr.transfer(nodes[addr].balance);
        node.balance = 0;
    }
}
