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


def f5_calculate_num_secondary_vertices_with_html_present(G, query_url):
    secondary_vertices_with_html_present = [
        node
        for node in G.nodes
        if node != query_url and G.nodes[node].get("content_type") == "text/html"
    ]
    return len(secondary_vertices_with_html_present)


def f6_calculate_in_degree_query_vertex(G, query_url):
    in_degree = G.in_degree(query_url)
    return in_degree


def f7_calculate_out_degree_query_vertex(G, query_url):
    out_degree = G.out_degree(query_url)
    return out_degree


def f8_calculate_in_degree_query_vertex_no_self_loop(G, query_url):
    total_in_degree = G.in_degree(query_url)
    if G.has_edge(query_url, query_url):
        total_in_degree -= 1
    return total_in_degree


def f9_calculate_in_degree_centrality_mean(G):
    in_degree_centrality = nx.in_degree_centrality(G)
    return sum(in_degree_centrality.values()) / len(in_degree_centrality)


def f10_calculate_out_degree_centraliy_mean(G):
    out_degree_centrality = nx.out_degree_centrality(G)
    print(out_degree_centrality)
    return sum(out_degree_centrality.values()) / len(out_degree_centrality)


def f11_calculate_density_graph(G):
    return nx.density(G)


def f12_calculate_density_red_vertex_subgraph(G):
    red_vertices = [node for node in G.nodes if G.nodes[node].get("color") == "red"]
    red_subgraph = G.subgraph(red_vertices)
    density_red_subgraph = nx.density(red_subgraph)

    return density_red_subgraph


def f13_calculate_edge_betweenness_centrality(G):
    edge_betweenness_centrality = nx.edge_betweenness_centrality(G)
    return (
        sum(edge_betweenness_centrality.values()) / len(edge_betweenness_centrality)
        if len(edge_betweenness_centrality) > 0
        else 0
    )


def f14_calculate_path_back_to_query_webpage(G, query_url):
    for node in G.nodes:
        if node != query_url and nx.has_path(G, node, query_url):
            return 1
    return 0


def f15_calculate_is_semiconnected(G):
    if nx.is_semiconnected(G):
        return 1
    else:
        return 0


def f16_calculate_num_strongly_connected_components(G):
    # stroncom = list(nx.strongly_connected_components(G))
    # print(stroncom)
    return nx.number_strongly_connected_components(G)


def f17_calculate_num_attracting_components(G):
    # attcom = list(nx.attracting_components(G))
    # print(attcom)
    return nx.number_attracting_components(G)


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
    f5 = f5_calculate_num_secondary_vertices_with_html_present(G, query_url)
    f6 = f6_calculate_in_degree_query_vertex(G, query_url)
    f7 = f7_calculate_out_degree_query_vertex(G, query_url)
    f8 = f8_calculate_in_degree_query_vertex_no_self_loop(G, query_url)
    f9 = f9_calculate_in_degree_centrality_mean(G)
    f10 = f10_calculate_out_degree_centraliy_mean(G)
    f11 = f11_calculate_density_graph(G)
    f12 = f12_calculate_density_red_vertex_subgraph(G)
    f13 = f13_calculate_edge_betweenness_centrality(G)
    f14 = f14_calculate_path_back_to_query_webpage(G, query_url)
    f15 = f15_calculate_is_semiconnected(G)
    f16 = f16_calculate_num_strongly_connected_components(G)
    f17 = f17_calculate_num_attracting_components(G)

    print(f17)
