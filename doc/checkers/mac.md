# MAC Checker

The ```mac``` checker checks if the value is a valid an IEEE 802 MAC-48, EUI-48, EUI-64, or a 20-octet IP over InfiniBand link-layer address. If the value is not a valid MAC address, the checker will return the ```NOT_MAC``` result. Here is an example:

```golang
type Network struct {
  HardwareAddress string `checkers:"mac"`
}

network := &Network{
  HardwareAddress: "00:00:5e:00:53:01",
}

_, valid := Check(network)
if !valid {
  // Send the mistakes back to the user
}
```

In your custom checkers, you can call the ```mac``` checker function ```IsMAc``` to validate the user input. Here is an example:

```golang
result := IsMac("00:00:5e:00:53:01")

if result != ResultValid {
  // Send the mistakes back to the user
}
```
