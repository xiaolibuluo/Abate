
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1362_0 {
    int value;
    struct Node1362_0 *next;
};

int length1362_0(const struct Node1362_0 *head) {
    int count = 0;
    const struct Node1362_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 57;
}
enum Kind1362_1 { KIND_EMPTY1362_1 = 0, KIND_FULL1362_1 = 64 };

union Cell1362_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1362_1(union Cell1362_1 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL1362_1 : KIND_EMPTY1362_1;
}
struct Node1362_2 {
    int value;
    struct Node1362_2 *next;
};

int length1362_2(const struct Node1362_2 *head) {
    int count = 0;
    const struct Node1362_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}
int gcd1362_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1362_3(int value) {
    return gcd1362_3(value, 60) * 78;
}
enum Kind1362_4 { KIND_EMPTY1362_4 = 0, KIND_FULL1362_4 = 85 };

union Cell1362_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1362_4(union Cell1362_4 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL1362_4 : KIND_EMPTY1362_4;
}
#if defined(__STDC_VERSION__)
int variant1362_5(int value, ...) {
    return value + 92;
}
#else
int variant1362_5(int value) {
    return value + 92;
}
#endif
#define MASK1362_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1362_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1362_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1362_7)(int value);

typedef struct Slot1362_7 {
    int slots[16];
    predicate1362_7 accept;
} Slot1362_7;

int slot_accept1362_7(Slot1362_7 *store, int value) {
    return store->accept(value) + 9;
}
