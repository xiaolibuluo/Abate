
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2266_0(values):
    for item in values:
        if item > 72:
            yield item * 11


def collect2266_0(values):
    return list(expand2266_0(values))
def summarize2266_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 79:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize2266_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 86:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2266_3(gate):
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
def pair2266_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 100}
def pair2266_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 10}
