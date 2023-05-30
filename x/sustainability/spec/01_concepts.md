<!--
order: 1
-->

# Concepts

## The SlashRedirect Mechanism

The slash redirect mechanism was designed to:

- Redirect Slashed Tokens: Instead of the traditional method of burning slashed tokens, the SlashRedirect mechanism allocates these tokens to a designated multisig address. This process ensures that the slashed tokens are not lost, but instead, are redirected for possible future utilization.
- Customized Slash Function: In order to achieve this, we've tailored the original slash function in the staking keeper module. This adaptation allows the seamless redirection of slashed tokens to the specified multisig address.
- Isolated Module Development: To prevent potential interference with other dependencies within the app, we've intentionally avoided direct modifications to the staking types. Instead, we've developed a new module. This module, independent yet integrated, enables the updating and reading of the current multisig address slated to receive the redirected slashed tokens.

Through these strategic modifications, the SlashRedirect mechanism introduces a novel approach to the allocation of slashed tokens, thereby enhancing the overall efficiency of the system.
