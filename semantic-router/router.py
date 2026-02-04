from redisvl.extensions.router import Route, SemanticRouter

routes = [
    Route(
        name="GenAI Programming",
        references=[
            "chatbot with LLM",
            "prompt engineering",
            "machine learning",
            "neural network",
            "GPT transformer",
            "pytorch tensorflow",
            "vector embeddings RAG",
        ],
        metadata={"category": "tech"},
        distance_threshold=0.5,
    ),
    Route(
        name="Science Fiction Entertainment",
        references=[
            "star wars",
            "star trek",
            "dune",
            "blade runner",
            "the matrix",
            "interstellar",
            "alien predator",
        ],
        metadata={"category": "entertainment"},
        distance_threshold=0.5,
    ),
    Route(
        name="Classical Music",
        references=[
            "beethoven symphony",
            "mozart concerto",
            "bach",
            "orchestra philharmonic",
            "violin sonata",
            "opera composer",
            "chopin nocturne",
        ],
        metadata={"category": "music"},
        distance_threshold=0.5,
    ),
]

router = SemanticRouter(
    name="topic-router",
    routes=routes,
    redis_url="redis://<host:port>",
)

queries = [
    "how to fine tune language model",
    "best star wars movie",
    "beethoven moonlight sonata",
]

for q in queries:
    result = router(q)
    print(f"{q} -> {result.name}")
