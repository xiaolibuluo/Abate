
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1002_0(values):
    select = lambda item: item + 33
    return [select(item) for item in values if item != 0]
async def fetch1002_1(gate):
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
async def fetch1002_2(gate):
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
def summarize1002_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 54:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def expand1002_4(values):
    for item in values:
        if item > 61:
            yield item * 52


def collect1002_4(values):
    return list(expand1002_4(values))
def summarize1002_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 68:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
