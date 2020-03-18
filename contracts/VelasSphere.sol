//TODO general
// Make the upgradable contract
// Check how ofter the node inside the majority. in case when it is offline for long time exclude it. it losses membershipFee

pragma solidity >=0.4.22 <0.6.0;

contract VelasSphere {
    uint membershipFee = 100000000000;
    // maybe temporary
    uint gracePeriod = 50;
    uint nodeCount;
    uint minNodesVoices;
    uint customerCount;
    uint poolCount;

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
        // 2/3 of 94 nodes. Need to be count?
        minNodesVoices = 62;
    }

    struct Location {
        uint pool; //Pool of 94 nodes joined together.
        uint place; //One of 94 position in that generation. Once all of 94 positon is busy in current generation, so need to move to next generation
    }

    struct Pool {
        uint poolID; // number of pool
        uint nodeCount;
    }

    mapping (uint => Pool) pools;

    struct Node {
        address addr;
        uint balance;
        bool active;
        Location location;
    }

    mapping(address => Node) nodes;

    struct Customer {
        Pricing pricing;
        // customer choose which pool can create invoices
        bool specificPool;
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
        Customer storage current = customers[msg.sender];
        current.pricing.keepPerByte = _keepPerByte;
        current.pricing.writePerByte = _writePerByte;
        current.pricing.GPUTPerCycle = _GPUTPerCycle;
        current.pricing.CPUTtPerCycle = _CPUTtPerCycle;
        current.pricing.isChanged = true;
    }

    function deposit() internal {
        Customer storage current = customers[msg.sender];
        require(msg.value > 0);
        current.balance += msg.value;
        if (current.registered == false) {
        current.pricing = defaultPricing;
            customerCount += 1;
        }
        current.registered = true;
    }

    //pull - user can define a specific pool. if he defines 0 then all pools
    //_places - user can define a specific places in a pool. if 0 all places
    function depositWithNodes(uint _pull, uint _places) public payable {
        deposit();
        changePool(_pull, _places);
    }

    function changePool(uint _pull, uint _places) public {
        Customer storage current = customers[msg.sender];
        current.location.place = _places;
        current.location.pool = _pull;
        current.specificPool = true;
    }

    struct Invoice {
        uint height_start;
        uint height_end;
        address user;
        Pricing pricing;
        Resources used;
        uint price;
        uint voices;
    }

    struct Resources {
        uint keepPerByte;
        uint writePerByte;
        uint GPUTPerCycle;
        uint CPUTtPerCycle;
    }

    mapping(address => Invoice) invoices;

    //node sends the invoice to decrease the balance of the customer
    function createInvoice(uint height_start, uint height_end, address user, uint keepPerByte, uint writePerByte, uint GPUTPerCycle, uint CPUTtPerCycle) public  {
            Node storage node = nodes[msg.sender];
            require(node.active);
            //check if node allowed to create invoice
            Customer storage customer = customers[user];
            if (customer.specificPool == true) {
                require(customer.location.pool == node.location.pool);
            }

            Invoice storage invoice = invoices[user];
            if (invoice.voices == 0) {
               invoice.height_start = height_start;
               invoice.height_end = height_end;
               invoice.user = user;
               invoice.used.keepPerByte = keepPerByte;
               invoice.used.writePerByte = writePerByte;
               invoice.used.GPUTPerCycle = GPUTPerCycle;
               invoice.used.CPUTtPerCycle = CPUTtPerCycle;
            } else {
               require(invoice.height_start == height_start);
               require(invoice.height_end == height_end);
               require(invoice.used.keepPerByte == keepPerByte);
               require(invoice.used.writePerByte == writePerByte);
               require(invoice.used.GPUTPerCycle == GPUTPerCycle);
               require(invoice.used.CPUTtPerCycle == CPUTtPerCycle);
            }

            invoice.voices += 1;
            uint price;
            price = calculatePrice(customer.pricing, keepPerByte, writePerByte, GPUTPerCycle, CPUTtPerCycle);

            //check if the grace period has past
            if (block.number >= height_end + gracePeriod) {
                if (invoice.voices >= minNodesVoices) {
                closeInvoice(user, price);
                return;
            }
                //grace period past, but there is not enough voices. Invoice is invalid
                delete invoices[user];
                return;

            }
            // invoice has 100% of voices before grace period
            if (invoice.voices == 94) {
                closeInvoice(user, price);
            }
    }

    function calculatePrice(Pricing storage pricing, uint keepPerByte, uint writePerByte, uint GPUTPerCycle, uint CPUTtPerCycle) internal returns (uint) {
        uint price;
        price += pricing.keepPerByte * keepPerByte;
        price += pricing.writePerByte * writePerByte;
        price += pricing.CPUTtPerCycle * CPUTtPerCycle;
        price += pricing.GPUTPerCycle * GPUTPerCycle;
        return price;
    }

    function closeInvoice(address user, uint price) internal {
        //Nodes cannot work more than user requested
        require(price <= customers[user].balance);
        customers[user].balance -= price;
        delete invoices[user];
    }

    function getNextBitPosition() internal returns (uint) {
        Pool storage pool = pools[poolCount];
        if (pool.nodeCount >= 94) {
               poolCount += 1;
        }

        uint position;
            position = 1 << pool.nodeCount;
            pool.nodeCount += 1;
           return position;

    }

    function registerNode(address addr) public payable {
        Node storage node = nodes[addr];
        //TODO need to check if it exists
        require(node.active == false);
        node.active = true;
        require(msg.value == membershipFee);

        node.location.place = getNextBitPosition();
        node.location.pool = poolCount;

        nodeCount += 1;
    }

    function withdraw(address payable addr) public {
        Node storage node = nodes[addr];
        require(node.balance > 0);
        addr.transfer(nodes[addr].balance);
        node.balance = 0;
    }

    struct PricePropose {
        address proposer;
        Pricing pricing;
        uint votes;
        uint height_start;
    }

    mapping (address => PricePropose) proposes;

    function createPricePropose(uint _height_start, uint _keepPerByte, uint _writePerByte, uint _GPUTPerCycle, uint _CPUTtPerCycle) public {
        //TODO validate _height_start ?
        PricePropose storage propose = proposes[msg.sender];
        //one propose for one address
        require(propose.votes == 0);
        //TODO check if node wasn't banned

        propose.pricing.keepPerByte = _keepPerByte;
        propose.pricing.writePerByte = _writePerByte;
        propose.pricing.GPUTPerCycle = _GPUTPerCycle;
        propose.pricing.CPUTtPerCycle = _CPUTtPerCycle;
        propose.height_start = _height_start;
        propose.votes +=1;
    }

    function voteForPrice(address priceProposer) public {
        require(msg.sender != priceProposer);
        PricePropose storage propose = proposes[priceProposer];
        //check if propose exists
        require(propose.votes == 0);
        //check gracePeriod for voting
        propose.votes +=1;

         if (block.number >= propose.height_start + gracePeriod) {
             uint _minNodesVoices;
            //TODO caculate _minNodesVoices
                if (propose.votes >= _minNodesVoices) {
                //TODO change default price
                return;
            }
                //TODO check if there was less then 1/3 of voices - ban proposer
                delete proposes[priceProposer];
                return;

            }
            // propose has 100% of voices before grace period
            if (propose.votes == nodeCount) {
            //TODO change default price
            }
    }
}
