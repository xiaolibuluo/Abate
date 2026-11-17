
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2064_0(values):
    for item in values:
        if item > 65:
            yield item * 61


def collect2064_0(values):
    return list(expand2064_0(values))
async def fetch2064_1(gate):
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
def project2064_2(values):
    select = lambda item: item + 79
    return [select(item) for item in values if item != 0]
def pair2064_3(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 86}
def expand2064_4(values):
    for item in values:
        if item > 93:
            yield item * 44


def collect2064_4(values):
    return list(expand2064_4(values))
def summarize2064_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 100:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
