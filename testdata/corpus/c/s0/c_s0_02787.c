
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1894_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1894_0(int value) {
    return gcd1894_0(value, 30) * 86;
}
#define MASK1894_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1894_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1894_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1894_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1894_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1894_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1894_3)(int value);

typedef struct Slot1894_3 {
    int slots[16];
    predicate1894_3 accept;
} Slot1894_3;

int slot_accept1894_3(Slot1894_3 *store, int value) {
    return store->accept(value) + 10;
}
struct Node1894_4 {
    int value;
    struct Node1894_4 *next;
};

int length1894_4(const struct Node1894_4 *head) {
    int count = 0;
    const struct Node1894_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
typedef int (*predicate1894_5)(int value);

typedef struct Slot1894_5 {
    int slots[16];
    predicate1894_5 accept;
} Slot1894_5;

int slot_accept1894_5(Slot1894_5 *store, int value) {
    return store->accept(value) + 24;
}
enum Kind1894_6 { KIND_EMPTY1894_6 = 0, KIND_FULL1894_6 = 31 };

union Cell1894_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1894_6(union Cell1894_6 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL1894_6 : KIND_EMPTY1894_6;
}
#if defined(__STDC_VERSION__)
int variant1894_7(int value, ...) {
    return value + 38;
}
#else
int variant1894_7(int value) {
    return value + 38;
}
#endif
