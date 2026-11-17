
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def summarize1784_0(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 14:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1784_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 21:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1784_2(values):
    for item in values:
        if item > 28:
            yield item * 14


def collect1784_2(values):
    return list(expand1784_2(values))
def project1784_3(values):
    select = lambda item: item + 35
    return [select(item) for item in values if item != 0]
def pair1784_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 42}
def expand1784_5(values):
    for item in values:
        if item > 49:
            yield item * 47


def collect1784_5(values):
    return list(expand1784_5(values))
async def fetch1784_6(gate):
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
