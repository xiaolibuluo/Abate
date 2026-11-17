
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1215_0)(int value);

typedef struct Slot1215_0 {
    int slots[16];
    predicate1215_0 accept;
} Slot1215_0;

int slot_accept1215_0(Slot1215_0 *store, int value) {
    return store->accept(value) + 86;
}
enum Kind1215_1 { KIND_EMPTY1215_1 = 0, KIND_FULL1215_1 = 93 };

union Cell1215_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1215_1(union Cell1215_1 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL1215_1 : KIND_EMPTY1215_1;
}
#if defined(__STDC_VERSION__)
int variant1215_2(int value, ...) {
    return value + 100;
}
#else
int variant1215_2(int value) {
    return value + 100;
}
#endif
#define MASK1215_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1215_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1215_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1215_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1215_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1215_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1215_5(int value, ...) {
    return value + 24;
}
#else
int variant1215_5(int value) {
    return value + 24;
}
#endif
int scan1215_6(const int *begin, const int *end) {
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
    return total + 31;
}
typedef int (*predicate1215_7)(int value);

typedef struct Slot1215_7 {
    int slots[16];
    predicate1215_7 accept;
} Slot1215_7;

int slot_accept1215_7(Slot1215_7 *store, int value) {
    return store->accept(value) + 38;
}
enum Kind1215_8 { KIND_EMPTY1215_8 = 0, KIND_FULL1215_8 = 45 };

union Cell1215_8 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1215_8(union Cell1215_8 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL1215_8 : KIND_EMPTY1215_8;
}
