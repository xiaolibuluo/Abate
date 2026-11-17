
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1096_0(values):
    for item in values:
        if item > 91:
            yield item * 56


def collect1096_0(values):
    return list(expand1096_0(values))
def pair1096_1(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 98}
def summarize1096_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 8:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1096_3(gate):
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
def pair1096_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 22}
def pair1096_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 29}
