
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node951_0 {
    int value;
    struct Node951_0 *next;
};

int length951_0(const struct Node951_0 *head) {
    int count = 0;
    const struct Node951_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
int gcd951_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale951_1(int value) {
    return gcd951_1(value, 28) * 56;
}
enum Kind951_2 { KIND_EMPTY951_2 = 0, KIND_FULL951_2 = 63 };

union Cell951_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind951_2(union Cell951_2 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL951_2 : KIND_EMPTY951_2;
}
typedef int (*predicate951_3)(int value);

typedef struct Slot951_3 {
    int slots[16];
    predicate951_3 accept;
} Slot951_3;

int slot_accept951_3(Slot951_3 *store, int value) {
    return store->accept(value) + 70;
}
#define MASK951_4(v) (((v) << 4) | ((v) & 0x0F))

int classify951_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK951_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
