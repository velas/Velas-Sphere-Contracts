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
    //TODO is it enough?
    uint votesToBanPermanently = 5;

    constructor() public {
        // 2/3 of 94 nodes. Need to be count?
        minNodesVoices = 62;
    }

    struct Node {
        address payable staking_addr;
        address mining_addr;
        uint balance;
        bool active;
    }

    struct NodeToBan {
         address addr;
         uint votes;
    }

    mapping(address => Node) nodes;
    mapping(address => NodeToBan) banned;

    struct Customer {
        uint balance;
        bool registered;
        mapping(address => NodeToBan) banned;
    }

    mapping(address => Customer) customers;

    function() external payable {
        deposit();
    }

    function getNodeCount() external view returns(uint256) {
        return nodeCount;
    }

    function banNode(address _node) external {
         Customer storage current = customers[msg.sender];
         //it's enough in this customer's map
         current.banned[_node].votes = 1;
         banned[_node].votes += 1;

         if (banned[_node].votes == votesToBanPermanently) {
             //TODO ban permanently
         }
    }

    function deposit() internal {
        Customer storage current = customers[msg.sender];
        require(msg.value > 0);
        current.balance += msg.value;
        if (current.registered == false) {
            customerCount += 1;
        }
        current.registered = true;
    }

    struct Invoice {
        uint deadline;
        bool isOpened;
        address user;
        uint price;
    }

    mapping(address => Invoice) invoices;

    function openInvoice(address addr, uint deadline) external {
        //TODO check customer
        require(addr == msg.sender);
        Invoice storage invoice = invoices[addr];
        invoice.deadline = deadline;
        invoice.isOpened = true;
    }

    //node sends the invoice to decrease the balance of the customer
    function createInvoice(address user, uint price) external  {
            Node storage node = nodes[msg.sender];
            require(node.active);
            //TODO check if node was permanently banned
            require(banned[node.staking_addr].votes < votesToBanPermanently);
            Customer storage customer = customers[user];
            //check if user banned node
            require(customer.banned[node.staking_addr].votes == 0);

            Invoice storage invoice = invoices[user];
            require(invoice.isOpened);

            invoice.price += price;
    }


    function closeInvoice(address user, uint price) internal {
        //Nodes cannot work more than user requested
        require(price <= customers[user].balance);
        customers[user].balance -= price;
        delete invoices[user];
    }

    function isRegistered(address mining_addr) internal returns (bool) {
        return nodes[mining_addr].active;
    }

    function registerNode(address payable staking_addr, address mining_addr) external payable {
        Node storage node = nodes[mining_addr];
        //TODO need to check if it exists
        require(node.active == false);
        node.active = true;
        require(msg.value == membershipFee);

        node.staking_addr = staking_addr;

        nodeCount += 1;
    }

    function withdraw(address payable addr) external {
        //TODO check staking signature
        Node storage node = nodes[addr];
        require(node.balance > 0);
        node.staking_addr.transfer(nodes[addr].balance);
        node.balance = 0;
    }

    function changeMiningAddr(address old_addr, address new_addr) external {
        //TODO check staking signature
        Node storage node = nodes[old_addr];
        node.mining_addr = new_addr;
    }

    mapping(address => mapping(bytes32 => uint)) directories;

    function registerDirectory(bytes32 directory) external {
        directories[msg.sender][directory] = 0;
    }

    function increaseDirectoryNonce(bytes32 directory) external {
        uint nonce = directories[msg.sender][directory];
        nonce++;
        directories[msg.sender][directory] = nonce;
    }

    function getDirectoryNonce(bytes32 directory) external view returns(uint) {
        return directories[msg.sender][directory];
    }
 }
