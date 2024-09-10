import networkx as nx
import matplotlib.pyplot as plt
import tldextract

query_webpage_url = "https://www.google.com"
level_1_neighbours = set(
    [
        "https://www.google.com",
        "https://www.youtube.com",
        "https://www.gmail.com",
        "https://www.google.com/maps",
    ]
)
level_2_neighbours = {
    "https://www.youtube.com": set(
        [
            "https://www.youtube.com/feed",
            "https://www.gmail.com",
            "https://www.google.com",
        ]
    ),
    "https://www.gmail.com": set(["https://www.youtube.com/watch?v=123456"]),
    "https://www.google.com/maps": set(["https://www.gmail.com/inbox"]),
}


query_webpage_domain = tldextract.extract(query_webpage_url).registered_domain

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
        if tldextract.extract(url).registered_domain == query_webpage_domain:
            G.add_node(url, color="red")
        else:
            G.add_node(url, color="blue")
        G.add_edge(query_webpage_url, url)

for level_1_url, level_2_urls in level_2_neighbours.items():
    for level_2_url in level_2_urls:
        if not G.has_node(level_2_url):
            if (
                tldextract.extract(level_2_url).registered_domain
                == query_webpage_domain
            ):
                G.add_node(level_2_url, color="red")
            else:
                G.add_node(level_2_url, color="green")
        G.add_edge(level_1_url, level_2_url)


node_colors = [G.nodes[node]["color"] for node in G.nodes]
plt.figure(figsize=(8, 6))
pos = nx.spring_layout(G)  # 使用spring布局
nx.draw(
    G,
    pos,
    with_labels=True,
    node_color=node_colors,
    node_size=500,
    font_size=10,
    font_color="black",
    edge_color="gray",
)
plt.title("Webpage Graph with Red, Blue, and Green Nodes")

plt.savefig("/home/shrugging/project/PhishDetect/phishgraph/output/webpage_graph.png")
plt.show()
