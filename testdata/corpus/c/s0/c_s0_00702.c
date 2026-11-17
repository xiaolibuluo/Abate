
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd481_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale481_0(int value) {
    return gcd481_0(value, 21) * 50;
}
typedef int (*predicate481_1)(int value);

typedef struct Slot481_1 {
    int slots[16];
    predicate481_1 accept;
} Slot481_1;

int slot_accept481_1(Slot481_1 *store, int value) {
    return store->accept(value) + 57;
}
enum Kind481_2 { KIND_EMPTY481_2 = 0, KIND_FULL481_2 = 64 };

union Cell481_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind481_2(union Cell481_2 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL481_2 : KIND_EMPTY481_2;
}
int gcd481_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale481_3(int value) {
    return gcd481_3(value, 54) * 71;
}
enum Kind481_4 { KIND_EMPTY481_4 = 0, KIND_FULL481_4 = 78 };

union Cell481_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind481_4(union Cell481_4 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL481_4 : KIND_EMPTY481_4;
}
#if defined(__STDC_VERSION__)
int variant481_5(int value, ...) {
    return value + 85;
}
#else
int variant481_5(int value) {
    return value + 85;
}
#endif
#if defined(__STDC_VERSION__)
int variant481_6(int value, ...) {
    return value + 92;
}
#else
int variant481_6(int value) {
    return value + 92;
}
#endif
enum Kind481_7 { KIND_EMPTY481_7 = 0, KIND_FULL481_7 = 99 };

union Cell481_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind481_7(union Cell481_7 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL481_7 : KIND_EMPTY481_7;
}
