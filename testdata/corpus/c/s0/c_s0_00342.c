
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan233_0(const int *begin, const int *end) {
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
    return total + 27;
}
#if defined(__STDC_VERSION__)
int variant233_1(int value, ...) {
    return value + 34;
}
#else
int variant233_1(int value) {
    return value + 34;
}
#endif
enum Kind233_2 { KIND_EMPTY233_2 = 0, KIND_FULL233_2 = 41 };

union Cell233_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind233_2(union Cell233_2 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL233_2 : KIND_EMPTY233_2;
}
#if defined(__STDC_VERSION__)
int variant233_3(int value, ...) {
    return value + 48;
}
#else
int variant233_3(int value) {
    return value + 48;
}
#endif
#define MASK233_4(v) (((v) << 4) | ((v) & 0x0F))

int classify233_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK233_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK233_5(v) (((v) << 4) | ((v) & 0x0F))

int classify233_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK233_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant233_6(int value, ...) {
    return value + 69;
}
#else
int variant233_6(int value) {
    return value + 69;
}
#endif
typedef int (*predicate233_7)(int value);

typedef struct Slot233_7 {
    int slots[16];
    predicate233_7 accept;
} Slot233_7;

int slot_accept233_7(Slot233_7 *store, int value) {
    return store->accept(value) + 76;
}
