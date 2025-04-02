function Measure-R2PO {
    <#
    .SYNOPSIS
        Measures latency and retrieves Azure subscriptions.

    .DESCRIPTION
        This function benchmarks the time taken to fetch the list of 
        available subscriptions.

    .NOTES
        - Ensure authentication via `Connect-AzAccount` before running.

    .OUTPUTS
        PSCustomObject containing subscription details and API call timing.

    .EXAMPLE
        Measure-R2PO
        Retrieves all subscriptions while measuring latency.
    #>

    [CmdletBinding()]
    param()

    try {
        # Start timing PowerShell Az Cmdlet
        $azCmdletStartTime = Get-Date

        # Get Azure Subscription(s)
        $subscriptions = Get-AzSubscription -ErrorAction Stop

        # Stop timing PowerShell Az Cmdlet
        $azCmdletEndTime = Get-Date
        $azCmdletTime = $azCmdletEndTime - $azCmdletStartTime

        # Output results
        return [PsCustomObject]@{
            AzCmdletCallTimeMs = $azCmdletTime.TotalMilliseconds
            SubscriptionCount = $subscriptions.Count
        }
    }
    catch {
        Write-Error "Error retrieving Azure subscriptions: $_"
        return $null
    }
}
