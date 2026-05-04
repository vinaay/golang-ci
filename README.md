# 01_04 Set Up CI for Go

Starter workflows for Go projects are compact, with just a few steps to check out code, set up Go for a specific version, and then run commonly used go commands to build the project and run tests.

But you can tune up the suggested workflow by customizing setup steps and adding additional CI capabilities.

## References

| Reference | Description |
|----------|-------------|
| [actions/setup-go on GitHub Marketplace](https://github.com/marketplace/actions/setup-go-environment) | GitHub Action for setting up Go environments |
| [Revive linter for Go on GitHub Marketplace](https://github.com/marketplace/actions/revive-action) | GitHub Action for running the Revive linter for Go code |
| [actions/checkout on GitHub Marketplace](https://github.com/marketplace/actions/checkout) | GitHub Action for checking out repository code |
| [Documentation for the Go project used for this lesson](./GO_PROJECT_DETAILS.md) | Documentation for the Go project used in this lesson |
| [The updated workflow for this lesson](./go-ci-workflow.yml) | The complete workflow file for this lesson |

## Lab: Using a Go Starter Workflow with GitHub Actions

In this lab, you’ll create and run a starter GitHub Actions workflow for a Go project. You’ll diagnose a workflow failure caused by a Go version mismatch, update the workflow to use the project’s `go.mod` file, and extend the pipeline by adding a Go linter.

By the end of this lab, you’ll understand how starter workflows work, how Go versions are configured in GitHub Actions, and how to make your workflow more resilient and maintainable.

### Prerequisites

Before starting this lab, make sure you have:

- A new GitHub repository
- The Go project files from this lesson committed to the repository

### Instructions

#### Step 1: Create a Go Workflow from a Starter Template

1. Open your repository in GitHub.
2. Select the **Actions** tab.
3. Review the suggested workflows.
4. Locate the starter workflow **"Go by GitHub Actions."**.
5. Select **Configure**.

GitHub creates a starter workflow that includes steps to:

- Check out the repository
- Set up Go
- Run `go build ...`
- Run `go test ...`

Note how the workflow uses a hard-coded Go version.

#### Step 2: Commit and Run the Workflow

1. Select **Commit changes**.
2. Commit the workflow to the `main` branch.

Because the workflow is configured to run on pushes to `main`, committing the file automatically starts a workflow run.

#### Step 3: Investigate the Workflow Failure

1. Return to the **Actions** tab.
2. Select the last workflow run.
3. Wait for the job to complete.
4. Notice that the workflow fails.
5. Open the failed job and review the logs.

You’ll see an error indicating that a Go package is not available in the configured Go version.
A follow-up log entry explains the real issue: a dependency requires **Go 1.25**, but the workflow is using **Go 1.20**.

This mismatch causes the build to fail.

> [!TIP]
> As the starter workflow changes over time, different versions may be used or the workflow may run successfully.  In either case, please continue with the lab as the remaining steps still apply.

#### Step 4: Update the Workflow to Use `go.mod`

Configure the workflow to read the Go version directly from the project’s `go.mod` file.

1. Open the workflow file for editing.
2. While you're updating the workflow file, update the action versions:

   - Change `actions/checkout` to the latest major version.
   - Change `actions/setup-go` to the latest major version.

3. In the **Set up Go** step:

   - Replace `go-version` with `go-version-file`.
   - Replace the hard-coded version value with `go.mod`.

Using `go-version-file` ensures the workflow always uses the Go version required by the project, even if that version changes later.

#### Step 5: Add Go Linting with Revive

Next, you’ll extend the workflow with a linter.

1. Locate the search bar labeled **Search Marketplace for Actions**.
2. Search for **Revive**.
3. Select the Revive action from the results.
4. Copy the installation snippet.
5. Paste the step into your workflow **after** the **Set up Go** step.
6. Clean up the pasted code:

   - Fix formatting
   - Remove optional parameters
   - Remove the entire `with:` block

This adds automated linting for Go code to your CI pipeline.

#### Step 6: Commit and Re-Run the Workflow

1. Select **Commit changes**.
2. Commit the updated workflow to the `main` branch.
3. Return to the **Actions** tab.
4. Wait for the workflow run to complete.

Wait for the workflow to complete.

#### Step 8: Review the Successful Run

1. Open the successful workflow run.
2. Review the summary page.
3. Notice the Revive linting summary.
4. Open the build job.
5. Expand the **Set up Go** step.

You should see that the Go version was automatically set to **1.25**, pulled directly from the `go.mod` file.

This confirms that the workflow is now aligned with the project’s dependencies and can adapt to future version changes.

This approach keeps your CI pipeline reliable, future-proof, and easier to maintain as your Go projects evolve.

<!-- FooterStart -->
---
[← 01_03 Set Up CI for Python](../01_03_ci_for_python/README.md) | [01_05 Challenge: Build a CI Workflow for a Python Project →](../01_05_challenge_ci_workflow/README.md)
<!-- FooterEnd -->
