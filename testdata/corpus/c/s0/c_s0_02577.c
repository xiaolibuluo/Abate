
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1745_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1745_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1745_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1745_1)(int value);

typedef struct Slot1745_1 {
    int slots[16];
    predicate1745_1 accept;
} Slot1745_1;

int slot_accept1745_1(Slot1745_1 *store, int value) {
    return store->accept(value) + 96;
}
struct Node1745_2 {
    int value;
    struct Node1745_2 *next;
};

int length1745_2(const struct Node1745_2 *head) {
    int count = 0;
    const struct Node1745_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 6;
}
struct Node1745_3 {
    int value;
    struct Node1745_3 *next;
};

int length1745_3(const struct Node1745_3 *head) {
    int count = 0;
    const struct Node1745_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
int gcd1745_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1745_4(int value) {
    return gcd1745_4(value, 68) * 20;
}
enum Kind1745_5 { KIND_EMPTY1745_5 = 0, KIND_FULL1745_5 = 27 };

union Cell1745_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1745_5(union Cell1745_5 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL1745_5 : KIND_EMPTY1745_5;
}
#define MASK1745_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1745_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1745_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
