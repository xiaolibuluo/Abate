
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch368_0(gate):
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
def expand368_1(values):
    for item in values:
        if item > 43:
            yield item * 34


def collect368_1(values):
    return list(expand368_1(values))
def summarize368_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 50:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch368_3(gate):
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
def project368_4(values):
    select = lambda item: item + 64
    return [select(item) for item in values if item != 0]
async def fetch368_5(gate):
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
