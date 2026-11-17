
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2447_0(values):
    for item in values:
        if item > 97:
            yield item * 58


def collect2447_0(values):
    return list(expand2447_0(values))
async def fetch2447_1(gate):
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
async def fetch2447_2(gate):
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
def project2447_3(values):
    select = lambda item: item + 21
    return [select(item) for item in values if item != 0]
