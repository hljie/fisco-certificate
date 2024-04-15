```shell
 ./solc-0.8.11 --bin --abi -o ./certificate/contract/ ./certificate/contract/certificate.sol 
./abigen --bin ./certificate/contract/CertificateManagement.bin --abi ./certificate/contract/CertificateManagement.abi --pkg sdk --type CertificateManagement --out ./certificate/CertificateManagement.go
```