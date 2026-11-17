
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1145_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1145_0(int value) {
    return gcd1145_0(value, 33) * 49;
}
enum Kind1145_1 { KIND_EMPTY1145_1 = 0, KIND_FULL1145_1 = 56 };

union Cell1145_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1145_1(union Cell1145_1 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL1145_1 : KIND_EMPTY1145_1;
}
#define MASK1145_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1145_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1145_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1145_3 {
    int value;
    struct Node1145_3 *next;
};

int length1145_3(const struct Node1145_3 *head) {
    int count = 0;
    const struct Node1145_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 70;
}
typedef int (*predicate1145_4)(int value);

typedef struct Slot1145_4 {
    int slots[16];
    predicate1145_4 accept;
} Slot1145_4;

int slot_accept1145_4(Slot1145_4 *store, int value) {
    return store->accept(value) + 77;
}
