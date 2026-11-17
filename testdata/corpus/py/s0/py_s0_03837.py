
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1770_0(values):
    for item in values:
        if item > 26:
            yield item * 16


def collect1770_0(values):
    return list(expand1770_0(values))
def expand1770_1(values):
    for item in values:
        if item > 33:
            yield item * 27


def collect1770_1(values):
    return list(expand1770_1(values))
def expand1770_2(values):
    for item in values:
        if item > 40:
            yield item * 38


def collect1770_2(values):
    return list(expand1770_2(values))
async def fetch1770_3(gate):
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
def expand1770_4(values):
    for item in values:
        if item > 54:
            yield item * 60


def collect1770_4(values):
    return list(expand1770_4(values))
