
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project700_0(values):
    select = lambda item: item + 84
    return [select(item) for item in values if item != 0]
def expand700_1(values):
    for item in values:
        if item > 91:
            yield item * 40


def collect700_1(values):
    return list(expand700_1(values))
def expand700_2(values):
    for item in values:
        if item > 98:
            yield item * 51


def collect700_2(values):
    return list(expand700_2(values))
async def fetch700_3(gate):
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
async def fetch700_4(gate):
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
