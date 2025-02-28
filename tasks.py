import subprocess
import webbrowser

from invoke import task


SERVICES = (
    "7080",  # Keycloak
    "9090",  # Prometheus
    "9411",  # Zipkin
    "9000",  # Sonarqube
    "3000/swagger",  # Restify Docs
)
BASE_URL = "http://localhost"


def _format_url(service: str) -> str:
    return f"{BASE_URL}:{service}"


@task
def run_servers(ctx, open_browser=True):
    result = subprocess.run(
        ["docker", "compose", "-f", "docker-compose.main.yml", "up", "-d"]
    )
    result.check_returncode()
    if open_browser:
        print("Opening broser with all services...")
        webbrowser.open(_format_url(SERVICES[0]), new=1)
        for service in SERVICES[1:]:
            webbrowser.open_new_tab(_format_url(service))


@task
def stop(ctx):
    result = subprocess.run(
        ["docker", "compose", "-f", "docker-compose.main.yml", "up", "-d"]
    )
    result.check_returncode()
    print("Services stopped")
