from bs4 import BeautifulSoup
import os


def get_hyperlinks_from_html_file(file_path):
    with open(file_path, "r", encoding="utf-8") as f:
        soup = BeautifulSoup(f, "html.parser")
        hyperlinks = []
        for a_tag in soup.find_all("a"):
            hyperlinks.append(a_tag.get("href"))
        return hyperlinks


hyperlinks = get_hyperlinks_from_html_file(
    "/home/jxlu/project/PhishDetect/PhishGraph/data/bf5a4bb3b5751a746457d9520e1c8d3506fcf0dc9337cac12860483ea655c9c1/bf5a4bb3b5751a746457d9520e1c8d3506fcf0dc9337cac12860483ea655c9c1.html"
)
print(hyperlinks)
