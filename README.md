# CHAINCODE_TEMPLATES

A small repository I use for hands-on trainigs session for Hyperledger chaincode development. For more
information about chaincode development please have a look at [Hyperledger-Fabric Chaincode](https://hyperledger-fabric.readthedocs.io/en/latest/chaincode4ade.html#).

## Content
### exercise folder
Contains a chaincode file and it's unit test. Logic in those files has been replaced by `// TODO` statements for trainigs purpose.
The chaincode can not be build and the unit tests are not running out of the box. The files can not be validated due to the missing code.  

### solution folder
Contains the solution files for the chaincode exercise. The chaincode file can be compiled by `go build`  and `go test` is successful as
well.

```sh
$ go build --tags nopkcs11 -v -o chaincode
$ go test
```

This chaincode can be deployed on Hyperledger Fabric peers with version >= 1.0.


## How to use it
1. Clone this repository e.g.
 ```sh
 $ git clone git@github.com:dpdornseifer/chaincode_templates.git
 ```

2. Make sure that the dependencies for the Hyperledger Fabric shim interface are available in the ```$GOPATH``` - e.g.
 ```sh
 $ go get -u github.com/hyperledger/fabric/core/chaincode/shim
 ```

3. Implement the missing parts marked by `// TODO`.
4. Make sure that the chaincode can be build e.g. execute:
 ```sh
 $ go build --tags nopkcs11 -v -o chaincode
 ```
 in the folder. You can start the binary by doing a
 ```sh
 $ chmod +x chaincode
 ./chaincode
 ```

5. Make sure that the unit tests can be executed successful e.g.
 ```sh
 $ go test
 ```

6. Congratulations, you just finished this exercise.
  
