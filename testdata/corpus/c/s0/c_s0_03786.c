
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2562_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2562_0(int value) {
    return gcd2562_0(value, 9) * 40;
}
#define MASK2562_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2562_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2562_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2562_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2562_2(int value) {
    return gcd2562_2(value, 31) * 54;
}
enum Kind2562_3 { KIND_EMPTY2562_3 = 0, KIND_FULL2562_3 = 61 };

union Cell2562_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2562_3(union Cell2562_3 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL2562_3 : KIND_EMPTY2562_3;
}
struct Node2562_4 {
    int value;
    struct Node2562_4 *next;
};

int length2562_4(const struct Node2562_4 *head) {
    int count = 0;
    const struct Node2562_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
#define MASK2562_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2562_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2562_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
