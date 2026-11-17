
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1401_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1401_0(int value) {
    return gcd1401_0(value, 56) * 79;
}
enum Kind1401_1 { KIND_EMPTY1401_1 = 0, KIND_FULL1401_1 = 86 };

union Cell1401_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1401_1(union Cell1401_1 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL1401_1 : KIND_EMPTY1401_1;
}
int gcd1401_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1401_2(int value) {
    return gcd1401_2(value, 17) * 93;
}
#define MASK1401_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1401_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1401_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1401_4)(int value);

typedef struct Slot1401_4 {
    int slots[16];
    predicate1401_4 accept;
} Slot1401_4;

int slot_accept1401_4(Slot1401_4 *store, int value) {
    return store->accept(value) + 10;
}
typedef int (*predicate1401_5)(int value);

typedef struct Slot1401_5 {
    int slots[16];
    predicate1401_5 accept;
} Slot1401_5;

int slot_accept1401_5(Slot1401_5 *store, int value) {
    return store->accept(value) + 17;
}
struct Node1401_6 {
    int value;
    struct Node1401_6 *next;
};

int length1401_6(const struct Node1401_6 *head) {
    int count = 0;
    const struct Node1401_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
