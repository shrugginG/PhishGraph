from bs4 import BeautifulSoup
import os


def get_hyperlinks_from_html_file(file_path):
    with open(file_path, "r", encoding="utf-8") as f:
        soup = BeautifulSoup(f, "html.parser")
        hyperlinks = []
        for a_tag in soup.find_all("a"):
            hyperlinks.append(a_tag.get("href"))
        return hyperlinks


def list_url_dirs(path):
    return os.listdir(path)


url_dirs = list_url_dirs("/home/shrugging/project/PhishDetect/phishgraph/data/")
print(os.listdir("/home/shrugging/project/PhishDetect/phishgraph/data/"))

for url_sha256 in url_dirs:
    hyperlinks = get_hyperlinks_from_html_file(
        f"/home/shrugging/project/PhishDetect/phishgraph/data/{url_sha256}/{url_sha256}.html"
    )
    print(hyperlinks)

# hyperlinks = get_hyperlinks_from_html_file(
#     "/home/shrugging/project/PhishDetect/phishgraph/data/bf5a4bb3b5751a746457d9520e1c8d3506fcf0dc9337cac12860483ea655c9c1/bf5a4bb3b5751a746457d9520e1c8d3506fcf0dc9337cac12860483ea655c9c1.html"
# )
# print(hyperlinks)
#
