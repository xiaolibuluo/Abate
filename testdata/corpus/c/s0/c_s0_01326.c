
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK908_0(v) (((v) << 4) | ((v) & 0x0F))

int classify908_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK908_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd908_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale908_1(int value) {
    return gcd908_1(value, 32) * 79;
}
struct Node908_2 {
    int value;
    struct Node908_2 *next;
};

int length908_2(const struct Node908_2 *head) {
    int count = 0;
    const struct Node908_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
int gcd908_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale908_3(int value) {
    return gcd908_3(value, 54) * 93;
}
typedef int (*predicate908_4)(int value);

typedef struct Slot908_4 {
    int slots[16];
    predicate908_4 accept;
} Slot908_4;

int slot_accept908_4(Slot908_4 *store, int value) {
    return store->accept(value) + 100;
}
enum Kind908_5 { KIND_EMPTY908_5 = 0, KIND_FULL908_5 = 10 };

union Cell908_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind908_5(union Cell908_5 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL908_5 : KIND_EMPTY908_5;
}
