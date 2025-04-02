BeforeAll {
    . $PSScriptRoot/Measure-R2PO.ps1
}

Describe "Measure-R2PO" {
    
    It "Returns a valid object with expected properties" {
        $result = Measure-R2PO

        $result | Should -Not -BeNullOrEmpty
        $result | Should -BeOfType [PSCustomObject]
        $result.PSObject.Properties.Name | Should -Contain "AzCmdletCallTimeMs"
        $result.PSObject.Properties.Name | Should -Contain "SubscriptionCount"
    }

    It "Measures execution time greater than zero" {
        $result = Measure-R2PO

        $result.AzCmdletCallTimeMs | Should -BeGreaterThan 0
    }

    It "Returns a valid number of subscriptions" {
        $result = Measure-R2PO

        $result.SubscriptionCount | Should -BeGreaterThan 0
    }
}
