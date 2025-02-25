import subprocess
from invoke import task

# TODO: Add start commands, open browser with all links, prometheus, backend etc


@task
def run_servers(ctx):
    subprocess.run(["docker", "compose", 
                   "-f", "docker-compose.main.yml", "up"])
