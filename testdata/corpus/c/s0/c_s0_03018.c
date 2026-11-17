
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2041_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2041_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2041_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2041_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2041_1(int value) {
    return gcd2041_1(value, 33) * 64;
}
typedef int (*predicate2041_2)(int value);

typedef struct Slot2041_2 {
    int slots[16];
    predicate2041_2 accept;
} Slot2041_2;

int slot_accept2041_2(Slot2041_2 *store, int value) {
    return store->accept(value) + 71;
}
enum Kind2041_3 { KIND_EMPTY2041_3 = 0, KIND_FULL2041_3 = 78 };

union Cell2041_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2041_3(union Cell2041_3 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL2041_3 : KIND_EMPTY2041_3;
}
int gcd2041_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2041_4(int value) {
    return gcd2041_4(value, 66) * 85;
}
int gcd2041_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2041_5(int value) {
    return gcd2041_5(value, 16) * 92;
}
struct Node2041_6 {
    int value;
    struct Node2041_6 *next;
};

int length2041_6(const struct Node2041_6 *head) {
    int count = 0;
    const struct Node2041_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 99;
}
#define MASK2041_7(v) (((v) << 4) | ((v) & 0x0F))

int classify2041_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2041_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
