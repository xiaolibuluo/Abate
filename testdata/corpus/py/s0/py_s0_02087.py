
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand953_0(values):
    for item in values:
        if item > 75:
            yield item * 31


def collect953_0(values):
    return list(expand953_0(values))
async def fetch953_1(gate):
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
def expand953_2(values):
    for item in values:
        if item > 89:
            yield item * 53


def collect953_2(values):
    return list(expand953_2(values))
def expand953_3(values):
    for item in values:
        if item > 96:
            yield item * 64


def collect953_3(values):
    return list(expand953_3(values))
def expand953_4(values):
    for item in values:
        if item > 6:
            yield item * 14


def collect953_4(values):
    return list(expand953_4(values))
def pair953_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 13}
def expand953_6(values):
    for item in values:
        if item > 20:
            yield item * 36


def collect953_6(values):
    return list(expand953_6(values))
