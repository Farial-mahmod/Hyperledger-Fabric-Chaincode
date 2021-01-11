<html>
<head>
<style>
 
.button {
background-color: #000000; /* Green */
border: none;
color: white;
padding: 12px 25px;
text-align: center;
text-decoration: none;
display: inline-block;
font-size: 13px;
margin: 3px 1px;
cursor: pointer;
}
 
</style>
</head>

<body>
<center>


<?php

//Timestamp of Dhaka

date_default_timezone_set('Asia/Dhaka');
$date = date('h: i: s');
echo "As of: " . date("M d, Y") . " Time:  ". $date. "<br><br><br>";


// login.php simply takes user credentials


// Processing inputs (name, id) taken from the user

$newowner = $_POST['newowner'];
$id = $_POST['newid'];
$newid = 'asset' . $id;
$newpoint = $_POST['newpoint'];
$update = $_POST['asset'];


//Adding new asset amount

$change = `peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile /home/bdtask/h/fabric-samples/test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles /home/bdtask/h/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles /home/bdtask/h/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"UpdateAsset","Args":["{$id}","{$owner}","{$update}"]}'`;
$commit = `peer chaincode query -C mychannel -n basic -c '{"function":"UpdateAsset","Args":["{$id}","{$owner}","{$update}"]}'`;

shell_exec($change);

shell_exec($commit);

sleep(2);


// to produce / generate a new asset
$createAsset = `peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile /home/bdtask/h/fabric-samples/test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles /home/bdtask/h/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles /home/bdtask/h/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"CreateAsset","Args":["{$newid}","{$newowner}","{$newpoint}"]}'`;
echo shell_exec($createAsset);
   
sleep(2);
   
print 'New ID has been generated.' . "<br>";
$read = "peer chaincode query -C mychannel -n basic -c '{\"Args\":[\"ReadAsset\",\"$newid\"]}'";
   
// New Asset is being printed
   echo shell_exec($read);
}

// to delete an existing asset
$DeleteAsset = `peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile /home/bdtask/h/fabric-samples/test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n basic --peerAddresses localhost:7051 --tlsRootCertFiles /home/bdtask/h/fabric-samples/test-network/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles /home/bdtask/h/fabric-samples/test-network/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"function":"DeleteAsset","Args":["{$newid}"]}'`;
echo shell_exec($DeleteAsset);
   
 

?>

<form method="POST" action="login.php">
<br><br><br>
<Button class="button">Another Transaction</Button>
</form>
</center>
</body>
</html>
