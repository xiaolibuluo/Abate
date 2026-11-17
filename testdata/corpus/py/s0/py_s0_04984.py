
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2312_0(gate):
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
def summarize2312_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 95:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2312_2(gate):
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
def project2312_3(values):
    select = lambda item: item + 12
    return [select(item) for item in values if item != 0]
def summarize2312_4(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 19:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
