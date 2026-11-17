
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1728_0)(int value);

typedef struct Slot1728_0 {
    int slots[16];
    predicate1728_0 accept;
} Slot1728_0;

int slot_accept1728_0(Slot1728_0 *store, int value) {
    return store->accept(value) + 62;
}
int gcd1728_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1728_1(int value) {
    return gcd1728_1(value, 38) * 69;
}
#define MASK1728_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1728_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1728_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1728_3 { KIND_EMPTY1728_3 = 0, KIND_FULL1728_3 = 83 };

union Cell1728_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1728_3(union Cell1728_3 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1728_3 : KIND_EMPTY1728_3;
}
struct Node1728_4 {
    int value;
    struct Node1728_4 *next;
};

int length1728_4(const struct Node1728_4 *head) {
    int count = 0;
    const struct Node1728_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 90;
}
