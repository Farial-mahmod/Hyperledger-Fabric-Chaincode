package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Functions to work with an Asset
type SmartContract struct {
	contractapi.Contract
}

// Construction of Asset
type Asset struct {
	ID             string `json:"ID"`
	
	Owner          string `json:"owner"`
	
	AppraisedValue int    `json:"appraisedValue"`
}

// Initializing the ledger with primary information
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {

	assets := []Asset{
	
		{ID: "asset1", Owner: "DennisRitchie", AppraisedValue: 10},
		
		{ID: "asset2", Owner: "BjarneStroustrup", AppraisedValue: 10},
		
		{ID: "asset3", Owner: "RasmusLerdorf", AppraisedValue: 10},
		
		{ID: "asset4", Owner: "GuidoRossum", AppraisedValue: 10},
		
		{ID: "asset5", Owner: "LinusTorvalds", AppraisedValue: 10},
		
		{ID: "asset6", Owner: "AlanTuring", AppraisedValue: 10},
		
		{ID: "asset7", Owner: "GeoffreyHinton", AppraisedValue: 10},


	for _, asset := range assets {
		assetJSON, err := json.Marshal(asset)
		if err != nil {
			return err
		}

		err = ctx.GetStub().PutState(asset.ID, assetJSON)
		if err != nil {
			return fmt.Errorf("failed to add to world state. %v", err)
		}
	}

	return nil
}

// Creating an asset according to given values
func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, id string, owner string, appraisedValue int) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("asset %s already exists", id)
	}

	asset := Asset{
		ID:             id,

		Owner:          owner,
		
		AppraisedValue: appraisedValue,
		
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// Retrieves values from world state
func (s *SmartContract) ReadAsset(ctx contractapi.TransactionContextInterface, id string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("asset %s does not exist", id)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

// Updates an asset
func (s *SmartContract) UpdateAsset(ctx contractapi.TransactionContextInterface, id string, owner string, appraisedValue int) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("asset %s does not exist", id)
	}

	// overwriting original asset with the new one
	asset := Asset{
		ID:             id,
		Owner:          owner,
		AppraisedValue: appraisedValue,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// Deletes asset
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.AssetExists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// If the asset exists - 'true' is returned
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// Changes ownership of an asset
func (s *SmartContract) TransferAsset(ctx contractapi.TransactionContextInterface, id string, newOwner string) error {
	asset, err := s.ReadAsset(ctx, id)
	if err != nil {
		return err
	}

	asset.Owner = newOwner
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// All the world state assets are retrieved
func (s *SmartContract) GetAllAssets(ctx contractapi.TransactionContextInterface) ([]*Asset, error) {
	
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*Asset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var asset Asset
		err = json.Unmarshal(queryResponse.Value, &asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, &asset)
	}

	return assets, nil
}
