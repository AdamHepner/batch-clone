[![Build Status](https://dev.azure.com/adam0942/adam/_apis/build/status/AdamHepner.batch-clone?branchName=master)](https://dev.azure.com/adam0942/adam/_build/latest?definitionId=1&branchName=master)

# batch-clone
A small utility to clone all repos of given GitHub user

# Prerequisites
Generate [Personal Access Token](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line)

# How to use it

- Download and extract the binary from the releases page
- `./batch-clone --username <your-github-username> --access-token <your-personal-access-token>`

Please note - if you've never cloned anything from GitHub, you'll be asked to confirm the server fingerprint first. If you don't care (be warned), you can first follow [this advice](https://serverfault.com/a/631149) of more clever people than me on how to trust github's key in the first place.