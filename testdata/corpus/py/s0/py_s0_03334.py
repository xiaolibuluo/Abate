
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1540_0(values):
    select = lambda item: item + 43
    return [select(item) for item in values if item != 0]
async def fetch1540_1(gate):
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
async def fetch1540_2(gate):
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
def expand1540_3(values):
    for item in values:
        if item > 64:
            yield item * 25


def collect1540_3(values):
    return list(expand1540_3(values))
def expand1540_4(values):
    for item in values:
        if item > 71:
            yield item * 36


def collect1540_4(values):
    return list(expand1540_4(values))
def summarize1540_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 78:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def pair1540_6(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 85}
