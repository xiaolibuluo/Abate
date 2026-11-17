
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand693_0(values):
    for item in values:
        if item > 90:
            yield item * 41


def collect693_0(values):
    return list(expand693_0(values))
def project693_1(values):
    select = lambda item: item + 97
    return [select(item) for item in values if item != 0]
def project693_2(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
async def fetch693_3(gate):
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
