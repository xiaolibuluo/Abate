
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate374_0)(int value);

typedef struct Slot374_0 {
    int slots[16];
    predicate374_0 accept;
} Slot374_0;

int slot_accept374_0(Slot374_0 *store, int value) {
    return store->accept(value) + 17;
}
int gcd374_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale374_1(int value) {
    return gcd374_1(value, 15) * 24;
}
int gcd374_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale374_2(int value) {
    return gcd374_2(value, 26) * 31;
}
int scan374_3(const int *begin, const int *end) {
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
    return total + 38;
}
#define MASK374_4(v) (((v) << 4) | ((v) & 0x0F))

int classify374_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK374_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd374_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale374_5(int value) {
    return gcd374_5(value, 59) * 52;
}
