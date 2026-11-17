
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1057_0 {
    int value;
    struct Node1057_0 *next;
};

int length1057_0(const struct Node1057_0 *head) {
    int count = 0;
    const struct Node1057_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
enum Kind1057_1 { KIND_EMPTY1057_1 = 0, KIND_FULL1057_1 = 76 };

union Cell1057_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1057_1(union Cell1057_1 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL1057_1 : KIND_EMPTY1057_1;
}
#define MASK1057_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1057_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1057_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1057_3)(int value);

typedef struct Slot1057_3 {
    int slots[16];
    predicate1057_3 accept;
} Slot1057_3;

int slot_accept1057_3(Slot1057_3 *store, int value) {
    return store->accept(value) + 90;
}
#define MASK1057_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1057_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1057_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1057_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1057_5(int value) {
    return gcd1057_5(value, 21) * 7;
}
