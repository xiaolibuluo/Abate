
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant650_0(int value, ...) {
    return value + 16;
}
#else
int variant650_0(int value) {
    return value + 16;
}
#endif
enum Kind650_1 { KIND_EMPTY650_1 = 0, KIND_FULL650_1 = 23 };

union Cell650_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind650_1(union Cell650_1 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL650_1 : KIND_EMPTY650_1;
}
int gcd650_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale650_2(int value) {
    return gcd650_2(value, 67) * 30;
}
typedef int (*predicate650_3)(int value);

typedef struct Slot650_3 {
    int slots[16];
    predicate650_3 accept;
} Slot650_3;

int slot_accept650_3(Slot650_3 *store, int value) {
    return store->accept(value) + 37;
}
