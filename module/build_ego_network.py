import networkx as nx
import matplotlib.pyplot as plt
import tldextract
import json
from postgrel import get_postgrel_links

query_webpage_url = "https://www.botanical-journeys-plant-guides.com/"

query_url_sha256 = "448ad8d45e8b8b9559610161b9f08484390b04cf1540c73aff1cb01484c70d8f"
base_path = "/home/jxlu/project/PhishDetect/PhishGraph/data"

with open(
    f"{base_path}/{query_url_sha256}/links.json",
    "r",
) as links_file:
    links = json.load(links_file)
level_1_neighbours = set(links)

records = get_postgrel_links()

level_2_neighbours = {}

for link in level_1_neighbours:
    if link not in records.keys():
        level_2_neighbours[link] = []
    else:
        level_2_neighbours[link] = records[link]


query_webpage_domain = tldextract.extract(query_webpage_url).fqdn

# red_nodes = [query_webpage_url] + level_1_same_domains_url + level_2_same_domains_url
# blue_nodes = [url for url in level_1_neighbours if url not in level_1_same_domains_url]
# green_nodes = [url for url in level_2_neighbours if url not in level_2_same_domains_url]
#
# print(red_nodes)
# print(blue_nodes)
# print(green_nodes)
G = nx.DiGraph()


G.add_node(query_webpage_url, color="red")

for url in level_1_neighbours:
    if url == query_webpage_url:
        G.add_edge(query_webpage_url, url)
    else:
        if tldextract.extract(url).fqdn == query_webpage_domain:
            G.add_node(url, color="red")
        else:
            G.add_node(url, color="blue")
        G.add_edge(query_webpage_url, url)

for level_1_url, level_2_urls in level_2_neighbours.items():
    for level_2_url in level_2_urls:
        if not G.has_node(level_2_url):
            if tldextract.extract(level_2_url).fqdn == query_webpage_domain:
                G.add_node(level_2_url, color="red")
            else:
                G.add_node(level_2_url, color="green")
        G.add_edge(level_1_url, level_2_url)


node_colors = [G.nodes[node]["color"] for node in G.nodes]
plt.figure(figsize=(12, 12))
pos = nx.spring_layout(G, seed=42)

edges = nx.draw_networkx_edges(G, pos, alpha=0.3, edge_color="pink")

nx.draw_networkx_nodes(G, pos, node_color=node_colors, node_size=50, alpha=0.8)

plt.title("Webpage Graph with Red, Blue, and Green Nodes")

plt.savefig(f"{base_path}/{query_url_sha256}/ego_network.png")
plt.show()
