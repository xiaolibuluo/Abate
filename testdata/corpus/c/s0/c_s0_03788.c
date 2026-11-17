
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2564_0 { KIND_EMPTY2564_0 = 0, KIND_FULL2564_0 = 66 };

union Cell2564_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2564_0(union Cell2564_0 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL2564_0 : KIND_EMPTY2564_0;
}
#define MASK2564_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2564_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2564_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2564_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2564_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2564_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2564_3)(int value);

typedef struct Slot2564_3 {
    int slots[16];
    predicate2564_3 accept;
} Slot2564_3;

int slot_accept2564_3(Slot2564_3 *store, int value) {
    return store->accept(value) + 87;
}
struct Node2564_4 {
    int value;
    struct Node2564_4 *next;
};

int length2564_4(const struct Node2564_4 *head) {
    int count = 0;
    const struct Node2564_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
int gcd2564_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2564_5(int value) {
    return gcd2564_5(value, 17) * 101;
}
