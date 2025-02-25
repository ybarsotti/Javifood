import subprocess
from invoke import task


@task
def generate_docs(ctx):
    result = subprocess.run(
        [
            "swag",
            "init",
            "-g",
            "cmd/main.go",
            "--parseInternal",
            "--output",
            "cmd/docs/",
        ]
    )
    if result.returncode != 0:
        print("Error generating docs. ", result.stderr)


@task
def format(ctx):
    subprocess.run(["gofmt", "-w", "."])


@task
def fix_imports(ctx):
    subprocess.run(["goimports", "-w", "."])


@task
def improve_indentation(ctx):
    subprocess.run(["golines", "-w", "."])


@task
def linter(ctx):
    subprocess.run(["golangci-lint", "run"])


@task
def check_errors(ctx):
    subprocess.run(["errcheck", "./..."])


@task
def exam_code(ctx):
    subprocess.run(["go", "vet", "./..."])


@task
def purify_code(ctx):
    holy_commands = [
        format,
        fix_imports,
        improve_indentation,
        check_errors,
        exam_code,
        linter,
    ]
    for command in holy_commands:
        command(ctx)
