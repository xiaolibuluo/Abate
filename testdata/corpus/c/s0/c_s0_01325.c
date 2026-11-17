
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant907_0(int value, ...) {
    return value + 59;
}
#else
int variant907_0(int value) {
    return value + 59;
}
#endif
#define MASK907_1(v) (((v) << 4) | ((v) & 0x0F))

int classify907_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK907_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK907_2(v) (((v) << 4) | ((v) & 0x0F))

int classify907_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK907_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate907_3)(int value);

typedef struct Slot907_3 {
    int slots[16];
    predicate907_3 accept;
} Slot907_3;

int slot_accept907_3(Slot907_3 *store, int value) {
    return store->accept(value) + 80;
}
enum Kind907_4 { KIND_EMPTY907_4 = 0, KIND_FULL907_4 = 87 };

union Cell907_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind907_4(union Cell907_4 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL907_4 : KIND_EMPTY907_4;
}
int gcd907_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale907_5(int value) {
    return gcd907_5(value, 69) * 94;
}
