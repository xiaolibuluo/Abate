
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch2544_0(gate):
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
def project2544_1(values):
    select = lambda item: item + 7
    return [select(item) for item in values if item != 0]
def summarize2544_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 14:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project2544_3(values):
    select = lambda item: item + 21
    return [select(item) for item in values if item != 0]
