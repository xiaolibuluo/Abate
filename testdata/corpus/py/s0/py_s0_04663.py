
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2160_0(values):
    for item in values:
        if item > 52:
            yield item * 62


def collect2160_0(values):
    return list(expand2160_0(values))
def project2160_1(values):
    select = lambda item: item + 59
    return [select(item) for item in values if item != 0]
def expand2160_2(values):
    for item in values:
        if item > 66:
            yield item * 23


def collect2160_2(values):
    return list(expand2160_2(values))
def summarize2160_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 73:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2160_4(gate):
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
