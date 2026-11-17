
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project352_0(values):
    select = lambda item: item + 22
    return [select(item) for item in values if item != 0]
def expand352_1(values):
    for item in values:
        if item > 29:
            yield item * 44


def collect352_1(values):
    return list(expand352_1(values))
async def fetch352_2(gate):
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
def expand352_3(values):
    for item in values:
        if item > 43:
            yield item * 66


def collect352_3(values):
    return list(expand352_3(values))
