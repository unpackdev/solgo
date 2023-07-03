# Contract Information Parser

The contract information parser in SolGo allows you to extract valuable information from Solidity contracts. This information can be used for various purposes such as documentation generation, analysis, or contract interaction. The parser can extract details such as comments, license, pragmas, imports, contract name, and implemented interfaces.

Here's an example of how to use the contract information parser:

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/txpull/solgo"
	"github.com/txpull/solgo/contracts"
)

var contract = `
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// Some additional comments that can be extracted

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol";

contract MyToken is Initializable, ERC20Upgradeable, AccessControlUpgradeable, PausableUpgradeable {
    using SafeERC20Upgradeable for IERC20Upgradeable;
}
`

func main() {
	parser, err := solgo.New(context.Background(), strings.NewReader(contract))
	if err != nil {
		panic(err)
	}

	// Register the contract information listener
	contractListener := contracts.NewContractListener(parser.GetParser())
	if err := parser.RegisterListener(solgo.ListenerContractInfo, contractListener); err != nil {
		panic(err)
	}

	if errs := parser.Parse(); len(errs) > 0 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return
	}

	jsonResponse, _ := json.MarshalIndent(contractListener.ToStruct(), "", "  ")
	fmt.Println(string(jsonResponse))
}
```


In this example, we have a Solidity contract represented by the contract variable. We create a new solgo.Parser by passing the contract source code as a strings.Reader to the solgo.New function.

Next, we register a contract information listener by creating a new instance of contracts.ContractListener and passing the parser's underlying antlr.Parser as an argument. We then register the listener using parser.RegisterListener with the solgo.ListenerContractInfo listener type.

After registering the listener, we call parser.Parse() to start the parsing process. If any errors occur during parsing, we print them to the console.

Finally, we retrieve the contract information using contractListener.ToStruct(), which returns a structured representation of the contract information. We marshal the result into JSON format for display.

The response from the example above would be:

```json
{
  "comments": [
    "// Some additional comments that can be extracted"
  ],
  "license": "MIT",
  "pragmas": [
    "solidity ^0.8.0"
  ],
  "imports": [
    "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol",
    "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol",
    "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol",
    "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol",
    "@openzeppelin/contracts-upgradeable/token/ERC20/utils/SafeERC20Upgradeable.sol"
  ],
  "name": "MyToken",
  "implements": [
    "Initializable",
    "ERC20Upgradeable",
    "AccessControlUpgradeable",
    "PausableUpgradeable",
    "SafeERC20Upgradeable"
  ]
}
```

