
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand1278_0(values):
    for item in values:
        if item > 32:
            yield item * 49


def collect1278_0(values):
    return list(expand1278_0(values))
async def fetch1278_1(gate):
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
def summarize1278_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 46:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
