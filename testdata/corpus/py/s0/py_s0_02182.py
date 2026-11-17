
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1001_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 20}
def pair1001_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 27}
async def fetch1001_2(gate):
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
def summarize1001_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 41:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1001_4(values):
    for item in values:
        if item > 48:
            yield item * 45


def collect1001_4(values):
    return list(expand1001_4(values))
def summarize1001_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 55:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
