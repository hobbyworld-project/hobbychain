// SPDX-License-Identifier: GPLv3
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract ERC721Delegate is Ownable, ERC721Enumerable, AccessControl {
	using Counters for Counters.Counter;
	using SafeMath for uint256;

	bytes32 public constant MANAGER_ROLE = keccak256("MANAGER_ROLE");
	bytes32 public constant PARTNER_ROLE = keccak256("PARTNER_ROLE");
	bytes32 public constant HOBBY_ROLE = keccak256("HOBBY_ROLE");

	error NotAnOwner();
	error NotSmartContract();
	error AlreadyRegistered();
	error Unregistered();
	error InvalidRecipient();
	error InvalidTokenId();
	error NothingToWithdraw();
	error NothingToDistribute();

	uint256 public numberOfVotingPositions = 1;
	string public baseURI;
	bool public paused = false;
	uint256 public maxMintAmount = 100;
	// uint256 public cost = 0.005 ether;
	uint256 public maxSupply = 26 + 10000;
	uint256[5] private halvingHeights = [
		9010284,
		18020568,
		27030852,
		36041136,
		45051420
	];

	uint256 public maxVote = 120;
	uint256 public minVote = 80;
	uint256 public heightDiff = 2 * 12343;

	uint256 candidateTime = 2 days;

	uint256 public blockTime = 7;
	uint256 public finalTotalReward = (500000000 * 10) ^ 18;

	uint256 humanId = 20000;
	uint256 mergeId = 10026;
	uint256 genesisId = 0;
	uint256[4] private levels = [1, 26, 10026, 20000];
	uint256[5] private awards = [
		0,
		(20000 * 10) ^ 18,
		(10000 * 10) ^ 18,
		(4973 * 10) ^ 18,
		0
	];

	enum NFTType {
		Default,
		Master,
		Slave,
		Common,
		Human
	}

	enum NodeStatus {
		Default,
		Init,
		Voting,
		Fully,
		Failed,
		Expired
	}

	enum NodeType {
		Candidate,
		Validator,
		Termina
	}

	struct NFTAttributes {
		bool isTransfer;
		bool isActive;
		bool isMint;
		bool isBurn;
		uint256 ActiveTime;
		NFTType level;
		uint256 weight;
		string nftURI;
		address beneficiary;
	}

	struct VotingPosition {
		address valAddr;
		uint256 deadline;
		uint256 startTime;
		uint256 endTime;
		uint256 lastRewardedEpoch;
		uint256 votes;
		uint256 pledgeAmount;
		NodeStatus status;
		NodeType nodeType;
		bool executed;
		bool canceled;
		// uint256[] tokenIds;
		mapping(address => uint256[]) NFTmembers;
		// mapping(address => uint256[]) members;
	}

	mapping(string => string) public aiDataMap;
	mapping(address => bool) public whitelisted;
	mapping(address => uint256) public rewardAmountMap;
	mapping(address => uint256) public releasedAmountMap;
	mapping(uint256 => NFTAttributes) internal _nftAttributes;
	mapping(address => VotingPosition) public votingPositionsValues;

	event Vote(
		address indexed _valAddr,
		address indexed _from,
		uint256 _weight,
		uint256[] _tokenIds
	);
	event Unvote(
		address indexed valAddr,
		address indexed _from,
		uint256[] _tokenIds
	);
	event CreateCandidate(address indexed valAddr, uint256 indexed amount);
	event BaseURI(string _uri);
	event Merge(
		uint256 indexed _id,
		uint256 indexed _weight,
		uint256[] _tokenIds
	);
	event ActiveToken(
		uint256 indexed tokenId,
		address indexed owner,
		NFTType tokenType,
		uint256 voteWeight
		// uint256 vestingAmount
	);
	event VoteFinish(address indexed valAddr);
	event Unbond(address valAddr);

	modifier only(address who) {
		require(msg.sender == who);
		_;
	}

	/// @dev only owner of _tokenId can call this function
	modifier onlyNftOwner(uint256 _tokenId) {
		if (ownerOf(_tokenId) != msg.sender) revert NotAnOwner();

		_;
	}

	event Deploy(address addr);

	// gov module address: 0x7b5Fe22B5446f7C62Ea27B8BD71CeF94e03f3dF2
	constructor(address _owner, address _gov) ERC721("Genesis NFT", "GNFT") {
		require(_owner != address(0), "constructor: _owner is 0x0");
		_transferOwnership(_owner);
		_setupRole(MANAGER_ROLE, _owner);
		_setupRole(HOBBY_ROLE, _gov);
		// _setupRole(HOBBY_ROLE, msg.sender);

		// Debug
		emit Deploy(address(this));
	}

	function supportsInterface(
		bytes4 interfaceId
	)
		public
		view
		override(AccessControl, ERC721Enumerable)
		returns (
			// override(ERC721, ERC721Enumerable, AccessControl)
			bool
		)
	{
		return super.supportsInterface(interfaceId);
	}

	function _baseURI() internal view virtual override returns (string memory) {
		return baseURI;
	}

	function tokenURI(
		uint256 tokenId
	) public view virtual override returns (string memory) {
		require(
			_exists(tokenId),
			"ERC721Metadata: URI query for nonexistent token"
		);

		string memory currentBaseURI = _baseURI();

		string memory Id = _nftAttributes[tokenId].nftURI;
		return
			bytes(currentBaseURI).length > 0
				? string(abi.encodePacked(currentBaseURI, Id))
				: "";
	}

	function transferFrom(
		address from,
		address to,
		uint256 tokenId
	) public virtual override(ERC721, IERC721) {
		require(
			_isApprovedOrOwner(_msgSender(), tokenId),
			"ERC721: caller is not token owner or approved"
		);

		require(tokenId < levels[3], "ERC721: tokenId err");

		require(
			_nftAttributes[tokenId].isTransfer,
			"ERC721: tokenId is not transfer"
		);

		_transfer(from, to, tokenId);
	}

	function createCandidate(
		address valAddr,
		uint256 amount
	) public onlyRole(MANAGER_ROLE) returns (uint256 votingPositionId) {
		votingPositionsValues[valAddr].valAddr = valAddr;
		votingPositionsValues[valAddr].startTime = block.timestamp;
		votingPositionsValues[valAddr].endTime =
			block.timestamp +
			candidateTime;
		votingPositionsValues[valAddr].pledgeAmount = amount;
		votingPositionsValues[valAddr].nodeType = NodeType.Candidate;
		emit CreateCandidate(valAddr, amount);
		return votingPositionId;
	}

	// Common NFT merge
	function merge(
		uint256[] memory _tokenIds,
		string memory _newURI
	) public returns (bool) {
		address from = msg.sender;
		// uint256 supply = totalSupply();
		uint256 tokenIdsLength = _tokenIds.length;

		require(tokenIdsLength <= maxSupply, "length less");

		uint256 weight = 0;
		for (uint256 i = 0; i < tokenIdsLength; i++) {
			require(
				ERC721.ownerOf(_tokenIds[i]) == from,
				"ERC721: token from owner"
			);

			require(
				_nftAttributes[_tokenIds[i]].level == NFTType.Common,
				"ERC721: level not Common"
			);

			weight += _nftAttributes[_tokenIds[i]].weight;
			_nftAttributes[_tokenIds[i]].isBurn = true;
			_burn(_tokenIds[i]);
		}

		mergeId += 1;
		_safeMint(from, mergeId);

		_nftAttributes[mergeId].nftURI = _newURI;
		_nftAttributes[mergeId].level = NFTType.Common;
		_nftAttributes[mergeId].isTransfer = false;
		_nftAttributes[mergeId].weight = weight;

		emit Merge(mergeId, weight, _tokenIds);

		return true;
	}

	function mintGenesis(
		address _to,
		string memory _newURI,
		uint256 _mintAmount
	) public payable onlyOwner {
		require(!paused);
		require(_mintAmount > 0);
		require(_mintAmount <= maxMintAmount);
		require(genesisId + _mintAmount <= maxSupply);

		for (uint256 i = 1; i <= _mintAmount; i++) {
			_safeMint(_to, genesisId + i);

			_nftAttributes[genesisId + i].nftURI = _newURI;

			if (genesisId + i == levels[0]) {
				_nftAttributes[genesisId + i].weight = 20;
				_nftAttributes[genesisId + i].level = NFTType.Master;
			} else if (
				(genesisId + i > levels[0]) && (genesisId + i <= levels[1])
			) {
				_nftAttributes[genesisId + i].weight = 10;
				_nftAttributes[genesisId + i].level = NFTType.Slave;
			} else if (
				(genesisId + i > levels[1]) && (genesisId + i <= levels[3])
			) {
				_nftAttributes[genesisId + i].weight = 5;
				_nftAttributes[genesisId + i].level = NFTType.Common;
			}

			_nftAttributes[genesisId + i].isTransfer = true;
		}

		genesisId += _mintAmount;
	}

	function getVotingMembers(
		address _valAddr,
		address _to
	) public view returns (uint256[] memory) {
		return votingPositionsValues[_valAddr].NFTmembers[_to];
	}

	function getNFTAttributes(
		uint256 _tokenId
	) public view returns (NFTAttributes memory) {
		return _nftAttributes[_tokenId];
	}

	function costPrice() public view virtual returns (uint256) {
		return cost(humanId);
	}

	function mint(address _to, string memory _newURI) public payable {
		uint256 supply = humanId;
		require(!paused);

		uint256 i = 1;

		if (msg.sender != owner()) {
			if (whitelisted[msg.sender] != true) {
				require(msg.value >= cost(supply + i));
			}
		}

		_safeMint(_to, supply + i);

		_nftAttributes[supply + i].nftURI = _newURI;

		_nftAttributes[supply + i].weight = 10;
		_nftAttributes[supply + i].level = NFTType.Human;
		_nftAttributes[supply + i].isTransfer = false;
		humanId = humanId + i;

		payable(address(0)).transfer(msg.value);
	}

	function cost(uint256 n) public view virtual returns (uint256) {
		return (3 * 10) ^ 18;
	}

	function calculateRemainingReward() public view returns (uint256) {
		uint256 currentBlockHeight = block.number;
		uint256 initialReward = finalTotalReward.div(2).div(halvingHeights[0]);
		uint256 elapsedBlocks = currentBlockHeight;

		uint256 halvingPeriods = findHalvingPeriod(currentBlockHeight);

		uint256 totalReward = initialReward / (2 ** halvingPeriods);

		uint256 distributedReward = (initialReward - totalReward) *
			elapsedBlocks;

		uint256 remainingReward = finalTotalReward - distributedReward;

		return remainingReward;
	}

	// Find the halving period where the current block height is
	function findHalvingPeriod(
		uint256 currentHeight
	) internal view returns (uint256) {
		for (uint256 i = 0; i < halvingHeights.length; i++) {
			if (currentHeight < halvingHeights[i]) {
				return i;
			}
		}
		return halvingHeights.length;
	}

	function vote(
		address _valAddr,
		uint256[] memory _tokenIds
	) public returns (bool) {
		require(
			block.timestamp < votingPositionsValues[_valAddr].endTime,
			"Vote expired"
		);
		require(
			votingPositionsValues[_valAddr].nodeType == NodeType.Candidate ||
				votingPositionsValues[_valAddr].nodeType == NodeType.Validator,
			"Node has been terminated"
		);

		address from = msg.sender;

		uint256 weightAll = 0;

		//
		for (uint i = 0; i < _tokenIds.length; i++) {
			uint256 weight = _nftAttributes[_tokenIds[i]].weight;
			require(
				votingPositionsValues[_valAddr].votes + weight <= maxVote,
				"Max votes reached"
			);

			require(
				_nftAttributes[_tokenIds[i]].isTransfer,
				"Not allowed to vote"
			);

			_nftAttributes[_tokenIds[i]].isTransfer = false;

			votingPositionsValues[_valAddr].votes += weight;
			votingPositionsValues[_valAddr].NFTmembers[from].push(_tokenIds[i]);

			weightAll += weight;
		}

		//
		if (votingPositionsValues[_valAddr].votes >= maxVote) {
			votingPositionsValues[_valAddr].nodeType = NodeType.Validator;
			emit VoteFinish(votingPositionsValues[_valAddr].valAddr);
		}

		emit Vote(
			votingPositionsValues[_valAddr].valAddr,
			from,
			weightAll,
			_tokenIds
		);

		return false;
	}

	function unvote(
		address _valAddr,
		uint256[] memory _tokenIds
	) public returns (bool) {
		address from = msg.sender;

		require(
			_tokenIds.length ==
				votingPositionsValues[_valAddr].NFTmembers[from].length,
			"_tokenIds do not match"
		);

		uint256 length = votingPositionsValues[_valAddr]
			.NFTmembers[from]
			.length;

		for (uint i = 0; i < length; i++) {
			_nftAttributes[_tokenIds[i]].isTransfer = true;

			votingPositionsValues[_valAddr].votes -= _nftAttributes[
				_tokenIds[i]
			].weight;
			votingPositionsValues[_valAddr].NFTmembers[from].pop();
		}

		delete votingPositionsValues[_valAddr].NFTmembers[from];

		emit Unvote(_valAddr, from, _tokenIds);

		return false;
	}

	function activeToken(uint256 _tokenId) public returns (uint256 tokenId) {
		address from = msg.sender;
		require(ERC721.ownerOf(_tokenId) == from, "ERC721: token from owner");

		require(
			!_nftAttributes[_tokenId].isActive,
			"ERC721: already activated "
		);
		_nftAttributes[_tokenId].isActive = true;
		_nftAttributes[_tokenId].beneficiary = from;

		emit ActiveToken(
			_tokenId,
			from,
			_nftAttributes[_tokenId].level,
			_nftAttributes[_tokenId].weight
		);
		return tokenId;
	}

	function setBaseURI(string memory _newBaseURI) public onlyOwner {
		baseURI = _newBaseURI;

		emit BaseURI(_newBaseURI);
	}

	function validatorStatus(
		address _valAddr
	) public view returns (NodeStatus) {
		// followMap[_to][from] = true;
		return votingPositionsValues[_valAddr].status;
	}

	function setValidatorStatus(
		address _valAddr,
		NodeStatus _status
	) public onlyRole(HOBBY_ROLE) {
		votingPositionsValues[_valAddr].status = _status;
	}

	function unbond(address _valAddr) public {
		require(
			votingPositionsValues[_valAddr].status == NodeStatus.Init ||
				votingPositionsValues[_valAddr].status == NodeStatus.Voting,
			"ERC721: unbond status"
		);

		if (votingPositionsValues[_valAddr].status == NodeStatus.Expired) {
			delete votingPositionsValues[_valAddr];
		}

		emit Unbond(_valAddr);
	}

	function setRewardAmounts(
		address[] memory wallets,
		uint256[] memory amounts
	) external onlyRole(HOBBY_ROLE) {
		require(
			wallets.length == amounts.length,
			"Arrays must be the same length"
		);

		for (uint256 i = 0; i < wallets.length; i++) {
			// super._transfer(msg.sender, wallets[i], amounts[i]);
			rewardAmountMap[wallets[i]] = amounts[i];
		}
	}

	function setRewardAmount(
		address _owner,
		uint256 _amount
	) external onlyRole(HOBBY_ROLE) {
		rewardAmountMap[_owner] = _amount;
	}

	function setReleasedAmount(
		address _owner,
		uint256 _amount
	) public onlyRole(HOBBY_ROLE) {
		releasedAmountMap[_owner] = _amount;
	}

	function setManager(address _owner) public onlyRole(MANAGER_ROLE) {
		_setupRole(MANAGER_ROLE, _owner);
	}

	function setAiData(
		string memory _key,
		string memory _value
	) public onlyRole(MANAGER_ROLE) {
		aiDataMap[_key] = _value;
	}

	// DEBUG

	function setHobby(address _owner) public onlyRole(MANAGER_ROLE) {
		_setupRole(HOBBY_ROLE, _owner);
	}
}
