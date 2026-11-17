
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1303_0 {
    int value;
    struct Node1303_0 *next;
};

int length1303_0(const struct Node1303_0 *head) {
    int count = 0;
    const struct Node1303_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
#if defined(__STDC_VERSION__)
int variant1303_1(int value, ...) {
    return value + 73;
}
#else
int variant1303_1(int value) {
    return value + 73;
}
#endif
typedef int (*predicate1303_2)(int value);

typedef struct Slot1303_2 {
    int slots[16];
    predicate1303_2 accept;
} Slot1303_2;

int slot_accept1303_2(Slot1303_2 *store, int value) {
    return store->accept(value) + 80;
}
enum Kind1303_3 { KIND_EMPTY1303_3 = 0, KIND_FULL1303_3 = 87 };

union Cell1303_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1303_3(union Cell1303_3 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL1303_3 : KIND_EMPTY1303_3;
}
int gcd1303_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1303_4(int value) {
    return gcd1303_4(value, 24) * 94;
}
#define MASK1303_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1303_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1303_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1303_6 {
    int value;
    struct Node1303_6 *next;
};

int length1303_6(const struct Node1303_6 *head) {
    int count = 0;
    const struct Node1303_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
