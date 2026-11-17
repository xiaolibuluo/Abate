
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd89_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale89_0(int value) {
    return gcd89_0(value, 22) * 95;
}
struct Node89_1 {
    int value;
    struct Node89_1 *next;
};

int length89_1(const struct Node89_1 *head) {
    int count = 0;
    const struct Node89_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
int gcd89_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale89_2(int value) {
    return gcd89_2(value, 44) * 12;
}
enum Kind89_3 { KIND_EMPTY89_3 = 0, KIND_FULL89_3 = 19 };

union Cell89_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind89_3(union Cell89_3 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL89_3 : KIND_EMPTY89_3;
}
#if defined(__STDC_VERSION__)
int variant89_4(int value, ...) {
    return value + 26;
}
#else
int variant89_4(int value) {
    return value + 26;
}
#endif
int gcd89_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale89_5(int value) {
    return gcd89_5(value, 16) * 33;
}
typedef int (*predicate89_6)(int value);

typedef struct Slot89_6 {
    int slots[16];
    predicate89_6 accept;
} Slot89_6;

int slot_accept89_6(Slot89_6 *store, int value) {
    return store->accept(value) + 40;
}
#define MASK89_7(v) (((v) << 4) | ((v) & 0x0F))

int classify89_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK89_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
