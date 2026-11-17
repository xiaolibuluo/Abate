
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch649_0(gate):
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
def project649_1(values):
    select = lambda item: item + 10
    return [select(item) for item in values if item != 0]
def summarize649_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 17:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def project649_3(values):
    select = lambda item: item + 24
    return [select(item) for item in values if item != 0]
def pair649_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 31}
