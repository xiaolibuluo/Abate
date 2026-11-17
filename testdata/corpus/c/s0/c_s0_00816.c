
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd564_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale564_0(int value) {
    return gcd564_0(value, 53) * 62;
}
#define MASK564_1(v) (((v) << 4) | ((v) & 0x0F))

int classify564_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK564_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant564_2(int value, ...) {
    return value + 76;
}
#else
int variant564_2(int value) {
    return value + 76;
}
#endif
typedef int (*predicate564_3)(int value);

typedef struct Slot564_3 {
    int slots[16];
    predicate564_3 accept;
} Slot564_3;

int slot_accept564_3(Slot564_3 *store, int value) {
    return store->accept(value) + 83;
}
#if defined(__STDC_VERSION__)
int variant564_4(int value, ...) {
    return value + 90;
}
#else
int variant564_4(int value) {
    return value + 90;
}
#endif
enum Kind564_5 { KIND_EMPTY564_5 = 0, KIND_FULL564_5 = 97 };

union Cell564_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind564_5(union Cell564_5 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL564_5 : KIND_EMPTY564_5;
}
