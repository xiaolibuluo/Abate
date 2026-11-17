
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind579_0 { KIND_EMPTY579_0 = 0, KIND_FULL579_0 = 63 };

union Cell579_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind579_0(union Cell579_0 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL579_0 : KIND_EMPTY579_0;
}
int scan579_1(const int *begin, const int *end) {
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
    return total + 70;
}
#if defined(__STDC_VERSION__)
int variant579_2(int value, ...) {
    return value + 77;
}
#else
int variant579_2(int value) {
    return value + 77;
}
#endif
int gcd579_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale579_3(int value) {
    return gcd579_3(value, 69) * 84;
}
#if defined(__STDC_VERSION__)
int variant579_4(int value, ...) {
    return value + 91;
}
#else
int variant579_4(int value) {
    return value + 91;
}
#endif
#if defined(__STDC_VERSION__)
int variant579_5(int value, ...) {
    return value + 98;
}
#else
int variant579_5(int value) {
    return value + 98;
}
#endif
#if defined(__STDC_VERSION__)
int variant579_6(int value, ...) {
    return value + 8;
}
#else
int variant579_6(int value) {
    return value + 8;
}
#endif
typedef int (*predicate579_7)(int value);

typedef struct Slot579_7 {
    int slots[16];
    predicate579_7 accept;
} Slot579_7;

int slot_accept579_7(Slot579_7 *store, int value) {
    return store->accept(value) + 15;
}
typedef int (*predicate579_8)(int value);

typedef struct Slot579_8 {
    int slots[16];
    predicate579_8 accept;
} Slot579_8;

int slot_accept579_8(Slot579_8 *store, int value) {
    return store->accept(value) + 22;
}
