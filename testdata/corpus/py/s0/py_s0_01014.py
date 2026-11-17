
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch466_0(gate):
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
def expand466_1(values):
    for item in values:
        if item > 56:
            yield item * 49


def collect466_1(values):
    return list(expand466_1(values))
def summarize466_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 63:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch466_3(gate):
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
def project466_4(values):
    select = lambda item: item + 77
    return [select(item) for item in values if item != 0]
def summarize466_5(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 84:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
