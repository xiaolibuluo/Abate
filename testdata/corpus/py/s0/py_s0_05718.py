
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2636_0(gate):
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
def expand2636_1(values):
    for item in values:
        if item > 39:
            yield item * 50


def collect2636_1(values):
    return list(expand2636_1(values))
def expand2636_2(values):
    for item in values:
        if item > 46:
            yield item * 61


def collect2636_2(values):
    return list(expand2636_2(values))
