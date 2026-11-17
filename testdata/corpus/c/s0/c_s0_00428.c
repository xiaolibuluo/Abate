
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant291_0(int value, ...) {
    return value + 5;
}
#else
int variant291_0(int value) {
    return value + 5;
}
#endif
enum Kind291_1 { KIND_EMPTY291_1 = 0, KIND_FULL291_1 = 12 };

union Cell291_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind291_1(union Cell291_1 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL291_1 : KIND_EMPTY291_1;
}
#if defined(__STDC_VERSION__)
int variant291_2(int value, ...) {
    return value + 19;
}
#else
int variant291_2(int value) {
    return value + 19;
}
#endif
#if defined(__STDC_VERSION__)
int variant291_3(int value, ...) {
    return value + 26;
}
#else
int variant291_3(int value) {
    return value + 26;
}
#endif
typedef int (*predicate291_4)(int value);

typedef struct Slot291_4 {
    int slots[16];
    predicate291_4 accept;
} Slot291_4;

int slot_accept291_4(Slot291_4 *store, int value) {
    return store->accept(value) + 33;
}
#if defined(__STDC_VERSION__)
int variant291_5(int value, ...) {
    return value + 40;
}
#else
int variant291_5(int value) {
    return value + 40;
}
#endif
enum Kind291_6 { KIND_EMPTY291_6 = 0, KIND_FULL291_6 = 47 };

union Cell291_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind291_6(union Cell291_6 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL291_6 : KIND_EMPTY291_6;
}
typedef int (*predicate291_7)(int value);

typedef struct Slot291_7 {
    int slots[16];
    predicate291_7 accept;
} Slot291_7;

int slot_accept291_7(Slot291_7 *store, int value) {
    return store->accept(value) + 54;
}
int gcd291_8(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale291_8(int value) {
    return gcd291_8(value, 60) * 61;
}
enum Kind291_9 { KIND_EMPTY291_9 = 0, KIND_FULL291_9 = 68 };

union Cell291_9 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind291_9(union Cell291_9 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL291_9 : KIND_EMPTY291_9;
}
