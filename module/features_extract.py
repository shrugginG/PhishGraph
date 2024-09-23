import networkx as nx


def f1_calculate_pct_red_vertices(G):
    read_vertices = [node for node in G.nodes if G.nodes[node]["color"] == "red"]
    total_nodes = G.number_of_nodes()
    pct_red = len(read_vertices) / total_nodes if total_nodes > 0 else 0
    return pct_red


def f2_calculate_pct_blue_vertices(G):
    blue_vertices = [node for node in G.nodes if G.nodes[node]["color"] == "blue"]
    total_nodes = G.number_of_nodes()
    pct_blue = len(blue_vertices) / total_nodes if total_nodes > 0 else 0
    return pct_blue


def f3_calculate_pct_green_vertices(G):
    green_vertices = [node for node in G.nodes if G.nodes[node]["color"] == "green"]
    total_nodes = G.number_of_nodes()
    pct_green = len(green_vertices) / total_nodes if total_nodes > 0 else 0
    return pct_green


def f4_calculate_pct_secondary_red_vertices(G, query_url):
    secondary_red_vertices = [
        node
        for node in G.nodes
        if G.nodes[node]["color"] == "red" and node != query_url
    ]

    total_nodes = G.number_of_nodes()
    pct_secondary_red = (
        len(secondary_red_vertices) / total_nodes if total_nodes > 0 else 0
    )

    return pct_secondary_red


if __name__ == "__main__":
    query_url = "https://www.example.com/"
    G = nx.read_graphml(
        "/home/jxlu/project/PhishDetect/PhishGraph/data/test/ego_network.graphml"
    )
    print(G)
    f1 = f1_calculate_pct_red_vertices(G)
    f2 = f2_calculate_pct_blue_vertices(G)
    f3 = f3_calculate_pct_green_vertices(G)
    f4 = f4_calculate_pct_secondary_red_vertices(G, query_url)
    print(f4)
