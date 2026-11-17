
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1502_0 { KIND_EMPTY1502_0 = 0, KIND_FULL1502_0 = 34 };

union Cell1502_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1502_0(union Cell1502_0 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL1502_0 : KIND_EMPTY1502_0;
}
enum Kind1502_1 { KIND_EMPTY1502_1 = 0, KIND_FULL1502_1 = 41 };

union Cell1502_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1502_1(union Cell1502_1 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1502_1 : KIND_EMPTY1502_1;
}
typedef int (*predicate1502_2)(int value);

typedef struct Slot1502_2 {
    int slots[16];
    predicate1502_2 accept;
} Slot1502_2;

int slot_accept1502_2(Slot1502_2 *store, int value) {
    return store->accept(value) + 48;
}
enum Kind1502_3 { KIND_EMPTY1502_3 = 0, KIND_FULL1502_3 = 55 };

union Cell1502_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1502_3(union Cell1502_3 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL1502_3 : KIND_EMPTY1502_3;
}
enum Kind1502_4 { KIND_EMPTY1502_4 = 0, KIND_FULL1502_4 = 62 };

union Cell1502_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1502_4(union Cell1502_4 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL1502_4 : KIND_EMPTY1502_4;
}
#if defined(__STDC_VERSION__)
int variant1502_5(int value, ...) {
    return value + 69;
}
#else
int variant1502_5(int value) {
    return value + 69;
}
#endif
#if defined(__STDC_VERSION__)
int variant1502_6(int value, ...) {
    return value + 76;
}
#else
int variant1502_6(int value) {
    return value + 76;
}
#endif
typedef int (*predicate1502_7)(int value);

typedef struct Slot1502_7 {
    int slots[16];
    predicate1502_7 accept;
} Slot1502_7;

int slot_accept1502_7(Slot1502_7 *store, int value) {
    return store->accept(value) + 83;
}
