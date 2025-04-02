function Measure-R2PO {
    [CmdletBinding()]
    param()

    try {
        $azCmdletStartTime = Get-Date
        $subscriptions = Get-AzSubscription -ErrorAction Stop
        $azCmdletEndTime = Get-Date
        $azCmdletTime = ($azCmdletEndTime - $azCmdletStartTime).TotalMilliseconds

        $result = @{
            SubscriptionsCallTimeMs = $azCmdletTime
            SubscriptionCount       = $subscriptions.Count
        }

        return $result | ConvertTo-Json -Depth 2
    }
    catch {
        Write-Error "Error retrieving Azure subscriptions: $_"
        return $null
    }
}
