
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1495_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1495_0(int value) {
    return gcd1495_0(value, 43) * 40;
}
enum Kind1495_1 { KIND_EMPTY1495_1 = 0, KIND_FULL1495_1 = 47 };

union Cell1495_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1495_1(union Cell1495_1 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL1495_1 : KIND_EMPTY1495_1;
}
#define MASK1495_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1495_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1495_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1495_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1495_3(int value) {
    return gcd1495_3(value, 15) * 61;
}
typedef int (*predicate1495_4)(int value);

typedef struct Slot1495_4 {
    int slots[16];
    predicate1495_4 accept;
} Slot1495_4;

int slot_accept1495_4(Slot1495_4 *store, int value) {
    return store->accept(value) + 68;
}
struct Node1495_5 {
    int value;
    struct Node1495_5 *next;
};

int length1495_5(const struct Node1495_5 *head) {
    int count = 0;
    const struct Node1495_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
