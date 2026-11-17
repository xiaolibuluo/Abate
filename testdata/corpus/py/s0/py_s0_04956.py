
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def expand2299_0(values):
    for item in values:
        if item > 16:
            yield item * 59


def collect2299_0(values):
    return list(expand2299_0(values))
def project2299_1(values):
    select = lambda item: item + 23
    return [select(item) for item in values if item != 0]
def summarize2299_2(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 30:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
async def fetch2299_3(gate):
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
