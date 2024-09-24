import networkx as nx
import matplotlib.pyplot as plt

G = nx.DiGraph()

query_url = "https://www.example.com/"

G.add_node(query_url, color="red")
G.add_node("https://www.example.com/login", color="red", content_type="text/html")
G.add_node("https://www.example.com/register", color="red", content_type="text/html")
G.add_node("https://www.external.com/news", color="blue")
G.add_node("https://www.another.com/about", color="blue")
G.add_node("https://www.example.com/help", color="red", content_type="text/html")
# G.add_node("https://sub.example.com/contact", color="red")
G.add_node("https://sub.external.com/", color="green")
G.add_node("https://www.another.com/support", color="green")

G.add_edges_from(
    [
        (query_url, "https://www.example.com/login"),  # query_url 指向内部链接
        (query_url, "https://www.example.com/register"),
        (query_url, "https://www.external.com/news"),  # query_url 指向外部链接
        (query_url, "https://www.another.com/about"),  # query_url 指向外部链接
        (
            "https://www.example.com/login",
            "https://www.example.com/help",
        ),
        (
            "https://www.external.com/news",
            "https://sub.external.com/",
        ),
        ("https://www.another.com/about", "https://www.another.com/support"),
    ]
)

nx.write_graphml(
    G, "/home/jxlu/project/PhishDetect/PhishGraph/data/test/ego_network.graphml"
)
node_colors = [G.nodes[node]["color"] for node in G.nodes]

pos = nx.spring_layout(G)
nx.draw(
    G,
    pos,
    with_labels=True,
    node_color=node_colors,
    node_size=100,
    font_size=8,
    font_color="black",
    font_weight="bold",
    edge_color="gray",
)
plt.savefig("/home/jxlu/project/PhishDetect/PhishGraph/data/test/ego_network.png")
plt.show()
