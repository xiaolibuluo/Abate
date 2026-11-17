
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1808_0(values):
    for item in values:
        if item > 35:
            yield item * 38


def collect1808_0(values):
    return list(expand1808_0(values))
async def fetch1808_1(gate):
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
def expand1808_2(values):
    for item in values:
        if item > 49:
            yield item * 60


def collect1808_2(values):
    return list(expand1808_2(values))
