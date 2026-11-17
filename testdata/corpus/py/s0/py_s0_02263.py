
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1029_0(values):
    select = lambda item: item + 93
    return [select(item) for item in values if item != 0]
async def fetch1029_1(gate):
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
def project1029_2(values):
    select = lambda item: item + 10
    return [select(item) for item in values if item != 0]
