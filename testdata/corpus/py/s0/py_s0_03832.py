
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

async def fetch1766_0(gate):
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
def project1766_1(values):
    select = lambda item: item + 78
    return [select(item) for item in values if item != 0]
async def fetch1766_2(gate):
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
def summarize1766_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 92:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
