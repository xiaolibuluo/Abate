
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2485_0(values):
    for item in values:
        if item > 9:
            yield item * 19


def collect2485_0(values):
    return list(expand2485_0(values))
async def fetch2485_1(gate):
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
