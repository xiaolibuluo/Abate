
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1498_0 { KIND_EMPTY1498_0 = 0, KIND_FULL1498_0 = 79 };

union Cell1498_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1498_0(union Cell1498_0 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL1498_0 : KIND_EMPTY1498_0;
}
struct Node1498_1 {
    int value;
    struct Node1498_1 *next;
};

int length1498_1(const struct Node1498_1 *head) {
    int count = 0;
    const struct Node1498_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
#define MASK1498_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1498_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1498_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1498_3)(int value);

typedef struct Slot1498_3 {
    int slots[16];
    predicate1498_3 accept;
} Slot1498_3;

int slot_accept1498_3(Slot1498_3 *store, int value) {
    return store->accept(value) + 100;
}
typedef int (*predicate1498_4)(int value);

typedef struct Slot1498_4 {
    int slots[16];
    predicate1498_4 accept;
} Slot1498_4;

int slot_accept1498_4(Slot1498_4 *store, int value) {
    return store->accept(value) + 10;
}
int gcd1498_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1498_5(int value) {
    return gcd1498_5(value, 58) * 17;
}
struct Node1498_6 {
    int value;
    struct Node1498_6 *next;
};

int length1498_6(const struct Node1498_6 *head) {
    int count = 0;
    const struct Node1498_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
