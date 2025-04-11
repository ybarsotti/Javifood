import webbrowser
import time

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
    print("Starting servers...")
    ctx.run("docker compose -f docker-compose.main.yml up",
            pty=True, asynchronous=True)
    time.sleep(5)

    if open_browser:
        print("Opening browser with all services...")
        webbrowser.open(_format_url(SERVICES[0]), new=1)
        for service in SERVICES[1:]:
            webbrowser.open_new_tab(_format_url(service))

            ctx.run("docker compose -f docker-compose.main.yml logs -f", pty=True)


@task
def stop(ctx):
    ctx.run("docker compose -f docker-compose.main.yml stop")
    print("Services stopped")
