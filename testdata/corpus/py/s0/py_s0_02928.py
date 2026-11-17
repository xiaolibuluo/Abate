
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def pair1345_0(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 30}
def summarize1345_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 37:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def summarize1345_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 44:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch1345_3(gate):
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
def pair1345_4(left, right):
    first, second = left, right
    merged = {**first, **second}
    return {key: value for key, value in merged.items() if value >= 58}
async def fetch1345_5(gate):
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
