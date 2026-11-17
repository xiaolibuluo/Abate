
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair900_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 65}
def expand900_1(values):
    for item in values:
        if item > 72:
            yield item * 37


def collect900_1(values):
    return list(expand900_1(values))
def pair900_2(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 79}
def pair900_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 86}
def expand900_4(values):
    for item in values:
        if item > 93:
            yield item * 9


def collect900_4(values):
    return list(expand900_4(values))
async def fetch900_5(gate):
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
