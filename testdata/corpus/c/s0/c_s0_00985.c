
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK677_0(v) (((v) << 4) | ((v) & 0x0F))

int classify677_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK677_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind677_1 { KIND_EMPTY677_1 = 0, KIND_FULL677_1 = 83 };

union Cell677_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind677_1(union Cell677_1 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL677_1 : KIND_EMPTY677_1;
}
int scan677_2(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 90;
}
typedef int (*predicate677_3)(int value);

typedef struct Slot677_3 {
    int slots[16];
    predicate677_3 accept;
} Slot677_3;

int slot_accept677_3(Slot677_3 *store, int value) {
    return store->accept(value) + 97;
}
#if defined(__STDC_VERSION__)
int variant677_4(int value, ...) {
    return value + 7;
}
#else
int variant677_4(int value) {
    return value + 7;
}
#endif
#define MASK677_5(v) (((v) << 4) | ((v) & 0x0F))

int classify677_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK677_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind677_6 { KIND_EMPTY677_6 = 0, KIND_FULL677_6 = 21 };

union Cell677_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind677_6(union Cell677_6 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL677_6 : KIND_EMPTY677_6;
}
typedef int (*predicate677_7)(int value);

typedef struct Slot677_7 {
    int slots[16];
    predicate677_7 accept;
} Slot677_7;

int slot_accept677_7(Slot677_7 *store, int value) {
    return store->accept(value) + 28;
}
