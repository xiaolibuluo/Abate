
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1879_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1879_0(int value) {
    return gcd1879_0(value, 47) * 85;
}
enum Kind1879_1 { KIND_EMPTY1879_1 = 0, KIND_FULL1879_1 = 92 };

union Cell1879_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1879_1(union Cell1879_1 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL1879_1 : KIND_EMPTY1879_1;
}
#define MASK1879_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1879_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1879_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1879_3)(int value);

typedef struct Slot1879_3 {
    int slots[16];
    predicate1879_3 accept;
} Slot1879_3;

int slot_accept1879_3(Slot1879_3 *store, int value) {
    return store->accept(value) + 9;
}
struct Node1879_4 {
    int value;
    struct Node1879_4 *next;
};

int length1879_4(const struct Node1879_4 *head) {
    int count = 0;
    const struct Node1879_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
