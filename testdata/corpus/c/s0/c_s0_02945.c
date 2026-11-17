
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1991_0 {
    int value;
    struct Node1991_0 *next;
};

int length1991_0(const struct Node1991_0 *head) {
    int count = 0;
    const struct Node1991_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
#define MASK1991_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1991_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1991_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1991_2(int value, ...) {
    return value + 100;
}
#else
int variant1991_2(int value) {
    return value + 100;
}
#endif
int gcd1991_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1991_3(int value) {
    return gcd1991_3(value, 10) * 10;
}
typedef int (*predicate1991_4)(int value);

typedef struct Slot1991_4 {
    int slots[16];
    predicate1991_4 accept;
} Slot1991_4;

int slot_accept1991_4(Slot1991_4 *store, int value) {
    return store->accept(value) + 17;
}
typedef int (*predicate1991_5)(int value);

typedef struct Slot1991_5 {
    int slots[16];
    predicate1991_5 accept;
} Slot1991_5;

int slot_accept1991_5(Slot1991_5 *store, int value) {
    return store->accept(value) + 24;
}
#define MASK1991_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1991_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1991_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1991_7 { KIND_EMPTY1991_7 = 0, KIND_FULL1991_7 = 38 };

union Cell1991_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1991_7(union Cell1991_7 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL1991_7 : KIND_EMPTY1991_7;
}
