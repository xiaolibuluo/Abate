
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2393_0(values):
    for item in values:
        if item > 74:
            yield item * 46


def collect2393_0(values):
    return list(expand2393_0(values))
def summarize2393_1(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 81:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2393_2(gate):
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
