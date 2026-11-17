
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand418_0(values):
    for item in values:
        if item > 7:
            yield item * 68


def collect418_0(values):
    return list(expand418_0(values))
async def fetch418_1(gate):
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
def expand418_2(values):
    for item in values:
        if item > 21:
            yield item * 29


def collect418_2(values):
    return list(expand418_2(values))
def project418_3(values):
    select = lambda item: item + 28
    return [select(item) for item in values if item != 0]
def expand418_4(values):
    for item in values:
        if item > 35:
            yield item * 51


def collect418_4(values):
    return list(expand418_4(values))
def project418_5(values):
    select = lambda item: item + 42
    return [select(item) for item in values if item != 0]
async def fetch418_6(gate):
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
