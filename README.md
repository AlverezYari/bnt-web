# bnt-web


## Project Components

- **bnt-web**: The main web application.
- **bnt-infa**: The infrastructure for the web application.


### bnt-web

The main web application. Built with GO v1.22.1, HTMX, & TailwindCSS.

### bnt-infa

The infrastructure for the web application. Built with OpenTofu, k3s, and hosted on Hetzner Cloud.

**Setup**

```bash
alias createkh='tmp_script=$(mktemp) && curl -sSL -o "${tmp_script}" https://raw.githubusercontent.com/kube-hetzner/terraform-hcloud-kube-hetzner/master/scripts/create.sh && chmod +x "${tmp_script}" && "${tmp_script}" && rm "${tmp_script}"'
```

