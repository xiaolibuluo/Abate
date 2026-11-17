
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1833_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1833_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1833_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1833_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1833_1(int value) {
    return gcd1833_1(value, 41) * 76;
}
typedef int (*predicate1833_2)(int value);

typedef struct Slot1833_2 {
    int slots[16];
    predicate1833_2 accept;
} Slot1833_2;

int slot_accept1833_2(Slot1833_2 *store, int value) {
    return store->accept(value) + 83;
}
struct Node1833_3 {
    int value;
    struct Node1833_3 *next;
};

int length1833_3(const struct Node1833_3 *head) {
    int count = 0;
    const struct Node1833_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 90;
}
int gcd1833_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1833_4(int value) {
    return gcd1833_4(value, 13) * 97;
}
enum Kind1833_5 { KIND_EMPTY1833_5 = 0, KIND_FULL1833_5 = 7 };

union Cell1833_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1833_5(union Cell1833_5 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL1833_5 : KIND_EMPTY1833_5;
}
