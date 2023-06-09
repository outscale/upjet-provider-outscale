from git import Repo
import json
import requests
from github import Github
import os
import subprocess
import argparse
from kubernetes import client, config
import time

def create_pull_request(owner_name, repo_name, title, description, head_branch, base_branch, git_token):
    git_pulls_api = "https://api.github.com/repos/{0}/{1}/pulls".format(
        owner_name,
        repo_name
    )

    headers = {
        "Authorization": "token {0}".format(git_token),
        "Content-Type": "application/json"
    }

    payload = {
        "title": title,
        "body": description,
        "head": head_branch,
        "base": base_branch,
    }

    r = requests.post(
        git_pulls_api,
        headers=headers,
        data=json.dumps(payload))
    if not r.ok:
        print("Request Failed: {0}".format(r.text))

# to change
def get_branch(remote, full_local_path, branch):
    Repo.clone_from(remote,full_local_path, branch="first-step")
    repo = Repo(full_local_path)
    git = repo.git
    git.checkout
    git.checkout("HEAD", b=branch)

def add_and_commit(full_local_path, commit_msg):
    repo = Repo(full_local_path)
    repo.index.commit(commit_msg)

def push(full_local_path, branch):
    repo = Repo(full_local_path)
    repo.remote()
    origin = repo.remote(name="origin").push(branch)

def apply(full_local_path, folder):
    repo = Repo(full_local_path)
    repo.git.execute(['git', 'apply', folder])

def get_release(repo, token):
    G = Github(token)
    repo = G.get_repo(repo)
    releases = repo.get_releases()
    return releases[2].title, releases[2].body

def create_issue(repo, token, title, body):
    G = Github(token)
    repo = G.get_repo(repo)
    repo.create_issue(title, body)

def set_config(name, email, full_local_path):
    repo = Repo(full_local_path)
    repo.config_writer().set_value("user", "name", name).release()
    repo.config_writer().set_value("user", "email", email).release()

def execute_bash_cmd(cmd, full_local_path, my_env):
    result = subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True, cwd=full_local_path, env=my_env )
    print(result.stdout)
    print(result.stderr)

def read_file(filename):
    with open(filename, 'r') as f:
        return f.read()

def main():

    name = os.getenv('git_username', "outscale")
    email = os.getenv('git_email',"vincent.baer@outscale.com")
    username = os.getenv('git_username',"outscale")
    full_local_path = os.getenv('git_full_local_path', "/home/outscale/crossplane")
    password = os.getenv('git_password',"token")
    branch = os.getenv('git_branch',"release-terraform")
    commit_msg = os.getenv('git_commit',"Update with a new version of terraform")
    owner_name = os.getenv('git_owner', "outscale")
    repo_name = os.getenv('git_repo', "upjet-provider-outscale")
    remote = f"https://{username}:{password}@github.com/{owner_name}/{repo_name}.git"
    terraform_version_file = os.getenv("terraform_version_file", "./terraform_version")
    
    title = os.getenv('git_pr_title',"Update uppjet with the integration of a new release of terraform")
    description = os.getenv('git_pr_description', "Update upjet with a new release of terraform")
    base_branch = os.getenv('git_base_branch',"main")
    watch_target_projet = os.getenv('git_watch_target_projet',"outscale/terraform-provider-outscale")
    parser = argparse.ArgumentParser()
    parser.add_argument("-a","--apply", help="ApplyUpgrade", action="store_true")
    parser.add_argument("-c", "--clone", help="Clone", action="store_true")
    parser.add_argument("-bp","--buildpush", help="BuildPush", action="store_true")
    parser.add_argument("-et", "--e2etest", help="E2eTest", action="store_true")
    parser.add_argument("-p", "--pullrequest", help="PullRequest", action="store_true")
    parser.add_argument("-g", "--get", help="GetTerraformVersion", action="store_true")
    parser.add_argument("-u", "--update", help="Update", action="store_true" )
    args = parser.parse_args()
    apply = args.apply
    update = args.update 
    buildpush = args.buildpush
    clone = args.clone
    get_terraform_version = args.get

    my_env = os.environ.copy()
    terraform_release_version = os.getenv("terraform_release_version", "v0.8.2")
    terraform_version = terraform_release_version.replace("v","")
    current_branch = "{0}-{1}".format(branch, terraform_version)
    current_commit_msg = "{0}-{1}".format(commit_msg, terraform_version)
    pr_title = "{0} {1}".format(title, terraform_version)
    pr_description = "{0} {1}".format(description, terraform_version)
    release_title, release_body = get_release(watch_target_projet,password)
    find_terraform_version = terraform_release_version.replace("v","")

    if clone:
        get_branch(remote, full_local_path, current_branch)
        set_config(name, email, full_local_path)

    elif apply:
        add_and_commit(full_local_path, commit_msg)
        push(full_local_path, current_branch)
        create_pull_request(
            owner_name,
            repo_name,
            pr_title,
            pr_description,
            current_branch,
            base_branch,
            password
        )
    elif update:
        my_env['TERRAFORM_PROVIDER_VERSION'] = find_terraform_version
        my_env['TERRAFORM_NATIVE_PROVIDER_BINARY'] = "terraform-provider-outscale_v{0}".format(find_terraform_version)
        execute_bash_cmd(["mv", "config", "config-save"], full_local_path, my_env)
        execute_bash_cmd(["rm", "-rf", "config"], full_local_path, my_env)
        execute_bash_cmd(["mkdir", "config"], full_local_path, my_env)
        execute_bash_cmd(["mv", "config-save/external_name.go", "config/external_name.go"], full_local_path, my_env)
        execute_bash_cmd(["mv", "config-save/provider.go", "config/provider.go"], full_local_path, my_env)
        execute_bash_cmd(["mv", "config-save/provider-metadata.yaml", "config/provider-metadata.yaml"], full_local_path, my_env)
        execute_bash_cmd(["mv", "config-save/schema.json", "config/schema.json"], full_local_path, my_env)
        execute_bash_cmd(["rm", "-rf", "config-save"], full_local_path, my_env)
        execute_bash_cmd(["git", "apply", "--reject", "--whitespace=fix", "patch/0001-add-commit-for-config.patch"], full_local_path, my_env )
        execute_bash_cmd(["make", "submodules"], full_local_path, my_env )
        execute_bash_cmd(["make", "build"], full_local_path, my_env)
        execute_bash_cmd(["make", "docker-buildx"], full_local_path, my_env)
    elif buildpush:
        my_env['TERRAFORM_PROVIDER_VERSION'] = find_terraform_version
        my_env['TERRAFORM_NATIVE_PROVIDER_BINARY'] = "terraform-provider-outscale_v{0}".format(find_terraform_version)
        execute_bash_cmd(["make", "submodules"], full_local_path, my_env )
        execute_bash_cmd(["make", "build"], full_local_path, my_env)
        execute_bash_cmd(["make", "docker-buildx"], full_local_path, my_env)
        execute_bash_cmd(["make", "docker-push"], full_local_path, my_env)
    elif get_terraform_version:
        release_title, release_body = get_release(watch_target_projet,password)
        print(release_title)
   
if __name__ == "__main__":
    main()
