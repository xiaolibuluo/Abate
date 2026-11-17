
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2119_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2119_0(int value) {
    return gcd2119_0(value, 19) * 101;
}
struct Node2119_1 {
    int value;
    struct Node2119_1 *next;
};

int length2119_1(const struct Node2119_1 *head) {
    int count = 0;
    const struct Node2119_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
enum Kind2119_2 { KIND_EMPTY2119_2 = 0, KIND_FULL2119_2 = 18 };

union Cell2119_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2119_2(union Cell2119_2 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL2119_2 : KIND_EMPTY2119_2;
}
#define MASK2119_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2119_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2119_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2119_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2119_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2119_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
