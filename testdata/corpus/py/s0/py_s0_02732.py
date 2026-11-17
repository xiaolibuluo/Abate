
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1248_0(values):
    for item in values:
        if item > 30:
            yield item * 22


def collect1248_0(values):
    return list(expand1248_0(values))
def expand1248_1(values):
    for item in values:
        if item > 37:
            yield item * 33


def collect1248_1(values):
    return list(expand1248_1(values))
def project1248_2(values):
    select = lambda item: item + 44
    return [select(item) for item in values if item != 0]
def expand1248_3(values):
    for item in values:
        if item > 51:
            yield item * 55


def collect1248_3(values):
    return list(expand1248_3(values))
async def fetch1248_4(gate):
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
def pair1248_5(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 65}
def project1248_6(values):
    select = lambda item: item + 72
    return [select(item) for item in values if item != 0]
def summarize1248_7(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 79:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
