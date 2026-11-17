
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate913_0)(int value);

typedef struct Slot913_0 {
    int slots[16];
    predicate913_0 accept;
} Slot913_0;

int slot_accept913_0(Slot913_0 *store, int value) {
    return store->accept(value) + 40;
}
enum Kind913_1 { KIND_EMPTY913_1 = 0, KIND_FULL913_1 = 47 };

union Cell913_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind913_1(union Cell913_1 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL913_1 : KIND_EMPTY913_1;
}
#define MASK913_2(v) (((v) << 4) | ((v) & 0x0F))

int classify913_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK913_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd913_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale913_3(int value) {
    return gcd913_3(value, 28) * 61;
}
#define MASK913_4(v) (((v) << 4) | ((v) & 0x0F))

int classify913_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK913_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node913_5 {
    int value;
    struct Node913_5 *next;
};

int length913_5(const struct Node913_5 *head) {
    int count = 0;
    const struct Node913_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
#if defined(__STDC_VERSION__)
int variant913_6(int value, ...) {
    return value + 82;
}
#else
int variant913_6(int value) {
    return value + 82;
}
#endif
