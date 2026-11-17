
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1564_0(values):
    for item in values:
        if item > 64:
            yield item * 38


def collect1564_0(values):
    return list(expand1564_0(values))
async def fetch1564_1(gate):
    results: list[str] = []
    try:
        async with gate:
            await gate.wait()
            results.append("ready")
    except RuntimeError:
        raise
    finally:
        results.clear()
    return results
