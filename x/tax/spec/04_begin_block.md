<!-- order: 4 -->

# Begin-Block

1. Get all the Taxes registered in params.Taxes and select only the valid taxes. If there is no valid tax, exit.
2. Create a map by `TaxSourceAddress` to handle the taxes for the same `TaxSourceAddress` together based on the same balance when calculating rates for the same `TaxSourceAddress`.
3. Collect taxes from `TaxSourceAddress` and send them to `CollectionAddress` according to the `Rate` of each `Tax`.
4. Write to metric about successful tax collection and emit events.