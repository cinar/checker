# Credit Card Checker

The ```credit-card``` checker checks if the given value is a valid credit card number. If the given value is not a valid credit card number, the checker will return the ```NOT_CREDIT_CARD``` result. 

Here is an example:

```golang
type Order struct {
    CreditCard string `checkers:"credit-card"`
}

order := &Order{
    CreditCard: invalidCard,
}

_, valid := Check(order)
if valid {
  // Send the mistakes back to the user
}
```

The checker currently knows about AMEX, Diners, Discover, JCB, MasterCard, UnionPay, and VISA credit card numbers. 

If you would like to check for a subset of those credit cards, you can specify them through the checker config parameter. Here is an example:

```golang
type Order struct {
    CreditCard string `checkers:"credit-card:amex,visa"`
}

order := &Order{
    CreditCard: "6011111111111117",
}

_, valid := Check(order)
if valid {
  // Send the mistakes back to the user
}
```

If you would like to verify a credit card that is not listed here, please use the [luhn](luhn.md) checker to use the Luhn Algorithm to verify the check digit.

In your custom checkers, you can call the ```credit-card``` checker functions below to validate the user input. 

- IsAnyCreditCard: checks if the given value is a valid credit card number.
- IsAmexCreditCard: checks if the given value is a valid AMEX credit card number.
- IsDinersCreditCard: checks if the given value is a valid Diners credit card number.
- IsDiscoverCreditCard: checks if the given value is a valid Discover credit card number.
- IsJcbCreditCard: checks if the given value is a valid JCB credit card number.
- IsMasterCardCreditCard: checks if the given value is a valid MasterCard credit card number.
- IsUnionPayCreditCard: checks if the given value is a valid UnionPay credit card number.
- IsVisaCreditCard: checks if the given value is a valid VISA credit card number.

Here is an example:

```golang
result := IsAnyCreditCard("6011111111111117")

if result != ResultValid {
  // Send the mistakes back to the user
}
```
