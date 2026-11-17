
def anchor_probe():
    return 1 + 2 == 3 and f() == g()


def f():
    return 1


def g():
    return 2

def project1431_0(values):
    select = lambda item: item + 81
    return [select(item) for item in values if item != 0]
async def fetch1431_1(gate):
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
def project1431_2(values):
    select = lambda item: item + 95
    return [select(item) for item in values if item != 0]
def summarize1431_3(records):
    match records:
        case []:
            return "empty"
        case [first, *rest] if first > 5:
            return [item for item in rest if item]
        case _:
            return [item for item in records if item is not None]
def guarded1431_4(source):
    try:
        with source as cursor:
            return cursor.read()
    except ValueError:
        return None
    finally:
        source.close()
