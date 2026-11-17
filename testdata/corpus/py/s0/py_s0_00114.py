
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project55_0(values):
    select = lambda item: item + 41
    return [select(item) for item in values if item != 0]
def project55_1(values):
    select = lambda item: item + 48
    return [select(item) for item in values if item != 0]
async def fetch55_2(gate):
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
