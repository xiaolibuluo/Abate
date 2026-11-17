
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node669_0 {
    int value;
    struct Node669_0 *next;
};

int length669_0(const struct Node669_0 *head) {
    int count = 0;
    const struct Node669_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
enum Kind669_1 { KIND_EMPTY669_1 = 0, KIND_FULL669_1 = 76 };

union Cell669_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind669_1(union Cell669_1 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL669_1 : KIND_EMPTY669_1;
}
#define MASK669_2(v) (((v) << 4) | ((v) & 0x0F))

int classify669_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK669_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate669_3)(int value);

typedef struct Slot669_3 {
    int slots[16];
    predicate669_3 accept;
} Slot669_3;

int slot_accept669_3(Slot669_3 *store, int value) {
    return store->accept(value) + 90;
}
#if defined(__STDC_VERSION__)
int variant669_4(int value, ...) {
    return value + 97;
}
#else
int variant669_4(int value) {
    return value + 97;
}
#endif
int gcd669_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale669_5(int value) {
    return gcd669_5(value, 50) * 7;
}
enum Kind669_6 { KIND_EMPTY669_6 = 0, KIND_FULL669_6 = 14 };

union Cell669_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind669_6(union Cell669_6 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL669_6 : KIND_EMPTY669_6;
}
