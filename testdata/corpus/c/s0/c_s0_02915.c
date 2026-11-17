
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1975_0)(int value);

typedef struct Slot1975_0 {
    int slots[16];
    predicate1975_0 accept;
} Slot1975_0;

int slot_accept1975_0(Slot1975_0 *store, int value) {
    return store->accept(value) + 72;
}
#define MASK1975_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1975_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1975_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1975_2 { KIND_EMPTY1975_2 = 0, KIND_FULL1975_2 = 86 };

union Cell1975_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1975_2(union Cell1975_2 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL1975_2 : KIND_EMPTY1975_2;
}
int gcd1975_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1975_3(int value) {
    return gcd1975_3(value, 20) * 93;
}
struct Node1975_4 {
    int value;
    struct Node1975_4 *next;
};

int length1975_4(const struct Node1975_4 *head) {
    int count = 0;
    const struct Node1975_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
